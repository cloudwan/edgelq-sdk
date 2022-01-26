// Code generated by protoc-gen-goten-access
// Resource: AlertingCondition
// DO NOT EDIT!!!

package alerting_condition_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	alerting_condition_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v3/alerting_condition"
	alerting_condition "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_condition"
)

var (
	_ = context.Context(nil)
	_ = fmt.GoStringer(nil)

	_ = codes.NotFound
	_ = status.Status{}

	_ = watch_type.WatchType_STATEFUL
	_ = gotenresource.ListQuery(nil)
)

type apiAlertingConditionAccess struct {
	client alerting_condition_client.AlertingConditionServiceClient
}

func NewApiAlertingConditionAccess(client alerting_condition_client.AlertingConditionServiceClient) alerting_condition.AlertingConditionAccess {
	return &apiAlertingConditionAccess{client: client}
}

func (a *apiAlertingConditionAccess) GetAlertingCondition(ctx context.Context, query *alerting_condition.GetQuery) (*alerting_condition.AlertingCondition, error) {
	request := &alerting_condition_client.GetAlertingConditionRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	return a.client.GetAlertingCondition(ctx, request)
}

func (a *apiAlertingConditionAccess) BatchGetAlertingConditions(ctx context.Context, refs []*alerting_condition.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &alerting_condition_client.BatchGetAlertingConditionsRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetAlertingConditions(ctx, request)
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

func (a *apiAlertingConditionAccess) QueryAlertingConditions(ctx context.Context, query *alerting_condition.ListQuery) (*alerting_condition.QueryResultSnapshot, error) {
	request := &alerting_condition_client.ListAlertingConditionsRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListAlertingConditions(ctx, request)
	if err != nil {
		return nil, err
	}
	return &alerting_condition.QueryResultSnapshot{
		AlertingConditions: resp.AlertingConditions,
		NextPageCursor:     resp.NextPageToken,
		PrevPageCursor:     resp.PrevPageToken,
	}, nil
}

func (a *apiAlertingConditionAccess) WatchAlertingCondition(ctx context.Context, query *alerting_condition.GetQuery, observerCb func(*alerting_condition.AlertingConditionChange) error) error {
	request := &alerting_condition_client.WatchAlertingConditionRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
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
	}
	if query.Pager != nil {
		request.OrderBy = query.Pager.OrderBy
		request.PageSize = int32(query.Pager.Limit)
		request.PageToken = query.Pager.Cursor
	}
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

	if previousRes == nil {
		var err error
		previousRes, err = a.GetAlertingCondition(ctx, &alerting_condition.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if previousRes != nil {
		updateRequest := &alerting_condition_client.UpdateAlertingConditionRequest{
			AlertingCondition: res,
		}
		_, err := a.client.UpdateAlertingCondition(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &alerting_condition_client.CreateAlertingConditionRequest{
			AlertingCondition: res,
		}
		_, err := a.client.CreateAlertingCondition(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiAlertingConditionAccess) DeleteAlertingCondition(ctx context.Context, ref *alerting_condition.Reference, opts ...gotenresource.DeleteOption) error {
	request := &alerting_condition_client.DeleteAlertingConditionRequest{
		Name: ref,
	}
	_, err := a.client.DeleteAlertingCondition(ctx, request)
	return err
}
