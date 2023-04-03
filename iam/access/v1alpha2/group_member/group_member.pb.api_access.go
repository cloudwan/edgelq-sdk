// Code generated by protoc-gen-goten-access
// Resource: GroupMember
// DO NOT EDIT!!!

package group_member_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	group_member_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha2/group_member"
	group_member "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/group_member"
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

type apiGroupMemberAccess struct {
	client group_member_client.GroupMemberServiceClient
}

func NewApiGroupMemberAccess(client group_member_client.GroupMemberServiceClient) group_member.GroupMemberAccess {
	return &apiGroupMemberAccess{client: client}
}

func (a *apiGroupMemberAccess) GetGroupMember(ctx context.Context, query *group_member.GetQuery) (*group_member.GroupMember, error) {
	request := &group_member_client.GetGroupMemberRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetGroupMember(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiGroupMemberAccess) BatchGetGroupMembers(ctx context.Context, refs []*group_member.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &group_member_client.BatchGetGroupMembersRequest{
		Names: refs,
	}
	fieldMask := batchGetOpts.GetFieldMask(group_member.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*group_member.GroupMember_FieldMask)
	}
	resp, err := a.client.BatchGetGroupMembers(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[group_member.Name]*group_member.GroupMember, len(refs))
	for _, resolvedRes := range resp.GetGroupMembers() {
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

func (a *apiGroupMemberAccess) QueryGroupMembers(ctx context.Context, query *group_member.ListQuery) (*group_member.QueryResultSnapshot, error) {
	request := &group_member_client.ListGroupMembersRequest{
		Filter:            query.Filter,
		FieldMask:         query.Mask,
		IncludePagingInfo: query.WithPagingInfo,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListGroupMembers(ctx, request)
	if err != nil {
		return nil, err
	}
	return &group_member.QueryResultSnapshot{
		GroupMembers:      resp.GroupMembers,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiGroupMemberAccess) WatchGroupMember(ctx context.Context, query *group_member.GetQuery, observerCb func(*group_member.GroupMemberChange) error) error {
	request := &group_member_client.WatchGroupMemberRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchGroupMember(ctx, request)
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

func (a *apiGroupMemberAccess) WatchGroupMembers(ctx context.Context, query *group_member.WatchQuery, observerCb func(*group_member.QueryResultChange) error) error {
	request := &group_member_client.WatchGroupMembersRequest{
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
	changesStream, initErr := a.client.WatchGroupMembers(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &group_member.QueryResultChange{
			Changes:      respChange.GroupMemberChanges,
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

func (a *apiGroupMemberAccess) SaveGroupMember(ctx context.Context, res *group_member.GroupMember, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetGroupMember(ctx, &group_member.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &group_member_client.UpdateGroupMemberRequest{
			GroupMember: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*group_member.GroupMember_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &group_member_client.UpdateGroupMemberRequest_CAS{
				ConditionalState: conditionalState.(*group_member.GroupMember),
				FieldMask:        mask.(*group_member.GroupMember_FieldMask),
			}
		}
		_, err := a.client.UpdateGroupMember(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &group_member_client.CreateGroupMemberRequest{
			GroupMember: res,
		}
		_, err := a.client.CreateGroupMember(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiGroupMemberAccess) DeleteGroupMember(ctx context.Context, ref *group_member.Reference, opts ...gotenresource.DeleteOption) error {
	request := &group_member_client.DeleteGroupMemberRequest{
		Name: ref,
	}
	_, err := a.client.DeleteGroupMember(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(group_member.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return group_member.AsAnyCastAccess(NewApiGroupMemberAccess(group_member_client.NewGroupMemberServiceClient(cc)))
	})
}
