// Code generated by protoc-gen-goten-resource
// Resource change: SecretChange
// DO NOT EDIT!!!

package secret

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/secrets/resources/v1/project"
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

func (c *SecretChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*SecretChange_Added_)
	return ok
}

func (c *SecretChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*SecretChange_Modified_)
	return ok
}

func (c *SecretChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*SecretChange_Current_)
	return ok
}

func (c *SecretChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*SecretChange_Removed_)
	return ok
}

func (c *SecretChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *SecretChange_Added_:
		return cType.Added.ViewIndex
	case *SecretChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *SecretChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *SecretChange_Removed_:
		return cType.Removed.ViewIndex
	case *SecretChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *SecretChange) GetSecret() *Secret {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *SecretChange_Added_:
		return cType.Added.Secret
	case *SecretChange_Modified_:
		return cType.Modified.Secret
	case *SecretChange_Current_:
		return cType.Current.Secret
	case *SecretChange_Removed_:
		return nil
	}
	return nil
}

func (c *SecretChange) GetRawResource() gotenresource.Resource {
	return c.GetSecret()
}

func (c *SecretChange) GetSecretName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *SecretChange_Added_:
		return cType.Added.Secret.GetName()
	case *SecretChange_Modified_:
		return cType.Modified.Name
	case *SecretChange_Current_:
		return cType.Current.Secret.GetName()
	case *SecretChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *SecretChange) GetRawName() gotenresource.Name {
	return c.GetSecretName()
}

func (c *SecretChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &SecretChange_Added_{
		Added: &SecretChange_Added{
			Secret:    snapshot.(*Secret),
			ViewIndex: int32(idx),
		},
	}
}

func (c *SecretChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &SecretChange_Modified_{
		Modified: &SecretChange_Modified{
			Name:              name.(*Name),
			Secret:            snapshot.(*Secret),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *SecretChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &SecretChange_Current_{
		Current: &SecretChange_Current{
			Secret: snapshot.(*Secret),
		},
	}
}

func (c *SecretChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &SecretChange_Removed_{
		Removed: &SecretChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
