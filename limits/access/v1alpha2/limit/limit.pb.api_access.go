// Code generated by protoc-gen-goten-access
// Resource: Limit
// DO NOT EDIT!!!

package limit_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	limit_client "github.com/cloudwan/edgelq-sdk/limits/client/v1alpha2/limit"
	limit "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/limit"
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

type apiLimitAccess struct {
	client limit_client.LimitServiceClient
}

func NewApiLimitAccess(client limit_client.LimitServiceClient) limit.LimitAccess {
	return &apiLimitAccess{client: client}
}

func (a *apiLimitAccess) GetLimit(ctx context.Context, query *limit.GetQuery) (*limit.Limit, error) {
	request := &limit_client.GetLimitRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetLimit(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiLimitAccess) BatchGetLimits(ctx context.Context, refs []*limit.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &limit_client.BatchGetLimitsRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetLimits(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[limit.Name]*limit.Limit, len(refs))
	for _, resolvedRes := range resp.GetLimits() {
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

func (a *apiLimitAccess) QueryLimits(ctx context.Context, query *limit.ListQuery) (*limit.QueryResultSnapshot, error) {
	request := &limit_client.ListLimitsRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListLimits(ctx, request)
	if err != nil {
		return nil, err
	}
	return &limit.QueryResultSnapshot{
		Limits:         resp.Limits,
		NextPageCursor: resp.NextPageToken,
		PrevPageCursor: resp.PrevPageToken,
	}, nil
}

func (a *apiLimitAccess) WatchLimit(ctx context.Context, query *limit.GetQuery, observerCb func(*limit.LimitChange) error) error {
	request := &limit_client.WatchLimitRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchLimit(ctx, request)
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

func (a *apiLimitAccess) WatchLimits(ctx context.Context, query *limit.WatchQuery, observerCb func(*limit.QueryResultChange) error) error {
	request := &limit_client.WatchLimitsRequest{
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
	changesStream, initErr := a.client.WatchLimits(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &limit.QueryResultChange{
			Changes:      respChange.LimitChanges,
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

func (a *apiLimitAccess) SaveLimit(ctx context.Context, res *limit.Limit, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetLimit(ctx, &limit.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &limit_client.UpdateLimitRequest{
			Limit: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*limit.Limit_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &limit_client.UpdateLimitRequest_CAS{
				ConditionalState: conditionalState.(*limit.Limit),
				FieldMask:        mask.(*limit.Limit_FieldMask),
			}
		}
		_, err := a.client.UpdateLimit(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		return fmt.Errorf("create operation on %s is prohibited", res.Name.AsReference().String())
	}
}

func (a *apiLimitAccess) DeleteLimit(ctx context.Context, ref *limit.Reference, opts ...gotenresource.DeleteOption) error {
	return fmt.Errorf("Delete operation on Limit is prohibited")
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(limit.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return limit.AsAnyCastAccess(NewApiLimitAccess(limit_client.NewLimitServiceClient(cc)))
	})
}
