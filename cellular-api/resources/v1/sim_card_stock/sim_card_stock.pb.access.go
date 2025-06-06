// Code generated by protoc-gen-goten-resource
// Resource: SimCardStock
// DO NOT EDIT!!!

package sim_card_stock

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
	"github.com/cloudwan/goten-sdk/types/watch_type"
)

// proto imports
import (
	contract "github.com/cloudwan/edgelq-sdk/cellular-api/resources/v1/contract"
	sim_card "github.com/cloudwan/edgelq-sdk/cellular-api/resources/v1/sim_card"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(context.Context)

	_ = codes.Internal
	_ = status.Status{}

	_ = watch_type.WatchType_STATEFUL
	_ = new(gotenobject.FieldPath)
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &contract.Contract{}
	_ = &sim_card.SimCard{}
	_ = &meta.Meta{}
)

type SimCardStockAccess interface {
	GetSimCardStock(context.Context, *GetQuery, ...gotenresource.GetOption) (*SimCardStock, error)
	BatchGetSimCardStocks(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QuerySimCardStocks(context.Context, *ListQuery, ...gotenresource.QueryOption) (*QueryResultSnapshot, error)
	WatchSimCardStock(context.Context, *GetQuery, func(*SimCardStockChange) error) error
	WatchSimCardStocks(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveSimCardStock(context.Context, *SimCardStock, ...gotenresource.SaveOption) error
	DeleteSimCardStock(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	SimCardStockAccess
}

func AsAnyCastAccess(access SimCardStockAccess) gotenresource.Access {
	return &anyCastAccess{SimCardStockAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery, opts ...gotenresource.GetOption) (gotenresource.Resource, error) {
	if asSimCardStockQuery, ok := q.(*GetQuery); ok {
		return a.GetSimCardStock(ctx, asSimCardStockQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected SimCardStock, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	if asSimCardStockQuery, ok := q.(*ListQuery); ok {
		return a.QuerySimCardStocks(ctx, asSimCardStockQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected SimCardStock, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for SimCardStock")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asSimCardStockQuery, ok := q.(*GetQuery); ok {
		return a.WatchSimCardStock(ctx, asSimCardStockQuery, func(change *SimCardStockChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected SimCardStock, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asSimCardStockQuery, ok := q.(*WatchQuery); ok {
		return a.WatchSimCardStocks(ctx, asSimCardStockQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected SimCardStock, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asSimCardStockRes, ok := res.(*SimCardStock); ok {
		return a.SaveSimCardStock(ctx, asSimCardStockRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected SimCardStock, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asSimCardStockRef, ok := ref.(*Reference); ok {
		return a.DeleteSimCardStock(ctx, asSimCardStockRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected SimCardStock, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	simCardStockRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asSimCardStockRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected SimCardStock, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			simCardStockRefs = append(simCardStockRefs, asSimCardStockRef)
		}
	}
	return a.BatchGetSimCardStocks(ctx, simCardStockRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
