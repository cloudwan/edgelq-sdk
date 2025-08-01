// Code generated by protoc-gen-goten-access
// Resource: AttestationDomain
// DO NOT EDIT!!!

package attestation_domain_access

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

	attestation_domain_client "github.com/cloudwan/edgelq-sdk/iam/client/v1/attestation_domain"
	attestation_domain "github.com/cloudwan/edgelq-sdk/iam/resources/v1/attestation_domain"
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

type apiAttestationDomainAccess struct {
	client attestation_domain_client.AttestationDomainServiceClient
}

func NewApiAttestationDomainAccess(client attestation_domain_client.AttestationDomainServiceClient) attestation_domain.AttestationDomainAccess {
	return &apiAttestationDomainAccess{client: client}
}

func (a *apiAttestationDomainAccess) GetAttestationDomain(ctx context.Context, query *attestation_domain.GetQuery, opts ...gotenresource.GetOption) (*attestation_domain.AttestationDomain, error) {
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
	request := &attestation_domain_client.GetAttestationDomainRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetAttestationDomain(ctx, request, callOpts...)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiAttestationDomainAccess) BatchGetAttestationDomains(ctx context.Context, refs []*attestation_domain.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	callHeaders := metadata.MD{}
	if batchGetOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	asNames := make([]*attestation_domain.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &attestation_domain_client.BatchGetAttestationDomainsRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(attestation_domain.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*attestation_domain.AttestationDomain_FieldMask)
	}
	resp, err := a.client.BatchGetAttestationDomains(ctx, request, callOpts...)
	if err != nil {
		return err
	}
	resultMap := make(map[attestation_domain.Name]*attestation_domain.AttestationDomain, len(refs))
	for _, resolvedRes := range resp.GetAttestationDomains() {
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

func (a *apiAttestationDomainAccess) QueryAttestationDomains(ctx context.Context, query *attestation_domain.ListQuery, opts ...gotenresource.QueryOption) (*attestation_domain.QueryResultSnapshot, error) {
	qOpts := gotenresource.MakeQueryOptions(opts)
	callHeaders := metadata.MD{}
	if qOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	request := &attestation_domain_client.ListAttestationDomainsRequest{
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
	resp, err := a.client.ListAttestationDomains(ctx, request)
	if err != nil {
		return nil, err
	}
	return &attestation_domain.QueryResultSnapshot{
		AttestationDomains: resp.AttestationDomains,
		NextPageCursor:     resp.NextPageToken,
		PrevPageCursor:     resp.PrevPageToken,
		TotalResultsCount:  resp.TotalResultsCount,
		CurrentOffset:      resp.CurrentOffset,
	}, nil
}

func (a *apiAttestationDomainAccess) WatchAttestationDomain(ctx context.Context, query *attestation_domain.GetQuery, observerCb func(*attestation_domain.AttestationDomainChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &attestation_domain_client.WatchAttestationDomainRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	changesStream, initErr := a.client.WatchAttestationDomain(ctx, request)
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

func (a *apiAttestationDomainAccess) WatchAttestationDomains(ctx context.Context, query *attestation_domain.WatchQuery, observerCb func(*attestation_domain.QueryResultChange) error) error {
	request := &attestation_domain_client.WatchAttestationDomainsRequest{
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

	changesStream, initErr := a.client.WatchAttestationDomains(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return status.Errorf(status.Code(err), "watch recv error: %s", err)
		}
		changesWithPaging := &attestation_domain.QueryResultChange{
			Changes:      respChange.AttestationDomainChanges,
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

func (a *apiAttestationDomainAccess) SaveAttestationDomain(ctx context.Context, res *attestation_domain.AttestationDomain, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	var resp *attestation_domain.AttestationDomain
	var err error
	if !saveOpts.OnlyCreate() {
		updateRequest := &attestation_domain_client.UpdateAttestationDomainRequest{
			AttestationDomain: res,
			AllowMissing:      !saveOpts.OnlyUpdate(),
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*attestation_domain.AttestationDomain_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &attestation_domain_client.UpdateAttestationDomainRequest_CAS{
				ConditionalState: conditionalState.(*attestation_domain.AttestationDomain),
				FieldMask:        mask.(*attestation_domain.AttestationDomain_FieldMask),
			}
		}
		resp, err = a.client.UpdateAttestationDomain(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &attestation_domain_client.CreateAttestationDomainRequest{
			AttestationDomain: res,
		}
		resp, err = a.client.CreateAttestationDomain(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiAttestationDomainAccess) DeleteAttestationDomain(ctx context.Context, ref *attestation_domain.Reference, opts ...gotenresource.DeleteOption) error {
	delOpts := gotenresource.MakeDeleteOptions(opts)
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &attestation_domain_client.DeleteAttestationDomainRequest{
		Name:         &ref.Name,
		AllowMissing: delOpts.AllowMissing(),
	}
	_, err := a.client.DeleteAttestationDomain(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *attestation_domain.Filter) (*attestation_domain.Filter, *attestation_domain.ParentName) {
	var withParentExtraction func(cnd attestation_domain.FilterCondition) attestation_domain.FilterCondition
	var resultParent *attestation_domain.ParentName
	var resultFilter *attestation_domain.Filter
	withParentExtraction = func(cnd attestation_domain.FilterCondition) attestation_domain.FilterCondition {
		switch tCnd := cnd.(type) {
		case *attestation_domain.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]attestation_domain.FilterCondition, 0)
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
				return attestation_domain.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *attestation_domain.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*attestation_domain.Name)
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
		resultFilter = &attestation_domain.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func GetApiAccessBuilder() *gotenaccess.ApiAccessBuilder {
	return gotenaccess.GetRegistry().FindApiAccessBuilder(attestation_domain.GetDescriptor())
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(attestation_domain.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return attestation_domain.AsAnyCastAccess(NewApiAttestationDomainAccess(attestation_domain_client.NewAttestationDomainServiceClient(cc)))
	})
}
