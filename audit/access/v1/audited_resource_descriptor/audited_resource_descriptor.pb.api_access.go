// Code generated by protoc-gen-goten-access
// Resource: AuditedResourceDescriptor
// DO NOT EDIT!!!

package audited_resource_descriptor_access

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

	audited_resource_descriptor_client "github.com/cloudwan/edgelq-sdk/audit/client/v1/audited_resource_descriptor"
	audited_resource_descriptor "github.com/cloudwan/edgelq-sdk/audit/resources/v1/audited_resource_descriptor"
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

type apiAuditedResourceDescriptorAccess struct {
	client audited_resource_descriptor_client.AuditedResourceDescriptorServiceClient
}

func NewApiAuditedResourceDescriptorAccess(client audited_resource_descriptor_client.AuditedResourceDescriptorServiceClient) audited_resource_descriptor.AuditedResourceDescriptorAccess {
	return &apiAuditedResourceDescriptorAccess{client: client}
}

func (a *apiAuditedResourceDescriptorAccess) GetAuditedResourceDescriptor(ctx context.Context, query *audited_resource_descriptor.GetQuery) (*audited_resource_descriptor.AuditedResourceDescriptor, error) {
	if !query.Reference.IsFullyQualified() {
		return nil, status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &audited_resource_descriptor_client.GetAuditedResourceDescriptorRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetAuditedResourceDescriptor(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiAuditedResourceDescriptorAccess) BatchGetAuditedResourceDescriptors(ctx context.Context, refs []*audited_resource_descriptor.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	asNames := make([]*audited_resource_descriptor.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &audited_resource_descriptor_client.BatchGetAuditedResourceDescriptorsRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(audited_resource_descriptor.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*audited_resource_descriptor.AuditedResourceDescriptor_FieldMask)
	}
	resp, err := a.client.BatchGetAuditedResourceDescriptors(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[audited_resource_descriptor.Name]*audited_resource_descriptor.AuditedResourceDescriptor, len(refs))
	for _, resolvedRes := range resp.GetAuditedResourceDescriptors() {
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

func (a *apiAuditedResourceDescriptorAccess) QueryAuditedResourceDescriptors(ctx context.Context, query *audited_resource_descriptor.ListQuery) (*audited_resource_descriptor.QueryResultSnapshot, error) {
	request := &audited_resource_descriptor_client.ListAuditedResourceDescriptorsRequest{
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
	resp, err := a.client.ListAuditedResourceDescriptors(ctx, request)
	if err != nil {
		return nil, err
	}
	return &audited_resource_descriptor.QueryResultSnapshot{
		AuditedResourceDescriptors: resp.AuditedResourceDescriptors,
		NextPageCursor:             resp.NextPageToken,
		PrevPageCursor:             resp.PrevPageToken,
		TotalResultsCount:          resp.TotalResultsCount,
		CurrentOffset:              resp.CurrentOffset,
	}, nil
}

func (a *apiAuditedResourceDescriptorAccess) WatchAuditedResourceDescriptor(ctx context.Context, query *audited_resource_descriptor.GetQuery, observerCb func(*audited_resource_descriptor.AuditedResourceDescriptorChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &audited_resource_descriptor_client.WatchAuditedResourceDescriptorRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchAuditedResourceDescriptor(ctx, request)
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

func (a *apiAuditedResourceDescriptorAccess) WatchAuditedResourceDescriptors(ctx context.Context, query *audited_resource_descriptor.WatchQuery, observerCb func(*audited_resource_descriptor.QueryResultChange) error) error {
	request := &audited_resource_descriptor_client.WatchAuditedResourceDescriptorsRequest{
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
	changesStream, initErr := a.client.WatchAuditedResourceDescriptors(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &audited_resource_descriptor.QueryResultChange{
			Changes:      respChange.AuditedResourceDescriptorChanges,
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

func (a *apiAuditedResourceDescriptorAccess) SaveAuditedResourceDescriptor(ctx context.Context, res *audited_resource_descriptor.AuditedResourceDescriptor, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetAuditedResourceDescriptor(ctx, &audited_resource_descriptor.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}
	var resp *audited_resource_descriptor.AuditedResourceDescriptor
	var err error
	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &audited_resource_descriptor_client.UpdateAuditedResourceDescriptorRequest{
			AuditedResourceDescriptor: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*audited_resource_descriptor.AuditedResourceDescriptor_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &audited_resource_descriptor_client.UpdateAuditedResourceDescriptorRequest_CAS{
				ConditionalState: conditionalState.(*audited_resource_descriptor.AuditedResourceDescriptor),
				FieldMask:        mask.(*audited_resource_descriptor.AuditedResourceDescriptor_FieldMask),
			}
		}
		resp, err = a.client.UpdateAuditedResourceDescriptor(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &audited_resource_descriptor_client.CreateAuditedResourceDescriptorRequest{
			AuditedResourceDescriptor: res,
		}
		resp, err = a.client.CreateAuditedResourceDescriptor(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiAuditedResourceDescriptorAccess) DeleteAuditedResourceDescriptor(ctx context.Context, ref *audited_resource_descriptor.Reference, opts ...gotenresource.DeleteOption) error {
	return fmt.Errorf("Delete operation on AuditedResourceDescriptor is prohibited")
}
func getParentAndFilter(fullFilter *audited_resource_descriptor.Filter) (*audited_resource_descriptor.Filter, *audited_resource_descriptor.ParentName) {
	var withParentExtraction func(cnd audited_resource_descriptor.FilterCondition) audited_resource_descriptor.FilterCondition
	var resultParent *audited_resource_descriptor.ParentName
	var resultFilter *audited_resource_descriptor.Filter
	withParentExtraction = func(cnd audited_resource_descriptor.FilterCondition) audited_resource_descriptor.FilterCondition {
		switch tCnd := cnd.(type) {
		case *audited_resource_descriptor.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]audited_resource_descriptor.FilterCondition, 0)
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
				return audited_resource_descriptor.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *audited_resource_descriptor.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*audited_resource_descriptor.Name)
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
		resultFilter = &audited_resource_descriptor.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(audited_resource_descriptor.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return audited_resource_descriptor.AsAnyCastAccess(NewApiAuditedResourceDescriptorAccess(audited_resource_descriptor_client.NewAuditedResourceDescriptorServiceClient(cc)))
	})
}
