// Code generated by protoc-gen-goten-resource
// Resource change: SimCardChange
// DO NOT EDIT!!!

package sim_card

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &iam_project.Project{}
	_ = &fieldmaskpb.FieldMask{}
)

func (c *SimCardChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*SimCardChange_Added_)
	return ok
}

func (c *SimCardChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*SimCardChange_Modified_)
	return ok
}

func (c *SimCardChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*SimCardChange_Current_)
	return ok
}

func (c *SimCardChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*SimCardChange_Removed_)
	return ok
}

func (c *SimCardChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *SimCardChange_Added_:
		return cType.Added.ViewIndex
	case *SimCardChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *SimCardChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *SimCardChange_Removed_:
		return cType.Removed.ViewIndex
	case *SimCardChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *SimCardChange) GetSimCard() *SimCard {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *SimCardChange_Added_:
		return cType.Added.SimCard
	case *SimCardChange_Modified_:
		return cType.Modified.SimCard
	case *SimCardChange_Current_:
		return cType.Current.SimCard
	case *SimCardChange_Removed_:
		return nil
	}
	return nil
}

func (c *SimCardChange) GetRawResource() gotenresource.Resource {
	return c.GetSimCard()
}

func (c *SimCardChange) GetSimCardName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *SimCardChange_Added_:
		return cType.Added.SimCard.GetName()
	case *SimCardChange_Modified_:
		return cType.Modified.Name
	case *SimCardChange_Current_:
		return cType.Current.SimCard.GetName()
	case *SimCardChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *SimCardChange) GetRawName() gotenresource.Name {
	return c.GetSimCardName()
}

func (c *SimCardChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &SimCardChange_Added_{
		Added: &SimCardChange_Added{
			SimCard:   snapshot.(*SimCard),
			ViewIndex: int32(idx),
		},
	}
}

func (c *SimCardChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &SimCardChange_Modified_{
		Modified: &SimCardChange_Modified{
			Name:              name.(*Name),
			SimCard:           snapshot.(*SimCard),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *SimCardChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &SimCardChange_Current_{
		Current: &SimCardChange_Current{
			SimCard: snapshot.(*SimCard),
		},
	}
}

func (c *SimCardChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &SimCardChange_Removed_{
		Removed: &SimCardChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
