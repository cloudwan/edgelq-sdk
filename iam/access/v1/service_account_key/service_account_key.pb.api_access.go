// Code generated by protoc-gen-goten-access
// Resource: ServiceAccountKey
// DO NOT EDIT!!!

package service_account_key_access

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

	service_account_key_client "github.com/cloudwan/edgelq-sdk/iam/client/v1/service_account_key"
	service_account_key "github.com/cloudwan/edgelq-sdk/iam/resources/v1/service_account_key"
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

type apiServiceAccountKeyAccess struct {
	client service_account_key_client.ServiceAccountKeyServiceClient
}

func NewApiServiceAccountKeyAccess(client service_account_key_client.ServiceAccountKeyServiceClient) service_account_key.ServiceAccountKeyAccess {
	return &apiServiceAccountKeyAccess{client: client}
}

func (a *apiServiceAccountKeyAccess) GetServiceAccountKey(ctx context.Context, query *service_account_key.GetQuery, opts ...gotenresource.GetOption) (*service_account_key.ServiceAccountKey, error) {
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
	request := &service_account_key_client.GetServiceAccountKeyRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetServiceAccountKey(ctx, request, callOpts...)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiServiceAccountKeyAccess) BatchGetServiceAccountKeys(ctx context.Context, refs []*service_account_key.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	callHeaders := metadata.MD{}
	if batchGetOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	asNames := make([]*service_account_key.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &service_account_key_client.BatchGetServiceAccountKeysRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(service_account_key.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*service_account_key.ServiceAccountKey_FieldMask)
	}
	resp, err := a.client.BatchGetServiceAccountKeys(ctx, request, callOpts...)
	if err != nil {
		return err
	}
	resultMap := make(map[service_account_key.Name]*service_account_key.ServiceAccountKey, len(refs))
	for _, resolvedRes := range resp.GetServiceAccountKeys() {
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

func (a *apiServiceAccountKeyAccess) QueryServiceAccountKeys(ctx context.Context, query *service_account_key.ListQuery, opts ...gotenresource.QueryOption) (*service_account_key.QueryResultSnapshot, error) {
	qOpts := gotenresource.MakeQueryOptions(opts)
	callHeaders := metadata.MD{}
	if qOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	request := &service_account_key_client.ListServiceAccountKeysRequest{
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
	resp, err := a.client.ListServiceAccountKeys(ctx, request)
	if err != nil {
		return nil, err
	}
	return &service_account_key.QueryResultSnapshot{
		ServiceAccountKeys: resp.ServiceAccountKeys,
		NextPageCursor:     resp.NextPageToken,
		PrevPageCursor:     resp.PrevPageToken,
		TotalResultsCount:  resp.TotalResultsCount,
		CurrentOffset:      resp.CurrentOffset,
	}, nil
}

func (a *apiServiceAccountKeyAccess) WatchServiceAccountKey(ctx context.Context, query *service_account_key.GetQuery, observerCb func(*service_account_key.ServiceAccountKeyChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &service_account_key_client.WatchServiceAccountKeyRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	changesStream, initErr := a.client.WatchServiceAccountKey(ctx, request)
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

func (a *apiServiceAccountKeyAccess) WatchServiceAccountKeys(ctx context.Context, query *service_account_key.WatchQuery, observerCb func(*service_account_key.QueryResultChange) error) error {
	request := &service_account_key_client.WatchServiceAccountKeysRequest{
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

	changesStream, initErr := a.client.WatchServiceAccountKeys(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return status.Errorf(status.Code(err), "watch recv error: %s", err)
		}
		changesWithPaging := &service_account_key.QueryResultChange{
			Changes:      respChange.ServiceAccountKeyChanges,
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

func (a *apiServiceAccountKeyAccess) SaveServiceAccountKey(ctx context.Context, res *service_account_key.ServiceAccountKey, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	var resp *service_account_key.ServiceAccountKey
	var err error
	if !saveOpts.OnlyCreate() {
		updateRequest := &service_account_key_client.UpdateServiceAccountKeyRequest{
			ServiceAccountKey: res,
			AllowMissing:      !saveOpts.OnlyUpdate(),
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*service_account_key.ServiceAccountKey_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &service_account_key_client.UpdateServiceAccountKeyRequest_CAS{
				ConditionalState: conditionalState.(*service_account_key.ServiceAccountKey),
				FieldMask:        mask.(*service_account_key.ServiceAccountKey_FieldMask),
			}
		}
		resp, err = a.client.UpdateServiceAccountKey(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &service_account_key_client.CreateServiceAccountKeyRequest{
			ServiceAccountKey: res,
		}
		resp, err = a.client.CreateServiceAccountKey(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiServiceAccountKeyAccess) DeleteServiceAccountKey(ctx context.Context, ref *service_account_key.Reference, opts ...gotenresource.DeleteOption) error {
	delOpts := gotenresource.MakeDeleteOptions(opts)
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &service_account_key_client.DeleteServiceAccountKeyRequest{
		Name:         &ref.Name,
		AllowMissing: delOpts.AllowMissing(),
	}
	_, err := a.client.DeleteServiceAccountKey(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *service_account_key.Filter) (*service_account_key.Filter, *service_account_key.ParentName) {
	var withParentExtraction func(cnd service_account_key.FilterCondition) service_account_key.FilterCondition
	var resultParent *service_account_key.ParentName
	var resultFilter *service_account_key.Filter
	withParentExtraction = func(cnd service_account_key.FilterCondition) service_account_key.FilterCondition {
		switch tCnd := cnd.(type) {
		case *service_account_key.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]service_account_key.FilterCondition, 0)
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
				return service_account_key.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *service_account_key.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*service_account_key.Name)
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
		resultFilter = &service_account_key.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func GetApiAccessBuilder() *gotenaccess.ApiAccessBuilder {
	return gotenaccess.GetRegistry().FindApiAccessBuilder(service_account_key.GetDescriptor())
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(service_account_key.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return service_account_key.AsAnyCastAccess(NewApiServiceAccountKeyAccess(service_account_key_client.NewServiceAccountKeyServiceClient(cc)))
	})
}
