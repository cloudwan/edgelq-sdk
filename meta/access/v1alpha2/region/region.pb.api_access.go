// Code generated by protoc-gen-goten-access
// Resource: Region
// DO NOT EDIT!!!

package region_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	region_client "github.com/cloudwan/edgelq-sdk/meta/client/v1alpha2/region"
	region "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/region"
)

var (
	_ = new(context.Context)
	_ = new(fmt.GoStringer)

	_ = new(grpc.ClientConnInterface)
	_ = codes.NotFound
	_ = status.Status{}

	_ = new(gotenaccess.Watcher)
	_ = watch_type.WatchType_STATEFUL
	_ = new(gotenresource.ListQuery)
)

type apiRegionAccess struct {
	client region_client.RegionServiceClient
}

func NewApiRegionAccess(client region_client.RegionServiceClient) region.RegionAccess {
	return &apiRegionAccess{client: client}
}

func (a *apiRegionAccess) GetRegion(ctx context.Context, query *region.GetQuery) (*region.Region, error) {
	request := &region_client.GetRegionRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetRegion(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiRegionAccess) BatchGetRegions(ctx context.Context, refs []*region.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &region_client.BatchGetRegionsRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetRegions(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[region.Name]*region.Region, len(refs))
	for _, resolvedRes := range resp.GetRegions() {
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

func (a *apiRegionAccess) QueryRegions(ctx context.Context, query *region.ListQuery) (*region.QueryResultSnapshot, error) {
	request := &region_client.ListRegionsRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListRegions(ctx, request)
	if err != nil {
		return nil, err
	}
	return &region.QueryResultSnapshot{
		Regions:        resp.Regions,
		NextPageCursor: resp.NextPageToken,
		PrevPageCursor: resp.PrevPageToken,
	}, nil
}

func (a *apiRegionAccess) WatchRegion(ctx context.Context, query *region.GetQuery, observerCb func(*region.RegionChange) error) error {
	request := &region_client.WatchRegionRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchRegion(ctx, request)
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

func (a *apiRegionAccess) WatchRegions(ctx context.Context, query *region.WatchQuery, observerCb func(*region.QueryResultChange) error) error {
	request := &region_client.WatchRegionsRequest{
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
	changesStream, initErr := a.client.WatchRegions(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &region.QueryResultChange{
			Changes:      respChange.RegionChanges,
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

func (a *apiRegionAccess) SaveRegion(ctx context.Context, res *region.Region, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetRegion(ctx, &region.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &region_client.UpdateRegionRequest{
			Region: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*region.Region_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &region_client.UpdateRegionRequest_CAS{
				ConditionalState: conditionalState.(*region.Region),
				FieldMask:        mask.(*region.Region_FieldMask),
			}
		}
		_, err := a.client.UpdateRegion(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &region_client.CreateRegionRequest{
			Region: res,
		}
		_, err := a.client.CreateRegion(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiRegionAccess) DeleteRegion(ctx context.Context, ref *region.Reference, opts ...gotenresource.DeleteOption) error {
	request := &region_client.DeleteRegionRequest{
		Name: ref,
	}
	_, err := a.client.DeleteRegion(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(region.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return region.AsAnyCastAccess(NewApiRegionAccess(region_client.NewRegionServiceClient(cc)))
	})
}
