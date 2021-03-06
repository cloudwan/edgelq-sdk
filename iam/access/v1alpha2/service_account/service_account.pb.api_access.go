// Code generated by protoc-gen-goten-access
// Resource: ServiceAccount
// DO NOT EDIT!!!

package service_account_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	service_account_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha2/service_account"
	service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/service_account"
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

type apiServiceAccountAccess struct {
	client service_account_client.ServiceAccountServiceClient
}

func NewApiServiceAccountAccess(client service_account_client.ServiceAccountServiceClient) service_account.ServiceAccountAccess {
	return &apiServiceAccountAccess{client: client}
}

func (a *apiServiceAccountAccess) GetServiceAccount(ctx context.Context, query *service_account.GetQuery) (*service_account.ServiceAccount, error) {
	request := &service_account_client.GetServiceAccountRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetServiceAccount(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiServiceAccountAccess) BatchGetServiceAccounts(ctx context.Context, refs []*service_account.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &service_account_client.BatchGetServiceAccountsRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetServiceAccounts(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[service_account.Name]*service_account.ServiceAccount, len(refs))
	for _, resolvedRes := range resp.GetServiceAccounts() {
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

func (a *apiServiceAccountAccess) QueryServiceAccounts(ctx context.Context, query *service_account.ListQuery) (*service_account.QueryResultSnapshot, error) {
	request := &service_account_client.ListServiceAccountsRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListServiceAccounts(ctx, request)
	if err != nil {
		return nil, err
	}
	return &service_account.QueryResultSnapshot{
		ServiceAccounts: resp.ServiceAccounts,
		NextPageCursor:  resp.NextPageToken,
		PrevPageCursor:  resp.PrevPageToken,
	}, nil
}

func (a *apiServiceAccountAccess) WatchServiceAccount(ctx context.Context, query *service_account.GetQuery, observerCb func(*service_account.ServiceAccountChange) error) error {
	request := &service_account_client.WatchServiceAccountRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchServiceAccount(ctx, request)
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

func (a *apiServiceAccountAccess) WatchServiceAccounts(ctx context.Context, query *service_account.WatchQuery, observerCb func(*service_account.QueryResultChange) error) error {
	request := &service_account_client.WatchServiceAccountsRequest{
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
	changesStream, initErr := a.client.WatchServiceAccounts(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &service_account.QueryResultChange{
			Changes:      respChange.ServiceAccountChanges,
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

func (a *apiServiceAccountAccess) SaveServiceAccount(ctx context.Context, res *service_account.ServiceAccount, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetServiceAccount(ctx, &service_account.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &service_account_client.UpdateServiceAccountRequest{
			ServiceAccount: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*service_account.ServiceAccount_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &service_account_client.UpdateServiceAccountRequest_CAS{
				ConditionalState: conditionalState.(*service_account.ServiceAccount),
				FieldMask:        mask.(*service_account.ServiceAccount_FieldMask),
			}
		}
		_, err := a.client.UpdateServiceAccount(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &service_account_client.CreateServiceAccountRequest{
			ServiceAccount: res,
		}
		_, err := a.client.CreateServiceAccount(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiServiceAccountAccess) DeleteServiceAccount(ctx context.Context, ref *service_account.Reference, opts ...gotenresource.DeleteOption) error {
	request := &service_account_client.DeleteServiceAccountRequest{
		Name: ref,
	}
	_, err := a.client.DeleteServiceAccount(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(service_account.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return service_account.AsAnyCastAccess(NewApiServiceAccountAccess(service_account_client.NewServiceAccountServiceClient(cc)))
	})
}
