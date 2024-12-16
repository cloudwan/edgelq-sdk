// Code generated by protoc-gen-goten-resource
// Resource: TimeSeriesCollectionRule
// DO NOT EDIT!!!

package time_series_collection_rule

import (
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

// make sure we're using proto imports
var (
	_ = &common.LabelDescriptor{}
	_ = &project.Project{}
	_ = &time_serie.Point{}
	_ = &time_series_forwarder_sink.TimeSeriesForwarderSink{}
	_ = &meta.Meta{}
)

const (
	NamePattern_Project = "projects/{project}/timeSeriesCollectionRules/{time_series_collection_rule}"
)

type NamePattern struct {
	Pattern gotenresource.NamePattern `firestore:"pattern"`
}

type NameBuilder struct {
	nameObj Name
}

func NewNameBuilder() *NameBuilder {
	return &NameBuilder{
		nameObj: Name{
			TimeSeriesCollectionRuleId: gotenresource.WildcardId,
			ParentName: ParentName{
				NamePattern: NamePattern{
					// Set default pattern - just first.
					Pattern: NamePattern_Project,
				},
			},
		},
	}
}

func (b *NameBuilder) Name() *Name {
	copied := b.nameObj
	return &copied
}

func (b *NameBuilder) Reference() *Reference {
	return b.nameObj.AsReference()
}

func (b *NameBuilder) Parent() *ParentName {
	copied := b.nameObj.ParentName
	return &copied
}

func (b *NameBuilder) ParentReference() *ParentReference {
	return b.nameObj.ParentName.AsReference()
}

func (b *NameBuilder) SetId(id string) *NameBuilder {
	b.nameObj.TimeSeriesCollectionRuleId = id
	return b
}

func (b *NameBuilder) SetProject(parent *project.Name) *NameBuilder {
	parentName := &b.nameObj.ParentName

	switch parent.Pattern {
	case project.NamePattern_Nil:
		parentName.Pattern = NamePattern_Project
	}
	parentName.ProjectId = parent.ProjectId
	return b
}

func (b *NameBuilder) SetProjectId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.ProjectId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" {
		parentName.Pattern = NamePattern_Project
	}
	return b
}
