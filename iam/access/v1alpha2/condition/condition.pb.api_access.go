// Code generated by protoc-gen-goten-access
// Resource: Condition
// DO NOT EDIT!!!

package condition_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	condition_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha2/condition"
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/condition"
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

type apiConditionAccess struct {
	client condition_client.ConditionServiceClient
}

func NewApiConditionAccess(client condition_client.ConditionServiceClient) condition.ConditionAccess {
	return &apiConditionAccess{client: client}
}

func (a *apiConditionAccess) GetCondition(ctx context.Context, query *condition.GetQuery) (*condition.Condition, error) {
	request := &condition_client.GetConditionRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetCondition(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiConditionAccess) BatchGetConditions(ctx context.Context, refs []*condition.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &condition_client.BatchGetConditionsRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetConditions(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[condition.Name]*condition.Condition, len(refs))
	for _, resolvedRes := range resp.GetConditions() {
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

func (a *apiConditionAccess) QueryConditions(ctx context.Context, query *condition.ListQuery) (*condition.QueryResultSnapshot, error) {
	request := &condition_client.ListConditionsRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListConditions(ctx, request)
	if err != nil {
		return nil, err
	}
	return &condition.QueryResultSnapshot{
		Conditions:     resp.Conditions,
		NextPageCursor: resp.NextPageToken,
		PrevPageCursor: resp.PrevPageToken,
	}, nil
}

func (a *apiConditionAccess) WatchCondition(ctx context.Context, query *condition.GetQuery, observerCb func(*condition.ConditionChange) error) error {
	request := &condition_client.WatchConditionRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchCondition(ctx, request)
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

func (a *apiConditionAccess) WatchConditions(ctx context.Context, query *condition.WatchQuery, observerCb func(*condition.QueryResultChange) error) error {
	request := &condition_client.WatchConditionsRequest{
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
	changesStream, initErr := a.client.WatchConditions(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &condition.QueryResultChange{
			Changes:      respChange.ConditionChanges,
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

func (a *apiConditionAccess) SaveCondition(ctx context.Context, res *condition.Condition, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetCondition(ctx, &condition.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &condition_client.UpdateConditionRequest{
			Condition: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*condition.Condition_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &condition_client.UpdateConditionRequest_CAS{
				ConditionalState: conditionalState.(*condition.Condition),
				FieldMask:        mask.(*condition.Condition_FieldMask),
			}
		}
		_, err := a.client.UpdateCondition(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &condition_client.CreateConditionRequest{
			Condition: res,
		}
		_, err := a.client.CreateCondition(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiConditionAccess) DeleteCondition(ctx context.Context, ref *condition.Reference, opts ...gotenresource.DeleteOption) error {
	request := &condition_client.DeleteConditionRequest{
		Name: ref,
	}
	_, err := a.client.DeleteCondition(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(condition.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return condition.AsAnyCastAccess(NewApiConditionAccess(condition_client.NewConditionServiceClient(cc)))
	})
}
