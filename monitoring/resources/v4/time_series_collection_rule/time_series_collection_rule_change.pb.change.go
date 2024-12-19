// Code generated by protoc-gen-goten-resource
// Resource change: TimeSeriesCollectionRuleChange
// DO NOT EDIT!!!

package time_series_collection_rule

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/project"
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

func (c *TimeSeriesCollectionRuleChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*TimeSeriesCollectionRuleChange_Added_)
	return ok
}

func (c *TimeSeriesCollectionRuleChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*TimeSeriesCollectionRuleChange_Modified_)
	return ok
}

func (c *TimeSeriesCollectionRuleChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*TimeSeriesCollectionRuleChange_Current_)
	return ok
}

func (c *TimeSeriesCollectionRuleChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*TimeSeriesCollectionRuleChange_Removed_)
	return ok
}

func (c *TimeSeriesCollectionRuleChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *TimeSeriesCollectionRuleChange_Added_:
		return cType.Added.ViewIndex
	case *TimeSeriesCollectionRuleChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *TimeSeriesCollectionRuleChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *TimeSeriesCollectionRuleChange_Removed_:
		return cType.Removed.ViewIndex
	case *TimeSeriesCollectionRuleChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *TimeSeriesCollectionRuleChange) GetTimeSeriesCollectionRule() *TimeSeriesCollectionRule {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *TimeSeriesCollectionRuleChange_Added_:
		return cType.Added.TimeSeriesCollectionRule
	case *TimeSeriesCollectionRuleChange_Modified_:
		return cType.Modified.TimeSeriesCollectionRule
	case *TimeSeriesCollectionRuleChange_Current_:
		return cType.Current.TimeSeriesCollectionRule
	case *TimeSeriesCollectionRuleChange_Removed_:
		return nil
	}
	return nil
}

func (c *TimeSeriesCollectionRuleChange) GetRawResource() gotenresource.Resource {
	return c.GetTimeSeriesCollectionRule()
}

func (c *TimeSeriesCollectionRuleChange) GetTimeSeriesCollectionRuleName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *TimeSeriesCollectionRuleChange_Added_:
		return cType.Added.TimeSeriesCollectionRule.GetName()
	case *TimeSeriesCollectionRuleChange_Modified_:
		return cType.Modified.Name
	case *TimeSeriesCollectionRuleChange_Current_:
		return cType.Current.TimeSeriesCollectionRule.GetName()
	case *TimeSeriesCollectionRuleChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *TimeSeriesCollectionRuleChange) GetRawName() gotenresource.Name {
	return c.GetTimeSeriesCollectionRuleName()
}

func (c *TimeSeriesCollectionRuleChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &TimeSeriesCollectionRuleChange_Added_{
		Added: &TimeSeriesCollectionRuleChange_Added{
			TimeSeriesCollectionRule: snapshot.(*TimeSeriesCollectionRule),
			ViewIndex:                int32(idx),
		},
	}
}

func (c *TimeSeriesCollectionRuleChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &TimeSeriesCollectionRuleChange_Modified_{
		Modified: &TimeSeriesCollectionRuleChange_Modified{
			Name:                     name.(*Name),
			TimeSeriesCollectionRule: snapshot.(*TimeSeriesCollectionRule),
			PreviousViewIndex:        int32(prevIdx),
			ViewIndex:                int32(newIdx),
		},
	}
}

func (c *TimeSeriesCollectionRuleChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &TimeSeriesCollectionRuleChange_Current_{
		Current: &TimeSeriesCollectionRuleChange_Current{
			TimeSeriesCollectionRule: snapshot.(*TimeSeriesCollectionRule),
		},
	}
}

func (c *TimeSeriesCollectionRuleChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &TimeSeriesCollectionRuleChange_Removed_{
		Removed: &TimeSeriesCollectionRuleChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}