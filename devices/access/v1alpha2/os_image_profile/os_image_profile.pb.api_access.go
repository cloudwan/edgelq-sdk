// Code generated by protoc-gen-goten-access
// Resource: OsImageProfile
// DO NOT EDIT!!!

package os_image_profile_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	os_image_profile_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha2/os_image_profile"
	os_image_profile "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/os_image_profile"
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

type apiOsImageProfileAccess struct {
	client os_image_profile_client.OsImageProfileServiceClient
}

func NewApiOsImageProfileAccess(client os_image_profile_client.OsImageProfileServiceClient) os_image_profile.OsImageProfileAccess {
	return &apiOsImageProfileAccess{client: client}
}

func (a *apiOsImageProfileAccess) GetOsImageProfile(ctx context.Context, query *os_image_profile.GetQuery) (*os_image_profile.OsImageProfile, error) {
	request := &os_image_profile_client.GetOsImageProfileRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetOsImageProfile(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiOsImageProfileAccess) BatchGetOsImageProfiles(ctx context.Context, refs []*os_image_profile.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &os_image_profile_client.BatchGetOsImageProfilesRequest{
		Names: refs,
	}
	fieldMask := batchGetOpts.GetFieldMask(os_image_profile.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*os_image_profile.OsImageProfile_FieldMask)
	}
	resp, err := a.client.BatchGetOsImageProfiles(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[os_image_profile.Name]*os_image_profile.OsImageProfile, len(refs))
	for _, resolvedRes := range resp.GetOsImageProfiles() {
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

func (a *apiOsImageProfileAccess) QueryOsImageProfiles(ctx context.Context, query *os_image_profile.ListQuery) (*os_image_profile.QueryResultSnapshot, error) {
	request := &os_image_profile_client.ListOsImageProfilesRequest{
		Filter:            query.Filter,
		FieldMask:         query.Mask,
		IncludePagingInfo: query.WithPagingInfo,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListOsImageProfiles(ctx, request)
	if err != nil {
		return nil, err
	}
	return &os_image_profile.QueryResultSnapshot{
		OsImageProfiles:   resp.OsImageProfiles,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiOsImageProfileAccess) WatchOsImageProfile(ctx context.Context, query *os_image_profile.GetQuery, observerCb func(*os_image_profile.OsImageProfileChange) error) error {
	request := &os_image_profile_client.WatchOsImageProfileRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchOsImageProfile(ctx, request)
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

func (a *apiOsImageProfileAccess) WatchOsImageProfiles(ctx context.Context, query *os_image_profile.WatchQuery, observerCb func(*os_image_profile.QueryResultChange) error) error {
	request := &os_image_profile_client.WatchOsImageProfilesRequest{
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
	changesStream, initErr := a.client.WatchOsImageProfiles(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &os_image_profile.QueryResultChange{
			Changes:      respChange.OsImageProfileChanges,
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

func (a *apiOsImageProfileAccess) SaveOsImageProfile(ctx context.Context, res *os_image_profile.OsImageProfile, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetOsImageProfile(ctx, &os_image_profile.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &os_image_profile_client.UpdateOsImageProfileRequest{
			OsImageProfile: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*os_image_profile.OsImageProfile_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &os_image_profile_client.UpdateOsImageProfileRequest_CAS{
				ConditionalState: conditionalState.(*os_image_profile.OsImageProfile),
				FieldMask:        mask.(*os_image_profile.OsImageProfile_FieldMask),
			}
		}
		_, err := a.client.UpdateOsImageProfile(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &os_image_profile_client.CreateOsImageProfileRequest{
			OsImageProfile: res,
		}
		_, err := a.client.CreateOsImageProfile(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiOsImageProfileAccess) DeleteOsImageProfile(ctx context.Context, ref *os_image_profile.Reference, opts ...gotenresource.DeleteOption) error {
	request := &os_image_profile_client.DeleteOsImageProfileRequest{
		Name: ref,
	}
	_, err := a.client.DeleteOsImageProfile(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(os_image_profile.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return os_image_profile.AsAnyCastAccess(NewApiOsImageProfileAccess(os_image_profile_client.NewOsImageProfileServiceClient(cc)))
	})
}
