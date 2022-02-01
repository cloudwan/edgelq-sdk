// Code generated by protoc-gen-goten-resource
// Resource: ConfigMap
// DO NOT EDIT!!!

package config_map

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
	project "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha/project"
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
)

// ensure the imports are used
var (
	_ = context.Context(nil)

	_ = codes.Internal
	_ = status.Status{}

	_ = watch_type.WatchType_STATEFUL
	_ = gotenobject.FieldPath(nil)
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
	_ = &ntt_meta.Meta{}
)

type ConfigMapAccess interface {
	GetConfigMap(context.Context, *GetQuery) (*ConfigMap, error)
	BatchGetConfigMaps(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryConfigMaps(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchConfigMap(context.Context, *GetQuery, func(*ConfigMapChange) error) error
	WatchConfigMaps(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveConfigMap(context.Context, *ConfigMap, ...gotenresource.SaveOption) error
	DeleteConfigMap(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	ConfigMapAccess
}

func AsAnyCastAccess(access ConfigMapAccess) gotenresource.Access {
	return &anyCastAccess{ConfigMapAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asConfigMapQuery, ok := q.(*GetQuery); ok {
		return a.GetConfigMap(ctx, asConfigMapQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ConfigMap, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asConfigMapQuery, ok := q.(*ListQuery); ok {
		return a.QueryConfigMaps(ctx, asConfigMapQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ConfigMap, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for ConfigMap")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asConfigMapQuery, ok := q.(*GetQuery); ok {
		return a.WatchConfigMap(ctx, asConfigMapQuery, func(change *ConfigMapChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ConfigMap, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asConfigMapQuery, ok := q.(*WatchQuery); ok {
		return a.WatchConfigMaps(ctx, asConfigMapQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ConfigMap, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asConfigMapRes, ok := res.(*ConfigMap); ok {
		return a.SaveConfigMap(ctx, asConfigMapRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ConfigMap, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asConfigMapRef, ok := ref.(*Reference); ok {
		return a.DeleteConfigMap(ctx, asConfigMapRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ConfigMap, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	configMapRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asConfigMapRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected ConfigMap, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			configMapRefs = append(configMapRefs, asConfigMapRef)
		}
	}
	return a.BatchGetConfigMaps(ctx, configMapRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
