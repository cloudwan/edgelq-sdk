// Code generated by protoc-gen-goten-object
// File: edgelq/meta/multi_region/proto/syncing_meta.proto
// DO NOT EDIT!!!

package syncing_meta

// proto imports
import ()

// make sure we're using proto imports
var ()

type SyncingMetaFieldPathBuilder struct{}

func NewSyncingMetaFieldPathBuilder() SyncingMetaFieldPathBuilder {
	return SyncingMetaFieldPathBuilder{}
}
func (SyncingMetaFieldPathBuilder) OwningRegion() SyncingMetaPathSelectorOwningRegion {
	return SyncingMetaPathSelectorOwningRegion{}
}
func (SyncingMetaFieldPathBuilder) Regions() SyncingMetaPathSelectorRegions {
	return SyncingMetaPathSelectorRegions{}
}

type SyncingMetaPathSelectorOwningRegion struct{}

func (SyncingMetaPathSelectorOwningRegion) FieldPath() *SyncingMeta_FieldTerminalPath {
	return &SyncingMeta_FieldTerminalPath{selector: SyncingMeta_FieldPathSelectorOwningRegion}
}

func (s SyncingMetaPathSelectorOwningRegion) WithValue(value string) *SyncingMeta_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*SyncingMeta_FieldTerminalPathValue)
}

func (s SyncingMetaPathSelectorOwningRegion) WithArrayOfValues(values []string) *SyncingMeta_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*SyncingMeta_FieldTerminalPathArrayOfValues)
}

type SyncingMetaPathSelectorRegions struct{}

func (SyncingMetaPathSelectorRegions) FieldPath() *SyncingMeta_FieldTerminalPath {
	return &SyncingMeta_FieldTerminalPath{selector: SyncingMeta_FieldPathSelectorRegions}
}

func (s SyncingMetaPathSelectorRegions) WithValue(value []string) *SyncingMeta_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*SyncingMeta_FieldTerminalPathValue)
}

func (s SyncingMetaPathSelectorRegions) WithArrayOfValues(values [][]string) *SyncingMeta_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*SyncingMeta_FieldTerminalPathArrayOfValues)
}

func (s SyncingMetaPathSelectorRegions) WithItemValue(value string) *SyncingMeta_FieldTerminalPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*SyncingMeta_FieldTerminalPathArrayItemValue)
}
