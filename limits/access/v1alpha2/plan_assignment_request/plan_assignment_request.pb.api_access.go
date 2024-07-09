// Code generated by protoc-gen-goten-access
// Resource: PlanAssignmentRequest
// DO NOT EDIT!!!

package plan_assignment_request_access

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

	plan_assignment_request_client "github.com/cloudwan/edgelq-sdk/limits/client/v1alpha2/plan_assignment_request"
	plan_assignment_request "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/plan_assignment_request"
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

type apiPlanAssignmentRequestAccess struct {
	client plan_assignment_request_client.PlanAssignmentRequestServiceClient
}

func NewApiPlanAssignmentRequestAccess(client plan_assignment_request_client.PlanAssignmentRequestServiceClient) plan_assignment_request.PlanAssignmentRequestAccess {
	return &apiPlanAssignmentRequestAccess{client: client}
}

func (a *apiPlanAssignmentRequestAccess) GetPlanAssignmentRequest(ctx context.Context, query *plan_assignment_request.GetQuery) (*plan_assignment_request.PlanAssignmentRequest, error) {
	if !query.Reference.IsFullyQualified() {
		return nil, status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &plan_assignment_request_client.GetPlanAssignmentRequestRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetPlanAssignmentRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiPlanAssignmentRequestAccess) BatchGetPlanAssignmentRequests(ctx context.Context, refs []*plan_assignment_request.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	asNames := make([]*plan_assignment_request.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &plan_assignment_request_client.BatchGetPlanAssignmentRequestsRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(plan_assignment_request.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*plan_assignment_request.PlanAssignmentRequest_FieldMask)
	}
	resp, err := a.client.BatchGetPlanAssignmentRequests(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[plan_assignment_request.Name]*plan_assignment_request.PlanAssignmentRequest, len(refs))
	for _, resolvedRes := range resp.GetPlanAssignmentRequests() {
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

func (a *apiPlanAssignmentRequestAccess) QueryPlanAssignmentRequests(ctx context.Context, query *plan_assignment_request.ListQuery) (*plan_assignment_request.QueryResultSnapshot, error) {
	request := &plan_assignment_request_client.ListPlanAssignmentRequestsRequest{
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
	resp, err := a.client.ListPlanAssignmentRequests(ctx, request)
	if err != nil {
		return nil, err
	}
	return &plan_assignment_request.QueryResultSnapshot{
		PlanAssignmentRequests: resp.PlanAssignmentRequests,
		NextPageCursor:         resp.NextPageToken,
		PrevPageCursor:         resp.PrevPageToken,
		TotalResultsCount:      resp.TotalResultsCount,
		CurrentOffset:          resp.CurrentOffset,
	}, nil
}

func (a *apiPlanAssignmentRequestAccess) WatchPlanAssignmentRequest(ctx context.Context, query *plan_assignment_request.GetQuery, observerCb func(*plan_assignment_request.PlanAssignmentRequestChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &plan_assignment_request_client.WatchPlanAssignmentRequestRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchPlanAssignmentRequest(ctx, request)
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

func (a *apiPlanAssignmentRequestAccess) WatchPlanAssignmentRequests(ctx context.Context, query *plan_assignment_request.WatchQuery, observerCb func(*plan_assignment_request.QueryResultChange) error) error {
	request := &plan_assignment_request_client.WatchPlanAssignmentRequestsRequest{
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
	changesStream, initErr := a.client.WatchPlanAssignmentRequests(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &plan_assignment_request.QueryResultChange{
			Changes:      respChange.PlanAssignmentRequestChanges,
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

func (a *apiPlanAssignmentRequestAccess) SavePlanAssignmentRequest(ctx context.Context, res *plan_assignment_request.PlanAssignmentRequest, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetPlanAssignmentRequest(ctx, &plan_assignment_request.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}
	var resp *plan_assignment_request.PlanAssignmentRequest
	var err error
	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &plan_assignment_request_client.UpdatePlanAssignmentRequestRequest{
			PlanAssignmentRequest: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*plan_assignment_request.PlanAssignmentRequest_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &plan_assignment_request_client.UpdatePlanAssignmentRequestRequest_CAS{
				ConditionalState: conditionalState.(*plan_assignment_request.PlanAssignmentRequest),
				FieldMask:        mask.(*plan_assignment_request.PlanAssignmentRequest_FieldMask),
			}
		}
		resp, err = a.client.UpdatePlanAssignmentRequest(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &plan_assignment_request_client.CreatePlanAssignmentRequestRequest{
			PlanAssignmentRequest: res,
		}
		resp, err = a.client.CreatePlanAssignmentRequest(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiPlanAssignmentRequestAccess) DeletePlanAssignmentRequest(ctx context.Context, ref *plan_assignment_request.Reference, opts ...gotenresource.DeleteOption) error {
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &plan_assignment_request_client.DeletePlanAssignmentRequestRequest{
		Name: &ref.Name,
	}
	_, err := a.client.DeletePlanAssignmentRequest(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *plan_assignment_request.Filter) (*plan_assignment_request.Filter, *plan_assignment_request.ParentName) {
	var withParentExtraction func(cnd plan_assignment_request.FilterCondition) plan_assignment_request.FilterCondition
	var resultParent *plan_assignment_request.ParentName
	var resultFilter *plan_assignment_request.Filter
	withParentExtraction = func(cnd plan_assignment_request.FilterCondition) plan_assignment_request.FilterCondition {
		switch tCnd := cnd.(type) {
		case *plan_assignment_request.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]plan_assignment_request.FilterCondition, 0)
				for _, subCnd := range tCnd.Conditions {
					if subCndNoParent := withParentExtraction(subCnd); subCndNoParent != nil {
						withoutParentCnds = append(withoutParentCnds, subCndNoParent)
					}
				}
				if len(withoutParentCnds) == 0 {
					return nil
				}
				return plan_assignment_request.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *plan_assignment_request.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*plan_assignment_request.Name)
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
		resultFilter = &plan_assignment_request.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(plan_assignment_request.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return plan_assignment_request.AsAnyCastAccess(NewApiPlanAssignmentRequestAccess(plan_assignment_request_client.NewPlanAssignmentRequestServiceClient(cc)))
	})
}
