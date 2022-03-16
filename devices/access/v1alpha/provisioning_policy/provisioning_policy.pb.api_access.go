// Code generated by protoc-gen-goten-access
// Resource: ProvisioningPolicy
// DO NOT EDIT!!!

package provisioning_policy_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	provisioning_policy_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha/provisioning_policy"
	provisioning_policy "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha/provisioning_policy"
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

type apiProvisioningPolicyAccess struct {
	client provisioning_policy_client.ProvisioningPolicyServiceClient
}

func NewApiProvisioningPolicyAccess(client provisioning_policy_client.ProvisioningPolicyServiceClient) provisioning_policy.ProvisioningPolicyAccess {
	return &apiProvisioningPolicyAccess{client: client}
}

func (a *apiProvisioningPolicyAccess) GetProvisioningPolicy(ctx context.Context, query *provisioning_policy.GetQuery) (*provisioning_policy.ProvisioningPolicy, error) {
	request := &provisioning_policy_client.GetProvisioningPolicyRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetProvisioningPolicy(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiProvisioningPolicyAccess) BatchGetProvisioningPolicys(ctx context.Context, refs []*provisioning_policy.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &provisioning_policy_client.BatchGetProvisioningPolicysRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetProvisioningPolicys(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[provisioning_policy.Name]*provisioning_policy.ProvisioningPolicy, len(refs))
	for _, resolvedRes := range resp.GetProvisioningPolicys() {
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

func (a *apiProvisioningPolicyAccess) QueryProvisioningPolicys(ctx context.Context, query *provisioning_policy.ListQuery) (*provisioning_policy.QueryResultSnapshot, error) {
	request := &provisioning_policy_client.ListProvisioningPolicysRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListProvisioningPolicys(ctx, request)
	if err != nil {
		return nil, err
	}
	return &provisioning_policy.QueryResultSnapshot{
		ProvisioningPolicys: resp.ProvisioningPolicys,
		NextPageCursor:      resp.NextPageToken,
		PrevPageCursor:      resp.PrevPageToken,
	}, nil
}

func (a *apiProvisioningPolicyAccess) WatchProvisioningPolicy(ctx context.Context, query *provisioning_policy.GetQuery, observerCb func(*provisioning_policy.ProvisioningPolicyChange) error) error {
	request := &provisioning_policy_client.WatchProvisioningPolicyRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchProvisioningPolicy(ctx, request)
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

func (a *apiProvisioningPolicyAccess) WatchProvisioningPolicys(ctx context.Context, query *provisioning_policy.WatchQuery, observerCb func(*provisioning_policy.QueryResultChange) error) error {
	request := &provisioning_policy_client.WatchProvisioningPolicysRequest{
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
	changesStream, initErr := a.client.WatchProvisioningPolicys(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &provisioning_policy.QueryResultChange{
			Changes:      respChange.ProvisioningPolicyChanges,
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

func (a *apiProvisioningPolicyAccess) SaveProvisioningPolicy(ctx context.Context, res *provisioning_policy.ProvisioningPolicy, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetProvisioningPolicy(ctx, &provisioning_policy.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &provisioning_policy_client.UpdateProvisioningPolicyRequest{
			ProvisioningPolicy: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*provisioning_policy.ProvisioningPolicy_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &provisioning_policy_client.UpdateProvisioningPolicyRequest_CAS{
				ConditionalState: conditionalState.(*provisioning_policy.ProvisioningPolicy),
				FieldMask:        mask.(*provisioning_policy.ProvisioningPolicy_FieldMask),
			}
		}
		_, err := a.client.UpdateProvisioningPolicy(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &provisioning_policy_client.CreateProvisioningPolicyRequest{
			ProvisioningPolicy: res,
		}
		_, err := a.client.CreateProvisioningPolicy(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiProvisioningPolicyAccess) DeleteProvisioningPolicy(ctx context.Context, ref *provisioning_policy.Reference, opts ...gotenresource.DeleteOption) error {
	request := &provisioning_policy_client.DeleteProvisioningPolicyRequest{
		Name: ref,
	}
	_, err := a.client.DeleteProvisioningPolicy(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(provisioning_policy.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return provisioning_policy.AsAnyCastAccess(NewApiProvisioningPolicyAccess(provisioning_policy_client.NewProvisioningPolicyServiceClient(cc)))
	})
}
