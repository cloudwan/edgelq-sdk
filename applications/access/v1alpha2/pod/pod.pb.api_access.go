// Code generated by protoc-gen-goten-access
// Resource: Pod
// DO NOT EDIT!!!

package pod_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	pod_client "github.com/cloudwan/edgelq-sdk/applications/client/v1alpha2/pod"
	pod "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha2/pod"
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

type apiPodAccess struct {
	client pod_client.PodServiceClient
}

func NewApiPodAccess(client pod_client.PodServiceClient) pod.PodAccess {
	return &apiPodAccess{client: client}
}

func (a *apiPodAccess) GetPod(ctx context.Context, query *pod.GetQuery) (*pod.Pod, error) {
	request := &pod_client.GetPodRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetPod(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiPodAccess) BatchGetPods(ctx context.Context, refs []*pod.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &pod_client.BatchGetPodsRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetPods(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[pod.Name]*pod.Pod, len(refs))
	for _, resolvedRes := range resp.GetPods() {
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

func (a *apiPodAccess) QueryPods(ctx context.Context, query *pod.ListQuery) (*pod.QueryResultSnapshot, error) {
	request := &pod_client.ListPodsRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListPods(ctx, request)
	if err != nil {
		return nil, err
	}
	return &pod.QueryResultSnapshot{
		Pods:           resp.Pods,
		NextPageCursor: resp.NextPageToken,
		PrevPageCursor: resp.PrevPageToken,
	}, nil
}

func (a *apiPodAccess) WatchPod(ctx context.Context, query *pod.GetQuery, observerCb func(*pod.PodChange) error) error {
	request := &pod_client.WatchPodRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchPod(ctx, request)
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

func (a *apiPodAccess) WatchPods(ctx context.Context, query *pod.WatchQuery, observerCb func(*pod.QueryResultChange) error) error {
	request := &pod_client.WatchPodsRequest{
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
	changesStream, initErr := a.client.WatchPods(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &pod.QueryResultChange{
			Changes:      respChange.PodChanges,
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

func (a *apiPodAccess) SavePod(ctx context.Context, res *pod.Pod, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetPod(ctx, &pod.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &pod_client.UpdatePodRequest{
			Pod: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*pod.Pod_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &pod_client.UpdatePodRequest_CAS{
				ConditionalState: conditionalState.(*pod.Pod),
				FieldMask:        mask.(*pod.Pod_FieldMask),
			}
		}
		_, err := a.client.UpdatePod(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &pod_client.CreatePodRequest{
			Pod: res,
		}
		_, err := a.client.CreatePod(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiPodAccess) DeletePod(ctx context.Context, ref *pod.Reference, opts ...gotenresource.DeleteOption) error {
	request := &pod_client.DeletePodRequest{
		Name: ref,
	}
	_, err := a.client.DeletePod(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(pod.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return pod.AsAnyCastAccess(NewApiPodAccess(pod_client.NewPodServiceClient(cc)))
	})
}
