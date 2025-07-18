// Code generated by protoc-gen-goten-access
// Resource: ServiceAccount
// DO NOT EDIT!!!

package service_account_access

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

	service_account_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha2/service_account"
	service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/service_account"
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

type apiServiceAccountAccess struct {
	client service_account_client.ServiceAccountServiceClient
}

func NewApiServiceAccountAccess(client service_account_client.ServiceAccountServiceClient) service_account.ServiceAccountAccess {
	return &apiServiceAccountAccess{client: client}
}

func (a *apiServiceAccountAccess) GetServiceAccount(ctx context.Context, query *service_account.GetQuery, opts ...gotenresource.GetOption) (*service_account.ServiceAccount, error) {
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
	request := &service_account_client.GetServiceAccountRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetServiceAccount(ctx, request, callOpts...)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiServiceAccountAccess) BatchGetServiceAccounts(ctx context.Context, refs []*service_account.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	callHeaders := metadata.MD{}
	if batchGetOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	asNames := make([]*service_account.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &service_account_client.BatchGetServiceAccountsRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(service_account.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*service_account.ServiceAccount_FieldMask)
	}
	resp, err := a.client.BatchGetServiceAccounts(ctx, request, callOpts...)
	if err != nil {
		return err
	}
	resultMap := make(map[service_account.Name]*service_account.ServiceAccount, len(refs))
	for _, resolvedRes := range resp.GetServiceAccounts() {
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

func (a *apiServiceAccountAccess) QueryServiceAccounts(ctx context.Context, query *service_account.ListQuery, opts ...gotenresource.QueryOption) (*service_account.QueryResultSnapshot, error) {
	qOpts := gotenresource.MakeQueryOptions(opts)
	callHeaders := metadata.MD{}
	if qOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	request := &service_account_client.ListServiceAccountsRequest{
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
	resp, err := a.client.ListServiceAccounts(ctx, request)
	if err != nil {
		return nil, err
	}
	return &service_account.QueryResultSnapshot{
		ServiceAccounts:   resp.ServiceAccounts,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiServiceAccountAccess) WatchServiceAccount(ctx context.Context, query *service_account.GetQuery, observerCb func(*service_account.ServiceAccountChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &service_account_client.WatchServiceAccountRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	changesStream, initErr := a.client.WatchServiceAccount(ctx, request)
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

func (a *apiServiceAccountAccess) WatchServiceAccounts(ctx context.Context, query *service_account.WatchQuery, observerCb func(*service_account.QueryResultChange) error) error {
	request := &service_account_client.WatchServiceAccountsRequest{
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

	changesStream, initErr := a.client.WatchServiceAccounts(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return status.Errorf(status.Code(err), "watch recv error: %s", err)
		}
		changesWithPaging := &service_account.QueryResultChange{
			Changes:      respChange.ServiceAccountChanges,
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

func (a *apiServiceAccountAccess) SaveServiceAccount(ctx context.Context, res *service_account.ServiceAccount, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	var resp *service_account.ServiceAccount
	var err error
	if !saveOpts.OnlyCreate() {
		updateRequest := &service_account_client.UpdateServiceAccountRequest{
			ServiceAccount: res,
			AllowMissing:   !saveOpts.OnlyUpdate(),
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*service_account.ServiceAccount_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &service_account_client.UpdateServiceAccountRequest_CAS{
				ConditionalState: conditionalState.(*service_account.ServiceAccount),
				FieldMask:        mask.(*service_account.ServiceAccount_FieldMask),
			}
		}
		resp, err = a.client.UpdateServiceAccount(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &service_account_client.CreateServiceAccountRequest{
			ServiceAccount: res,
		}
		resp, err = a.client.CreateServiceAccount(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiServiceAccountAccess) DeleteServiceAccount(ctx context.Context, ref *service_account.Reference, opts ...gotenresource.DeleteOption) error {
	delOpts := gotenresource.MakeDeleteOptions(opts)
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &service_account_client.DeleteServiceAccountRequest{
		Name:         &ref.Name,
		AllowMissing: delOpts.AllowMissing(),
	}
	_, err := a.client.DeleteServiceAccount(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *service_account.Filter) (*service_account.Filter, *service_account.ParentName) {
	var withParentExtraction func(cnd service_account.FilterCondition) service_account.FilterCondition
	var resultParent *service_account.ParentName
	var resultFilter *service_account.Filter
	withParentExtraction = func(cnd service_account.FilterCondition) service_account.FilterCondition {
		switch tCnd := cnd.(type) {
		case *service_account.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]service_account.FilterCondition, 0)
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
				return service_account.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *service_account.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*service_account.Name)
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
		resultFilter = &service_account.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func GetApiAccessBuilder() *gotenaccess.ApiAccessBuilder {
	return gotenaccess.GetRegistry().FindApiAccessBuilder(service_account.GetDescriptor())
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(service_account.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return service_account.AsAnyCastAccess(NewApiServiceAccountAccess(service_account_client.NewServiceAccountServiceClient(cc)))
	})
}
