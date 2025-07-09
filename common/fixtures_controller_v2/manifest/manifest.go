package manifest

import (
	"fmt"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
	"github.com/cloudwan/goten-sdk/runtime/resource/sharding"
)

type ManifestDependency struct {
	Name            string
	ApiAccess       *gotenaccess.ApiAccessBuilder
	Filter          gotenresource.Filter
	Parent          gotenresource.Name
	ExcFilter       gotenresource.Filter
	ShardingLabel   string
	DisableSharding bool
}

type ValueRepl func(dst gotenresource.Resource, depGetter func(name string) gotenresource.Resource)

type Manifest struct {
	ApiAccess          *gotenaccess.ApiAccessBuilder
	GroupName          string
	Filter             gotenresource.Filter
	Parent             gotenresource.Name
	ShardingLabel      string
	ShardingDec        *sharding.Decorator
	RestrictedRegionId string
	CreateForEach      []*ManifestDependency
	CommonReplacements []ValueRepl
	Resources          []struct {
		Resource          gotenresource.Resource
		LocalReplacements []ValueRepl
	}
	DiffMask gotenobject.FieldMask
}

func (ds *ManifestDependency) GetUniqueKey() string {
	typeName := ds.ApiAccess.GetDescriptor().GetResourceTypeName()
	filterStr := "<nil>"
	if ds.Filter != nil {
		filterStr = ds.Filter.String()
	}
	parentStr := "<nil>"
	if ds.Parent != nil {
		parentStr = ds.Parent.String()
	}
	excFilterStr := "<nil>"
	if ds.ExcFilter != nil {
		excFilterStr = ds.ExcFilter.String()
	}
	return fmt.Sprintf("%v.%s.%s.%s.%s.%s.%s.%s",
		ds.DisableSharding, typeName.FullyQualifiedTypeName(), typeName.Version(), ds.Name, filterStr, parentStr, excFilterStr, ds.ShardingLabel)
}
