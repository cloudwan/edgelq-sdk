// Code generated by protoc-gen-goten-resource
// Resource change: TsEntryChange
// DO NOT EDIT!!!

package ts_entry

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ts_condition "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/ts_condition"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &ts_condition.TsCondition{}
	_ = &fieldmaskpb.FieldMask{}
)

func (c *TsEntryChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*TsEntryChange_Added_)
	return ok
}

func (c *TsEntryChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*TsEntryChange_Modified_)
	return ok
}

func (c *TsEntryChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*TsEntryChange_Current_)
	return ok
}

func (c *TsEntryChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*TsEntryChange_Removed_)
	return ok
}

func (c *TsEntryChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *TsEntryChange_Added_:
		return cType.Added.ViewIndex
	case *TsEntryChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *TsEntryChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *TsEntryChange_Removed_:
		return cType.Removed.ViewIndex
	case *TsEntryChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *TsEntryChange) GetTsEntry() *TsEntry {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *TsEntryChange_Added_:
		return cType.Added.TsEntry
	case *TsEntryChange_Modified_:
		return cType.Modified.TsEntry
	case *TsEntryChange_Current_:
		return cType.Current.TsEntry
	case *TsEntryChange_Removed_:
		return nil
	}
	return nil
}

func (c *TsEntryChange) GetRawResource() gotenresource.Resource {
	return c.GetTsEntry()
}

func (c *TsEntryChange) GetTsEntryName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *TsEntryChange_Added_:
		return cType.Added.TsEntry.GetName()
	case *TsEntryChange_Modified_:
		return cType.Modified.Name
	case *TsEntryChange_Current_:
		return cType.Current.TsEntry.GetName()
	case *TsEntryChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *TsEntryChange) GetRawName() gotenresource.Name {
	return c.GetTsEntryName()
}

func (c *TsEntryChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &TsEntryChange_Added_{
		Added: &TsEntryChange_Added{
			TsEntry:   snapshot.(*TsEntry),
			ViewIndex: int32(idx),
		},
	}
}

func (c *TsEntryChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &TsEntryChange_Modified_{
		Modified: &TsEntryChange_Modified{
			Name:              name.(*Name),
			TsEntry:           snapshot.(*TsEntry),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *TsEntryChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &TsEntryChange_Current_{
		Current: &TsEntryChange_Current{
			TsEntry: snapshot.(*TsEntry),
		},
	}
}

func (c *TsEntryChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &TsEntryChange_Removed_{
		Removed: &TsEntryChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
