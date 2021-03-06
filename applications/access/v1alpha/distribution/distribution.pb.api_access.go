// Code generated by protoc-gen-goten-access
// Resource: Distribution
// DO NOT EDIT!!!

package distribution_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	distribution_client "github.com/cloudwan/edgelq-sdk/applications/client/v1alpha/distribution"
	distribution "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha/distribution"
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

type apiDistributionAccess struct {
	client distribution_client.DistributionServiceClient
}

func NewApiDistributionAccess(client distribution_client.DistributionServiceClient) distribution.DistributionAccess {
	return &apiDistributionAccess{client: client}
}

func (a *apiDistributionAccess) GetDistribution(ctx context.Context, query *distribution.GetQuery) (*distribution.Distribution, error) {
	request := &distribution_client.GetDistributionRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetDistribution(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiDistributionAccess) BatchGetDistributions(ctx context.Context, refs []*distribution.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &distribution_client.BatchGetDistributionsRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetDistributions(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[distribution.Name]*distribution.Distribution, len(refs))
	for _, resolvedRes := range resp.GetDistributions() {
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

func (a *apiDistributionAccess) QueryDistributions(ctx context.Context, query *distribution.ListQuery) (*distribution.QueryResultSnapshot, error) {
	request := &distribution_client.ListDistributionsRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListDistributions(ctx, request)
	if err != nil {
		return nil, err
	}
	return &distribution.QueryResultSnapshot{
		Distributions:  resp.Distributions,
		NextPageCursor: resp.NextPageToken,
		PrevPageCursor: resp.PrevPageToken,
	}, nil
}

func (a *apiDistributionAccess) WatchDistribution(ctx context.Context, query *distribution.GetQuery, observerCb func(*distribution.DistributionChange) error) error {
	request := &distribution_client.WatchDistributionRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchDistribution(ctx, request)
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

func (a *apiDistributionAccess) WatchDistributions(ctx context.Context, query *distribution.WatchQuery, observerCb func(*distribution.QueryResultChange) error) error {
	request := &distribution_client.WatchDistributionsRequest{
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
	changesStream, initErr := a.client.WatchDistributions(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &distribution.QueryResultChange{
			Changes:      respChange.DistributionChanges,
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

func (a *apiDistributionAccess) SaveDistribution(ctx context.Context, res *distribution.Distribution, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetDistribution(ctx, &distribution.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &distribution_client.UpdateDistributionRequest{
			Distribution: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*distribution.Distribution_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &distribution_client.UpdateDistributionRequest_CAS{
				ConditionalState: conditionalState.(*distribution.Distribution),
				FieldMask:        mask.(*distribution.Distribution_FieldMask),
			}
		}
		_, err := a.client.UpdateDistribution(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &distribution_client.CreateDistributionRequest{
			Distribution: res,
		}
		_, err := a.client.CreateDistribution(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiDistributionAccess) DeleteDistribution(ctx context.Context, ref *distribution.Reference, opts ...gotenresource.DeleteOption) error {
	request := &distribution_client.DeleteDistributionRequest{
		Name: ref,
	}
	_, err := a.client.DeleteDistribution(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(distribution.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return distribution.AsAnyCastAccess(NewApiDistributionAccess(distribution_client.NewDistributionServiceClient(cc)))
	})
}
