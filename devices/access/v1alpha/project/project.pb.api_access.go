// Code generated by protoc-gen-goten-access
// Resource: Project
// DO NOT EDIT!!!

package project_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	project_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha/project"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha/project"
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

type apiProjectAccess struct {
	client project_client.ProjectServiceClient
}

func NewApiProjectAccess(client project_client.ProjectServiceClient) project.ProjectAccess {
	return &apiProjectAccess{client: client}
}

func (a *apiProjectAccess) GetProject(ctx context.Context, query *project.GetQuery) (*project.Project, error) {
	request := &project_client.GetProjectRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetProject(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiProjectAccess) BatchGetProjects(ctx context.Context, refs []*project.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &project_client.BatchGetProjectsRequest{
		Names: refs,
	}
	fieldMask := batchGetOpts.GetFieldMask(project.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*project.Project_FieldMask)
	}
	resp, err := a.client.BatchGetProjects(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[project.Name]*project.Project, len(refs))
	for _, resolvedRes := range resp.GetProjects() {
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

func (a *apiProjectAccess) QueryProjects(ctx context.Context, query *project.ListQuery) (*project.QueryResultSnapshot, error) {
	request := &project_client.ListProjectsRequest{
		Filter:            query.Filter,
		FieldMask:         query.Mask,
		IncludePagingInfo: query.WithPagingInfo,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListProjects(ctx, request)
	if err != nil {
		return nil, err
	}
	return &project.QueryResultSnapshot{
		Projects:          resp.Projects,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiProjectAccess) WatchProject(ctx context.Context, query *project.GetQuery, observerCb func(*project.ProjectChange) error) error {
	request := &project_client.WatchProjectRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchProject(ctx, request)
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

func (a *apiProjectAccess) WatchProjects(ctx context.Context, query *project.WatchQuery, observerCb func(*project.QueryResultChange) error) error {
	request := &project_client.WatchProjectsRequest{
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
	changesStream, initErr := a.client.WatchProjects(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &project.QueryResultChange{
			Changes:      respChange.ProjectChanges,
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

func (a *apiProjectAccess) SaveProject(ctx context.Context, res *project.Project, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetProject(ctx, &project.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &project_client.UpdateProjectRequest{
			Project: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*project.Project_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &project_client.UpdateProjectRequest_CAS{
				ConditionalState: conditionalState.(*project.Project),
				FieldMask:        mask.(*project.Project_FieldMask),
			}
		}
		_, err := a.client.UpdateProject(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &project_client.CreateProjectRequest{
			Project: res,
		}
		_, err := a.client.CreateProject(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiProjectAccess) DeleteProject(ctx context.Context, ref *project.Reference, opts ...gotenresource.DeleteOption) error {
	request := &project_client.DeleteProjectRequest{
		Name: ref,
	}
	_, err := a.client.DeleteProject(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(project.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return project.AsAnyCastAccess(NewApiProjectAccess(project_client.NewProjectServiceClient(cc)))
	})
}
