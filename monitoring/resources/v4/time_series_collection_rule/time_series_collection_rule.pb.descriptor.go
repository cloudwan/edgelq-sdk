// Code generated by protoc-gen-goten-resource
// Resource: TimeSeriesCollectionRule
// DO NOT EDIT!!!

package time_series_collection_rule

import (
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/common"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/project"
	time_serie "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/time_serie"
	time_series_forwarder_sink "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/time_series_forwarder_sink"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(gotenobject.FieldPath)
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &common.LabelDescriptor{}
	_ = &project.Project{}
	_ = &time_serie.Point{}
	_ = &time_series_forwarder_sink.TimeSeriesForwarderSink{}
	_ = &meta.Meta{}
)

var (
	descriptor *Descriptor
)

func (r *TimeSeriesCollectionRule) GetRawName() gotenresource.Name {
	return r.GetName()
}

func (r *TimeSeriesCollectionRule) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

func (r *TimeSeriesCollectionRule) EnsureMetadata() *meta.Meta {
	if r.Metadata == nil {
		r.Metadata = &meta.Meta{}
	}
	if r.Metadata.Lifecycle == nil {
		r.Metadata.Lifecycle = &meta.Lifecycle{}
	}
	return r.Metadata
}

type Descriptor struct {
	nameDescriptor *gotenresource.NameDescriptor
	typeName       *gotenresource.TypeName
}

func GetDescriptor() *Descriptor {
	return descriptor
}

func (d *Descriptor) NewResource() gotenresource.Resource {
	return &TimeSeriesCollectionRule{}
}

func (d *Descriptor) NewResourceName() gotenresource.Name {
	return NewNameBuilder().Name()
}

func (d *Descriptor) NewGetQuery() gotenresource.GetQuery {
	return &GetQuery{}
}

func (d *Descriptor) NewListQuery() gotenresource.ListQuery {
	return &ListQuery{}
}

func (d *Descriptor) NewSearchQuery() gotenresource.SearchQuery {
	return nil
}

func (d *Descriptor) NewWatchQuery() gotenresource.WatchQuery {
	return &WatchQuery{}
}

func (d *Descriptor) NewResourceCursor() gotenresource.Cursor {
	return &PagerCursor{}
}

func (d *Descriptor) NewResourceFilter() gotenresource.Filter {
	return &Filter{}
}

func (d *Descriptor) NewResourceOrderBy() gotenresource.OrderBy {
	return &OrderBy{}
}

func (d *Descriptor) NewResourcePager() gotenresource.PagerQuery {
	return MakePagerQuery(nil, nil, 100, true)
}

func (d *Descriptor) NewResourceFieldMask() gotenobject.FieldMask {
	return &TimeSeriesCollectionRule_FieldMask{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return &TimeSeriesCollectionRuleChange{}
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return &QueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return &QueryResultChange{}
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(TimeSeriesCollectionRuleList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return make(TimeSeriesCollectionRuleChangeList, size, reserved)
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(TimeSeriesCollectionRuleNameList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(TimeSeriesCollectionRuleReferenceList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return make(TimeSeriesCollectionRuleParentNameList, size, reserved)
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return make(TimeSeriesCollectionRuleParentReferenceList, size, reserved)
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(TimeSeriesCollectionRuleMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return make(TimeSeriesCollectionRuleChangeMap, reserved)
}

func (d *Descriptor) GetResourceTypeName() *gotenresource.TypeName {
	return d.typeName
}

func (d *Descriptor) GetNameDescriptor() *gotenresource.NameDescriptor {
	return d.nameDescriptor
}

func (d *Descriptor) CanBeParentless() bool {
	return false
}

func (d *Descriptor) GetParentResDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		project.GetDescriptor(),
	}
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParseTimeSeriesCollectionRule_FieldPath(raw)
}

func (d *Descriptor) ParseResourceName(nameStr string) (gotenresource.Name, error) {
	return ParseName(nameStr)
}

func (d *Descriptor) SupportsMetadata() bool {
	return true
}

func (d *Descriptor) SupportsDbConstraints() bool {
	return true
}

func initTimeSeriesCollectionRuleDescriptor() {
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"TimeSeriesCollectionRule", "TimeSeriesCollectionRules", "monitoring.edgelq.com", "v4"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&TimeSeriesCollectionRule_FieldTerminalPath{selector: TimeSeriesCollectionRule_FieldPathSelectorName},
			"pattern", "timeSeriesCollectionRuleId",
			[]string{"projectId"},
			[]gotenresource.NamePattern{NamePattern_Project}),
	}
	gotenresource.GetRegistry().RegisterDescriptor(descriptor)
}

func init() {
	initTimeSeriesCollectionRuleDescriptor()
}