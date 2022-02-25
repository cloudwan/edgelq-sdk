// Code generated by protoc-gen-goten-access
// Resource: User
// DO NOT EDIT!!!

package user_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	user_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha2/user"
	user "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/user"
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

type apiUserAccess struct {
	client user_client.UserServiceClient
}

func NewApiUserAccess(client user_client.UserServiceClient) user.UserAccess {
	return &apiUserAccess{client: client}
}

func (a *apiUserAccess) GetUser(ctx context.Context, query *user.GetQuery) (*user.User, error) {
	request := &user_client.GetUserRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	return a.client.GetUser(ctx, request)
}

func (a *apiUserAccess) BatchGetUsers(ctx context.Context, refs []*user.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &user_client.BatchGetUsersRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetUsers(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[user.Name]*user.User, len(refs))
	for _, resolvedRes := range resp.GetUsers() {
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

func (a *apiUserAccess) QueryUsers(ctx context.Context, query *user.ListQuery) (*user.QueryResultSnapshot, error) {
	request := &user_client.ListUsersRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListUsers(ctx, request)
	if err != nil {
		return nil, err
	}
	return &user.QueryResultSnapshot{
		Users:          resp.Users,
		NextPageCursor: resp.NextPageToken,
		PrevPageCursor: resp.PrevPageToken,
	}, nil
}

func (a *apiUserAccess) WatchUser(ctx context.Context, query *user.GetQuery, observerCb func(*user.UserChange) error) error {
	request := &user_client.WatchUserRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchUser(ctx, request)
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

func (a *apiUserAccess) WatchUsers(ctx context.Context, query *user.WatchQuery, observerCb func(*user.QueryResultChange) error) error {
	request := &user_client.WatchUsersRequest{
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
	changesStream, initErr := a.client.WatchUsers(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &user.QueryResultChange{
			Changes:      respChange.UserChanges,
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

func (a *apiUserAccess) SaveUser(ctx context.Context, res *user.User, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetUser(ctx, &user.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &user_client.UpdateUserRequest{
			User: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*user.User_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &user_client.UpdateUserRequest_CAS{
				ConditionalState: conditionalState.(*user.User),
				FieldMask:        mask.(*user.User_FieldMask),
			}
		}
		_, err := a.client.UpdateUser(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &user_client.CreateUserRequest{
			User: res,
		}
		_, err := a.client.CreateUser(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiUserAccess) DeleteUser(ctx context.Context, ref *user.Reference, opts ...gotenresource.DeleteOption) error {
	request := &user_client.DeleteUserRequest{
		Name: ref,
	}
	_, err := a.client.DeleteUser(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(user.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return user.AsAnyCastAccess(NewApiUserAccess(user_client.NewUserServiceClient(cc)))
	})
}
