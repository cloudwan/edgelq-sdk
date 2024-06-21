// Code generated by protoc-gen-goten-access
// Resource: ProjectInvitation
// DO NOT EDIT!!!

package project_invitation_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
	gotenfilter "github.com/cloudwan/goten-sdk/runtime/resource/filter"
	"github.com/cloudwan/goten-sdk/types/watch_type"

	project_invitation_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha2/project_invitation"
	project_invitation "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project_invitation"
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
	_ = gotenfilter.Eq
)

type apiProjectInvitationAccess struct {
	client project_invitation_client.ProjectInvitationServiceClient
}

func NewApiProjectInvitationAccess(client project_invitation_client.ProjectInvitationServiceClient) project_invitation.ProjectInvitationAccess {
	return &apiProjectInvitationAccess{client: client}
}

func (a *apiProjectInvitationAccess) GetProjectInvitation(ctx context.Context, query *project_invitation.GetQuery) (*project_invitation.ProjectInvitation, error) {
	if !query.Reference.IsFullyQualified() {
		return nil, status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &project_invitation_client.GetProjectInvitationRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetProjectInvitation(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiProjectInvitationAccess) BatchGetProjectInvitations(ctx context.Context, refs []*project_invitation.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	asNames := make([]*project_invitation.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &project_invitation_client.BatchGetProjectInvitationsRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(project_invitation.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*project_invitation.ProjectInvitation_FieldMask)
	}
	resp, err := a.client.BatchGetProjectInvitations(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[project_invitation.Name]*project_invitation.ProjectInvitation, len(refs))
	for _, resolvedRes := range resp.GetProjectInvitations() {
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

func (a *apiProjectInvitationAccess) QueryProjectInvitations(ctx context.Context, query *project_invitation.ListQuery) (*project_invitation.QueryResultSnapshot, error) {
	request := &project_invitation_client.ListProjectInvitationsRequest{
		Filter:            query.Filter,
		FieldMask:         query.Mask,
		IncludePagingInfo: query.WithPagingInfo,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	if query.Filter != nil && query.Filter.GetCondition() != nil {
		request.Filter, request.Parent = getParentAndFilter(query.Filter)
	}
	resp, err := a.client.ListProjectInvitations(ctx, request)
	if err != nil {
		return nil, err
	}
	return &project_invitation.QueryResultSnapshot{
		ProjectInvitations: resp.ProjectInvitations,
		NextPageCursor:     resp.NextPageToken,
		PrevPageCursor:     resp.PrevPageToken,
		TotalResultsCount:  resp.TotalResultsCount,
		CurrentOffset:      resp.CurrentOffset,
	}, nil
}

func (a *apiProjectInvitationAccess) WatchProjectInvitation(ctx context.Context, query *project_invitation.GetQuery, observerCb func(*project_invitation.ProjectInvitationChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &project_invitation_client.WatchProjectInvitationRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchProjectInvitation(ctx, request)
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

func (a *apiProjectInvitationAccess) WatchProjectInvitations(ctx context.Context, query *project_invitation.WatchQuery, observerCb func(*project_invitation.QueryResultChange) error) error {
	request := &project_invitation_client.WatchProjectInvitationsRequest{
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
	if query.Filter != nil && query.Filter.GetCondition() != nil {
		request.Filter, request.Parent = getParentAndFilter(query.Filter)
	}
	changesStream, initErr := a.client.WatchProjectInvitations(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &project_invitation.QueryResultChange{
			Changes:      respChange.ProjectInvitationChanges,
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

func (a *apiProjectInvitationAccess) SaveProjectInvitation(ctx context.Context, res *project_invitation.ProjectInvitation, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetProjectInvitation(ctx, &project_invitation.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}
	var resp *project_invitation.ProjectInvitation
	var err error
	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &project_invitation_client.UpdateProjectInvitationRequest{
			ProjectInvitation: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*project_invitation.ProjectInvitation_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &project_invitation_client.UpdateProjectInvitationRequest_CAS{
				ConditionalState: conditionalState.(*project_invitation.ProjectInvitation),
				FieldMask:        mask.(*project_invitation.ProjectInvitation_FieldMask),
			}
		}
		resp, err = a.client.UpdateProjectInvitation(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &project_invitation_client.CreateProjectInvitationRequest{
			ProjectInvitation: res,
		}
		resp, err = a.client.CreateProjectInvitation(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiProjectInvitationAccess) DeleteProjectInvitation(ctx context.Context, ref *project_invitation.Reference, opts ...gotenresource.DeleteOption) error {
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &project_invitation_client.DeleteProjectInvitationRequest{
		Name: &ref.Name,
	}
	_, err := a.client.DeleteProjectInvitation(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *project_invitation.Filter) (*project_invitation.Filter, *project_invitation.ParentName) {
	var withParentExtraction func(cnd project_invitation.FilterCondition) project_invitation.FilterCondition
	var resultParent *project_invitation.ParentName
	var resultFilter *project_invitation.Filter
	withParentExtraction = func(cnd project_invitation.FilterCondition) project_invitation.FilterCondition {
		switch tCnd := cnd.(type) {
		case *project_invitation.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]project_invitation.FilterCondition, 0)
				for _, subCnd := range tCnd.Conditions {
					if subCndNoParent := withParentExtraction(subCnd); subCndNoParent != nil {
						withoutParentCnds = append(withoutParentCnds, subCndNoParent)
					}
				}
				if len(withoutParentCnds) == 0 {
					return nil
				}
				if len(withoutParentCnds) == 1 {
					return withoutParentCnds[0]
				}
				return project_invitation.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *project_invitation.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*project_invitation.Name)
				if nameValue != nil && nameValue.ParentName.IsSpecified() {
					resultParent = &nameValue.ParentName
					if nameValue.IsFullyQualified() {
						return tCnd
					}
					return nil
				}
			}
			return tCnd
		default:
			return tCnd
		}
	}
	cndWithoutParent := withParentExtraction(fullFilter.GetCondition())
	if cndWithoutParent != nil {
		resultFilter = &project_invitation.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(project_invitation.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return project_invitation.AsAnyCastAccess(NewApiProjectInvitationAccess(project_invitation_client.NewProjectInvitationServiceClient(cc)))
	})
}
