// Code generated by protoc-gen-goten-resource
// Resource change: BucketChange
// DO NOT EDIT!!!

package bucket

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &fieldmaskpb.FieldMask{}
	_ = &meta_service.Service{}
)

func (c *BucketChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*BucketChange_Added_)
	return ok
}

func (c *BucketChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*BucketChange_Modified_)
	return ok
}

func (c *BucketChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*BucketChange_Current_)
	return ok
}

func (c *BucketChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*BucketChange_Removed_)
	return ok
}

func (c *BucketChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *BucketChange_Added_:
		return cType.Added.ViewIndex
	case *BucketChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *BucketChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *BucketChange_Removed_:
		return cType.Removed.ViewIndex
	case *BucketChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *BucketChange) GetBucket() *Bucket {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *BucketChange_Added_:
		return cType.Added.Bucket
	case *BucketChange_Modified_:
		return cType.Modified.Bucket
	case *BucketChange_Current_:
		return cType.Current.Bucket
	case *BucketChange_Removed_:
		return nil
	}
	return nil
}

func (c *BucketChange) GetRawResource() gotenresource.Resource {
	return c.GetBucket()
}

func (c *BucketChange) GetBucketName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *BucketChange_Added_:
		return cType.Added.Bucket.GetName()
	case *BucketChange_Modified_:
		return cType.Modified.Name
	case *BucketChange_Current_:
		return cType.Current.Bucket.GetName()
	case *BucketChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *BucketChange) GetRawName() gotenresource.Name {
	return c.GetBucketName()
}

func (c *BucketChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &BucketChange_Added_{
		Added: &BucketChange_Added{
			Bucket:    snapshot.(*Bucket),
			ViewIndex: int32(idx),
		},
	}
}

func (c *BucketChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &BucketChange_Modified_{
		Modified: &BucketChange_Modified{
			Name:              name.(*Name),
			Bucket:            snapshot.(*Bucket),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *BucketChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &BucketChange_Current_{
		Current: &BucketChange_Current{
			Bucket: snapshot.(*Bucket),
		},
	}
}

func (c *BucketChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &BucketChange_Removed_{
		Removed: &BucketChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}