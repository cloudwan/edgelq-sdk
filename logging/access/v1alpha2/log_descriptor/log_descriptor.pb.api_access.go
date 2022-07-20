// Code generated by protoc-gen-goten-access
// Resource: LogDescriptor
// DO NOT EDIT!!!

package log_descriptor_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	log_descriptor_client "github.com/cloudwan/edgelq-sdk/logging/client/v1alpha2/log_descriptor"
	log_descriptor "github.com/cloudwan/edgelq-sdk/logging/resources/v1alpha2/log_descriptor"
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

type apiLogDescriptorAccess struct {
	client log_descriptor_client.LogDescriptorServiceClient
}

func NewApiLogDescriptorAccess(client log_descriptor_client.LogDescriptorServiceClient) log_descriptor.LogDescriptorAccess {
	return &apiLogDescriptorAccess{client: client}
}

func (a *apiLogDescriptorAccess) GetLogDescriptor(ctx context.Context, query *log_descriptor.GetQuery) (*log_descriptor.LogDescriptor, error) {
	request := &log_descriptor_client.GetLogDescriptorRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetLogDescriptor(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiLogDescriptorAccess) BatchGetLogDescriptors(ctx context.Context, refs []*log_descriptor.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &log_descriptor_client.BatchGetLogDescriptorsRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetLogDescriptors(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[log_descriptor.Name]*log_descriptor.LogDescriptor, len(refs))
	for _, resolvedRes := range resp.GetLogDescriptors() {
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

func (a *apiLogDescriptorAccess) QueryLogDescriptors(ctx context.Context, query *log_descriptor.ListQuery) (*log_descriptor.QueryResultSnapshot, error) {
	request := &log_descriptor_client.ListLogDescriptorsRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListLogDescriptors(ctx, request)
	if err != nil {
		return nil, err
	}
	return &log_descriptor.QueryResultSnapshot{
		LogDescriptors: resp.LogDescriptors,
		NextPageCursor: resp.NextPageToken,
		PrevPageCursor: resp.PrevPageToken,
	}, nil
}

func (a *apiLogDescriptorAccess) WatchLogDescriptor(ctx context.Context, query *log_descriptor.GetQuery, observerCb func(*log_descriptor.LogDescriptorChange) error) error {
	request := &log_descriptor_client.WatchLogDescriptorRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchLogDescriptor(ctx, request)
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

func (a *apiLogDescriptorAccess) WatchLogDescriptors(ctx context.Context, query *log_descriptor.WatchQuery, observerCb func(*log_descriptor.QueryResultChange) error) error {
	request := &log_descriptor_client.WatchLogDescriptorsRequest{
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
	changesStream, initErr := a.client.WatchLogDescriptors(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &log_descriptor.QueryResultChange{
			Changes:      respChange.LogDescriptorChanges,
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

func (a *apiLogDescriptorAccess) SaveLogDescriptor(ctx context.Context, res *log_descriptor.LogDescriptor, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetLogDescriptor(ctx, &log_descriptor.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &log_descriptor_client.UpdateLogDescriptorRequest{
			LogDescriptor: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*log_descriptor.LogDescriptor_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &log_descriptor_client.UpdateLogDescriptorRequest_CAS{
				ConditionalState: conditionalState.(*log_descriptor.LogDescriptor),
				FieldMask:        mask.(*log_descriptor.LogDescriptor_FieldMask),
			}
		}
		_, err := a.client.UpdateLogDescriptor(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &log_descriptor_client.CreateLogDescriptorRequest{
			LogDescriptor: res,
		}
		_, err := a.client.CreateLogDescriptor(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiLogDescriptorAccess) DeleteLogDescriptor(ctx context.Context, ref *log_descriptor.Reference, opts ...gotenresource.DeleteOption) error {
	request := &log_descriptor_client.DeleteLogDescriptorRequest{
		Name: ref,
	}
	_, err := a.client.DeleteLogDescriptor(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(log_descriptor.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return log_descriptor.AsAnyCastAccess(NewApiLogDescriptorAccess(log_descriptor_client.NewLogDescriptorServiceClient(cc)))
	})
}