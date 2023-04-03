// Code generated by protoc-gen-goten-access
// Resource: PlanAssignment
// DO NOT EDIT!!!

package plan_assignment_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	plan_assignment_client "github.com/cloudwan/edgelq-sdk/limits/client/v1alpha2/plan_assignment"
	plan_assignment "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/plan_assignment"
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

type apiPlanAssignmentAccess struct {
	client plan_assignment_client.PlanAssignmentServiceClient
}

func NewApiPlanAssignmentAccess(client plan_assignment_client.PlanAssignmentServiceClient) plan_assignment.PlanAssignmentAccess {
	return &apiPlanAssignmentAccess{client: client}
}

func (a *apiPlanAssignmentAccess) GetPlanAssignment(ctx context.Context, query *plan_assignment.GetQuery) (*plan_assignment.PlanAssignment, error) {
	request := &plan_assignment_client.GetPlanAssignmentRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetPlanAssignment(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiPlanAssignmentAccess) BatchGetPlanAssignments(ctx context.Context, refs []*plan_assignment.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &plan_assignment_client.BatchGetPlanAssignmentsRequest{
		Names: refs,
	}
	fieldMask := batchGetOpts.GetFieldMask(plan_assignment.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*plan_assignment.PlanAssignment_FieldMask)
	}
	resp, err := a.client.BatchGetPlanAssignments(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[plan_assignment.Name]*plan_assignment.PlanAssignment, len(refs))
	for _, resolvedRes := range resp.GetPlanAssignments() {
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

func (a *apiPlanAssignmentAccess) QueryPlanAssignments(ctx context.Context, query *plan_assignment.ListQuery) (*plan_assignment.QueryResultSnapshot, error) {
	request := &plan_assignment_client.ListPlanAssignmentsRequest{
		Filter:            query.Filter,
		FieldMask:         query.Mask,
		IncludePagingInfo: query.WithPagingInfo,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListPlanAssignments(ctx, request)
	if err != nil {
		return nil, err
	}
	return &plan_assignment.QueryResultSnapshot{
		PlanAssignments:   resp.PlanAssignments,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiPlanAssignmentAccess) WatchPlanAssignment(ctx context.Context, query *plan_assignment.GetQuery, observerCb func(*plan_assignment.PlanAssignmentChange) error) error {
	request := &plan_assignment_client.WatchPlanAssignmentRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchPlanAssignment(ctx, request)
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

func (a *apiPlanAssignmentAccess) WatchPlanAssignments(ctx context.Context, query *plan_assignment.WatchQuery, observerCb func(*plan_assignment.QueryResultChange) error) error {
	request := &plan_assignment_client.WatchPlanAssignmentsRequest{
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
	changesStream, initErr := a.client.WatchPlanAssignments(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &plan_assignment.QueryResultChange{
			Changes:      respChange.PlanAssignmentChanges,
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

func (a *apiPlanAssignmentAccess) SavePlanAssignment(ctx context.Context, res *plan_assignment.PlanAssignment, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetPlanAssignment(ctx, &plan_assignment.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &plan_assignment_client.UpdatePlanAssignmentRequest{
			PlanAssignment: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*plan_assignment.PlanAssignment_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &plan_assignment_client.UpdatePlanAssignmentRequest_CAS{
				ConditionalState: conditionalState.(*plan_assignment.PlanAssignment),
				FieldMask:        mask.(*plan_assignment.PlanAssignment_FieldMask),
			}
		}
		_, err := a.client.UpdatePlanAssignment(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &plan_assignment_client.CreatePlanAssignmentRequest{
			PlanAssignment: res,
		}
		_, err := a.client.CreatePlanAssignment(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiPlanAssignmentAccess) DeletePlanAssignment(ctx context.Context, ref *plan_assignment.Reference, opts ...gotenresource.DeleteOption) error {
	request := &plan_assignment_client.DeletePlanAssignmentRequest{
		Name: ref,
	}
	_, err := a.client.DeletePlanAssignment(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(plan_assignment.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return plan_assignment.AsAnyCastAccess(NewApiPlanAssignmentAccess(plan_assignment_client.NewPlanAssignmentServiceClient(cc)))
	})
}
