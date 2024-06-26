// Code generated by protoc-gen-goten-resource
// Resource change: ConfigMapChange
// DO NOT EDIT!!!

package config_map

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha2/project"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
	_ = &fieldmaskpb.FieldMask{}
)

func (c *ConfigMapChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ConfigMapChange_Added_)
	return ok
}

func (c *ConfigMapChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ConfigMapChange_Modified_)
	return ok
}

func (c *ConfigMapChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ConfigMapChange_Current_)
	return ok
}

func (c *ConfigMapChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ConfigMapChange_Removed_)
	return ok
}

func (c *ConfigMapChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ConfigMapChange_Added_:
		return cType.Added.ViewIndex
	case *ConfigMapChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *ConfigMapChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ConfigMapChange_Removed_:
		return cType.Removed.ViewIndex
	case *ConfigMapChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *ConfigMapChange) GetConfigMap() *ConfigMap {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ConfigMapChange_Added_:
		return cType.Added.ConfigMap
	case *ConfigMapChange_Modified_:
		return cType.Modified.ConfigMap
	case *ConfigMapChange_Current_:
		return cType.Current.ConfigMap
	case *ConfigMapChange_Removed_:
		return nil
	}
	return nil
}

func (c *ConfigMapChange) GetRawResource() gotenresource.Resource {
	return c.GetConfigMap()
}

func (c *ConfigMapChange) GetConfigMapName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ConfigMapChange_Added_:
		return cType.Added.ConfigMap.GetName()
	case *ConfigMapChange_Modified_:
		return cType.Modified.Name
	case *ConfigMapChange_Current_:
		return cType.Current.ConfigMap.GetName()
	case *ConfigMapChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *ConfigMapChange) GetRawName() gotenresource.Name {
	return c.GetConfigMapName()
}

func (c *ConfigMapChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &ConfigMapChange_Added_{
		Added: &ConfigMapChange_Added{
			ConfigMap: snapshot.(*ConfigMap),
			ViewIndex: int32(idx),
		},
	}
}

func (c *ConfigMapChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &ConfigMapChange_Modified_{
		Modified: &ConfigMapChange_Modified{
			Name:              name.(*Name),
			ConfigMap:         snapshot.(*ConfigMap),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *ConfigMapChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &ConfigMapChange_Current_{
		Current: &ConfigMapChange_Current{
			ConfigMap: snapshot.(*ConfigMap),
		},
	}
}

func (c *ConfigMapChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &ConfigMapChange_Removed_{
		Removed: &ConfigMapChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
