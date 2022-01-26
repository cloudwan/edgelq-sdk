// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha/common.proto
// DO NOT EDIT!!!

package iam_common

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	role "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/role"
	service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/service_account"
	user "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/user"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// ensure the imports are used
var (
	_ = fmt.Stringer(nil)
	_ = sort.Interface(nil)

	_ = proto.Message(nil)
	_ = fieldmaskpb.FieldMask{}

	_ = gotenobject.FieldPath(nil)
)

// make sure we're using proto imports
var (
	_ = &role.Role{}
	_ = &service_account.ServiceAccount{}
	_ = &user.User{}
	_ = &timestamp.Timestamp{}
)

func (o *Actor) GotenObjectExt() {}

func (o *Actor) MakeFullFieldMask() *Actor_FieldMask {
	return FullActor_FieldMask()
}

func (o *Actor) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullActor_FieldMask()
}

func (o *Actor) MakeDiffFieldMask(other *Actor) *Actor_FieldMask {
	if o == nil && other == nil {
		return &Actor_FieldMask{}
	}
	if o == nil || other == nil {
		return FullActor_FieldMask()
	}

	res := &Actor_FieldMask{}
	if o.GetUser().String() != other.GetUser().String() {
		res.Paths = append(res.Paths, &Actor_FieldTerminalPath{selector: Actor_FieldPathSelectorUser})
	}
	if o.GetServiceAccount().String() != other.GetServiceAccount().String() {
		res.Paths = append(res.Paths, &Actor_FieldTerminalPath{selector: Actor_FieldPathSelectorServiceAccount})
	}
	return res
}

func (o *Actor) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Actor))
}

func (o *Actor) Clone() *Actor {
	if o == nil {
		return nil
	}
	result := &Actor{}
	if o.User == nil {
		result.User = nil
	} else if data, err := o.User.ProtoString(); err != nil {
		panic(err)
	} else {
		result.User = &user.Reference{}
		if err := result.User.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	if o.ServiceAccount == nil {
		result.ServiceAccount = nil
	} else if data, err := o.ServiceAccount.ProtoString(); err != nil {
		panic(err)
	} else {
		result.ServiceAccount = &service_account.Reference{}
		if err := result.ServiceAccount.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	return result
}

func (o *Actor) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Actor) Merge(source *Actor) {
	if source.GetUser() != nil {
		if data, err := source.GetUser().ProtoString(); err != nil {
			panic(err)
		} else {
			o.User = &user.Reference{}
			if err := o.User.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.User = nil
	}
	if source.GetServiceAccount() != nil {
		if data, err := source.GetServiceAccount().ProtoString(); err != nil {
			panic(err)
		} else {
			o.ServiceAccount = &service_account.Reference{}
			if err := o.ServiceAccount.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.ServiceAccount = nil
	}
}

func (o *Actor) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Actor))
}

func (o *Invitation) GotenObjectExt() {}

func (o *Invitation) MakeFullFieldMask() *Invitation_FieldMask {
	return FullInvitation_FieldMask()
}

func (o *Invitation) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullInvitation_FieldMask()
}

func (o *Invitation) MakeDiffFieldMask(other *Invitation) *Invitation_FieldMask {
	if o == nil && other == nil {
		return &Invitation_FieldMask{}
	}
	if o == nil || other == nil {
		return FullInvitation_FieldMask()
	}

	res := &Invitation_FieldMask{}
	if o.GetInviteeEmail() != other.GetInviteeEmail() {
		res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorInviteeEmail})
	}
	{
		subMask := o.GetInviterActor().MakeDiffFieldMask(other.GetInviterActor())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorInviterActor})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Invitation_FieldSubPath{selector: Invitation_FieldPathSelectorInviterActor, subPath: subpath})
			}
		}
	}
	if o.GetInviterFullName() != other.GetInviterFullName() {
		res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorInviterFullName})
	}
	if o.GetInviterEmail() != other.GetInviterEmail() {
		res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorInviterEmail})
	}

	if len(o.GetRoles()) == len(other.GetRoles()) {
		for i, lValue := range o.GetRoles() {
			rValue := other.GetRoles()[i]
			if lValue.String() != rValue.String() {
				res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorRoles})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorRoles})
	}
	if !proto.Equal(o.GetExpirationDate(), other.GetExpirationDate()) {
		res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorExpirationDate})
	}
	if o.GetState() != other.GetState() {
		res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorState})
	}
	return res
}

func (o *Invitation) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Invitation))
}

func (o *Invitation) Clone() *Invitation {
	if o == nil {
		return nil
	}
	result := &Invitation{}
	result.InviteeEmail = o.InviteeEmail
	result.InviterActor = o.InviterActor.Clone()
	result.InviterFullName = o.InviterFullName
	result.InviterEmail = o.InviterEmail
	result.Roles = make([]*role.Reference, len(o.Roles))
	for i, sourceValue := range o.Roles {
		if sourceValue == nil {
			result.Roles[i] = nil
		} else if data, err := sourceValue.ProtoString(); err != nil {
			panic(err)
		} else {
			result.Roles[i] = &role.Reference{}
			if err := result.Roles[i].ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	}
	result.ExpirationDate = proto.Clone(o.ExpirationDate).(*timestamp.Timestamp)
	result.State = o.State
	return result
}

func (o *Invitation) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Invitation) Merge(source *Invitation) {
	o.InviteeEmail = source.GetInviteeEmail()
	if source.GetInviterActor() != nil {
		if o.InviterActor == nil {
			o.InviterActor = new(Actor)
		}
		o.InviterActor.Merge(source.GetInviterActor())
	}
	o.InviterFullName = source.GetInviterFullName()
	o.InviterEmail = source.GetInviterEmail()
	for _, sourceValue := range source.GetRoles() {
		exists := false
		for _, currentValue := range o.Roles {
			leftProtoStr, _ := currentValue.ProtoString()
			rightProtoStr, _ := sourceValue.ProtoString()
			if leftProtoStr == rightProtoStr {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *role.Reference
			if sourceValue != nil {
				if data, err := sourceValue.ProtoString(); err != nil {
					panic(err)
				} else {
					newDstElement = &role.Reference{}
					if err := newDstElement.ParseProtoString(data); err != nil {
						panic(err)
					}
				}
			}
			o.Roles = append(o.Roles, newDstElement)
		}
	}

	if source.GetExpirationDate() != nil {
		if o.ExpirationDate == nil {
			o.ExpirationDate = new(timestamp.Timestamp)
		}
		proto.Merge(o.ExpirationDate, source.GetExpirationDate())
	}
	o.State = source.GetState()
}

func (o *Invitation) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Invitation))
}
