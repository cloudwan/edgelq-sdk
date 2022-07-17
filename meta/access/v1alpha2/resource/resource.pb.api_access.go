// Code generated by protoc-gen-goten-access
// Resource: Resource
// DO NOT EDIT!!!

package resource_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	resource_client "github.com/cloudwan/edgelq-sdk/meta/client/v1alpha2/resource"
	resource "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/resource"
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

type apiResourceAccess struct {
	client resource_client.ResourceServiceClient
}

func NewApiResourceAccess(client resource_client.ResourceServiceClient) resource.ResourceAccess {
	return &apiResourceAccess{client: client}
}

func (a *apiResourceAccess) GetResource(ctx context.Context, query *resource.GetQuery) (*resource.Resource, error) {
	request := &resource_client.GetResourceRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetResource(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiResourceAccess) BatchGetResources(ctx context.Context, refs []*resource.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &resource_client.BatchGetResourcesRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetResources(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[resource.Name]*resource.Resource, len(refs))
	for _, resolvedRes := range resp.GetResources() {
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

func (a *apiResourceAccess) QueryResources(ctx context.Context, query *resource.ListQuery) (*resource.QueryResultSnapshot, error) {
	request := &resource_client.ListResourcesRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListResources(ctx, request)
	if err != nil {
		return nil, err
	}
	return &resource.QueryResultSnapshot{
		Resources:      resp.Resources,
		NextPageCursor: resp.NextPageToken,
		PrevPageCursor: resp.PrevPageToken,
	}, nil
}

func (a *apiResourceAccess) WatchResource(ctx context.Context, query *resource.GetQuery, observerCb func(*resource.ResourceChange) error) error {
	request := &resource_client.WatchResourceRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchResource(ctx, request)
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

func (a *apiResourceAccess) WatchResources(ctx context.Context, query *resource.WatchQuery, observerCb func(*resource.QueryResultChange) error) error {
	request := &resource_client.WatchResourcesRequest{
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
	changesStream, initErr := a.client.WatchResources(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &resource.QueryResultChange{
			Changes:      respChange.ResourceChanges,
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

func (a *apiResourceAccess) SaveResource(ctx context.Context, res *resource.Resource, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetResource(ctx, &resource.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &resource_client.UpdateResourceRequest{
			Resource: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*resource.Resource_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &resource_client.UpdateResourceRequest_CAS{
				ConditionalState: conditionalState.(*resource.Resource),
				FieldMask:        mask.(*resource.Resource_FieldMask),
			}
		}
		_, err := a.client.UpdateResource(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &resource_client.CreateResourceRequest{
			Resource: res,
		}
		_, err := a.client.CreateResource(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiResourceAccess) DeleteResource(ctx context.Context, ref *resource.Reference, opts ...gotenresource.DeleteOption) error {
	request := &resource_client.DeleteResourceRequest{
		Name: ref,
	}
	_, err := a.client.DeleteResource(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(resource.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return resource.AsAnyCastAccess(NewApiResourceAccess(resource_client.NewResourceServiceClient(cc)))
	})
}
