package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/encoding/protojson"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
	"github.com/cloudwan/goten-sdk/types/view"
	"github.com/cloudwan/goten-sdk/types/watch_type"

	adevice "github.com/cloudwan/edgelq-sdk/devices/access/v1/device"
	cdevice "github.com/cloudwan/edgelq-sdk/devices/client/v1/device"
	rdevice "github.com/cloudwan/edgelq-sdk/devices/resources/v1/device"

	"github.com/cloudwan/edgelq-sdk/examples/utils"
)

/* This example focuses on Watcher component.
 *
 * Based on provided parameters (endpoint, credentials, projectID) it watches Device resources
 * from Devices service in real-time - printing every addition, deletion and modification (diff).
 *
 * In order to use execute example, you need either user account (prepare access token, like one from web browser),
 * or credentials for service account (JSON format, as defined as in edgelq-sdk/common/api/credentials.proto -
 * ServiceAccount). Of course, user or service account will need permission to watch devices in project you
 * select, so this is pre-requirement.
 */

func main() {
	ctx := context.Background()

	// ------------------------- Get params and just verify they were provided -------------------------
	endpoint := pflag.StringP("endpoint", "e", "", "API endpoint (for example, devices.stg01b.edgelq.com:443)")
	projectId := pflag.StringP("project", "p", "", "Name of the project to watch devices from")
	accessToken := pflag.StringP("access-token", "a", "", "Active token for your user")
	credsFile := pflag.StringP("credentials", "c", "", "Path to service account credentials file")
	pflag.Parse()
	if *projectId == "" {
		pflag.Usage()
		panic(fmt.Errorf("no project ID was provided"))
	}
	if *endpoint == "" {
		pflag.Usage()
		panic(fmt.Errorf("no endpoint was provided"))
	}
	if *accessToken == "" && *credsFile == "" {
		pflag.Usage()
		panic(fmt.Errorf("access token OR credentials file are necessary in order to execute this example"))
	}

	// ------------------------- Create GRPC connection -------------------------
	grpcConn := utils.Dial(ctx, *endpoint, *accessToken, *credsFile)
	devicesClient := cdevice.NewDeviceServiceClient(grpcConn)

	// ------------------------- Create Watcher -------------------------
	// We are interested in devices from selected project (as provided in argument),
	// we dont care about region, so we put a wildcard there.
	cfg := &adevice.WatcherConfig{
		WatcherConfig: gotenaccess.NewWatcherConfig(),
		WatchType:     watch_type.WatchType_STATELESS,
		View:          view.View_FULL,
	}
	filter := &adevice.WatcherFilterParams{
		Parent: rdevice.NewNameBuilder().SetProjectId(*projectId).
			SetRegionId(gotenresource.WildcardId).Parent(),
	}
	watcher := adevice.NewWatcher(devicesClient, cfg, filter)

	// ------------------------- Run Watcher on separate goroutine -------------------------
	go func() {
		err := watcher.Run(ctx)
		_, _ = os.Stderr.WriteString(fmt.Sprintf("Watch finished with error: %s\n", err))
		os.Exit(1)
	}()

	// ------------------------- Observe in real-time changes -------------------------
	for {
		select {
		case <-ctx.Done():
			_, _ = os.Stdout.WriteString(fmt.Sprintf("watch-devices finished with err=%s\n", ctx.Err()))
			os.Exit(1)

		case evt := <-watcher.Events():
			// If we get this event, then we will later get Resync with fresh devices snapshot
			if evt.LostSync() {
				_, _ = os.Stdout.WriteString(fmt.Sprintf("watch-devices has lost sync, will need to recover\n\n"))
			}
			// We get this for the first iteration and whenever we lost sync
			if evt.Resync() {
				_, _ = os.Stdout.WriteString(fmt.Sprintf("watch-devices sync received\n\n"))
			}

			// Just print...
			for _, change := range evt.Changes {
				if change.IsAdd() {
					jsonData, err := protojson.Marshal(change.GetAdded())
					if err != nil {
						_, _ = os.Stdout.WriteString(fmt.Sprintf("\nerror marshaling device %s: %s\n", change.GetName(), err))
					} else {
						_, _ = os.Stdout.WriteString(fmt.Sprintf("\nadded new device %s: %s\n", change.GetName(), string(jsonData)))
					}
				} else if change.IsModify() {
					currentState := change.GetCurrent()
					previousState := change.GetPrevious()
					updateMask := currentState.MakeDiffFieldMask(previousState)
					oldValues := updateMask.Project(previousState)
					newValues := updateMask.Project(currentState)

					oldValuesJson, err := protojson.Marshal(oldValues)
					if err != nil {
						_, _ = os.Stdout.WriteString(fmt.Sprintf("\nerror marshaling old values for device %s: %s\n", change.GetName(), err))
						continue
					}
					newValuesJson, err := protojson.Marshal(newValues)
					if err != nil {
						_, _ = os.Stdout.WriteString(fmt.Sprintf("\nerror marshaling new values for device %s: %s\n", change.GetName(), err))
						continue
					}
					_, _ = os.Stdout.WriteString(fmt.Sprintf("\ndevice %s has been modified:\nOld values: %s\nNewValues: %s\n",
						change.GetName(), string(oldValuesJson), string(newValuesJson)))
				} else if change.IsDelete() {
					// If we had specified a filter, delete could also mean that device has left "view".
					_, _ = os.Stdout.WriteString(fmt.Sprintf("\ndeleted device %s\n", change.GetName()))
				}
			}
		}
	}
}
