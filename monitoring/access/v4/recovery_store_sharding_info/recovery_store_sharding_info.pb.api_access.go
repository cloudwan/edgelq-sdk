// Code generated by protoc-gen-goten-access
// Resource: RecoveryStoreShardingInfo
// DO NOT EDIT!!!

package recovery_store_sharding_info_access

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

	recovery_store_sharding_info_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v4/recovery_store_sharding_info"
	recovery_store_sharding_info "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/recovery_store_sharding_info"
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

type apiRecoveryStoreShardingInfoAccess struct {
	client recovery_store_sharding_info_client.RecoveryStoreShardingInfoServiceClient
}

func NewApiRecoveryStoreShardingInfoAccess(client recovery_store_sharding_info_client.RecoveryStoreShardingInfoServiceClient) recovery_store_sharding_info.RecoveryStoreShardingInfoAccess {
	return &apiRecoveryStoreShardingInfoAccess{client: client}
}

func (a *apiRecoveryStoreShardingInfoAccess) GetRecoveryStoreShardingInfo(ctx context.Context, query *recovery_store_sharding_info.GetQuery) (*recovery_store_sharding_info.RecoveryStoreShardingInfo, error) {
	if !query.Reference.IsFullyQualified() {
		return nil, status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &recovery_store_sharding_info_client.GetRecoveryStoreShardingInfoRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetRecoveryStoreShardingInfo(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiRecoveryStoreShardingInfoAccess) BatchGetRecoveryStoreShardingInfos(ctx context.Context, refs []*recovery_store_sharding_info.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	asNames := make([]*recovery_store_sharding_info.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &recovery_store_sharding_info_client.BatchGetRecoveryStoreShardingInfosRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(recovery_store_sharding_info.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*recovery_store_sharding_info.RecoveryStoreShardingInfo_FieldMask)
	}
	resp, err := a.client.BatchGetRecoveryStoreShardingInfos(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[recovery_store_sharding_info.Name]*recovery_store_sharding_info.RecoveryStoreShardingInfo, len(refs))
	for _, resolvedRes := range resp.GetRecoveryStoreShardingInfos() {
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

func (a *apiRecoveryStoreShardingInfoAccess) QueryRecoveryStoreShardingInfos(ctx context.Context, query *recovery_store_sharding_info.ListQuery) (*recovery_store_sharding_info.QueryResultSnapshot, error) {
	request := &recovery_store_sharding_info_client.ListRecoveryStoreShardingInfosRequest{
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
	resp, err := a.client.ListRecoveryStoreShardingInfos(ctx, request)
	if err != nil {
		return nil, err
	}
	return &recovery_store_sharding_info.QueryResultSnapshot{
		RecoveryStoreShardingInfos: resp.RecoveryStoreShardingInfos,
		NextPageCursor:             resp.NextPageToken,
		PrevPageCursor:             resp.PrevPageToken,
		TotalResultsCount:          resp.TotalResultsCount,
		CurrentOffset:              resp.CurrentOffset,
	}, nil
}

func (a *apiRecoveryStoreShardingInfoAccess) WatchRecoveryStoreShardingInfo(ctx context.Context, query *recovery_store_sharding_info.GetQuery, observerCb func(*recovery_store_sharding_info.RecoveryStoreShardingInfoChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &recovery_store_sharding_info_client.WatchRecoveryStoreShardingInfoRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchRecoveryStoreShardingInfo(ctx, request)
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

func (a *apiRecoveryStoreShardingInfoAccess) WatchRecoveryStoreShardingInfos(ctx context.Context, query *recovery_store_sharding_info.WatchQuery, observerCb func(*recovery_store_sharding_info.QueryResultChange) error) error {
	request := &recovery_store_sharding_info_client.WatchRecoveryStoreShardingInfosRequest{
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
	changesStream, initErr := a.client.WatchRecoveryStoreShardingInfos(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &recovery_store_sharding_info.QueryResultChange{
			Changes:      respChange.RecoveryStoreShardingInfoChanges,
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

func (a *apiRecoveryStoreShardingInfoAccess) SaveRecoveryStoreShardingInfo(ctx context.Context, res *recovery_store_sharding_info.RecoveryStoreShardingInfo, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetRecoveryStoreShardingInfo(ctx, &recovery_store_sharding_info.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}
	var resp *recovery_store_sharding_info.RecoveryStoreShardingInfo
	var err error
	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &recovery_store_sharding_info_client.UpdateRecoveryStoreShardingInfoRequest{
			RecoveryStoreShardingInfo: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*recovery_store_sharding_info.RecoveryStoreShardingInfo_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &recovery_store_sharding_info_client.UpdateRecoveryStoreShardingInfoRequest_CAS{
				ConditionalState: conditionalState.(*recovery_store_sharding_info.RecoveryStoreShardingInfo),
				FieldMask:        mask.(*recovery_store_sharding_info.RecoveryStoreShardingInfo_FieldMask),
			}
		}
		resp, err = a.client.UpdateRecoveryStoreShardingInfo(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &recovery_store_sharding_info_client.CreateRecoveryStoreShardingInfoRequest{
			RecoveryStoreShardingInfo: res,
		}
		resp, err = a.client.CreateRecoveryStoreShardingInfo(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiRecoveryStoreShardingInfoAccess) DeleteRecoveryStoreShardingInfo(ctx context.Context, ref *recovery_store_sharding_info.Reference, opts ...gotenresource.DeleteOption) error {
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &recovery_store_sharding_info_client.DeleteRecoveryStoreShardingInfoRequest{
		Name: &ref.Name,
	}
	_, err := a.client.DeleteRecoveryStoreShardingInfo(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *recovery_store_sharding_info.Filter) (*recovery_store_sharding_info.Filter, *recovery_store_sharding_info.ParentName) {
	var withParentExtraction func(cnd recovery_store_sharding_info.FilterCondition) recovery_store_sharding_info.FilterCondition
	var resultParent *recovery_store_sharding_info.ParentName
	var resultFilter *recovery_store_sharding_info.Filter
	withParentExtraction = func(cnd recovery_store_sharding_info.FilterCondition) recovery_store_sharding_info.FilterCondition {
		switch tCnd := cnd.(type) {
		case *recovery_store_sharding_info.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]recovery_store_sharding_info.FilterCondition, 0)
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
				return recovery_store_sharding_info.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *recovery_store_sharding_info.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*recovery_store_sharding_info.Name)
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
		resultFilter = &recovery_store_sharding_info.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(recovery_store_sharding_info.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return recovery_store_sharding_info.AsAnyCastAccess(NewApiRecoveryStoreShardingInfoAccess(recovery_store_sharding_info_client.NewRecoveryStoreShardingInfoServiceClient(cc)))
	})
}
