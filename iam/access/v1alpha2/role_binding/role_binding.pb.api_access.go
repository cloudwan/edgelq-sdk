// Code generated by protoc-gen-goten-access
// Resource: RoleBinding
// DO NOT EDIT!!!

package role_binding_access

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

	role_binding_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha2/role_binding"
	role_binding "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/role_binding"
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

type apiRoleBindingAccess struct {
	client role_binding_client.RoleBindingServiceClient
}

func NewApiRoleBindingAccess(client role_binding_client.RoleBindingServiceClient) role_binding.RoleBindingAccess {
	return &apiRoleBindingAccess{client: client}
}

func (a *apiRoleBindingAccess) GetRoleBinding(ctx context.Context, query *role_binding.GetQuery) (*role_binding.RoleBinding, error) {
	if !query.Reference.IsFullyQualified() {
		return nil, status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &role_binding_client.GetRoleBindingRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetRoleBinding(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiRoleBindingAccess) BatchGetRoleBindings(ctx context.Context, refs []*role_binding.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	asNames := make([]*role_binding.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &role_binding_client.BatchGetRoleBindingsRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(role_binding.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*role_binding.RoleBinding_FieldMask)
	}
	resp, err := a.client.BatchGetRoleBindings(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[role_binding.Name]*role_binding.RoleBinding, len(refs))
	for _, resolvedRes := range resp.GetRoleBindings() {
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

func (a *apiRoleBindingAccess) QueryRoleBindings(ctx context.Context, query *role_binding.ListQuery) (*role_binding.QueryResultSnapshot, error) {
	request := &role_binding_client.ListRoleBindingsRequest{
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
	resp, err := a.client.ListRoleBindings(ctx, request)
	if err != nil {
		return nil, err
	}
	return &role_binding.QueryResultSnapshot{
		RoleBindings:      resp.RoleBindings,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiRoleBindingAccess) WatchRoleBinding(ctx context.Context, query *role_binding.GetQuery, observerCb func(*role_binding.RoleBindingChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &role_binding_client.WatchRoleBindingRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchRoleBinding(ctx, request)
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

func (a *apiRoleBindingAccess) WatchRoleBindings(ctx context.Context, query *role_binding.WatchQuery, observerCb func(*role_binding.QueryResultChange) error) error {
	request := &role_binding_client.WatchRoleBindingsRequest{
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
	changesStream, initErr := a.client.WatchRoleBindings(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &role_binding.QueryResultChange{
			Changes:      respChange.RoleBindingChanges,
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

func (a *apiRoleBindingAccess) SaveRoleBinding(ctx context.Context, res *role_binding.RoleBinding, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetRoleBinding(ctx, &role_binding.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}
	var resp *role_binding.RoleBinding
	var err error
	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &role_binding_client.UpdateRoleBindingRequest{
			RoleBinding: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*role_binding.RoleBinding_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &role_binding_client.UpdateRoleBindingRequest_CAS{
				ConditionalState: conditionalState.(*role_binding.RoleBinding),
				FieldMask:        mask.(*role_binding.RoleBinding_FieldMask),
			}
		}
		resp, err = a.client.UpdateRoleBinding(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &role_binding_client.CreateRoleBindingRequest{
			RoleBinding: res,
		}
		resp, err = a.client.CreateRoleBinding(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiRoleBindingAccess) DeleteRoleBinding(ctx context.Context, ref *role_binding.Reference, opts ...gotenresource.DeleteOption) error {
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &role_binding_client.DeleteRoleBindingRequest{
		Name: &ref.Name,
	}
	_, err := a.client.DeleteRoleBinding(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *role_binding.Filter) (*role_binding.Filter, *role_binding.ParentName) {
	var withParentExtraction func(cnd role_binding.FilterCondition) role_binding.FilterCondition
	var resultParent *role_binding.ParentName
	var resultFilter *role_binding.Filter
	withParentExtraction = func(cnd role_binding.FilterCondition) role_binding.FilterCondition {
		switch tCnd := cnd.(type) {
		case *role_binding.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]role_binding.FilterCondition, 0)
				for _, subCnd := range tCnd.Conditions {
					if subCndNoParent := withParentExtraction(subCnd); subCndNoParent != nil {
						withoutParentCnds = append(withoutParentCnds, subCndNoParent)
					}
				}
				if len(withoutParentCnds) == 0 {
					return nil
				}
				return role_binding.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *role_binding.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*role_binding.Name)
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
		resultFilter = &role_binding.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(role_binding.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return role_binding.AsAnyCastAccess(NewApiRoleBindingAccess(role_binding_client.NewRoleBindingServiceClient(cc)))
	})
}
