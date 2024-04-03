// Code generated by protoc-gen-goten-access
// Resource: TpmAttestationCert
// DO NOT EDIT!!!

package tpm_attestation_cert_access

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

	tpm_attestation_cert_client "github.com/cloudwan/edgelq-sdk/ztp/client/v1/tpm_attestation_cert"
	tpm_attestation_cert "github.com/cloudwan/edgelq-sdk/ztp/resources/v1/tpm_attestation_cert"
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

type apiTpmAttestationCertAccess struct {
	client tpm_attestation_cert_client.TpmAttestationCertServiceClient
}

func NewApiTpmAttestationCertAccess(client tpm_attestation_cert_client.TpmAttestationCertServiceClient) tpm_attestation_cert.TpmAttestationCertAccess {
	return &apiTpmAttestationCertAccess{client: client}
}

func (a *apiTpmAttestationCertAccess) GetTpmAttestationCert(ctx context.Context, query *tpm_attestation_cert.GetQuery) (*tpm_attestation_cert.TpmAttestationCert, error) {
	if !query.Reference.IsFullyQualified() {
		return nil, status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &tpm_attestation_cert_client.GetTpmAttestationCertRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetTpmAttestationCert(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiTpmAttestationCertAccess) BatchGetTpmAttestationCerts(ctx context.Context, refs []*tpm_attestation_cert.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	asNames := make([]*tpm_attestation_cert.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &tpm_attestation_cert_client.BatchGetTpmAttestationCertsRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(tpm_attestation_cert.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*tpm_attestation_cert.TpmAttestationCert_FieldMask)
	}
	resp, err := a.client.BatchGetTpmAttestationCerts(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[tpm_attestation_cert.Name]*tpm_attestation_cert.TpmAttestationCert, len(refs))
	for _, resolvedRes := range resp.GetTpmAttestationCerts() {
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

func (a *apiTpmAttestationCertAccess) QueryTpmAttestationCerts(ctx context.Context, query *tpm_attestation_cert.ListQuery) (*tpm_attestation_cert.QueryResultSnapshot, error) {
	request := &tpm_attestation_cert_client.ListTpmAttestationCertsRequest{
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
	resp, err := a.client.ListTpmAttestationCerts(ctx, request)
	if err != nil {
		return nil, err
	}
	return &tpm_attestation_cert.QueryResultSnapshot{
		TpmAttestationCerts: resp.TpmAttestationCerts,
		NextPageCursor:      resp.NextPageToken,
		PrevPageCursor:      resp.PrevPageToken,
		TotalResultsCount:   resp.TotalResultsCount,
		CurrentOffset:       resp.CurrentOffset,
	}, nil
}

func (a *apiTpmAttestationCertAccess) WatchTpmAttestationCert(ctx context.Context, query *tpm_attestation_cert.GetQuery, observerCb func(*tpm_attestation_cert.TpmAttestationCertChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &tpm_attestation_cert_client.WatchTpmAttestationCertRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchTpmAttestationCert(ctx, request)
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

func (a *apiTpmAttestationCertAccess) WatchTpmAttestationCerts(ctx context.Context, query *tpm_attestation_cert.WatchQuery, observerCb func(*tpm_attestation_cert.QueryResultChange) error) error {
	request := &tpm_attestation_cert_client.WatchTpmAttestationCertsRequest{
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
	changesStream, initErr := a.client.WatchTpmAttestationCerts(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &tpm_attestation_cert.QueryResultChange{
			Changes:      respChange.TpmAttestationCertChanges,
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

func (a *apiTpmAttestationCertAccess) SaveTpmAttestationCert(ctx context.Context, res *tpm_attestation_cert.TpmAttestationCert, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetTpmAttestationCert(ctx, &tpm_attestation_cert.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}
	var resp *tpm_attestation_cert.TpmAttestationCert
	var err error
	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &tpm_attestation_cert_client.UpdateTpmAttestationCertRequest{
			TpmAttestationCert: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*tpm_attestation_cert.TpmAttestationCert_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &tpm_attestation_cert_client.UpdateTpmAttestationCertRequest_CAS{
				ConditionalState: conditionalState.(*tpm_attestation_cert.TpmAttestationCert),
				FieldMask:        mask.(*tpm_attestation_cert.TpmAttestationCert_FieldMask),
			}
		}
		resp, err = a.client.UpdateTpmAttestationCert(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &tpm_attestation_cert_client.CreateTpmAttestationCertRequest{
			TpmAttestationCert: res,
		}
		resp, err = a.client.CreateTpmAttestationCert(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiTpmAttestationCertAccess) DeleteTpmAttestationCert(ctx context.Context, ref *tpm_attestation_cert.Reference, opts ...gotenresource.DeleteOption) error {
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &tpm_attestation_cert_client.DeleteTpmAttestationCertRequest{
		Name: &ref.Name,
	}
	_, err := a.client.DeleteTpmAttestationCert(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *tpm_attestation_cert.Filter) (*tpm_attestation_cert.Filter, *tpm_attestation_cert.ParentName) {
	var withParentExtraction func(cnd tpm_attestation_cert.FilterCondition) tpm_attestation_cert.FilterCondition
	var resultParent *tpm_attestation_cert.ParentName
	var resultFilter *tpm_attestation_cert.Filter
	withParentExtraction = func(cnd tpm_attestation_cert.FilterCondition) tpm_attestation_cert.FilterCondition {
		switch tCnd := cnd.(type) {
		case *tpm_attestation_cert.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]tpm_attestation_cert.FilterCondition, 0)
				for _, subCnd := range tCnd.Conditions {
					if subCndNoParent := withParentExtraction(subCnd); subCndNoParent != nil {
						withoutParentCnds = append(withoutParentCnds, subCndNoParent)
					}
				}
				if len(withoutParentCnds) == 0 {
					return nil
				}
				return tpm_attestation_cert.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *tpm_attestation_cert.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*tpm_attestation_cert.Name)
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
		resultFilter = &tpm_attestation_cert.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(tpm_attestation_cert.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return tpm_attestation_cert.AsAnyCastAccess(NewApiTpmAttestationCertAccess(tpm_attestation_cert_client.NewTpmAttestationCertServiceClient(cc)))
	})
}