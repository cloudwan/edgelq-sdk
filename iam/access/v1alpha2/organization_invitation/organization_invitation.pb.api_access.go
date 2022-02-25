// Code generated by protoc-gen-goten-access
// Resource: OrganizationInvitation
// DO NOT EDIT!!!

package organization_invitation_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	organization_invitation_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha2/organization_invitation"
	organization_invitation "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization_invitation"
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

type apiOrganizationInvitationAccess struct {
	client organization_invitation_client.OrganizationInvitationServiceClient
}

func NewApiOrganizationInvitationAccess(client organization_invitation_client.OrganizationInvitationServiceClient) organization_invitation.OrganizationInvitationAccess {
	return &apiOrganizationInvitationAccess{client: client}
}

func (a *apiOrganizationInvitationAccess) GetOrganizationInvitation(ctx context.Context, query *organization_invitation.GetQuery) (*organization_invitation.OrganizationInvitation, error) {
	request := &organization_invitation_client.GetOrganizationInvitationRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	return a.client.GetOrganizationInvitation(ctx, request)
}

func (a *apiOrganizationInvitationAccess) BatchGetOrganizationInvitations(ctx context.Context, refs []*organization_invitation.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &organization_invitation_client.BatchGetOrganizationInvitationsRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetOrganizationInvitations(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[organization_invitation.Name]*organization_invitation.OrganizationInvitation, len(refs))
	for _, resolvedRes := range resp.GetOrganizationInvitations() {
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

func (a *apiOrganizationInvitationAccess) QueryOrganizationInvitations(ctx context.Context, query *organization_invitation.ListQuery) (*organization_invitation.QueryResultSnapshot, error) {
	request := &organization_invitation_client.ListOrganizationInvitationsRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListOrganizationInvitations(ctx, request)
	if err != nil {
		return nil, err
	}
	return &organization_invitation.QueryResultSnapshot{
		OrganizationInvitations: resp.OrganizationInvitations,
		NextPageCursor:          resp.NextPageToken,
		PrevPageCursor:          resp.PrevPageToken,
	}, nil
}

func (a *apiOrganizationInvitationAccess) WatchOrganizationInvitation(ctx context.Context, query *organization_invitation.GetQuery, observerCb func(*organization_invitation.OrganizationInvitationChange) error) error {
	request := &organization_invitation_client.WatchOrganizationInvitationRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchOrganizationInvitation(ctx, request)
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

func (a *apiOrganizationInvitationAccess) WatchOrganizationInvitations(ctx context.Context, query *organization_invitation.WatchQuery, observerCb func(*organization_invitation.QueryResultChange) error) error {
	request := &organization_invitation_client.WatchOrganizationInvitationsRequest{
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
	changesStream, initErr := a.client.WatchOrganizationInvitations(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &organization_invitation.QueryResultChange{
			Changes:      respChange.OrganizationInvitationChanges,
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

func (a *apiOrganizationInvitationAccess) SaveOrganizationInvitation(ctx context.Context, res *organization_invitation.OrganizationInvitation, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetOrganizationInvitation(ctx, &organization_invitation.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &organization_invitation_client.UpdateOrganizationInvitationRequest{
			OrganizationInvitation: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*organization_invitation.OrganizationInvitation_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &organization_invitation_client.UpdateOrganizationInvitationRequest_CAS{
				ConditionalState: conditionalState.(*organization_invitation.OrganizationInvitation),
				FieldMask:        mask.(*organization_invitation.OrganizationInvitation_FieldMask),
			}
		}
		_, err := a.client.UpdateOrganizationInvitation(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &organization_invitation_client.CreateOrganizationInvitationRequest{
			OrganizationInvitation: res,
		}
		_, err := a.client.CreateOrganizationInvitation(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiOrganizationInvitationAccess) DeleteOrganizationInvitation(ctx context.Context, ref *organization_invitation.Reference, opts ...gotenresource.DeleteOption) error {
	request := &organization_invitation_client.DeleteOrganizationInvitationRequest{
		Name: ref,
	}
	_, err := a.client.DeleteOrganizationInvitation(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(organization_invitation.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return organization_invitation.AsAnyCastAccess(NewApiOrganizationInvitationAccess(organization_invitation_client.NewOrganizationInvitationServiceClient(cc)))
	})
}
