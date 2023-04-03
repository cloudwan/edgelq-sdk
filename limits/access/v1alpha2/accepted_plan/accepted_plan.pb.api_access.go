// Code generated by protoc-gen-goten-access
// Resource: AcceptedPlan
// DO NOT EDIT!!!

package accepted_plan_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	accepted_plan_client "github.com/cloudwan/edgelq-sdk/limits/client/v1alpha2/accepted_plan"
	accepted_plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/accepted_plan"
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

type apiAcceptedPlanAccess struct {
	client accepted_plan_client.AcceptedPlanServiceClient
}

func NewApiAcceptedPlanAccess(client accepted_plan_client.AcceptedPlanServiceClient) accepted_plan.AcceptedPlanAccess {
	return &apiAcceptedPlanAccess{client: client}
}

func (a *apiAcceptedPlanAccess) GetAcceptedPlan(ctx context.Context, query *accepted_plan.GetQuery) (*accepted_plan.AcceptedPlan, error) {
	request := &accepted_plan_client.GetAcceptedPlanRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetAcceptedPlan(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiAcceptedPlanAccess) BatchGetAcceptedPlans(ctx context.Context, refs []*accepted_plan.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &accepted_plan_client.BatchGetAcceptedPlansRequest{
		Names: refs,
	}
	fieldMask := batchGetOpts.GetFieldMask(accepted_plan.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*accepted_plan.AcceptedPlan_FieldMask)
	}
	resp, err := a.client.BatchGetAcceptedPlans(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[accepted_plan.Name]*accepted_plan.AcceptedPlan, len(refs))
	for _, resolvedRes := range resp.GetAcceptedPlans() {
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

func (a *apiAcceptedPlanAccess) QueryAcceptedPlans(ctx context.Context, query *accepted_plan.ListQuery) (*accepted_plan.QueryResultSnapshot, error) {
	request := &accepted_plan_client.ListAcceptedPlansRequest{
		Filter:            query.Filter,
		FieldMask:         query.Mask,
		IncludePagingInfo: query.WithPagingInfo,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListAcceptedPlans(ctx, request)
	if err != nil {
		return nil, err
	}
	return &accepted_plan.QueryResultSnapshot{
		AcceptedPlans:     resp.AcceptedPlans,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiAcceptedPlanAccess) WatchAcceptedPlan(ctx context.Context, query *accepted_plan.GetQuery, observerCb func(*accepted_plan.AcceptedPlanChange) error) error {
	request := &accepted_plan_client.WatchAcceptedPlanRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchAcceptedPlan(ctx, request)
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

func (a *apiAcceptedPlanAccess) WatchAcceptedPlans(ctx context.Context, query *accepted_plan.WatchQuery, observerCb func(*accepted_plan.QueryResultChange) error) error {
	request := &accepted_plan_client.WatchAcceptedPlansRequest{
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
	changesStream, initErr := a.client.WatchAcceptedPlans(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &accepted_plan.QueryResultChange{
			Changes:      respChange.AcceptedPlanChanges,
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

func (a *apiAcceptedPlanAccess) SaveAcceptedPlan(ctx context.Context, res *accepted_plan.AcceptedPlan, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetAcceptedPlan(ctx, &accepted_plan.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &accepted_plan_client.UpdateAcceptedPlanRequest{
			AcceptedPlan: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*accepted_plan.AcceptedPlan_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &accepted_plan_client.UpdateAcceptedPlanRequest_CAS{
				ConditionalState: conditionalState.(*accepted_plan.AcceptedPlan),
				FieldMask:        mask.(*accepted_plan.AcceptedPlan_FieldMask),
			}
		}
		_, err := a.client.UpdateAcceptedPlan(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &accepted_plan_client.CreateAcceptedPlanRequest{
			AcceptedPlan: res,
		}
		_, err := a.client.CreateAcceptedPlan(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiAcceptedPlanAccess) DeleteAcceptedPlan(ctx context.Context, ref *accepted_plan.Reference, opts ...gotenresource.DeleteOption) error {
	request := &accepted_plan_client.DeleteAcceptedPlanRequest{
		Name: ref,
	}
	_, err := a.client.DeleteAcceptedPlan(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(accepted_plan.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return accepted_plan.AsAnyCastAccess(NewApiAcceptedPlanAccess(accepted_plan_client.NewAcceptedPlanServiceClient(cc)))
	})
}
