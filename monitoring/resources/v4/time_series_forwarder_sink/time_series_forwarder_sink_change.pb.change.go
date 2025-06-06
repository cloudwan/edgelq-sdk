// Code generated by protoc-gen-goten-resource
// Resource change: TimeSeriesForwarderSinkChange
// DO NOT EDIT!!!

package time_series_forwarder_sink

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

func (c *TimeSeriesForwarderSinkChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*TimeSeriesForwarderSinkChange_Added_)
	return ok
}

func (c *TimeSeriesForwarderSinkChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*TimeSeriesForwarderSinkChange_Modified_)
	return ok
}

func (c *TimeSeriesForwarderSinkChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*TimeSeriesForwarderSinkChange_Current_)
	return ok
}

func (c *TimeSeriesForwarderSinkChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*TimeSeriesForwarderSinkChange_Removed_)
	return ok
}

func (c *TimeSeriesForwarderSinkChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *TimeSeriesForwarderSinkChange_Added_:
		return cType.Added.ViewIndex
	case *TimeSeriesForwarderSinkChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *TimeSeriesForwarderSinkChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *TimeSeriesForwarderSinkChange_Removed_:
		return cType.Removed.ViewIndex
	case *TimeSeriesForwarderSinkChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *TimeSeriesForwarderSinkChange) GetTimeSeriesForwarderSink() *TimeSeriesForwarderSink {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *TimeSeriesForwarderSinkChange_Added_:
		return cType.Added.TimeSeriesForwarderSink
	case *TimeSeriesForwarderSinkChange_Modified_:
		return cType.Modified.TimeSeriesForwarderSink
	case *TimeSeriesForwarderSinkChange_Current_:
		return cType.Current.TimeSeriesForwarderSink
	case *TimeSeriesForwarderSinkChange_Removed_:
		return nil
	}
	return nil
}

func (c *TimeSeriesForwarderSinkChange) GetRawResource() gotenresource.Resource {
	return c.GetTimeSeriesForwarderSink()
}

func (c *TimeSeriesForwarderSinkChange) GetTimeSeriesForwarderSinkName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *TimeSeriesForwarderSinkChange_Added_:
		return cType.Added.TimeSeriesForwarderSink.GetName()
	case *TimeSeriesForwarderSinkChange_Modified_:
		return cType.Modified.Name
	case *TimeSeriesForwarderSinkChange_Current_:
		return cType.Current.TimeSeriesForwarderSink.GetName()
	case *TimeSeriesForwarderSinkChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *TimeSeriesForwarderSinkChange) GetRawName() gotenresource.Name {
	return c.GetTimeSeriesForwarderSinkName()
}

func (c *TimeSeriesForwarderSinkChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &TimeSeriesForwarderSinkChange_Added_{
		Added: &TimeSeriesForwarderSinkChange_Added{
			TimeSeriesForwarderSink: snapshot.(*TimeSeriesForwarderSink),
			ViewIndex:               int32(idx),
		},
	}
}

func (c *TimeSeriesForwarderSinkChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &TimeSeriesForwarderSinkChange_Modified_{
		Modified: &TimeSeriesForwarderSinkChange_Modified{
			Name:                    name.(*Name),
			TimeSeriesForwarderSink: snapshot.(*TimeSeriesForwarderSink),
			PreviousViewIndex:       int32(prevIdx),
			ViewIndex:               int32(newIdx),
		},
	}
}

func (c *TimeSeriesForwarderSinkChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &TimeSeriesForwarderSinkChange_Current_{
		Current: &TimeSeriesForwarderSinkChange_Current{
			TimeSeriesForwarderSink: snapshot.(*TimeSeriesForwarderSink),
		},
	}
}

func (c *TimeSeriesForwarderSinkChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &TimeSeriesForwarderSinkChange_Removed_{
		Removed: &TimeSeriesForwarderSinkChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
