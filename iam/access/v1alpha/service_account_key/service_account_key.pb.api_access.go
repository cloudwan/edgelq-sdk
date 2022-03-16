// Code generated by protoc-gen-goten-access
// Resource: ServiceAccountKey
// DO NOT EDIT!!!

package service_account_key_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	service_account_key_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha/service_account_key"
	service_account_key "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/service_account_key"
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

type apiServiceAccountKeyAccess struct {
	client service_account_key_client.ServiceAccountKeyServiceClient
}

func NewApiServiceAccountKeyAccess(client service_account_key_client.ServiceAccountKeyServiceClient) service_account_key.ServiceAccountKeyAccess {
	return &apiServiceAccountKeyAccess{client: client}
}

func (a *apiServiceAccountKeyAccess) GetServiceAccountKey(ctx context.Context, query *service_account_key.GetQuery) (*service_account_key.ServiceAccountKey, error) {
	request := &service_account_key_client.GetServiceAccountKeyRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetServiceAccountKey(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiServiceAccountKeyAccess) BatchGetServiceAccountKeys(ctx context.Context, refs []*service_account_key.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &service_account_key_client.BatchGetServiceAccountKeysRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetServiceAccountKeys(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[service_account_key.Name]*service_account_key.ServiceAccountKey, len(refs))
	for _, resolvedRes := range resp.GetServiceAccountKeys() {
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

func (a *apiServiceAccountKeyAccess) QueryServiceAccountKeys(ctx context.Context, query *service_account_key.ListQuery) (*service_account_key.QueryResultSnapshot, error) {
	request := &service_account_key_client.ListServiceAccountKeysRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListServiceAccountKeys(ctx, request)
	if err != nil {
		return nil, err
	}
	return &service_account_key.QueryResultSnapshot{
		ServiceAccountKeys: resp.ServiceAccountKeys,
		NextPageCursor:     resp.NextPageToken,
		PrevPageCursor:     resp.PrevPageToken,
	}, nil
}

func (a *apiServiceAccountKeyAccess) WatchServiceAccountKey(ctx context.Context, query *service_account_key.GetQuery, observerCb func(*service_account_key.ServiceAccountKeyChange) error) error {
	request := &service_account_key_client.WatchServiceAccountKeyRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchServiceAccountKey(ctx, request)
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

func (a *apiServiceAccountKeyAccess) WatchServiceAccountKeys(ctx context.Context, query *service_account_key.WatchQuery, observerCb func(*service_account_key.QueryResultChange) error) error {
	request := &service_account_key_client.WatchServiceAccountKeysRequest{
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
	changesStream, initErr := a.client.WatchServiceAccountKeys(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &service_account_key.QueryResultChange{
			Changes:      respChange.ServiceAccountKeyChanges,
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

func (a *apiServiceAccountKeyAccess) SaveServiceAccountKey(ctx context.Context, res *service_account_key.ServiceAccountKey, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetServiceAccountKey(ctx, &service_account_key.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &service_account_key_client.UpdateServiceAccountKeyRequest{
			ServiceAccountKey: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*service_account_key.ServiceAccountKey_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &service_account_key_client.UpdateServiceAccountKeyRequest_CAS{
				ConditionalState: conditionalState.(*service_account_key.ServiceAccountKey),
				FieldMask:        mask.(*service_account_key.ServiceAccountKey_FieldMask),
			}
		}
		_, err := a.client.UpdateServiceAccountKey(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &service_account_key_client.CreateServiceAccountKeyRequest{
			ServiceAccountKey: res,
		}
		_, err := a.client.CreateServiceAccountKey(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiServiceAccountKeyAccess) DeleteServiceAccountKey(ctx context.Context, ref *service_account_key.Reference, opts ...gotenresource.DeleteOption) error {
	request := &service_account_key_client.DeleteServiceAccountKeyRequest{
		Name: ref,
	}
	_, err := a.client.DeleteServiceAccountKey(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(service_account_key.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return service_account_key.AsAnyCastAccess(NewApiServiceAccountKeyAccess(service_account_key_client.NewServiceAccountKeyServiceClient(cc)))
	})
}
