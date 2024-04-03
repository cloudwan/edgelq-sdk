// Code generated by protoc-gen-goten-access
// Resource: MetricDescriptor
// DO NOT EDIT!!!

package metric_descriptor_access

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

	metric_descriptor_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v4/metric_descriptor"
	metric_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/metric_descriptor"
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

type apiMetricDescriptorAccess struct {
	client metric_descriptor_client.MetricDescriptorServiceClient
}

func NewApiMetricDescriptorAccess(client metric_descriptor_client.MetricDescriptorServiceClient) metric_descriptor.MetricDescriptorAccess {
	return &apiMetricDescriptorAccess{client: client}
}

func (a *apiMetricDescriptorAccess) GetMetricDescriptor(ctx context.Context, query *metric_descriptor.GetQuery) (*metric_descriptor.MetricDescriptor, error) {
	if !query.Reference.IsFullyQualified() {
		return nil, status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &metric_descriptor_client.GetMetricDescriptorRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetMetricDescriptor(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiMetricDescriptorAccess) BatchGetMetricDescriptors(ctx context.Context, refs []*metric_descriptor.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	asNames := make([]*metric_descriptor.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &metric_descriptor_client.BatchGetMetricDescriptorsRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(metric_descriptor.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*metric_descriptor.MetricDescriptor_FieldMask)
	}
	resp, err := a.client.BatchGetMetricDescriptors(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[metric_descriptor.Name]*metric_descriptor.MetricDescriptor, len(refs))
	for _, resolvedRes := range resp.GetMetricDescriptors() {
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

func (a *apiMetricDescriptorAccess) QueryMetricDescriptors(ctx context.Context, query *metric_descriptor.ListQuery) (*metric_descriptor.QueryResultSnapshot, error) {
	request := &metric_descriptor_client.ListMetricDescriptorsRequest{
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
	resp, err := a.client.ListMetricDescriptors(ctx, request)
	if err != nil {
		return nil, err
	}
	return &metric_descriptor.QueryResultSnapshot{
		MetricDescriptors: resp.MetricDescriptors,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiMetricDescriptorAccess) WatchMetricDescriptor(ctx context.Context, query *metric_descriptor.GetQuery, observerCb func(*metric_descriptor.MetricDescriptorChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &metric_descriptor_client.WatchMetricDescriptorRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchMetricDescriptor(ctx, request)
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

func (a *apiMetricDescriptorAccess) WatchMetricDescriptors(ctx context.Context, query *metric_descriptor.WatchQuery, observerCb func(*metric_descriptor.QueryResultChange) error) error {
	request := &metric_descriptor_client.WatchMetricDescriptorsRequest{
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
	changesStream, initErr := a.client.WatchMetricDescriptors(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &metric_descriptor.QueryResultChange{
			Changes:      respChange.MetricDescriptorChanges,
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

func (a *apiMetricDescriptorAccess) SaveMetricDescriptor(ctx context.Context, res *metric_descriptor.MetricDescriptor, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetMetricDescriptor(ctx, &metric_descriptor.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}
	var resp *metric_descriptor.MetricDescriptor
	var err error
	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &metric_descriptor_client.UpdateMetricDescriptorRequest{
			MetricDescriptor: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*metric_descriptor.MetricDescriptor_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &metric_descriptor_client.UpdateMetricDescriptorRequest_CAS{
				ConditionalState: conditionalState.(*metric_descriptor.MetricDescriptor),
				FieldMask:        mask.(*metric_descriptor.MetricDescriptor_FieldMask),
			}
		}
		resp, err = a.client.UpdateMetricDescriptor(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &metric_descriptor_client.CreateMetricDescriptorRequest{
			MetricDescriptor: res,
		}
		resp, err = a.client.CreateMetricDescriptor(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiMetricDescriptorAccess) DeleteMetricDescriptor(ctx context.Context, ref *metric_descriptor.Reference, opts ...gotenresource.DeleteOption) error {
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &metric_descriptor_client.DeleteMetricDescriptorRequest{
		Name: &ref.Name,
	}
	_, err := a.client.DeleteMetricDescriptor(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *metric_descriptor.Filter) (*metric_descriptor.Filter, *metric_descriptor.ParentName) {
	var withParentExtraction func(cnd metric_descriptor.FilterCondition) metric_descriptor.FilterCondition
	var resultParent *metric_descriptor.ParentName
	var resultFilter *metric_descriptor.Filter
	withParentExtraction = func(cnd metric_descriptor.FilterCondition) metric_descriptor.FilterCondition {
		switch tCnd := cnd.(type) {
		case *metric_descriptor.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]metric_descriptor.FilterCondition, 0)
				for _, subCnd := range tCnd.Conditions {
					if subCndNoParent := withParentExtraction(subCnd); subCndNoParent != nil {
						withoutParentCnds = append(withoutParentCnds, subCndNoParent)
					}
				}
				if len(withoutParentCnds) == 0 {
					return nil
				}
				return metric_descriptor.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *metric_descriptor.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*metric_descriptor.Name)
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
		resultFilter = &metric_descriptor.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(metric_descriptor.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return metric_descriptor.AsAnyCastAccess(NewApiMetricDescriptorAccess(metric_descriptor_client.NewMetricDescriptorServiceClient(cc)))
	})
}