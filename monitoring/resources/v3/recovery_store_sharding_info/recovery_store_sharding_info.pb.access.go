// Code generated by protoc-gen-goten-resource
// Resource: RecoveryStoreShardingInfo
// DO NOT EDIT!!!

package recovery_store_sharding_info

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	_ = &ntt_meta.Meta{}
	_ = &duration.Duration{}
	_ = &timestamp.Timestamp{}
)

type RecoveryStoreShardingInfoAccess interface {
	GetRecoveryStoreShardingInfo(context.Context, *GetQuery) (*RecoveryStoreShardingInfo, error)
	BatchGetRecoveryStoreShardingInfos(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryRecoveryStoreShardingInfos(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchRecoveryStoreShardingInfo(context.Context, *GetQuery, func(*RecoveryStoreShardingInfoChange) error) error
	WatchRecoveryStoreShardingInfos(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveRecoveryStoreShardingInfo(context.Context, *RecoveryStoreShardingInfo, ...gotenresource.SaveOption) error
	DeleteRecoveryStoreShardingInfo(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	RecoveryStoreShardingInfoAccess
}

func AsAnyCastAccess(access RecoveryStoreShardingInfoAccess) gotenresource.Access {
	return &anyCastAccess{RecoveryStoreShardingInfoAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asRecoveryStoreShardingInfoQuery, ok := q.(*GetQuery); ok {
		return a.GetRecoveryStoreShardingInfo(ctx, asRecoveryStoreShardingInfoQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected RecoveryStoreShardingInfo, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asRecoveryStoreShardingInfoQuery, ok := q.(*ListQuery); ok {
		return a.QueryRecoveryStoreShardingInfos(ctx, asRecoveryStoreShardingInfoQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected RecoveryStoreShardingInfo, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.QueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for RecoveryStoreShardingInfo")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asRecoveryStoreShardingInfoQuery, ok := q.(*GetQuery); ok {
		return a.WatchRecoveryStoreShardingInfo(ctx, asRecoveryStoreShardingInfoQuery, func(change *RecoveryStoreShardingInfoChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected RecoveryStoreShardingInfo, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asRecoveryStoreShardingInfoQuery, ok := q.(*WatchQuery); ok {
		return a.WatchRecoveryStoreShardingInfos(ctx, asRecoveryStoreShardingInfoQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected RecoveryStoreShardingInfo, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asRecoveryStoreShardingInfoRes, ok := res.(*RecoveryStoreShardingInfo); ok {
		return a.SaveRecoveryStoreShardingInfo(ctx, asRecoveryStoreShardingInfoRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected RecoveryStoreShardingInfo, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asRecoveryStoreShardingInfoRef, ok := ref.(*Reference); ok {
		return a.DeleteRecoveryStoreShardingInfo(ctx, asRecoveryStoreShardingInfoRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected RecoveryStoreShardingInfo, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	recoveryStoreShardingInfoRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asRecoveryStoreShardingInfoRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected RecoveryStoreShardingInfo, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			recoveryStoreShardingInfoRefs = append(recoveryStoreShardingInfoRefs, asRecoveryStoreShardingInfoRef)
		}
	}
	return a.BatchGetRecoveryStoreShardingInfos(ctx, recoveryStoreShardingInfoRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}