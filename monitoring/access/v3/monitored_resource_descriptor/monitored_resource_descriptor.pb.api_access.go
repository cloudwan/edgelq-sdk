// Code generated by protoc-gen-goten-access
// Resource: MonitoredResourceDescriptor
// DO NOT EDIT!!!

package monitored_resource_descriptor_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	monitored_resource_descriptor_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v3/monitored_resource_descriptor"
	monitored_resource_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/monitored_resource_descriptor"
)

var (
	_ = context.Context(nil)
	_ = fmt.GoStringer(nil)

	_ = grpc.ClientConnInterface(nil)
	_ = codes.NotFound
	_ = status.Status{}

	_ = gotenaccess.Watcher(nil)
	_ = watch_type.WatchType_STATEFUL
	_ = gotenresource.ListQuery(nil)
)

type apiMonitoredResourceDescriptorAccess struct {
	client monitored_resource_descriptor_client.MonitoredResourceDescriptorServiceClient
}

func NewApiMonitoredResourceDescriptorAccess(client monitored_resource_descriptor_client.MonitoredResourceDescriptorServiceClient) monitored_resource_descriptor.MonitoredResourceDescriptorAccess {
	return &apiMonitoredResourceDescriptorAccess{client: client}
}

func (a *apiMonitoredResourceDescriptorAccess) GetMonitoredResourceDescriptor(ctx context.Context, query *monitored_resource_descriptor.GetQuery) (*monitored_resource_descriptor.MonitoredResourceDescriptor, error) {
	request := &monitored_resource_descriptor_client.GetMonitoredResourceDescriptorRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetMonitoredResourceDescriptor(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiMonitoredResourceDescriptorAccess) BatchGetMonitoredResourceDescriptors(ctx context.Context, refs []*monitored_resource_descriptor.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &monitored_resource_descriptor_client.BatchGetMonitoredResourceDescriptorsRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetMonitoredResourceDescriptors(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[monitored_resource_descriptor.Name]*monitored_resource_descriptor.MonitoredResourceDescriptor, len(refs))
	for _, resolvedRes := range resp.GetMonitoredResourceDescriptors() {
		resultMap[*resolvedRes.GetName()] = resolvedRes
	}
	for _, ref := range refs {
		resolvedRes := resultMap[ref.Name]
		if resolvedRes != nil {
			ref.Resolve(resolvedRes)
		}
	}
	if batchGetOpts.MustResolveAll() && len(resp.GetMissing()) > 0 {
		return status.Errorf(codes.NotFound, "Number of references not found: %d", len(resp.GetMissing()))
	}
	return nil
}

func (a *apiMonitoredResourceDescriptorAccess) QueryMonitoredResourceDescriptors(ctx context.Context, query *monitored_resource_descriptor.ListQuery) (*monitored_resource_descriptor.QueryResultSnapshot, error) {
	request := &monitored_resource_descriptor_client.ListMonitoredResourceDescriptorsRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListMonitoredResourceDescriptors(ctx, request)
	if err != nil {
		return nil, err
	}
	return &monitored_resource_descriptor.QueryResultSnapshot{
		MonitoredResourceDescriptors: resp.MonitoredResourceDescriptors,
		NextPageCursor:               resp.NextPageToken,
		PrevPageCursor:               resp.PrevPageToken,
	}, nil
}

func (a *apiMonitoredResourceDescriptorAccess) WatchMonitoredResourceDescriptor(ctx context.Context, query *monitored_resource_descriptor.GetQuery, observerCb func(*monitored_resource_descriptor.MonitoredResourceDescriptorChange) error) error {
	request := &monitored_resource_descriptor_client.WatchMonitoredResourceDescriptorRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchMonitoredResourceDescriptor(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		resp, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		change := resp.GetChange()
		if err := observerCb(change); err != nil {
			return err
		}
	}
}

func (a *apiMonitoredResourceDescriptorAccess) WatchMonitoredResourceDescriptors(ctx context.Context, query *monitored_resource_descriptor.WatchQuery, observerCb func(*monitored_resource_descriptor.QueryResultChange) error) error {
	request := &monitored_resource_descriptor_client.WatchMonitoredResourceDescriptorsRequest{
		Filter:       query.Filter,
		FieldMask:    query.Mask,
		MaxChunkSize: int32(query.ChunkSize),
		Type:         query.WatchType,
		ResumeToken:  query.ResumeToken,
	}
	if query.Pager != nil {
		request.OrderBy = query.Pager.OrderBy
		request.PageSize = int32(query.Pager.Limit)
		request.PageToken = query.Pager.Cursor
	}
	changesStream, initErr := a.client.WatchMonitoredResourceDescriptors(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &monitored_resource_descriptor.QueryResultChange{
			Changes:      respChange.MonitoredResourceDescriptorChanges,
			IsCurrent:    respChange.IsCurrent,
			IsHardReset:  respChange.IsHardReset,
			IsSoftReset:  respChange.IsSoftReset,
			ResumeToken:  respChange.ResumeToken,
			SnapshotSize: respChange.SnapshotSize,
		}
		if respChange.PageTokenChange != nil {
			changesWithPaging.PrevPageCursor = respChange.PageTokenChange.PrevPageToken
			changesWithPaging.NextPageCursor = respChange.PageTokenChange.NextPageToken
		}
		if err := observerCb(changesWithPaging); err != nil {
			return err
		}
	}
}

func (a *apiMonitoredResourceDescriptorAccess) SaveMonitoredResourceDescriptor(ctx context.Context, res *monitored_resource_descriptor.MonitoredResourceDescriptor, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetMonitoredResourceDescriptor(ctx, &monitored_resource_descriptor.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &monitored_resource_descriptor_client.UpdateMonitoredResourceDescriptorRequest{
			MonitoredResourceDescriptor: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*monitored_resource_descriptor.MonitoredResourceDescriptor_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &monitored_resource_descriptor_client.UpdateMonitoredResourceDescriptorRequest_CAS{
				ConditionalState: conditionalState.(*monitored_resource_descriptor.MonitoredResourceDescriptor),
				FieldMask:        mask.(*monitored_resource_descriptor.MonitoredResourceDescriptor_FieldMask),
			}
		}
		_, err := a.client.UpdateMonitoredResourceDescriptor(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &monitored_resource_descriptor_client.CreateMonitoredResourceDescriptorRequest{
			MonitoredResourceDescriptor: res,
		}
		_, err := a.client.CreateMonitoredResourceDescriptor(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiMonitoredResourceDescriptorAccess) DeleteMonitoredResourceDescriptor(ctx context.Context, ref *monitored_resource_descriptor.Reference, opts ...gotenresource.DeleteOption) error {
	request := &monitored_resource_descriptor_client.DeleteMonitoredResourceDescriptorRequest{
		Name: ref,
	}
	_, err := a.client.DeleteMonitoredResourceDescriptor(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(monitored_resource_descriptor.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return monitored_resource_descriptor.AsAnyCastAccess(NewApiMonitoredResourceDescriptorAccess(monitored_resource_descriptor_client.NewMonitoredResourceDescriptorServiceClient(cc)))
	})
}
