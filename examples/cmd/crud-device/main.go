package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/pflag"

	"github.com/cloudwan/goten-sdk/types/meta"
	"github.com/cloudwan/goten-sdk/types/view"

	cdevice "github.com/cloudwan/edgelq-sdk/devices/client/v1/device"
	rdevice "github.com/cloudwan/edgelq-sdk/devices/resources/v1/device"

	"github.com/cloudwan/edgelq-sdk/examples/utils"
)

/* This example focuses on simple create-list-update-get-delete of Device resource.
 *
 * It needs simple parameters (endpoint, credentials, projectID, regionID). Based on those,
 * it creates device, makes some list query (that should contain your device), then just
 * executes some update, get, finally delete.
 *
 * You need to provide project ID and regionID that actually exist in environment you want to use.
 *
 * In order to execute this example, you need either user account (prepare access token, like one from web browser),
 * or credentials for service account (JSON format, as defined as in edgelq-sdk/common/api/credentials.proto -
 * ServiceAccount). Of course, user or service account will need permission to manage Pods in project you
 * select, so this is pre-requirement.
 *
 * Of course, Device onboarding and record creation is not as simple as that, but its good to show how to use
 * Golang library for EdgeLQ. However, as we are showing API example rather than providing Device study, we can
 * keep our Device extremely simple.
 */

func main() {
	ctx := context.Background()

	// ------------------------- Get params and just verify they were provided -------------------------
	endpoint := pflag.StringP("endpoint", "e", "", "API endpoint (for example, devices.stg01b.edgelq.com:443)")
	projectId := pflag.StringP("project", "p", "", "Project ID for this demo")
	regionId := pflag.StringP("region", "r", "", "Region ID for this demo")
	accessToken := pflag.StringP("access-token", "a", "", "Active token for your user")
	credsFile := pflag.StringP("credentials", "c", "", "Path to service account credentials file")
	pflag.Parse()
	if *projectId == "" {
		pflag.Usage()
		panic(fmt.Errorf("no project ID was provided"))
	}
	if *regionId == "" {
		pflag.Usage()
		panic(fmt.Errorf("no region ID was provided"))
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

	// ------------------------- Lets begin! -------------------------
	// server will generate device ID, we need to provide parent only
	createdDevice, err := devicesClient.CreateDevice(ctx, &cdevice.CreateDeviceRequest{
		Parent: rdevice.NewNameBuilder().SetProjectId(*projectId).SetRegionId(*regionId).Parent(),
		Device: &rdevice.Device{
			Metadata: &meta.Meta{
				Labels: map[string]string{"device_type": "test"},
			},
			DisplayName: "Example Device",
		},
	})
	if err != nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("failed to create device: %s\n", err))
		os.Exit(1)
	}
	_, _ = os.Stderr.WriteString(fmt.Sprintf("device with name %s has been created\n", createdDevice.GetName()))

	// get all devices like that
	listDevicesResp, err := devicesClient.ListDevices(ctx, &cdevice.ListDevicesRequest{
		Parent: rdevice.NewNameBuilder().SetProjectId(*projectId).SetRegionId(*regionId).Parent(),
		Filter: rdevice.NewFilterBuilder().Where().Metadata().Labels().WithKey("device_type").Eq("test").Filter(),
		View:   view.View_FULL,
	})
	if err != nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("failed to list devices: %s\n", err))
		os.Exit(1)
	}

	// in theory, if by any chance there are thousands of devices matching parent/filter in request,
	// we may not find our device at all! Server applies paging automatically even if we did not specify
	// page size in the request.
	for _, obtainedDevice := range listDevicesResp.GetDevices() {
		if obtainedDevice.GetName().GotenEqual(createdDevice.GetName()) {
			_, _ = os.Stderr.WriteString(fmt.Sprintf("we found our device!\n"))
		} else {
			_, _ = os.Stderr.WriteString(fmt.Sprintf("found device %s with same test label\n", obtainedDevice.GetName()))
		}
	}

	// update our device. Since we are updating only field path metadata.labels.other_label, we should
	// not have cleared previous label
	updatedDevice, err := devicesClient.UpdateDevice(ctx, &cdevice.UpdateDeviceRequest{
		Device: &rdevice.Device{
			Name: createdDevice.GetName(),
			Metadata: &meta.Meta{
				Labels: map[string]string{"other_label": "other"},
			},
		},
		UpdateMask: &rdevice.Device_FieldMask{Paths: []rdevice.Device_FieldPath{
			rdevice.NewDeviceFieldPathBuilder().Metadata().Labels().WithKey("other_label").FieldPath(),
		}},
	})
	if err != nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("failed to update device: %s\n", err))
		os.Exit(1)
	}
	if updatedDevice.GetMetadata().GetLabels()["device_type"] != "test" {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("label device_type is incorrect, updated device looks like: %s\n", updatedDevice))
		os.Exit(1)
	}
	if updatedDevice.GetMetadata().GetLabels()["other_label"] != "other" {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("label other_label is incorrect, updated device looks like: %s\n", updatedDevice))
		os.Exit(1)
	}

	// lets read device, but only selected label... view NAME puts "name" and "displayName" in field
	// mask, but thats all we should receive.
	obtainedDevice, err := devicesClient.GetDevice(ctx, &cdevice.GetDeviceRequest{
		Name: updatedDevice.GetName(),
		View: view.View_NAME,
		FieldMask: &rdevice.Device_FieldMask{Paths: []rdevice.Device_FieldPath{
			rdevice.NewDeviceFieldPathBuilder().Metadata().Labels().WithKey("other_label").FieldPath(),
		}},
	})
	if err != nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("failed to read device: %s\n", err))
		os.Exit(1)
	}
	if obtainedDevice.GetMetadata().GetLabels()["device_type"] != "" {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("label device_type is not empty, obtained device looks like: %s\n", obtainedDevice))
		os.Exit(1)
	}
	if obtainedDevice.GetMetadata().GetLabels()["other_label"] != "other" {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("label other_label is incorrect, obtained device looks like: %s\n", obtainedDevice))
		os.Exit(1)
	}

	// just delete
	_, err = devicesClient.DeleteDevice(ctx, &cdevice.DeleteDeviceRequest{Name: createdDevice.GetName()})
	if err != nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("failed to delete device: %s\n", err))
		os.Exit(1)
	}
}
