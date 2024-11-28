// Code generated by protoc-gen-goten-access
// Resource: AlertingCondition
// DO NOT EDIT!!!

package alerting_condition_access

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

	alerting_condition_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v4/alerting_condition"
	alerting_condition "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/alerting_condition"
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

type apiAlertingConditionAccess struct {
	client alerting_condition_client.AlertingConditionServiceClient
}

func NewApiAlertingConditionAccess(client alerting_condition_client.AlertingConditionServiceClient) alerting_condition.AlertingConditionAccess {
	return &apiAlertingConditionAccess{client: client}
}

func (a *apiAlertingConditionAccess) GetAlertingCondition(ctx context.Context, query *alerting_condition.GetQuery, opts ...gotenresource.GetOption) (*alerting_condition.AlertingCondition, error) {
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
	request := &alerting_condition_client.GetAlertingConditionRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetAlertingCondition(ctx, request, callOpts...)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiAlertingConditionAccess) BatchGetAlertingConditions(ctx context.Context, refs []*alerting_condition.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	callHeaders := metadata.MD{}
	if batchGetOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	asNames := make([]*alerting_condition.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &alerting_condition_client.BatchGetAlertingConditionsRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(alerting_condition.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*alerting_condition.AlertingCondition_FieldMask)
	}
	resp, err := a.client.BatchGetAlertingConditions(ctx, request, callOpts...)
	if err != nil {
		return err
	}
	resultMap := make(map[alerting_condition.Name]*alerting_condition.AlertingCondition, len(refs))
	for _, resolvedRes := range resp.GetAlertingConditions() {
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

func (a *apiAlertingConditionAccess) QueryAlertingConditions(ctx context.Context, query *alerting_condition.ListQuery, opts ...gotenresource.QueryOption) (*alerting_condition.QueryResultSnapshot, error) {
	qOpts := gotenresource.MakeQueryOptions(opts)
	callHeaders := metadata.MD{}
	if qOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	request := &alerting_condition_client.ListAlertingConditionsRequest{
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
	resp, err := a.client.ListAlertingConditions(ctx, request)
	if err != nil {
		return nil, err
	}
	return &alerting_condition.QueryResultSnapshot{
		AlertingConditions: resp.AlertingConditions,
		NextPageCursor:     resp.NextPageToken,
		PrevPageCursor:     resp.PrevPageToken,
		TotalResultsCount:  resp.TotalResultsCount,
		CurrentOffset:      resp.CurrentOffset,
	}, nil
}

func (a *apiAlertingConditionAccess) SearchAlertingConditions(ctx context.Context, query *alerting_condition.SearchQuery, opts ...gotenresource.QueryOption) (*alerting_condition.QueryResultSnapshot, error) {
	qOpts := gotenresource.MakeQueryOptions(opts)
	callHeaders := metadata.MD{}
	if qOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	request := &alerting_condition_client.SearchAlertingConditionsRequest{
		Phrase:    query.Phrase,
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	if query.Filter != nil && query.Filter.GetCondition() != nil {
		request.Filter, request.Parent = getParentAndFilter(query.Filter)
	}
	resp, err := a.client.SearchAlertingConditions(ctx, request, callOpts...)
	if err != nil {
		return nil, err
	}
	return &alerting_condition.QueryResultSnapshot{
		AlertingConditions: resp.AlertingConditions,
		NextPageCursor:     resp.NextPageToken,
		PrevPageCursor:     resp.PrevPageToken,
		CurrentOffset:      resp.CurrentOffset,
		TotalResultsCount:  resp.TotalResultsCount,
	}, nil
}

func (a *apiAlertingConditionAccess) WatchAlertingCondition(ctx context.Context, query *alerting_condition.GetQuery, observerCb func(*alerting_condition.AlertingConditionChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &alerting_condition_client.WatchAlertingConditionRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	changesStream, initErr := a.client.WatchAlertingCondition(ctx, request)
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

func (a *apiAlertingConditionAccess) WatchAlertingConditions(ctx context.Context, query *alerting_condition.WatchQuery, observerCb func(*alerting_condition.QueryResultChange) error) error {
	request := &alerting_condition_client.WatchAlertingConditionsRequest{
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

	changesStream, initErr := a.client.WatchAlertingConditions(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &alerting_condition.QueryResultChange{
			Changes:      respChange.AlertingConditionChanges,
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

func (a *apiAlertingConditionAccess) SaveAlertingCondition(ctx context.Context, res *alerting_condition.AlertingCondition, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetAlertingCondition(ctx, &alerting_condition.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}
	var resp *alerting_condition.AlertingCondition
	var err error
	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &alerting_condition_client.UpdateAlertingConditionRequest{
			AlertingCondition: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*alerting_condition.AlertingCondition_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &alerting_condition_client.UpdateAlertingConditionRequest_CAS{
				ConditionalState: conditionalState.(*alerting_condition.AlertingCondition),
				FieldMask:        mask.(*alerting_condition.AlertingCondition_FieldMask),
			}
		}
		resp, err = a.client.UpdateAlertingCondition(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &alerting_condition_client.CreateAlertingConditionRequest{
			AlertingCondition: res,
		}
		resp, err = a.client.CreateAlertingCondition(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiAlertingConditionAccess) DeleteAlertingCondition(ctx context.Context, ref *alerting_condition.Reference, opts ...gotenresource.DeleteOption) error {
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &alerting_condition_client.DeleteAlertingConditionRequest{
		Name: &ref.Name,
	}
	_, err := a.client.DeleteAlertingCondition(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *alerting_condition.Filter) (*alerting_condition.Filter, *alerting_condition.ParentName) {
	var withParentExtraction func(cnd alerting_condition.FilterCondition) alerting_condition.FilterCondition
	var resultParent *alerting_condition.ParentName
	var resultFilter *alerting_condition.Filter
	withParentExtraction = func(cnd alerting_condition.FilterCondition) alerting_condition.FilterCondition {
		switch tCnd := cnd.(type) {
		case *alerting_condition.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]alerting_condition.FilterCondition, 0)
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
				return alerting_condition.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *alerting_condition.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*alerting_condition.Name)
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
		resultFilter = &alerting_condition.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(alerting_condition.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return alerting_condition.AsAnyCastAccess(NewApiAlertingConditionAccess(alerting_condition_client.NewAlertingConditionServiceClient(cc)))
	})
}
