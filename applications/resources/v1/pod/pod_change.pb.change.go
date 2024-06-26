// Code generated by protoc-gen-goten-resource
// Resource change: PodChange
// DO NOT EDIT!!!

package pod

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/applications/resources/v1/project"
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

func (c *PodChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*PodChange_Added_)
	return ok
}

func (c *PodChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*PodChange_Modified_)
	return ok
}

func (c *PodChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*PodChange_Current_)
	return ok
}

func (c *PodChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*PodChange_Removed_)
	return ok
}

func (c *PodChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *PodChange_Added_:
		return cType.Added.ViewIndex
	case *PodChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *PodChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *PodChange_Removed_:
		return cType.Removed.ViewIndex
	case *PodChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *PodChange) GetPod() *Pod {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *PodChange_Added_:
		return cType.Added.Pod
	case *PodChange_Modified_:
		return cType.Modified.Pod
	case *PodChange_Current_:
		return cType.Current.Pod
	case *PodChange_Removed_:
		return nil
	}
	return nil
}

func (c *PodChange) GetRawResource() gotenresource.Resource {
	return c.GetPod()
}

func (c *PodChange) GetPodName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *PodChange_Added_:
		return cType.Added.Pod.GetName()
	case *PodChange_Modified_:
		return cType.Modified.Name
	case *PodChange_Current_:
		return cType.Current.Pod.GetName()
	case *PodChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *PodChange) GetRawName() gotenresource.Name {
	return c.GetPodName()
}

func (c *PodChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &PodChange_Added_{
		Added: &PodChange_Added{
			Pod:       snapshot.(*Pod),
			ViewIndex: int32(idx),
		},
	}
}

func (c *PodChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &PodChange_Modified_{
		Modified: &PodChange_Modified{
			Name:              name.(*Name),
			Pod:               snapshot.(*Pod),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *PodChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &PodChange_Current_{
		Current: &PodChange_Current{
			Pod: snapshot.(*Pod),
		},
	}
}

func (c *PodChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &PodChange_Removed_{
		Removed: &PodChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
