// Code generated by protoc-gen-goten-access
// Resource: Deployment
// DO NOT EDIT!!!

package deployment_access

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

	deployment_client "github.com/cloudwan/edgelq-sdk/meta/client/v1alpha2/deployment"
	deployment "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/deployment"
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

type apiDeploymentAccess struct {
	client deployment_client.DeploymentServiceClient
}

func NewApiDeploymentAccess(client deployment_client.DeploymentServiceClient) deployment.DeploymentAccess {
	return &apiDeploymentAccess{client: client}
}

func (a *apiDeploymentAccess) GetDeployment(ctx context.Context, query *deployment.GetQuery, opts ...gotenresource.GetOption) (*deployment.Deployment, error) {
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
	request := &deployment_client.GetDeploymentRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetDeployment(ctx, request, callOpts...)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiDeploymentAccess) BatchGetDeployments(ctx context.Context, refs []*deployment.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	callHeaders := metadata.MD{}
	if batchGetOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	asNames := make([]*deployment.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &deployment_client.BatchGetDeploymentsRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(deployment.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*deployment.Deployment_FieldMask)
	}
	resp, err := a.client.BatchGetDeployments(ctx, request, callOpts...)
	if err != nil {
		return err
	}
	resultMap := make(map[deployment.Name]*deployment.Deployment, len(refs))
	for _, resolvedRes := range resp.GetDeployments() {
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

func (a *apiDeploymentAccess) QueryDeployments(ctx context.Context, query *deployment.ListQuery, opts ...gotenresource.QueryOption) (*deployment.QueryResultSnapshot, error) {
	qOpts := gotenresource.MakeQueryOptions(opts)
	callHeaders := metadata.MD{}
	if qOpts.GetSkipCache() {
		callHeaders["cache-control"] = []string{"no-cache"}
	}
	callOpts := []grpc.CallOption{}
	if len(callHeaders) > 0 {
		callOpts = append(callOpts, grpc.Header(&callHeaders))
	}
	request := &deployment_client.ListDeploymentsRequest{
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
	resp, err := a.client.ListDeployments(ctx, request)
	if err != nil {
		return nil, err
	}
	return &deployment.QueryResultSnapshot{
		Deployments:       resp.Deployments,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiDeploymentAccess) WatchDeployment(ctx context.Context, query *deployment.GetQuery, observerCb func(*deployment.DeploymentChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &deployment_client.WatchDeploymentRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	changesStream, initErr := a.client.WatchDeployment(ctx, request)
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

func (a *apiDeploymentAccess) WatchDeployments(ctx context.Context, query *deployment.WatchQuery, observerCb func(*deployment.QueryResultChange) error) error {
	request := &deployment_client.WatchDeploymentsRequest{
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

	changesStream, initErr := a.client.WatchDeployments(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return status.Errorf(status.Code(err), "watch recv error: %s", err)
		}
		changesWithPaging := &deployment.QueryResultChange{
			Changes:      respChange.DeploymentChanges,
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

func (a *apiDeploymentAccess) SaveDeployment(ctx context.Context, res *deployment.Deployment, opts ...gotenresource.SaveOption) error {
	return status.Errorf(codes.Internal, "save operation on %s does not exist", res.Name.AsReference().String())
}

func (a *apiDeploymentAccess) DeleteDeployment(ctx context.Context, ref *deployment.Reference, opts ...gotenresource.DeleteOption) error {
	return status.Errorf(codes.Internal, "Delete operation on Deployment is prohibited")
}
func getParentAndFilter(fullFilter *deployment.Filter) (*deployment.Filter, *deployment.ParentName) {
	var withParentExtraction func(cnd deployment.FilterCondition) deployment.FilterCondition
	var resultParent *deployment.ParentName
	var resultFilter *deployment.Filter
	withParentExtraction = func(cnd deployment.FilterCondition) deployment.FilterCondition {
		switch tCnd := cnd.(type) {
		case *deployment.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]deployment.FilterCondition, 0)
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
				return deployment.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *deployment.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*deployment.Name)
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
		resultFilter = &deployment.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func GetApiAccessBuilder() *gotenaccess.ApiAccessBuilder {
	return gotenaccess.GetRegistry().FindApiAccessBuilder(deployment.GetDescriptor())
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(deployment.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return deployment.AsAnyCastAccess(NewApiDeploymentAccess(deployment_client.NewDeploymentServiceClient(cc)))
	})
}
