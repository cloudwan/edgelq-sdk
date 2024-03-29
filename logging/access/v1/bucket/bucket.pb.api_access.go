// Code generated by protoc-gen-goten-access
// Resource: Bucket
// DO NOT EDIT!!!

package bucket_access

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

	bucket_client "github.com/cloudwan/edgelq-sdk/logging/client/v1/bucket"
	bucket "github.com/cloudwan/edgelq-sdk/logging/resources/v1/bucket"
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

type apiBucketAccess struct {
	client bucket_client.BucketServiceClient
}

func NewApiBucketAccess(client bucket_client.BucketServiceClient) bucket.BucketAccess {
	return &apiBucketAccess{client: client}
}

func (a *apiBucketAccess) GetBucket(ctx context.Context, query *bucket.GetQuery) (*bucket.Bucket, error) {
	if !query.Reference.IsFullyQualified() {
		return nil, status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &bucket_client.GetBucketRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetBucket(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiBucketAccess) BatchGetBuckets(ctx context.Context, refs []*bucket.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	asNames := make([]*bucket.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &bucket_client.BatchGetBucketsRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(bucket.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*bucket.Bucket_FieldMask)
	}
	resp, err := a.client.BatchGetBuckets(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[bucket.Name]*bucket.Bucket, len(refs))
	for _, resolvedRes := range resp.GetBuckets() {
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

func (a *apiBucketAccess) QueryBuckets(ctx context.Context, query *bucket.ListQuery) (*bucket.QueryResultSnapshot, error) {
	request := &bucket_client.ListBucketsRequest{
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
	resp, err := a.client.ListBuckets(ctx, request)
	if err != nil {
		return nil, err
	}
	return &bucket.QueryResultSnapshot{
		Buckets:           resp.Buckets,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiBucketAccess) WatchBucket(ctx context.Context, query *bucket.GetQuery, observerCb func(*bucket.BucketChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &bucket_client.WatchBucketRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchBucket(ctx, request)
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

func (a *apiBucketAccess) WatchBuckets(ctx context.Context, query *bucket.WatchQuery, observerCb func(*bucket.QueryResultChange) error) error {
	request := &bucket_client.WatchBucketsRequest{
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
	if query.Filter != nil && query.Filter.GetCondition() != nil {
		request.Filter, request.Parent = getParentAndFilter(query.Filter)
	}
	changesStream, initErr := a.client.WatchBuckets(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &bucket.QueryResultChange{
			Changes:      respChange.BucketChanges,
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

func (a *apiBucketAccess) SaveBucket(ctx context.Context, res *bucket.Bucket, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetBucket(ctx, &bucket.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}
	var resp *bucket.Bucket
	var err error
	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &bucket_client.UpdateBucketRequest{
			Bucket: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*bucket.Bucket_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &bucket_client.UpdateBucketRequest_CAS{
				ConditionalState: conditionalState.(*bucket.Bucket),
				FieldMask:        mask.(*bucket.Bucket_FieldMask),
			}
		}
		resp, err = a.client.UpdateBucket(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &bucket_client.CreateBucketRequest{
			Bucket: res,
		}
		resp, err = a.client.CreateBucket(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiBucketAccess) DeleteBucket(ctx context.Context, ref *bucket.Reference, opts ...gotenresource.DeleteOption) error {
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &bucket_client.DeleteBucketRequest{
		Name: &ref.Name,
	}
	_, err := a.client.DeleteBucket(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *bucket.Filter) (*bucket.Filter, *bucket.ParentName) {
	var withParentExtraction func(cnd bucket.FilterCondition) bucket.FilterCondition
	var resultParent *bucket.ParentName
	var resultFilter *bucket.Filter
	withParentExtraction = func(cnd bucket.FilterCondition) bucket.FilterCondition {
		switch tCnd := cnd.(type) {
		case *bucket.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]bucket.FilterCondition, 0)
				for _, subCnd := range tCnd.Conditions {
					if subCndNoParent := withParentExtraction(subCnd); subCndNoParent != nil {
						withoutParentCnds = append(withoutParentCnds, subCndNoParent)
					}
				}
				if len(withoutParentCnds) == 0 {
					return nil
				}
				return bucket.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *bucket.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*bucket.Name)
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
		resultFilter = &bucket.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(bucket.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return bucket.AsAnyCastAccess(NewApiBucketAccess(bucket_client.NewBucketServiceClient(cc)))
	})
}
