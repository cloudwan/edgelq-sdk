// Code generated by protoc-gen-goten-access
// Resource: Group
// DO NOT EDIT!!!

package group_access

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

	group_client "github.com/cloudwan/edgelq-sdk/iam/client/v1/group"
	group "github.com/cloudwan/edgelq-sdk/iam/resources/v1/group"
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

type apiGroupAccess struct {
	client group_client.GroupServiceClient
}

func NewApiGroupAccess(client group_client.GroupServiceClient) group.GroupAccess {
	return &apiGroupAccess{client: client}
}

func (a *apiGroupAccess) GetGroup(ctx context.Context, query *group.GetQuery) (*group.Group, error) {
	if !query.Reference.IsFullyQualified() {
		return nil, status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &group_client.GetGroupRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetGroup(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiGroupAccess) BatchGetGroups(ctx context.Context, refs []*group.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	asNames := make([]*group.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &group_client.BatchGetGroupsRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(group.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*group.Group_FieldMask)
	}
	resp, err := a.client.BatchGetGroups(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[group.Name]*group.Group, len(refs))
	for _, resolvedRes := range resp.GetGroups() {
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

func (a *apiGroupAccess) QueryGroups(ctx context.Context, query *group.ListQuery) (*group.QueryResultSnapshot, error) {
	request := &group_client.ListGroupsRequest{
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
	resp, err := a.client.ListGroups(ctx, request)
	if err != nil {
		return nil, err
	}
	return &group.QueryResultSnapshot{
		Groups:            resp.Groups,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiGroupAccess) WatchGroup(ctx context.Context, query *group.GetQuery, observerCb func(*group.GroupChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &group_client.WatchGroupRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchGroup(ctx, request)
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

func (a *apiGroupAccess) WatchGroups(ctx context.Context, query *group.WatchQuery, observerCb func(*group.QueryResultChange) error) error {
	request := &group_client.WatchGroupsRequest{
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
	changesStream, initErr := a.client.WatchGroups(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &group.QueryResultChange{
			Changes:      respChange.GroupChanges,
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

func (a *apiGroupAccess) SaveGroup(ctx context.Context, res *group.Group, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetGroup(ctx, &group.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}
	var resp *group.Group
	var err error
	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &group_client.UpdateGroupRequest{
			Group: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*group.Group_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &group_client.UpdateGroupRequest_CAS{
				ConditionalState: conditionalState.(*group.Group),
				FieldMask:        mask.(*group.Group_FieldMask),
			}
		}
		resp, err = a.client.UpdateGroup(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &group_client.CreateGroupRequest{
			Group: res,
		}
		resp, err = a.client.CreateGroup(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiGroupAccess) DeleteGroup(ctx context.Context, ref *group.Reference, opts ...gotenresource.DeleteOption) error {
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &group_client.DeleteGroupRequest{
		Name: &ref.Name,
	}
	_, err := a.client.DeleteGroup(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *group.Filter) (*group.Filter, *group.ParentName) {
	var withParentExtraction func(cnd group.FilterCondition) group.FilterCondition
	var resultParent *group.ParentName
	var resultFilter *group.Filter
	withParentExtraction = func(cnd group.FilterCondition) group.FilterCondition {
		switch tCnd := cnd.(type) {
		case *group.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]group.FilterCondition, 0)
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
				return group.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *group.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*group.Name)
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
		resultFilter = &group.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(group.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return group.AsAnyCastAccess(NewApiGroupAccess(group_client.NewGroupServiceClient(cc)))
	})
}
