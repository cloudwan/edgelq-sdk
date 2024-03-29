// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha2/user.proto
// DO NOT EDIT!!!

package user

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	meta "github.com/cloudwan/goten-sdk/types/meta"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// ensure the imports are used
var (
	_ = new(fmt.Stringer)
	_ = new(sort.Interface)

	_ = new(proto.Message)
	_ = googlefieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &timestamppb.Timestamp{}
	_ = &meta.Meta{}
)

func (o *User) GotenObjectExt() {}

func (o *User) MakeFullFieldMask() *User_FieldMask {
	return FullUser_FieldMask()
}

func (o *User) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullUser_FieldMask()
}

func (o *User) MakeDiffFieldMask(other *User) *User_FieldMask {
	if o == nil && other == nil {
		return &User_FieldMask{}
	}
	if o == nil || other == nil {
		return FullUser_FieldMask()
	}

	res := &User_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &User_FieldTerminalPath{selector: User_FieldPathSelectorName})
	}
	if o.GetFullName() != other.GetFullName() {
		res.Paths = append(res.Paths, &User_FieldTerminalPath{selector: User_FieldPathSelectorFullName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &User_FieldTerminalPath{selector: User_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &User_FieldSubPath{selector: User_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetEmail() != other.GetEmail() {
		res.Paths = append(res.Paths, &User_FieldTerminalPath{selector: User_FieldPathSelectorEmail})
	}
	if o.GetEmailVerified() != other.GetEmailVerified() {
		res.Paths = append(res.Paths, &User_FieldTerminalPath{selector: User_FieldPathSelectorEmailVerified})
	}
	{
		subMask := o.GetAuthInfo().MakeDiffFieldMask(other.GetAuthInfo())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &User_FieldTerminalPath{selector: User_FieldPathSelectorAuthInfo})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &User_FieldSubPath{selector: User_FieldPathSelectorAuthInfo, subPath: subpath})
			}
		}
	}

	if len(o.GetSettings()) == len(other.GetSettings()) {
		for i, lValue := range o.GetSettings() {
			rValue := other.GetSettings()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &User_FieldTerminalPath{selector: User_FieldPathSelectorSettings})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &User_FieldTerminalPath{selector: User_FieldPathSelectorSettings})
	}
	if !proto.Equal(o.GetRefreshedTime(), other.GetRefreshedTime()) {
		res.Paths = append(res.Paths, &User_FieldTerminalPath{selector: User_FieldPathSelectorRefreshedTime})
	}
	return res
}

func (o *User) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*User))
}

func (o *User) Clone() *User {
	if o == nil {
		return nil
	}
	result := &User{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &Name{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.FullName = o.FullName
	result.Metadata = o.Metadata.Clone()
	result.Email = o.Email
	result.EmailVerified = o.EmailVerified
	result.AuthInfo = o.AuthInfo.Clone()
	result.Settings = map[string]string{}
	for key, sourceValue := range o.Settings {
		result.Settings[key] = sourceValue
	}
	result.RefreshedTime = proto.Clone(o.RefreshedTime).(*timestamppb.Timestamp)
	return result
}

func (o *User) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *User) Merge(source *User) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &Name{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
	o.FullName = source.GetFullName()
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
	o.Email = source.GetEmail()
	o.EmailVerified = source.GetEmailVerified()
	if source.GetAuthInfo() != nil {
		if o.AuthInfo == nil {
			o.AuthInfo = new(User_AuthInfo)
		}
		o.AuthInfo.Merge(source.GetAuthInfo())
	}
	if source.GetSettings() != nil {
		if o.Settings == nil {
			o.Settings = make(map[string]string, len(source.GetSettings()))
		}
		for key, sourceValue := range source.GetSettings() {
			o.Settings[key] = sourceValue
		}
	}
	if source.GetRefreshedTime() != nil {
		if o.RefreshedTime == nil {
			o.RefreshedTime = new(timestamppb.Timestamp)
		}
		proto.Merge(o.RefreshedTime, source.GetRefreshedTime())
	}
}

func (o *User) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*User))
}

func (o *User_AuthInfo) GotenObjectExt() {}

func (o *User_AuthInfo) MakeFullFieldMask() *User_AuthInfo_FieldMask {
	return FullUser_AuthInfo_FieldMask()
}

func (o *User_AuthInfo) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullUser_AuthInfo_FieldMask()
}

func (o *User_AuthInfo) MakeDiffFieldMask(other *User_AuthInfo) *User_AuthInfo_FieldMask {
	if o == nil && other == nil {
		return &User_AuthInfo_FieldMask{}
	}
	if o == nil || other == nil {
		return FullUser_AuthInfo_FieldMask()
	}

	res := &User_AuthInfo_FieldMask{}
	if o.GetProvider() != other.GetProvider() {
		res.Paths = append(res.Paths, &UserAuthInfo_FieldTerminalPath{selector: UserAuthInfo_FieldPathSelectorProvider})
	}
	if o.GetId() != other.GetId() {
		res.Paths = append(res.Paths, &UserAuthInfo_FieldTerminalPath{selector: UserAuthInfo_FieldPathSelectorId})
	}
	return res
}

func (o *User_AuthInfo) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*User_AuthInfo))
}

func (o *User_AuthInfo) Clone() *User_AuthInfo {
	if o == nil {
		return nil
	}
	result := &User_AuthInfo{}
	result.Provider = o.Provider
	result.Id = o.Id
	return result
}

func (o *User_AuthInfo) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *User_AuthInfo) Merge(source *User_AuthInfo) {
	o.Provider = source.GetProvider()
	o.Id = source.GetId()
}

func (o *User_AuthInfo) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*User_AuthInfo))
}
