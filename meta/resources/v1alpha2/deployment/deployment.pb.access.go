// Code generated by protoc-gen-goten-resource
// Resource: Deployment
// DO NOT EDIT!!!

package deployment

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
	region "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/region"
	service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
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
	_ = &region.Region{}
	_ = &service.Service{}
	_ = &meta.Meta{}
)

type DeploymentAccess interface {
	GetDeployment(context.Context, *GetQuery) (*Deployment, error)
	BatchGetDeployments(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryDeployments(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchDeployment(context.Context, *GetQuery, func(*DeploymentChange) error) error
	WatchDeployments(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveDeployment(context.Context, *Deployment, ...gotenresource.SaveOption) error
	DeleteDeployment(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	DeploymentAccess
}

func AsAnyCastAccess(access DeploymentAccess) gotenresource.Access {
	return &anyCastAccess{DeploymentAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asDeploymentQuery, ok := q.(*GetQuery); ok {
		return a.GetDeployment(ctx, asDeploymentQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Deployment, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asDeploymentQuery, ok := q.(*ListQuery); ok {
		return a.QueryDeployments(ctx, asDeploymentQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Deployment, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.QueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for Deployment")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asDeploymentQuery, ok := q.(*GetQuery); ok {
		return a.WatchDeployment(ctx, asDeploymentQuery, func(change *DeploymentChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Deployment, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asDeploymentQuery, ok := q.(*WatchQuery); ok {
		return a.WatchDeployments(ctx, asDeploymentQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Deployment, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asDeploymentRes, ok := res.(*Deployment); ok {
		return a.SaveDeployment(ctx, asDeploymentRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Deployment, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asDeploymentRef, ok := ref.(*Reference); ok {
		return a.DeleteDeployment(ctx, asDeploymentRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Deployment, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	deploymentRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asDeploymentRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Deployment, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			deploymentRefs = append(deploymentRefs, asDeploymentRef)
		}
	}
	return a.BatchGetDeployments(ctx, deploymentRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
