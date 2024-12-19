// Code generated by protoc-gen-goten-access
// Resource: TimeSeriesForwarderSink
// DO NOT EDIT!!!

package time_series_forwarder_sink_access

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
	gotenfilter "github.com/cloudwan/goten-sdk/runtime/resource/filter"
	"github.com/cloudwan/goten-sdk/types/watch_type"

	time_series_forwarder_sink_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v4/time_series_forwarder_sink"
	time_series_forwarder_sink "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/time_series_forwarder_sink"
)

var (
	_ = new(context.Context)

	_ = metadata.MD{}
	_ = new(grpc.ClientConnInterface)
	_ = codes.NotFound
	_ = status.Status{}

	_ = new(gotenaccess.Watcher)
	_ = watch_type.WatchType_STATEFUL
	_ = new(gotenresource.ListQuery)
	_ = gotenfilter.Eq
)

type apiTimeSeriesForwarderSinkAccess struct {
	client time_series_forwarder_sink_client.TimeSeriesForwarderSinkServiceClient
}

func NewApiTimeSeriesForwarderSinkAccess(client time_series_forwarder_sink_client.TimeSeriesForwarderSinkServiceClient) time_series_forwarder_sink.TimeSeriesForwarderSinkAccess {
	return &apiTimeSeriesForwarderSinkAccess{client: client}
}

func (a *apiTimeSeriesForwarderSinkAccess) GetTimeSeriesForwarderSink(ctx context.Context, query *time_series_forwarder_sink.GetQuery, opts ...gotenresource.GetOption) (*time_series_forwarder_sink.TimeSeriesForwarderSink, error) {
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
	request := &time_series_forwarder_sink_client.GetTimeSeriesForwarderSinkRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetTimeSeriesForwarderSink(ctx, request, callOpts...)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiTimeSeriesForwarderSinkAccess) BatchGetTimeSeriesForwarderSinks(ctx context.Context, refs []*time_series_forwarder_sink.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	callHeaders := metadata.MD{}
	if batchGetOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	asNames := make([]*time_series_forwarder_sink.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &time_series_forwarder_sink_client.BatchGetTimeSeriesForwarderSinksRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(time_series_forwarder_sink.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*time_series_forwarder_sink.TimeSeriesForwarderSink_FieldMask)
	}
	resp, err := a.client.BatchGetTimeSeriesForwarderSinks(ctx, request, callOpts...)
	if err != nil {
		return err
	}
	resultMap := make(map[time_series_forwarder_sink.Name]*time_series_forwarder_sink.TimeSeriesForwarderSink, len(refs))
	for _, resolvedRes := range resp.GetTimeSeriesForwarderSinks() {
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

func (a *apiTimeSeriesForwarderSinkAccess) QueryTimeSeriesForwarderSinks(ctx context.Context, query *time_series_forwarder_sink.ListQuery, opts ...gotenresource.QueryOption) (*time_series_forwarder_sink.QueryResultSnapshot, error) {
	qOpts := gotenresource.MakeQueryOptions(opts)
	callHeaders := metadata.MD{}
	if qOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	request := &time_series_forwarder_sink_client.ListTimeSeriesForwarderSinksRequest{
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
	resp, err := a.client.ListTimeSeriesForwarderSinks(ctx, request)
	if err != nil {
		return nil, err
	}
	return &time_series_forwarder_sink.QueryResultSnapshot{
		TimeSeriesForwarderSinks: resp.TimeSeriesForwarderSinks,
		NextPageCursor:           resp.NextPageToken,
		PrevPageCursor:           resp.PrevPageToken,
		TotalResultsCount:        resp.TotalResultsCount,
		CurrentOffset:            resp.CurrentOffset,
	}, nil
}

func (a *apiTimeSeriesForwarderSinkAccess) WatchTimeSeriesForwarderSink(ctx context.Context, query *time_series_forwarder_sink.GetQuery, observerCb func(*time_series_forwarder_sink.TimeSeriesForwarderSinkChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &time_series_forwarder_sink_client.WatchTimeSeriesForwarderSinkRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	changesStream, initErr := a.client.WatchTimeSeriesForwarderSink(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		resp, err := changesStream.Recv()
		if err != nil {
			return status.Errorf(status.Code(err), "watch recv error: %s", err)
		}
		change := resp.GetChange()
		if err := observerCb(change); err != nil {
			return err
		}
	}
}

func (a *apiTimeSeriesForwarderSinkAccess) WatchTimeSeriesForwarderSinks(ctx context.Context, query *time_series_forwarder_sink.WatchQuery, observerCb func(*time_series_forwarder_sink.QueryResultChange) error) error {
	request := &time_series_forwarder_sink_client.WatchTimeSeriesForwarderSinksRequest{
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

	changesStream, initErr := a.client.WatchTimeSeriesForwarderSinks(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return status.Errorf(status.Code(err), "watch recv error: %s", err)
		}
		changesWithPaging := &time_series_forwarder_sink.QueryResultChange{
			Changes:      respChange.TimeSeriesForwarderSinkChanges,
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

func (a *apiTimeSeriesForwarderSinkAccess) SaveTimeSeriesForwarderSink(ctx context.Context, res *time_series_forwarder_sink.TimeSeriesForwarderSink, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	var resp *time_series_forwarder_sink.TimeSeriesForwarderSink
	var err error
	if !saveOpts.OnlyCreate() {
		updateRequest := &time_series_forwarder_sink_client.UpdateTimeSeriesForwarderSinkRequest{
			TimeSeriesForwarderSink: res,
			AllowMissing:            !saveOpts.OnlyUpdate(),
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*time_series_forwarder_sink.TimeSeriesForwarderSink_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &time_series_forwarder_sink_client.UpdateTimeSeriesForwarderSinkRequest_CAS{
				ConditionalState: conditionalState.(*time_series_forwarder_sink.TimeSeriesForwarderSink),
				FieldMask:        mask.(*time_series_forwarder_sink.TimeSeriesForwarderSink_FieldMask),
			}
		}
		resp, err = a.client.UpdateTimeSeriesForwarderSink(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &time_series_forwarder_sink_client.CreateTimeSeriesForwarderSinkRequest{
			TimeSeriesForwarderSink: res,
		}
		resp, err = a.client.CreateTimeSeriesForwarderSink(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiTimeSeriesForwarderSinkAccess) DeleteTimeSeriesForwarderSink(ctx context.Context, ref *time_series_forwarder_sink.Reference, _ ...gotenresource.DeleteOption) error {
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &time_series_forwarder_sink_client.DeleteTimeSeriesForwarderSinkRequest{
		Name: &ref.Name,
	}
	_, err := a.client.DeleteTimeSeriesForwarderSink(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *time_series_forwarder_sink.Filter) (*time_series_forwarder_sink.Filter, *time_series_forwarder_sink.ParentName) {
	var withParentExtraction func(cnd time_series_forwarder_sink.FilterCondition) time_series_forwarder_sink.FilterCondition
	var resultParent *time_series_forwarder_sink.ParentName
	var resultFilter *time_series_forwarder_sink.Filter
	withParentExtraction = func(cnd time_series_forwarder_sink.FilterCondition) time_series_forwarder_sink.FilterCondition {
		switch tCnd := cnd.(type) {
		case *time_series_forwarder_sink.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]time_series_forwarder_sink.FilterCondition, 0)
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
				return time_series_forwarder_sink.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *time_series_forwarder_sink.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*time_series_forwarder_sink.Name)
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
		resultFilter = &time_series_forwarder_sink.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(time_series_forwarder_sink.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return time_series_forwarder_sink.AsAnyCastAccess(NewApiTimeSeriesForwarderSinkAccess(time_series_forwarder_sink_client.NewTimeSeriesForwarderSinkServiceClient(cc)))
	})
}