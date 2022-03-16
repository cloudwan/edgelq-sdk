// Code generated by protoc-gen-goten-access
// Resource: Role
// DO NOT EDIT!!!

package role_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	role_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha2/role"
	role "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/role"
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

type apiRoleAccess struct {
	client role_client.RoleServiceClient
}

func NewApiRoleAccess(client role_client.RoleServiceClient) role.RoleAccess {
	return &apiRoleAccess{client: client}
}

func (a *apiRoleAccess) GetRole(ctx context.Context, query *role.GetQuery) (*role.Role, error) {
	request := &role_client.GetRoleRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetRole(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiRoleAccess) BatchGetRoles(ctx context.Context, refs []*role.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &role_client.BatchGetRolesRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetRoles(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[role.Name]*role.Role, len(refs))
	for _, resolvedRes := range resp.GetRoles() {
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

func (a *apiRoleAccess) QueryRoles(ctx context.Context, query *role.ListQuery) (*role.QueryResultSnapshot, error) {
	request := &role_client.ListRolesRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListRoles(ctx, request)
	if err != nil {
		return nil, err
	}
	return &role.QueryResultSnapshot{
		Roles:          resp.Roles,
		NextPageCursor: resp.NextPageToken,
		PrevPageCursor: resp.PrevPageToken,
	}, nil
}

func (a *apiRoleAccess) WatchRole(ctx context.Context, query *role.GetQuery, observerCb func(*role.RoleChange) error) error {
	request := &role_client.WatchRoleRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchRole(ctx, request)
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

func (a *apiRoleAccess) WatchRoles(ctx context.Context, query *role.WatchQuery, observerCb func(*role.QueryResultChange) error) error {
	request := &role_client.WatchRolesRequest{
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
	changesStream, initErr := a.client.WatchRoles(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &role.QueryResultChange{
			Changes:      respChange.RoleChanges,
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

func (a *apiRoleAccess) SaveRole(ctx context.Context, res *role.Role, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetRole(ctx, &role.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &role_client.UpdateRoleRequest{
			Role: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*role.Role_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &role_client.UpdateRoleRequest_CAS{
				ConditionalState: conditionalState.(*role.Role),
				FieldMask:        mask.(*role.Role_FieldMask),
			}
		}
		_, err := a.client.UpdateRole(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &role_client.CreateRoleRequest{
			Role: res,
		}
		_, err := a.client.CreateRole(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiRoleAccess) DeleteRole(ctx context.Context, ref *role.Reference, opts ...gotenresource.DeleteOption) error {
	request := &role_client.DeleteRoleRequest{
		Name: ref,
	}
	_, err := a.client.DeleteRole(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(role.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return role.AsAnyCastAccess(NewApiRoleAccess(role_client.NewRoleServiceClient(cc)))
	})
}
