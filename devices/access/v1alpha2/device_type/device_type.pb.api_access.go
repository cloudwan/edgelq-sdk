// Code generated by protoc-gen-goten-access
// Resource: DeviceType
// DO NOT EDIT!!!

package device_type_access

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

	device_type_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha2/device_type"
	device_type "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/device_type"
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

type apiDeviceTypeAccess struct {
	client device_type_client.DeviceTypeServiceClient
}

func NewApiDeviceTypeAccess(client device_type_client.DeviceTypeServiceClient) device_type.DeviceTypeAccess {
	return &apiDeviceTypeAccess{client: client}
}

func (a *apiDeviceTypeAccess) GetDeviceType(ctx context.Context, query *device_type.GetQuery) (*device_type.DeviceType, error) {
	if !query.Reference.IsFullyQualified() {
		return nil, status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &device_type_client.GetDeviceTypeRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetDeviceType(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiDeviceTypeAccess) BatchGetDeviceTypes(ctx context.Context, refs []*device_type.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	asNames := make([]*device_type.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &device_type_client.BatchGetDeviceTypesRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(device_type.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*device_type.DeviceType_FieldMask)
	}
	resp, err := a.client.BatchGetDeviceTypes(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[device_type.Name]*device_type.DeviceType, len(refs))
	for _, resolvedRes := range resp.GetDeviceTypes() {
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

func (a *apiDeviceTypeAccess) QueryDeviceTypes(ctx context.Context, query *device_type.ListQuery) (*device_type.QueryResultSnapshot, error) {
	request := &device_type_client.ListDeviceTypesRequest{
		Filter:            query.Filter,
		FieldMask:         query.Mask,
		IncludePagingInfo: query.WithPagingInfo,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListDeviceTypes(ctx, request)
	if err != nil {
		return nil, err
	}
	return &device_type.QueryResultSnapshot{
		DeviceTypes:       resp.DeviceTypes,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiDeviceTypeAccess) WatchDeviceType(ctx context.Context, query *device_type.GetQuery, observerCb func(*device_type.DeviceTypeChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &device_type_client.WatchDeviceTypeRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchDeviceType(ctx, request)
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

func (a *apiDeviceTypeAccess) WatchDeviceTypes(ctx context.Context, query *device_type.WatchQuery, observerCb func(*device_type.QueryResultChange) error) error {
	request := &device_type_client.WatchDeviceTypesRequest{
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
	changesStream, initErr := a.client.WatchDeviceTypes(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &device_type.QueryResultChange{
			Changes:      respChange.DeviceTypeChanges,
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

func (a *apiDeviceTypeAccess) SaveDeviceType(ctx context.Context, res *device_type.DeviceType, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetDeviceType(ctx, &device_type.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}
	var resp *device_type.DeviceType
	var err error
	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &device_type_client.UpdateDeviceTypeRequest{
			DeviceType: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*device_type.DeviceType_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &device_type_client.UpdateDeviceTypeRequest_CAS{
				ConditionalState: conditionalState.(*device_type.DeviceType),
				FieldMask:        mask.(*device_type.DeviceType_FieldMask),
			}
		}
		resp, err = a.client.UpdateDeviceType(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &device_type_client.CreateDeviceTypeRequest{
			DeviceType: res,
		}
		resp, err = a.client.CreateDeviceType(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiDeviceTypeAccess) DeleteDeviceType(ctx context.Context, ref *device_type.Reference, opts ...gotenresource.DeleteOption) error {
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &device_type_client.DeleteDeviceTypeRequest{
		Name: &ref.Name,
	}
	_, err := a.client.DeleteDeviceType(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(device_type.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return device_type.AsAnyCastAccess(NewApiDeviceTypeAccess(device_type_client.NewDeviceTypeServiceClient(cc)))
	})
}
