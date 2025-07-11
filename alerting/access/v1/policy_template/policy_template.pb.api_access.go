// Code generated by protoc-gen-goten-access
// Resource: PolicyTemplate
// DO NOT EDIT!!!

package policy_template_access

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

	policy_template_client "github.com/cloudwan/edgelq-sdk/alerting/client/v1/policy_template"
	policy_template "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/policy_template"
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

type apiPolicyTemplateAccess struct {
	client policy_template_client.PolicyTemplateServiceClient
}

func NewApiPolicyTemplateAccess(client policy_template_client.PolicyTemplateServiceClient) policy_template.PolicyTemplateAccess {
	return &apiPolicyTemplateAccess{client: client}
}

func (a *apiPolicyTemplateAccess) GetPolicyTemplate(ctx context.Context, query *policy_template.GetQuery, opts ...gotenresource.GetOption) (*policy_template.PolicyTemplate, error) {
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
	request := &policy_template_client.GetPolicyTemplateRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetPolicyTemplate(ctx, request, callOpts...)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiPolicyTemplateAccess) BatchGetPolicyTemplates(ctx context.Context, refs []*policy_template.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	callHeaders := metadata.MD{}
	if batchGetOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	asNames := make([]*policy_template.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &policy_template_client.BatchGetPolicyTemplatesRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(policy_template.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*policy_template.PolicyTemplate_FieldMask)
	}
	resp, err := a.client.BatchGetPolicyTemplates(ctx, request, callOpts...)
	if err != nil {
		return err
	}
	resultMap := make(map[policy_template.Name]*policy_template.PolicyTemplate, len(refs))
	for _, resolvedRes := range resp.GetPolicyTemplates() {
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

func (a *apiPolicyTemplateAccess) QueryPolicyTemplates(ctx context.Context, query *policy_template.ListQuery, opts ...gotenresource.QueryOption) (*policy_template.QueryResultSnapshot, error) {
	qOpts := gotenresource.MakeQueryOptions(opts)
	callHeaders := metadata.MD{}
	if qOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	request := &policy_template_client.ListPolicyTemplatesRequest{
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
	resp, err := a.client.ListPolicyTemplates(ctx, request)
	if err != nil {
		return nil, err
	}
	return &policy_template.QueryResultSnapshot{
		PolicyTemplates:   resp.PolicyTemplates,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiPolicyTemplateAccess) SearchPolicyTemplates(ctx context.Context, query *policy_template.SearchQuery, opts ...gotenresource.QueryOption) (*policy_template.QueryResultSnapshot, error) {
	qOpts := gotenresource.MakeQueryOptions(opts)
	callHeaders := metadata.MD{}
	if qOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	request := &policy_template_client.SearchPolicyTemplatesRequest{
		Phrase:    query.Phrase,
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	if query.Filter != nil && query.Filter.GetCondition() != nil {
		request.Filter, request.Parent = getParentAndFilter(query.Filter)
	}
	resp, err := a.client.SearchPolicyTemplates(ctx, request, callOpts...)
	if err != nil {
		return nil, err
	}
	return &policy_template.QueryResultSnapshot{
		PolicyTemplates:   resp.PolicyTemplates,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		CurrentOffset:     resp.CurrentOffset,
		TotalResultsCount: resp.TotalResultsCount,
	}, nil
}

func (a *apiPolicyTemplateAccess) WatchPolicyTemplate(ctx context.Context, query *policy_template.GetQuery, observerCb func(*policy_template.PolicyTemplateChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &policy_template_client.WatchPolicyTemplateRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	changesStream, initErr := a.client.WatchPolicyTemplate(ctx, request)
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

func (a *apiPolicyTemplateAccess) WatchPolicyTemplates(ctx context.Context, query *policy_template.WatchQuery, observerCb func(*policy_template.QueryResultChange) error) error {
	request := &policy_template_client.WatchPolicyTemplatesRequest{
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

	changesStream, initErr := a.client.WatchPolicyTemplates(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return status.Errorf(status.Code(err), "watch recv error: %s", err)
		}
		changesWithPaging := &policy_template.QueryResultChange{
			Changes:      respChange.PolicyTemplateChanges,
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

func (a *apiPolicyTemplateAccess) SavePolicyTemplate(ctx context.Context, res *policy_template.PolicyTemplate, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	var resp *policy_template.PolicyTemplate
	var err error
	if !saveOpts.OnlyCreate() {
		updateRequest := &policy_template_client.UpdatePolicyTemplateRequest{
			PolicyTemplate: res,
			AllowMissing:   !saveOpts.OnlyUpdate(),
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*policy_template.PolicyTemplate_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &policy_template_client.UpdatePolicyTemplateRequest_CAS{
				ConditionalState: conditionalState.(*policy_template.PolicyTemplate),
				FieldMask:        mask.(*policy_template.PolicyTemplate_FieldMask),
			}
		}
		resp, err = a.client.UpdatePolicyTemplate(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &policy_template_client.CreatePolicyTemplateRequest{
			PolicyTemplate: res,
		}
		resp, err = a.client.CreatePolicyTemplate(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiPolicyTemplateAccess) DeletePolicyTemplate(ctx context.Context, ref *policy_template.Reference, opts ...gotenresource.DeleteOption) error {
	delOpts := gotenresource.MakeDeleteOptions(opts)
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &policy_template_client.DeletePolicyTemplateRequest{
		Name:         &ref.Name,
		AllowMissing: delOpts.AllowMissing(),
	}
	_, err := a.client.DeletePolicyTemplate(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *policy_template.Filter) (*policy_template.Filter, *policy_template.ParentName) {
	var withParentExtraction func(cnd policy_template.FilterCondition) policy_template.FilterCondition
	var resultParent *policy_template.ParentName
	var resultFilter *policy_template.Filter
	withParentExtraction = func(cnd policy_template.FilterCondition) policy_template.FilterCondition {
		switch tCnd := cnd.(type) {
		case *policy_template.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]policy_template.FilterCondition, 0)
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
				return policy_template.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *policy_template.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*policy_template.Name)
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
		resultFilter = &policy_template.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func GetApiAccessBuilder() *gotenaccess.ApiAccessBuilder {
	return gotenaccess.GetRegistry().FindApiAccessBuilder(policy_template.GetDescriptor())
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(policy_template.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return policy_template.AsAnyCastAccess(NewApiPolicyTemplateAccess(policy_template_client.NewPolicyTemplateServiceClient(cc)))
	})
}
