// Code generated by protoc-gen-goten-access
// Resource: Permission
// DO NOT EDIT!!!

package permission_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	permission_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha2/permission"
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/permission"
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

type apiPermissionAccess struct {
	client permission_client.PermissionServiceClient
}

func NewApiPermissionAccess(client permission_client.PermissionServiceClient) permission.PermissionAccess {
	return &apiPermissionAccess{client: client}
}

func (a *apiPermissionAccess) GetPermission(ctx context.Context, query *permission.GetQuery) (*permission.Permission, error) {
	request := &permission_client.GetPermissionRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetPermission(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiPermissionAccess) BatchGetPermissions(ctx context.Context, refs []*permission.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &permission_client.BatchGetPermissionsRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetPermissions(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[permission.Name]*permission.Permission, len(refs))
	for _, resolvedRes := range resp.GetPermissions() {
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

func (a *apiPermissionAccess) QueryPermissions(ctx context.Context, query *permission.ListQuery) (*permission.QueryResultSnapshot, error) {
	request := &permission_client.ListPermissionsRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListPermissions(ctx, request)
	if err != nil {
		return nil, err
	}
	return &permission.QueryResultSnapshot{
		Permissions:    resp.Permissions,
		NextPageCursor: resp.NextPageToken,
		PrevPageCursor: resp.PrevPageToken,
	}, nil
}

func (a *apiPermissionAccess) WatchPermission(ctx context.Context, query *permission.GetQuery, observerCb func(*permission.PermissionChange) error) error {
	request := &permission_client.WatchPermissionRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchPermission(ctx, request)
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

func (a *apiPermissionAccess) WatchPermissions(ctx context.Context, query *permission.WatchQuery, observerCb func(*permission.QueryResultChange) error) error {
	request := &permission_client.WatchPermissionsRequest{
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
	changesStream, initErr := a.client.WatchPermissions(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &permission.QueryResultChange{
			Changes:      respChange.PermissionChanges,
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

func (a *apiPermissionAccess) SavePermission(ctx context.Context, res *permission.Permission, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetPermission(ctx, &permission.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &permission_client.UpdatePermissionRequest{
			Permission: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*permission.Permission_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &permission_client.UpdatePermissionRequest_CAS{
				ConditionalState: conditionalState.(*permission.Permission),
				FieldMask:        mask.(*permission.Permission_FieldMask),
			}
		}
		_, err := a.client.UpdatePermission(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &permission_client.CreatePermissionRequest{
			Permission: res,
		}
		_, err := a.client.CreatePermission(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiPermissionAccess) DeletePermission(ctx context.Context, ref *permission.Reference, opts ...gotenresource.DeleteOption) error {
	request := &permission_client.DeletePermissionRequest{
		Name: ref,
	}
	_, err := a.client.DeletePermission(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(permission.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return permission.AsAnyCastAccess(NewApiPermissionAccess(permission_client.NewPermissionServiceClient(cc)))
	})
}
