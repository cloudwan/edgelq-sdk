// Code generated by protoc-gen-goten-access
// Resource: MethodDescriptor
// DO NOT EDIT!!!

package method_descriptor_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	method_descriptor_client "github.com/cloudwan/edgelq-sdk/audit/client/v1alpha/method_descriptor"
	method_descriptor "github.com/cloudwan/edgelq-sdk/audit/resources/v1alpha/method_descriptor"
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

type apiMethodDescriptorAccess struct {
	client method_descriptor_client.MethodDescriptorServiceClient
}

func NewApiMethodDescriptorAccess(client method_descriptor_client.MethodDescriptorServiceClient) method_descriptor.MethodDescriptorAccess {
	return &apiMethodDescriptorAccess{client: client}
}

func (a *apiMethodDescriptorAccess) GetMethodDescriptor(ctx context.Context, query *method_descriptor.GetQuery) (*method_descriptor.MethodDescriptor, error) {
	request := &method_descriptor_client.GetMethodDescriptorRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetMethodDescriptor(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiMethodDescriptorAccess) BatchGetMethodDescriptors(ctx context.Context, refs []*method_descriptor.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &method_descriptor_client.BatchGetMethodDescriptorsRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetMethodDescriptors(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[method_descriptor.Name]*method_descriptor.MethodDescriptor, len(refs))
	for _, resolvedRes := range resp.GetMethodDescriptors() {
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

func (a *apiMethodDescriptorAccess) QueryMethodDescriptors(ctx context.Context, query *method_descriptor.ListQuery) (*method_descriptor.QueryResultSnapshot, error) {
	request := &method_descriptor_client.ListMethodDescriptorsRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListMethodDescriptors(ctx, request)
	if err != nil {
		return nil, err
	}
	return &method_descriptor.QueryResultSnapshot{
		MethodDescriptors: resp.MethodDescriptors,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
	}, nil
}

func (a *apiMethodDescriptorAccess) WatchMethodDescriptor(ctx context.Context, query *method_descriptor.GetQuery, observerCb func(*method_descriptor.MethodDescriptorChange) error) error {
	request := &method_descriptor_client.WatchMethodDescriptorRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchMethodDescriptor(ctx, request)
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

func (a *apiMethodDescriptorAccess) WatchMethodDescriptors(ctx context.Context, query *method_descriptor.WatchQuery, observerCb func(*method_descriptor.QueryResultChange) error) error {
	request := &method_descriptor_client.WatchMethodDescriptorsRequest{
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
	changesStream, initErr := a.client.WatchMethodDescriptors(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &method_descriptor.QueryResultChange{
			Changes:      respChange.MethodDescriptorChanges,
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

func (a *apiMethodDescriptorAccess) SaveMethodDescriptor(ctx context.Context, res *method_descriptor.MethodDescriptor, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetMethodDescriptor(ctx, &method_descriptor.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &method_descriptor_client.UpdateMethodDescriptorRequest{
			MethodDescriptor: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*method_descriptor.MethodDescriptor_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &method_descriptor_client.UpdateMethodDescriptorRequest_CAS{
				ConditionalState: conditionalState.(*method_descriptor.MethodDescriptor),
				FieldMask:        mask.(*method_descriptor.MethodDescriptor_FieldMask),
			}
		}
		_, err := a.client.UpdateMethodDescriptor(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &method_descriptor_client.CreateMethodDescriptorRequest{
			MethodDescriptor: res,
		}
		_, err := a.client.CreateMethodDescriptor(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiMethodDescriptorAccess) DeleteMethodDescriptor(ctx context.Context, ref *method_descriptor.Reference, opts ...gotenresource.DeleteOption) error {
	return fmt.Errorf("Delete operation on MethodDescriptor is prohibited")
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(method_descriptor.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return method_descriptor.AsAnyCastAccess(NewApiMethodDescriptorAccess(method_descriptor_client.NewMethodDescriptorServiceClient(cc)))
	})
}
