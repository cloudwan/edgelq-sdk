// Code generated by protoc-gen-goten-access
// Resource: LimitPool
// DO NOT EDIT!!!

package limit_pool_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
	gotenfilter "github.com/cloudwan/goten-sdk/runtime/resource/filter"
	"github.com/cloudwan/goten-sdk/types/watch_type"

	limit_pool_client "github.com/cloudwan/edgelq-sdk/limits/client/v1/limit_pool"
	limit_pool "github.com/cloudwan/edgelq-sdk/limits/resources/v1/limit_pool"
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
	_ = gotenfilter.Eq
)

type apiLimitPoolAccess struct {
	client limit_pool_client.LimitPoolServiceClient
}

func NewApiLimitPoolAccess(client limit_pool_client.LimitPoolServiceClient) limit_pool.LimitPoolAccess {
	return &apiLimitPoolAccess{client: client}
}

func (a *apiLimitPoolAccess) GetLimitPool(ctx context.Context, query *limit_pool.GetQuery) (*limit_pool.LimitPool, error) {
	if !query.Reference.IsFullyQualified() {
		return nil, status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &limit_pool_client.GetLimitPoolRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetLimitPool(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiLimitPoolAccess) BatchGetLimitPools(ctx context.Context, refs []*limit_pool.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	asNames := make([]*limit_pool.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &limit_pool_client.BatchGetLimitPoolsRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(limit_pool.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*limit_pool.LimitPool_FieldMask)
	}
	resp, err := a.client.BatchGetLimitPools(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[limit_pool.Name]*limit_pool.LimitPool, len(refs))
	for _, resolvedRes := range resp.GetLimitPools() {
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

func (a *apiLimitPoolAccess) QueryLimitPools(ctx context.Context, query *limit_pool.ListQuery) (*limit_pool.QueryResultSnapshot, error) {
	request := &limit_pool_client.ListLimitPoolsRequest{
		Filter:            query.Filter,
		FieldMask:         query.Mask,
		IncludePagingInfo: query.WithPagingInfo,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	if query.Filter != nil && query.Filter.GetCondition() != nil {
		request.Filter, request.Parent = getParentAndFilter(query.Filter)
	}
	resp, err := a.client.ListLimitPools(ctx, request)
	if err != nil {
		return nil, err
	}
	return &limit_pool.QueryResultSnapshot{
		LimitPools:        resp.LimitPools,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiLimitPoolAccess) WatchLimitPool(ctx context.Context, query *limit_pool.GetQuery, observerCb func(*limit_pool.LimitPoolChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &limit_pool_client.WatchLimitPoolRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchLimitPool(ctx, request)
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

func (a *apiLimitPoolAccess) WatchLimitPools(ctx context.Context, query *limit_pool.WatchQuery, observerCb func(*limit_pool.QueryResultChange) error) error {
	request := &limit_pool_client.WatchLimitPoolsRequest{
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
	if query.Filter != nil && query.Filter.GetCondition() != nil {
		request.Filter, request.Parent = getParentAndFilter(query.Filter)
	}
	changesStream, initErr := a.client.WatchLimitPools(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &limit_pool.QueryResultChange{
			Changes:      respChange.LimitPoolChanges,
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

func (a *apiLimitPoolAccess) SaveLimitPool(ctx context.Context, res *limit_pool.LimitPool, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetLimitPool(ctx, &limit_pool.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}
	var resp *limit_pool.LimitPool
	var err error
	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &limit_pool_client.UpdateLimitPoolRequest{
			LimitPool: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*limit_pool.LimitPool_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &limit_pool_client.UpdateLimitPoolRequest_CAS{
				ConditionalState: conditionalState.(*limit_pool.LimitPool),
				FieldMask:        mask.(*limit_pool.LimitPool_FieldMask),
			}
		}
		resp, err = a.client.UpdateLimitPool(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("create operation on %s is prohibited", res.Name.AsReference().String())
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiLimitPoolAccess) DeleteLimitPool(ctx context.Context, ref *limit_pool.Reference, opts ...gotenresource.DeleteOption) error {
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &limit_pool_client.DeleteLimitPoolRequest{
		Name: &ref.Name,
	}
	_, err := a.client.DeleteLimitPool(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *limit_pool.Filter) (*limit_pool.Filter, *limit_pool.ParentName) {
	var withParentExtraction func(cnd limit_pool.FilterCondition) limit_pool.FilterCondition
	var resultParent *limit_pool.ParentName
	var resultFilter *limit_pool.Filter
	withParentExtraction = func(cnd limit_pool.FilterCondition) limit_pool.FilterCondition {
		switch tCnd := cnd.(type) {
		case *limit_pool.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]limit_pool.FilterCondition, 0)
				for _, subCnd := range tCnd.Conditions {
					if subCndNoParent := withParentExtraction(subCnd); subCndNoParent != nil {
						withoutParentCnds = append(withoutParentCnds, subCndNoParent)
					}
				}
				if len(withoutParentCnds) == 0 {
					return nil
				}
				if len(withoutParentCnds) == 1 {
					return withoutParentCnds[0]
				}
				return limit_pool.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *limit_pool.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*limit_pool.Name)
				if nameValue != nil && nameValue.ParentName.IsSpecified() {
					resultParent = &nameValue.ParentName
					if nameValue.IsFullyQualified() {
						return tCnd
					}
					return nil
				}
			}
			return tCnd
		default:
			return tCnd
		}
	}
	cndWithoutParent := withParentExtraction(fullFilter.GetCondition())
	if cndWithoutParent != nil {
		resultFilter = &limit_pool.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(limit_pool.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return limit_pool.AsAnyCastAccess(NewApiLimitPoolAccess(limit_pool_client.NewLimitPoolServiceClient(cc)))
	})
}
