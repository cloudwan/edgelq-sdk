// Code generated by protoc-gen-goten-access
// Resource: MetricDescriptor
// DO NOT EDIT!!!

package metric_descriptor_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	metric_descriptor_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v3/metric_descriptor"
	metric_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/metric_descriptor"
)

var (
	_ = context.Context(nil)
	_ = fmt.GoStringer(nil)

	_ = codes.NotFound
	_ = status.Status{}

	_ = watch_type.WatchType_STATEFUL
	_ = gotenresource.ListQuery(nil)
)

type apiMetricDescriptorAccess struct {
	client metric_descriptor_client.MetricDescriptorServiceClient
}

func NewApiMetricDescriptorAccess(client metric_descriptor_client.MetricDescriptorServiceClient) metric_descriptor.MetricDescriptorAccess {
	return &apiMetricDescriptorAccess{client: client}
}

func (a *apiMetricDescriptorAccess) GetMetricDescriptor(ctx context.Context, query *metric_descriptor.GetQuery) (*metric_descriptor.MetricDescriptor, error) {
	request := &metric_descriptor_client.GetMetricDescriptorRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	return a.client.GetMetricDescriptor(ctx, request)
}

func (a *apiMetricDescriptorAccess) BatchGetMetricDescriptors(ctx context.Context, refs []*metric_descriptor.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &metric_descriptor_client.BatchGetMetricDescriptorsRequest{
		Names: refs,
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
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListMetricDescriptors(ctx, request)
	if err != nil {
		return nil, err
	}
	return &metric_descriptor.QueryResultSnapshot{
		MetricDescriptors: resp.MetricDescriptors,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
	}, nil
}

func (a *apiMetricDescriptorAccess) WatchMetricDescriptor(ctx context.Context, query *metric_descriptor.GetQuery, observerCb func(*metric_descriptor.MetricDescriptorChange) error) error {
	request := &metric_descriptor_client.WatchMetricDescriptorRequest{
		Name:      query.Reference,
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

	if previousRes == nil {
		var err error
		previousRes, err = a.GetMetricDescriptor(ctx, &metric_descriptor.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if previousRes != nil {
		updateRequest := &metric_descriptor_client.UpdateMetricDescriptorRequest{
			MetricDescriptor: res,
		}
		_, err := a.client.UpdateMetricDescriptor(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &metric_descriptor_client.CreateMetricDescriptorRequest{
			MetricDescriptor: res,
		}
		_, err := a.client.CreateMetricDescriptor(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiMetricDescriptorAccess) DeleteMetricDescriptor(ctx context.Context, ref *metric_descriptor.Reference, opts ...gotenresource.DeleteOption) error {
	request := &metric_descriptor_client.DeleteMetricDescriptorRequest{
		Name: ref,
	}
	_, err := a.client.DeleteMetricDescriptor(ctx, request)
	return err
}
