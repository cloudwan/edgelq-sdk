// Code generated by protoc-gen-goten-access
// Resource: OsVersion
// DO NOT EDIT!!!

package os_version_access

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

	os_version_client "github.com/cloudwan/edgelq-sdk/devices/client/v1/os_version"
	os_version "github.com/cloudwan/edgelq-sdk/devices/resources/v1/os_version"
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

type apiOsVersionAccess struct {
	client os_version_client.OsVersionServiceClient
}

func NewApiOsVersionAccess(client os_version_client.OsVersionServiceClient) os_version.OsVersionAccess {
	return &apiOsVersionAccess{client: client}
}

func (a *apiOsVersionAccess) GetOsVersion(ctx context.Context, query *os_version.GetQuery) (*os_version.OsVersion, error) {
	if !query.Reference.IsFullyQualified() {
		return nil, status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &os_version_client.GetOsVersionRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetOsVersion(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiOsVersionAccess) BatchGetOsVersions(ctx context.Context, refs []*os_version.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	asNames := make([]*os_version.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &os_version_client.BatchGetOsVersionsRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(os_version.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*os_version.OsVersion_FieldMask)
	}
	resp, err := a.client.BatchGetOsVersions(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[os_version.Name]*os_version.OsVersion, len(refs))
	for _, resolvedRes := range resp.GetOsVersions() {
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

func (a *apiOsVersionAccess) QueryOsVersions(ctx context.Context, query *os_version.ListQuery) (*os_version.QueryResultSnapshot, error) {
	request := &os_version_client.ListOsVersionsRequest{
		Filter:            query.Filter,
		FieldMask:         query.Mask,
		IncludePagingInfo: query.WithPagingInfo,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListOsVersions(ctx, request)
	if err != nil {
		return nil, err
	}
	return &os_version.QueryResultSnapshot{
		OsVersions:        resp.OsVersions,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiOsVersionAccess) WatchOsVersion(ctx context.Context, query *os_version.GetQuery, observerCb func(*os_version.OsVersionChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &os_version_client.WatchOsVersionRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchOsVersion(ctx, request)
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

func (a *apiOsVersionAccess) WatchOsVersions(ctx context.Context, query *os_version.WatchQuery, observerCb func(*os_version.QueryResultChange) error) error {
	request := &os_version_client.WatchOsVersionsRequest{
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
	changesStream, initErr := a.client.WatchOsVersions(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &os_version.QueryResultChange{
			Changes:      respChange.OsVersionChanges,
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

func (a *apiOsVersionAccess) SaveOsVersion(ctx context.Context, res *os_version.OsVersion, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetOsVersion(ctx, &os_version.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}
	var resp *os_version.OsVersion
	var err error
	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &os_version_client.UpdateOsVersionRequest{
			OsVersion: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*os_version.OsVersion_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &os_version_client.UpdateOsVersionRequest_CAS{
				ConditionalState: conditionalState.(*os_version.OsVersion),
				FieldMask:        mask.(*os_version.OsVersion_FieldMask),
			}
		}
		resp, err = a.client.UpdateOsVersion(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &os_version_client.CreateOsVersionRequest{
			OsVersion: res,
		}
		resp, err = a.client.CreateOsVersion(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiOsVersionAccess) DeleteOsVersion(ctx context.Context, ref *os_version.Reference, opts ...gotenresource.DeleteOption) error {
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &os_version_client.DeleteOsVersionRequest{
		Name: &ref.Name,
	}
	_, err := a.client.DeleteOsVersion(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(os_version.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return os_version.AsAnyCastAccess(NewApiOsVersionAccess(os_version_client.NewOsVersionServiceClient(cc)))
	})
}
