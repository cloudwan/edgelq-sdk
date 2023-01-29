// Code generated by protoc-gen-goten-access
// Resource: AlertingPolicy
// DO NOT EDIT!!!

package alerting_policy_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	alerting_policy_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v3/alerting_policy"
	alerting_policy "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_policy"
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

type apiAlertingPolicyAccess struct {
	client alerting_policy_client.AlertingPolicyServiceClient
}

func NewApiAlertingPolicyAccess(client alerting_policy_client.AlertingPolicyServiceClient) alerting_policy.AlertingPolicyAccess {
	return &apiAlertingPolicyAccess{client: client}
}

func (a *apiAlertingPolicyAccess) GetAlertingPolicy(ctx context.Context, query *alerting_policy.GetQuery) (*alerting_policy.AlertingPolicy, error) {
	request := &alerting_policy_client.GetAlertingPolicyRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetAlertingPolicy(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiAlertingPolicyAccess) BatchGetAlertingPolicies(ctx context.Context, refs []*alerting_policy.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &alerting_policy_client.BatchGetAlertingPoliciesRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetAlertingPolicies(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[alerting_policy.Name]*alerting_policy.AlertingPolicy, len(refs))
	for _, resolvedRes := range resp.GetAlertingPolicies() {
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

func (a *apiAlertingPolicyAccess) QueryAlertingPolicies(ctx context.Context, query *alerting_policy.ListQuery) (*alerting_policy.QueryResultSnapshot, error) {
	request := &alerting_policy_client.ListAlertingPoliciesRequest{
		Filter:            query.Filter,
		FieldMask:         query.Mask,
		IncludePagingInfo: query.WithPagingInfo,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListAlertingPolicies(ctx, request)
	if err != nil {
		return nil, err
	}
	return &alerting_policy.QueryResultSnapshot{
		AlertingPolicies:  resp.AlertingPolicies,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiAlertingPolicyAccess) SearchAlertingPolicies(ctx context.Context, query *alerting_policy.SearchQuery) (*alerting_policy.QueryResultSnapshot, error) {
	request := &alerting_policy_client.SearchAlertingPoliciesRequest{
		Phrase:    query.Phrase,
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.SearchAlertingPolicies(ctx, request)
	if err != nil {
		return nil, err
	}
	return &alerting_policy.QueryResultSnapshot{
		AlertingPolicies:  resp.AlertingPolicies,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		CurrentOffset:     resp.CurrentOffset,
		TotalResultsCount: resp.TotalResultsCount,
	}, nil
}

func (a *apiAlertingPolicyAccess) WatchAlertingPolicy(ctx context.Context, query *alerting_policy.GetQuery, observerCb func(*alerting_policy.AlertingPolicyChange) error) error {
	request := &alerting_policy_client.WatchAlertingPolicyRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchAlertingPolicy(ctx, request)
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

func (a *apiAlertingPolicyAccess) WatchAlertingPolicies(ctx context.Context, query *alerting_policy.WatchQuery, observerCb func(*alerting_policy.QueryResultChange) error) error {
	request := &alerting_policy_client.WatchAlertingPoliciesRequest{
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
	changesStream, initErr := a.client.WatchAlertingPolicies(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &alerting_policy.QueryResultChange{
			Changes:      respChange.AlertingPolicyChanges,
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

func (a *apiAlertingPolicyAccess) SaveAlertingPolicy(ctx context.Context, res *alerting_policy.AlertingPolicy, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetAlertingPolicy(ctx, &alerting_policy.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &alerting_policy_client.UpdateAlertingPolicyRequest{
			AlertingPolicy: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*alerting_policy.AlertingPolicy_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &alerting_policy_client.UpdateAlertingPolicyRequest_CAS{
				ConditionalState: conditionalState.(*alerting_policy.AlertingPolicy),
				FieldMask:        mask.(*alerting_policy.AlertingPolicy_FieldMask),
			}
		}
		_, err := a.client.UpdateAlertingPolicy(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &alerting_policy_client.CreateAlertingPolicyRequest{
			AlertingPolicy: res,
		}
		_, err := a.client.CreateAlertingPolicy(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiAlertingPolicyAccess) DeleteAlertingPolicy(ctx context.Context, ref *alerting_policy.Reference, opts ...gotenresource.DeleteOption) error {
	request := &alerting_policy_client.DeleteAlertingPolicyRequest{
		Name: ref,
	}
	_, err := a.client.DeleteAlertingPolicy(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(alerting_policy.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return alerting_policy.AsAnyCastAccess(NewApiAlertingPolicyAccess(alerting_policy_client.NewAlertingPolicyServiceClient(cc)))
	})
}
