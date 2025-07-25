// Code generated by protoc-gen-goten-access
// Resource: CustomizedImage
// DO NOT EDIT!!!

package customized_image_access

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

	customized_image_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha2/customized_image"
	customized_image "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/customized_image"
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

type apiCustomizedImageAccess struct {
	client customized_image_client.CustomizedImageServiceClient
}

func NewApiCustomizedImageAccess(client customized_image_client.CustomizedImageServiceClient) customized_image.CustomizedImageAccess {
	return &apiCustomizedImageAccess{client: client}
}

func (a *apiCustomizedImageAccess) GetCustomizedImage(ctx context.Context, query *customized_image.GetQuery, opts ...gotenresource.GetOption) (*customized_image.CustomizedImage, error) {
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
	request := &customized_image_client.GetCustomizedImageRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetCustomizedImage(ctx, request, callOpts...)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiCustomizedImageAccess) BatchGetCustomizedImages(ctx context.Context, refs []*customized_image.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	callHeaders := metadata.MD{}
	if batchGetOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	asNames := make([]*customized_image.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &customized_image_client.BatchGetCustomizedImagesRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(customized_image.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*customized_image.CustomizedImage_FieldMask)
	}
	resp, err := a.client.BatchGetCustomizedImages(ctx, request, callOpts...)
	if err != nil {
		return err
	}
	resultMap := make(map[customized_image.Name]*customized_image.CustomizedImage, len(refs))
	for _, resolvedRes := range resp.GetCustomizedImages() {
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

func (a *apiCustomizedImageAccess) QueryCustomizedImages(ctx context.Context, query *customized_image.ListQuery, opts ...gotenresource.QueryOption) (*customized_image.QueryResultSnapshot, error) {
	qOpts := gotenresource.MakeQueryOptions(opts)
	callHeaders := metadata.MD{}
	if qOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	request := &customized_image_client.ListCustomizedImagesRequest{
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
	resp, err := a.client.ListCustomizedImages(ctx, request)
	if err != nil {
		return nil, err
	}
	return &customized_image.QueryResultSnapshot{
		CustomizedImages:  resp.CustomizedImages,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiCustomizedImageAccess) WatchCustomizedImage(ctx context.Context, query *customized_image.GetQuery, observerCb func(*customized_image.CustomizedImageChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &customized_image_client.WatchCustomizedImageRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	changesStream, initErr := a.client.WatchCustomizedImage(ctx, request)
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

func (a *apiCustomizedImageAccess) WatchCustomizedImages(ctx context.Context, query *customized_image.WatchQuery, observerCb func(*customized_image.QueryResultChange) error) error {
	request := &customized_image_client.WatchCustomizedImagesRequest{
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

	changesStream, initErr := a.client.WatchCustomizedImages(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return status.Errorf(status.Code(err), "watch recv error: %s", err)
		}
		changesWithPaging := &customized_image.QueryResultChange{
			Changes:      respChange.CustomizedImageChanges,
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

func (a *apiCustomizedImageAccess) SaveCustomizedImage(ctx context.Context, res *customized_image.CustomizedImage, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	var resp *customized_image.CustomizedImage
	var err error
	if !saveOpts.OnlyCreate() {
		updateRequest := &customized_image_client.UpdateCustomizedImageRequest{
			CustomizedImage: res,
			AllowMissing:    !saveOpts.OnlyUpdate(),
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*customized_image.CustomizedImage_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &customized_image_client.UpdateCustomizedImageRequest_CAS{
				ConditionalState: conditionalState.(*customized_image.CustomizedImage),
				FieldMask:        mask.(*customized_image.CustomizedImage_FieldMask),
			}
		}
		resp, err = a.client.UpdateCustomizedImage(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &customized_image_client.CreateCustomizedImageRequest{
			CustomizedImage: res,
		}
		resp, err = a.client.CreateCustomizedImage(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiCustomizedImageAccess) DeleteCustomizedImage(ctx context.Context, ref *customized_image.Reference, opts ...gotenresource.DeleteOption) error {
	delOpts := gotenresource.MakeDeleteOptions(opts)
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &customized_image_client.DeleteCustomizedImageRequest{
		Name:         &ref.Name,
		AllowMissing: delOpts.AllowMissing(),
	}
	_, err := a.client.DeleteCustomizedImage(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *customized_image.Filter) (*customized_image.Filter, *customized_image.ParentName) {
	var withParentExtraction func(cnd customized_image.FilterCondition) customized_image.FilterCondition
	var resultParent *customized_image.ParentName
	var resultFilter *customized_image.Filter
	withParentExtraction = func(cnd customized_image.FilterCondition) customized_image.FilterCondition {
		switch tCnd := cnd.(type) {
		case *customized_image.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]customized_image.FilterCondition, 0)
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
				return customized_image.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *customized_image.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*customized_image.Name)
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
		resultFilter = &customized_image.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func GetApiAccessBuilder() *gotenaccess.ApiAccessBuilder {
	return gotenaccess.GetRegistry().FindApiAccessBuilder(customized_image.GetDescriptor())
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(customized_image.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return customized_image.AsAnyCastAccess(NewApiCustomizedImageAccess(customized_image_client.NewCustomizedImageServiceClient(cc)))
	})
}
