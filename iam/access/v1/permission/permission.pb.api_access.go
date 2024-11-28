// Code generated by protoc-gen-goten-access
// Resource: Permission
// DO NOT EDIT!!!

package permission_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
	gotenfilter "github.com/cloudwan/goten-sdk/runtime/resource/filter"
	"github.com/cloudwan/goten-sdk/types/watch_type"

	permission_client "github.com/cloudwan/edgelq-sdk/iam/client/v1/permission"
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1/permission"
)

var (
	_ = new(context.Context)
	_ = new(fmt.GoStringer)

	_ = metadata.MD{}
	_ = new(grpc.ClientConnInterface)
	_ = codes.NotFound
	_ = status.Status{}

	_ = new(gotenaccess.Watcher)
	_ = watch_type.WatchType_STATEFUL
	_ = new(gotenresource.ListQuery)
	_ = gotenfilter.Eq
)

type apiPermissionAccess struct {
	client permission_client.PermissionServiceClient
}

func NewApiPermissionAccess(client permission_client.PermissionServiceClient) permission.PermissionAccess {
	return &apiPermissionAccess{client: client}
}

func (a *apiPermissionAccess) GetPermission(ctx context.Context, query *permission.GetQuery, opts ...gotenresource.GetOption) (*permission.Permission, error) {
	getOpts := gotenresource.MakeGetOptions(opts)
	callHeaders := metadata.MD{}
	if getOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	if !query.Reference.IsFullyQualified() {
		return nil, status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &permission_client.GetPermissionRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetPermission(ctx, request, callOpts...)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiPermissionAccess) BatchGetPermissions(ctx context.Context, refs []*permission.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	callHeaders := metadata.MD{}
	if batchGetOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	asNames := make([]*permission.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &permission_client.BatchGetPermissionsRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(permission.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*permission.Permission_FieldMask)
	}
	resp, err := a.client.BatchGetPermissions(ctx, request, callOpts...)
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

func (a *apiPermissionAccess) QueryPermissions(ctx context.Context, query *permission.ListQuery, opts ...gotenresource.QueryOption) (*permission.QueryResultSnapshot, error) {
	qOpts := gotenresource.MakeQueryOptions(opts)
	callHeaders := metadata.MD{}
	if qOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	request := &permission_client.ListPermissionsRequest{
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
	resp, err := a.client.ListPermissions(ctx, request)
	if err != nil {
		return nil, err
	}
	return &permission.QueryResultSnapshot{
		Permissions:       resp.Permissions,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiPermissionAccess) WatchPermission(ctx context.Context, query *permission.GetQuery, observerCb func(*permission.PermissionChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &permission_client.WatchPermissionRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

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
		StartingTime: query.StartingTime,
	}
	if query.Pager != nil {
		request.OrderBy = query.Pager.OrderBy
		request.PageSize = int32(query.Pager.Limit)
		request.PageToken = query.Pager.Cursor
	}
	if query.Filter != nil && query.Filter.GetCondition() != nil {
		request.Filter, request.Parent = getParentAndFilter(query.Filter)
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

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
	var resp *permission.Permission
	var err error
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
		resp, err = a.client.UpdatePermission(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &permission_client.CreatePermissionRequest{
			Permission: res,
		}
		resp, err = a.client.CreatePermission(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiPermissionAccess) DeletePermission(ctx context.Context, ref *permission.Reference, opts ...gotenresource.DeleteOption) error {
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &permission_client.DeletePermissionRequest{
		Name: &ref.Name,
	}
	_, err := a.client.DeletePermission(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *permission.Filter) (*permission.Filter, *permission.ParentName) {
	var withParentExtraction func(cnd permission.FilterCondition) permission.FilterCondition
	var resultParent *permission.ParentName
	var resultFilter *permission.Filter
	withParentExtraction = func(cnd permission.FilterCondition) permission.FilterCondition {
		switch tCnd := cnd.(type) {
		case *permission.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]permission.FilterCondition, 0)
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
				return permission.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *permission.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*permission.Name)
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
		resultFilter = &permission.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(permission.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return permission.AsAnyCastAccess(NewApiPermissionAccess(permission_client.NewPermissionServiceClient(cc)))
	})
}
