// Code generated by protoc-gen-goten-access
// Resource: Alert
// DO NOT EDIT!!!

package alert_access

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

	alert_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v3/alert"
	alert "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alert"
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

type apiAlertAccess struct {
	client alert_client.AlertServiceClient
}

func NewApiAlertAccess(client alert_client.AlertServiceClient) alert.AlertAccess {
	return &apiAlertAccess{client: client}
}

func (a *apiAlertAccess) GetAlert(ctx context.Context, query *alert.GetQuery) (*alert.Alert, error) {
	if !query.Reference.IsFullyQualified() {
		return nil, status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &alert_client.GetAlertRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetAlert(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiAlertAccess) BatchGetAlerts(ctx context.Context, refs []*alert.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	asNames := make([]*alert.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &alert_client.BatchGetAlertsRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(alert.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*alert.Alert_FieldMask)
	}
	resp, err := a.client.BatchGetAlerts(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[alert.Name]*alert.Alert, len(refs))
	for _, resolvedRes := range resp.GetAlerts() {
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

func (a *apiAlertAccess) QueryAlerts(ctx context.Context, query *alert.ListQuery) (*alert.QueryResultSnapshot, error) {
	request := &alert_client.ListAlertsRequest{
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
	resp, err := a.client.ListAlerts(ctx, request)
	if err != nil {
		return nil, err
	}
	return &alert.QueryResultSnapshot{
		Alerts:            resp.Alerts,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiAlertAccess) WatchAlert(ctx context.Context, query *alert.GetQuery, observerCb func(*alert.AlertChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &alert_client.WatchAlertRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchAlert(ctx, request)
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

func (a *apiAlertAccess) WatchAlerts(ctx context.Context, query *alert.WatchQuery, observerCb func(*alert.QueryResultChange) error) error {
	request := &alert_client.WatchAlertsRequest{
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
	changesStream, initErr := a.client.WatchAlerts(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &alert.QueryResultChange{
			Changes:      respChange.AlertChanges,
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

func (a *apiAlertAccess) SaveAlert(ctx context.Context, res *alert.Alert, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetAlert(ctx, &alert.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}
	var resp *alert.Alert
	var err error
	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &alert_client.UpdateAlertRequest{
			Alert: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*alert.Alert_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &alert_client.UpdateAlertRequest_CAS{
				ConditionalState: conditionalState.(*alert.Alert),
				FieldMask:        mask.(*alert.Alert_FieldMask),
			}
		}
		resp, err = a.client.UpdateAlert(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &alert_client.CreateAlertRequest{
			Alert: res,
		}
		resp, err = a.client.CreateAlert(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiAlertAccess) DeleteAlert(ctx context.Context, ref *alert.Reference, opts ...gotenresource.DeleteOption) error {
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &alert_client.DeleteAlertRequest{
		Name: &ref.Name,
	}
	_, err := a.client.DeleteAlert(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *alert.Filter) (*alert.Filter, *alert.ParentName) {
	var withParentExtraction func(cnd alert.FilterCondition) alert.FilterCondition
	var resultParent *alert.ParentName
	var resultFilter *alert.Filter
	withParentExtraction = func(cnd alert.FilterCondition) alert.FilterCondition {
		switch tCnd := cnd.(type) {
		case *alert.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]alert.FilterCondition, 0)
				for _, subCnd := range tCnd.Conditions {
					if subCndNoParent := withParentExtraction(subCnd); subCndNoParent != nil {
						withoutParentCnds = append(withoutParentCnds, subCndNoParent)
					}
				}
				if len(withoutParentCnds) == 0 {
					return nil
				}
				return alert.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *alert.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*alert.Name)
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
		resultFilter = &alert.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(alert.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return alert.AsAnyCastAccess(NewApiAlertAccess(alert_client.NewAlertServiceClient(cc)))
	})
}
