// Code generated by protoc-gen-goten-object
// File: edgelq/applications/proto/v1/config_map.proto
// DO NOT EDIT!!!

package config_map

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/applications/resources/v1/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(fmt.Stringer)
	_ = new(sort.Interface)

	_ = new(proto.Message)
	_ = googlefieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
	_ = &meta.Meta{}
)

func (o *ConfigMap) GotenObjectExt() {}

func (o *ConfigMap) MakeFullFieldMask() *ConfigMap_FieldMask {
	return FullConfigMap_FieldMask()
}

func (o *ConfigMap) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullConfigMap_FieldMask()
}

func (o *ConfigMap) MakeDiffFieldMask(other *ConfigMap) *ConfigMap_FieldMask {
	if o == nil && other == nil {
		return &ConfigMap_FieldMask{}
	}
	if o == nil || other == nil {
		return FullConfigMap_FieldMask()
	}

	res := &ConfigMap_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ConfigMap_FieldSubPath{selector: ConfigMap_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorDisplayName})
	}
	if o.GetDescription() != other.GetDescription() {
		res.Paths = append(res.Paths, &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorDescription})
	}

	if len(o.GetData()) == len(other.GetData()) {
		for i, lValue := range o.GetData() {
			rValue := other.GetData()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorData})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorData})
	}

	if len(o.GetBinaryData()) == len(other.GetBinaryData()) {
		for i, lValue := range o.GetBinaryData() {
			rValue := other.GetBinaryData()[i]
			if string(lValue) != string(rValue) {
				res.Paths = append(res.Paths, &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorBinaryData})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorBinaryData})
	}
	return res
}

func (o *ConfigMap) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ConfigMap))
}

func (o *ConfigMap) Clone() *ConfigMap {
	if o == nil {
		return nil
	}
	result := &ConfigMap{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &Name{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Metadata = o.Metadata.Clone()
	result.DisplayName = o.DisplayName
	result.Description = o.Description
	result.Data = map[string]string{}
	for key, sourceValue := range o.Data {
		result.Data[key] = sourceValue
	}
	result.BinaryData = map[string][]byte{}
	for key, sourceValue := range o.BinaryData {
		result.BinaryData[key] = sourceValue
	}
	return result
}

func (o *ConfigMap) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ConfigMap) Merge(source *ConfigMap) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &Name{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
	o.DisplayName = source.GetDisplayName()
	o.Description = source.GetDescription()
	if source.GetData() != nil {
		if o.Data == nil {
			o.Data = make(map[string]string, len(source.GetData()))
		}
		for key, sourceValue := range source.GetData() {
			o.Data[key] = sourceValue
		}
	}
	if source.GetBinaryData() != nil {
		if o.BinaryData == nil {
			o.BinaryData = make(map[string][]byte, len(source.GetBinaryData()))
		}
		for key, sourceValue := range source.GetBinaryData() {
			o.BinaryData[key] = sourceValue
		}
	}
}

func (o *ConfigMap) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ConfigMap))
}
