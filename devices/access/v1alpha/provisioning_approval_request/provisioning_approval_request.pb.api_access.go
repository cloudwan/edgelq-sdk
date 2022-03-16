// Code generated by protoc-gen-goten-access
// Resource: ProvisioningApprovalRequest
// DO NOT EDIT!!!

package provisioning_approval_request_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	provisioning_approval_request_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha/provisioning_approval_request"
	provisioning_approval_request "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha/provisioning_approval_request"
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

type apiProvisioningApprovalRequestAccess struct {
	client provisioning_approval_request_client.ProvisioningApprovalRequestServiceClient
}

func NewApiProvisioningApprovalRequestAccess(client provisioning_approval_request_client.ProvisioningApprovalRequestServiceClient) provisioning_approval_request.ProvisioningApprovalRequestAccess {
	return &apiProvisioningApprovalRequestAccess{client: client}
}

func (a *apiProvisioningApprovalRequestAccess) GetProvisioningApprovalRequest(ctx context.Context, query *provisioning_approval_request.GetQuery) (*provisioning_approval_request.ProvisioningApprovalRequest, error) {
	request := &provisioning_approval_request_client.GetProvisioningApprovalRequestRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetProvisioningApprovalRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiProvisioningApprovalRequestAccess) BatchGetProvisioningApprovalRequests(ctx context.Context, refs []*provisioning_approval_request.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &provisioning_approval_request_client.BatchGetProvisioningApprovalRequestsRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetProvisioningApprovalRequests(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[provisioning_approval_request.Name]*provisioning_approval_request.ProvisioningApprovalRequest, len(refs))
	for _, resolvedRes := range resp.GetProvisioningApprovalRequests() {
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

func (a *apiProvisioningApprovalRequestAccess) QueryProvisioningApprovalRequests(ctx context.Context, query *provisioning_approval_request.ListQuery) (*provisioning_approval_request.QueryResultSnapshot, error) {
	request := &provisioning_approval_request_client.ListProvisioningApprovalRequestsRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListProvisioningApprovalRequests(ctx, request)
	if err != nil {
		return nil, err
	}
	return &provisioning_approval_request.QueryResultSnapshot{
		ProvisioningApprovalRequests: resp.ProvisioningApprovalRequests,
		NextPageCursor:               resp.NextPageToken,
		PrevPageCursor:               resp.PrevPageToken,
	}, nil
}

func (a *apiProvisioningApprovalRequestAccess) WatchProvisioningApprovalRequest(ctx context.Context, query *provisioning_approval_request.GetQuery, observerCb func(*provisioning_approval_request.ProvisioningApprovalRequestChange) error) error {
	request := &provisioning_approval_request_client.WatchProvisioningApprovalRequestRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchProvisioningApprovalRequest(ctx, request)
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

func (a *apiProvisioningApprovalRequestAccess) WatchProvisioningApprovalRequests(ctx context.Context, query *provisioning_approval_request.WatchQuery, observerCb func(*provisioning_approval_request.QueryResultChange) error) error {
	request := &provisioning_approval_request_client.WatchProvisioningApprovalRequestsRequest{
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
	changesStream, initErr := a.client.WatchProvisioningApprovalRequests(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &provisioning_approval_request.QueryResultChange{
			Changes:      respChange.ProvisioningApprovalRequestChanges,
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

func (a *apiProvisioningApprovalRequestAccess) SaveProvisioningApprovalRequest(ctx context.Context, res *provisioning_approval_request.ProvisioningApprovalRequest, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetProvisioningApprovalRequest(ctx, &provisioning_approval_request.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &provisioning_approval_request_client.UpdateProvisioningApprovalRequestRequest{
			ProvisioningApprovalRequest: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*provisioning_approval_request.ProvisioningApprovalRequest_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &provisioning_approval_request_client.UpdateProvisioningApprovalRequestRequest_CAS{
				ConditionalState: conditionalState.(*provisioning_approval_request.ProvisioningApprovalRequest),
				FieldMask:        mask.(*provisioning_approval_request.ProvisioningApprovalRequest_FieldMask),
			}
		}
		_, err := a.client.UpdateProvisioningApprovalRequest(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &provisioning_approval_request_client.CreateProvisioningApprovalRequestRequest{
			ProvisioningApprovalRequest: res,
		}
		_, err := a.client.CreateProvisioningApprovalRequest(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiProvisioningApprovalRequestAccess) DeleteProvisioningApprovalRequest(ctx context.Context, ref *provisioning_approval_request.Reference, opts ...gotenresource.DeleteOption) error {
	request := &provisioning_approval_request_client.DeleteProvisioningApprovalRequestRequest{
		Name: ref,
	}
	_, err := a.client.DeleteProvisioningApprovalRequest(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(provisioning_approval_request.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return provisioning_approval_request.AsAnyCastAccess(NewApiProvisioningApprovalRequestAccess(provisioning_approval_request_client.NewProvisioningApprovalRequestServiceClient(cc)))
	})
}
