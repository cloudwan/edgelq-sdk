// Code generated by protoc-gen-goten-resource
// Resource change: ProjectChange
// DO NOT EDIT!!!

package project

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &fieldmaskpb.FieldMask{}
)

func (c *ProjectChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProjectChange_Added_)
	return ok
}

func (c *ProjectChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProjectChange_Modified_)
	return ok
}

func (c *ProjectChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProjectChange_Current_)
	return ok
}

func (c *ProjectChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProjectChange_Removed_)
	return ok
}

func (c *ProjectChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ProjectChange_Added_:
		return cType.Added.ViewIndex
	case *ProjectChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *ProjectChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ProjectChange_Removed_:
		return cType.Removed.ViewIndex
	case *ProjectChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *ProjectChange) GetProject() *Project {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ProjectChange_Added_:
		return cType.Added.Project
	case *ProjectChange_Modified_:
		return cType.Modified.Project
	case *ProjectChange_Current_:
		return cType.Current.Project
	case *ProjectChange_Removed_:
		return nil
	}
	return nil
}

func (c *ProjectChange) GetRawResource() gotenresource.Resource {
	return c.GetProject()
}

func (c *ProjectChange) GetProjectName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ProjectChange_Added_:
		return cType.Added.Project.GetName()
	case *ProjectChange_Modified_:
		return cType.Modified.Name
	case *ProjectChange_Current_:
		return cType.Current.Project.GetName()
	case *ProjectChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *ProjectChange) GetRawName() gotenresource.Name {
	return c.GetProjectName()
}

func (c *ProjectChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &ProjectChange_Added_{
		Added: &ProjectChange_Added{
			Project:   snapshot.(*Project),
			ViewIndex: int32(idx),
		},
	}
}

func (c *ProjectChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &ProjectChange_Modified_{
		Modified: &ProjectChange_Modified{
			Name:              name.(*Name),
			Project:           snapshot.(*Project),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *ProjectChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &ProjectChange_Current_{
		Current: &ProjectChange_Current{
			Project: snapshot.(*Project),
		},
	}
}

func (c *ProjectChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &ProjectChange_Removed_{
		Removed: &ProjectChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
