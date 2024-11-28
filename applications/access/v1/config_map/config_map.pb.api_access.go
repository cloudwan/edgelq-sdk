// Code generated by protoc-gen-goten-access
// Resource: ConfigMap
// DO NOT EDIT!!!

package config_map_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
	gotenfilter "github.com/cloudwan/goten-sdk/runtime/resource/filter"
	"github.com/cloudwan/goten-sdk/types/watch_type"

	config_map_client "github.com/cloudwan/edgelq-sdk/applications/client/v1/config_map"
	config_map "github.com/cloudwan/edgelq-sdk/applications/resources/v1/config_map"
)

var (
	_ = new(context.Context)
	_ = new(fmt.GoStringer)

	_ = metadata.MD{}
	_ = new(grpc.ClientConnInterface)
	_ = codes.NotFound
	_ = status.Status{}

	_ = new(gotenaccess.Watcher)
	_ = watch_type.WatchType_STATEFUL
	_ = new(gotenresource.ListQuery)
	_ = gotenfilter.Eq
)

type apiConfigMapAccess struct {
	client config_map_client.ConfigMapServiceClient
}

func NewApiConfigMapAccess(client config_map_client.ConfigMapServiceClient) config_map.ConfigMapAccess {
	return &apiConfigMapAccess{client: client}
}

func (a *apiConfigMapAccess) GetConfigMap(ctx context.Context, query *config_map.GetQuery, opts ...gotenresource.GetOption) (*config_map.ConfigMap, error) {
	getOpts := gotenresource.MakeGetOptions(opts)
	callHeaders := metadata.MD{}
	if getOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	if !query.Reference.IsFullyQualified() {
		return nil, status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &config_map_client.GetConfigMapRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetConfigMap(ctx, request, callOpts...)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiConfigMapAccess) BatchGetConfigMaps(ctx context.Context, refs []*config_map.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	callHeaders := metadata.MD{}
	if batchGetOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	asNames := make([]*config_map.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &config_map_client.BatchGetConfigMapsRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(config_map.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*config_map.ConfigMap_FieldMask)
	}
	resp, err := a.client.BatchGetConfigMaps(ctx, request, callOpts...)
	if err != nil {
		return err
	}
	resultMap := make(map[config_map.Name]*config_map.ConfigMap, len(refs))
	for _, resolvedRes := range resp.GetConfigMaps() {
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

func (a *apiConfigMapAccess) QueryConfigMaps(ctx context.Context, query *config_map.ListQuery, opts ...gotenresource.QueryOption) (*config_map.QueryResultSnapshot, error) {
	qOpts := gotenresource.MakeQueryOptions(opts)
	callHeaders := metadata.MD{}
	if qOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	request := &config_map_client.ListConfigMapsRequest{
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
	resp, err := a.client.ListConfigMaps(ctx, request)
	if err != nil {
		return nil, err
	}
	return &config_map.QueryResultSnapshot{
		ConfigMaps:        resp.ConfigMaps,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiConfigMapAccess) WatchConfigMap(ctx context.Context, query *config_map.GetQuery, observerCb func(*config_map.ConfigMapChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &config_map_client.WatchConfigMapRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	changesStream, initErr := a.client.WatchConfigMap(ctx, request)
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

func (a *apiConfigMapAccess) WatchConfigMaps(ctx context.Context, query *config_map.WatchQuery, observerCb func(*config_map.QueryResultChange) error) error {
	request := &config_map_client.WatchConfigMapsRequest{
		Filter:       query.Filter,
		FieldMask:    query.Mask,
		MaxChunkSize: int32(query.ChunkSize),
		Type:         query.WatchType,
		ResumeToken:  query.ResumeToken,
		StartingTime: query.StartingTime,
	}
	if query.Pager != nil {
		request.OrderBy = query.Pager.OrderBy
		request.PageSize = int32(query.Pager.Limit)
		request.PageToken = query.Pager.Cursor
	}
	if query.Filter != nil && query.Filter.GetCondition() != nil {
		request.Filter, request.Parent = getParentAndFilter(query.Filter)
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	changesStream, initErr := a.client.WatchConfigMaps(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &config_map.QueryResultChange{
			Changes:      respChange.ConfigMapChanges,
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

func (a *apiConfigMapAccess) SaveConfigMap(ctx context.Context, res *config_map.ConfigMap, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetConfigMap(ctx, &config_map.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}
	var resp *config_map.ConfigMap
	var err error
	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &config_map_client.UpdateConfigMapRequest{
			ConfigMap: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*config_map.ConfigMap_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &config_map_client.UpdateConfigMapRequest_CAS{
				ConditionalState: conditionalState.(*config_map.ConfigMap),
				FieldMask:        mask.(*config_map.ConfigMap_FieldMask),
			}
		}
		resp, err = a.client.UpdateConfigMap(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &config_map_client.CreateConfigMapRequest{
			ConfigMap: res,
		}
		resp, err = a.client.CreateConfigMap(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiConfigMapAccess) DeleteConfigMap(ctx context.Context, ref *config_map.Reference, opts ...gotenresource.DeleteOption) error {
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &config_map_client.DeleteConfigMapRequest{
		Name: &ref.Name,
	}
	_, err := a.client.DeleteConfigMap(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *config_map.Filter) (*config_map.Filter, *config_map.ParentName) {
	var withParentExtraction func(cnd config_map.FilterCondition) config_map.FilterCondition
	var resultParent *config_map.ParentName
	var resultFilter *config_map.Filter
	withParentExtraction = func(cnd config_map.FilterCondition) config_map.FilterCondition {
		switch tCnd := cnd.(type) {
		case *config_map.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]config_map.FilterCondition, 0)
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
				return config_map.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *config_map.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*config_map.Name)
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
		resultFilter = &config_map.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(config_map.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return config_map.AsAnyCastAccess(NewApiConfigMapAccess(config_map_client.NewConfigMapServiceClient(cc)))
	})
}
