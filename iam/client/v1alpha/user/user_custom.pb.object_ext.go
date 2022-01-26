// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha/user_custom.proto
// DO NOT EDIT!!!

package user_client

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	user "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/user"
	view "github.com/cloudwan/goten-sdk/runtime/api/view"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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
	_ = &user.User{}
	_ = &field_mask.FieldMask{}
	_ = view.View(0)
)

func (o *GetUserByEmailRequest) GotenObjectExt() {}

func (o *GetUserByEmailRequest) MakeFullFieldMask() *GetUserByEmailRequest_FieldMask {
	return FullGetUserByEmailRequest_FieldMask()
}

func (o *GetUserByEmailRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullGetUserByEmailRequest_FieldMask()
}

func (o *GetUserByEmailRequest) MakeDiffFieldMask(other *GetUserByEmailRequest) *GetUserByEmailRequest_FieldMask {
	if o == nil && other == nil {
		return &GetUserByEmailRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullGetUserByEmailRequest_FieldMask()
	}

	res := &GetUserByEmailRequest_FieldMask{}
	if o.GetEmail() != other.GetEmail() {
		res.Paths = append(res.Paths, &GetUserByEmailRequest_FieldTerminalPath{selector: GetUserByEmailRequest_FieldPathSelectorEmail})
	}
	if !proto.Equal(o.GetFieldMask(), other.GetFieldMask()) {
		res.Paths = append(res.Paths, &GetUserByEmailRequest_FieldTerminalPath{selector: GetUserByEmailRequest_FieldPathSelectorFieldMask})
	}
	if o.GetView() != other.GetView() {
		res.Paths = append(res.Paths, &GetUserByEmailRequest_FieldTerminalPath{selector: GetUserByEmailRequest_FieldPathSelectorView})
	}
	if o.GetSkipCache() != other.GetSkipCache() {
		res.Paths = append(res.Paths, &GetUserByEmailRequest_FieldTerminalPath{selector: GetUserByEmailRequest_FieldPathSelectorSkipCache})
	}
	return res
}

func (o *GetUserByEmailRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*GetUserByEmailRequest))
}

func (o *GetUserByEmailRequest) Clone() *GetUserByEmailRequest {
	if o == nil {
		return nil
	}
	result := &GetUserByEmailRequest{}
	result.Email = o.Email
	result.FieldMask = proto.Clone(o.FieldMask).(*user.User_FieldMask)
	result.View = o.View
	result.SkipCache = o.SkipCache
	return result
}

func (o *GetUserByEmailRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *GetUserByEmailRequest) Merge(source *GetUserByEmailRequest) {
	o.Email = source.GetEmail()
	if source.GetFieldMask() != nil {
		if o.FieldMask == nil {
			o.FieldMask = new(user.User_FieldMask)
		}
		mergedMask := fieldmaskpb.Union(source.GetFieldMask().ToProtoFieldMask(), o.FieldMask.ToProtoFieldMask())
		if err := o.FieldMask.FromProtoFieldMask(mergedMask); err != nil {
			panic(err)
		}
	}
	o.View = source.GetView()
	o.SkipCache = source.GetSkipCache()
}

func (o *GetUserByEmailRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*GetUserByEmailRequest))
}

func (o *BatchGetUsersByEmailRequest) GotenObjectExt() {}

func (o *BatchGetUsersByEmailRequest) MakeFullFieldMask() *BatchGetUsersByEmailRequest_FieldMask {
	return FullBatchGetUsersByEmailRequest_FieldMask()
}

func (o *BatchGetUsersByEmailRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullBatchGetUsersByEmailRequest_FieldMask()
}

func (o *BatchGetUsersByEmailRequest) MakeDiffFieldMask(other *BatchGetUsersByEmailRequest) *BatchGetUsersByEmailRequest_FieldMask {
	if o == nil && other == nil {
		return &BatchGetUsersByEmailRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullBatchGetUsersByEmailRequest_FieldMask()
	}

	res := &BatchGetUsersByEmailRequest_FieldMask{}

	if len(o.GetEmails()) == len(other.GetEmails()) {
		for i, lValue := range o.GetEmails() {
			rValue := other.GetEmails()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &BatchGetUsersByEmailRequest_FieldTerminalPath{selector: BatchGetUsersByEmailRequest_FieldPathSelectorEmails})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &BatchGetUsersByEmailRequest_FieldTerminalPath{selector: BatchGetUsersByEmailRequest_FieldPathSelectorEmails})
	}
	if !proto.Equal(o.GetFieldMask(), other.GetFieldMask()) {
		res.Paths = append(res.Paths, &BatchGetUsersByEmailRequest_FieldTerminalPath{selector: BatchGetUsersByEmailRequest_FieldPathSelectorFieldMask})
	}
	if o.GetView() != other.GetView() {
		res.Paths = append(res.Paths, &BatchGetUsersByEmailRequest_FieldTerminalPath{selector: BatchGetUsersByEmailRequest_FieldPathSelectorView})
	}
	if o.GetSkipCache() != other.GetSkipCache() {
		res.Paths = append(res.Paths, &BatchGetUsersByEmailRequest_FieldTerminalPath{selector: BatchGetUsersByEmailRequest_FieldPathSelectorSkipCache})
	}
	return res
}

func (o *BatchGetUsersByEmailRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*BatchGetUsersByEmailRequest))
}

func (o *BatchGetUsersByEmailRequest) Clone() *BatchGetUsersByEmailRequest {
	if o == nil {
		return nil
	}
	result := &BatchGetUsersByEmailRequest{}
	result.Emails = make([]string, len(o.Emails))
	for i, sourceValue := range o.Emails {
		result.Emails[i] = sourceValue
	}
	result.FieldMask = proto.Clone(o.FieldMask).(*user.User_FieldMask)
	result.View = o.View
	result.SkipCache = o.SkipCache
	return result
}

func (o *BatchGetUsersByEmailRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *BatchGetUsersByEmailRequest) Merge(source *BatchGetUsersByEmailRequest) {
	for _, sourceValue := range source.GetEmails() {
		exists := false
		for _, currentValue := range o.Emails {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.Emails = append(o.Emails, newDstElement)
		}
	}

	if source.GetFieldMask() != nil {
		if o.FieldMask == nil {
			o.FieldMask = new(user.User_FieldMask)
		}
		mergedMask := fieldmaskpb.Union(source.GetFieldMask().ToProtoFieldMask(), o.FieldMask.ToProtoFieldMask())
		if err := o.FieldMask.FromProtoFieldMask(mergedMask); err != nil {
			panic(err)
		}
	}
	o.View = source.GetView()
	o.SkipCache = source.GetSkipCache()
}

func (o *BatchGetUsersByEmailRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*BatchGetUsersByEmailRequest))
}

func (o *BatchGetUsersByEmailResponse) GotenObjectExt() {}

func (o *BatchGetUsersByEmailResponse) MakeFullFieldMask() *BatchGetUsersByEmailResponse_FieldMask {
	return FullBatchGetUsersByEmailResponse_FieldMask()
}

func (o *BatchGetUsersByEmailResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullBatchGetUsersByEmailResponse_FieldMask()
}

func (o *BatchGetUsersByEmailResponse) MakeDiffFieldMask(other *BatchGetUsersByEmailResponse) *BatchGetUsersByEmailResponse_FieldMask {
	if o == nil && other == nil {
		return &BatchGetUsersByEmailResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullBatchGetUsersByEmailResponse_FieldMask()
	}

	res := &BatchGetUsersByEmailResponse_FieldMask{}

	if len(o.GetUsers()) == len(other.GetUsers()) {
		for i, lValue := range o.GetUsers() {
			rValue := other.GetUsers()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &BatchGetUsersByEmailResponse_FieldTerminalPath{selector: BatchGetUsersByEmailResponse_FieldPathSelectorUsers})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &BatchGetUsersByEmailResponse_FieldTerminalPath{selector: BatchGetUsersByEmailResponse_FieldPathSelectorUsers})
	}

	if len(o.GetMissing()) == len(other.GetMissing()) {
		for i, lValue := range o.GetMissing() {
			rValue := other.GetMissing()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &BatchGetUsersByEmailResponse_FieldTerminalPath{selector: BatchGetUsersByEmailResponse_FieldPathSelectorMissing})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &BatchGetUsersByEmailResponse_FieldTerminalPath{selector: BatchGetUsersByEmailResponse_FieldPathSelectorMissing})
	}
	return res
}

func (o *BatchGetUsersByEmailResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*BatchGetUsersByEmailResponse))
}

func (o *BatchGetUsersByEmailResponse) Clone() *BatchGetUsersByEmailResponse {
	if o == nil {
		return nil
	}
	result := &BatchGetUsersByEmailResponse{}
	result.Users = make([]*user.User, len(o.Users))
	for i, sourceValue := range o.Users {
		result.Users[i] = sourceValue.Clone()
	}
	result.Missing = make([]string, len(o.Missing))
	for i, sourceValue := range o.Missing {
		result.Missing[i] = sourceValue
	}
	return result
}

func (o *BatchGetUsersByEmailResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *BatchGetUsersByEmailResponse) Merge(source *BatchGetUsersByEmailResponse) {
	for _, sourceValue := range source.GetUsers() {
		exists := false
		for _, currentValue := range o.Users {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *user.User
			if sourceValue != nil {
				newDstElement = new(user.User)
				newDstElement.Merge(sourceValue)
			}
			o.Users = append(o.Users, newDstElement)
		}
	}

	for _, sourceValue := range source.GetMissing() {
		exists := false
		for _, currentValue := range o.Missing {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.Missing = append(o.Missing, newDstElement)
		}
	}

}

func (o *BatchGetUsersByEmailResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*BatchGetUsersByEmailResponse))
}

func (o *RefreshUserFromIdTokenRequest) GotenObjectExt() {}

func (o *RefreshUserFromIdTokenRequest) MakeFullFieldMask() *RefreshUserFromIdTokenRequest_FieldMask {
	return FullRefreshUserFromIdTokenRequest_FieldMask()
}

func (o *RefreshUserFromIdTokenRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRefreshUserFromIdTokenRequest_FieldMask()
}

func (o *RefreshUserFromIdTokenRequest) MakeDiffFieldMask(other *RefreshUserFromIdTokenRequest) *RefreshUserFromIdTokenRequest_FieldMask {
	if o == nil && other == nil {
		return &RefreshUserFromIdTokenRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRefreshUserFromIdTokenRequest_FieldMask()
	}

	res := &RefreshUserFromIdTokenRequest_FieldMask{}
	if o.GetIdToken() != other.GetIdToken() {
		res.Paths = append(res.Paths, &RefreshUserFromIdTokenRequest_FieldTerminalPath{selector: RefreshUserFromIdTokenRequest_FieldPathSelectorIdToken})
	}
	return res
}

func (o *RefreshUserFromIdTokenRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*RefreshUserFromIdTokenRequest))
}

func (o *RefreshUserFromIdTokenRequest) Clone() *RefreshUserFromIdTokenRequest {
	if o == nil {
		return nil
	}
	result := &RefreshUserFromIdTokenRequest{}
	result.IdToken = o.IdToken
	return result
}

func (o *RefreshUserFromIdTokenRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *RefreshUserFromIdTokenRequest) Merge(source *RefreshUserFromIdTokenRequest) {
	o.IdToken = source.GetIdToken()
}

func (o *RefreshUserFromIdTokenRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*RefreshUserFromIdTokenRequest))
}

func (o *RefreshUserFromIdTokenResponse) GotenObjectExt() {}

func (o *RefreshUserFromIdTokenResponse) MakeFullFieldMask() *RefreshUserFromIdTokenResponse_FieldMask {
	return FullRefreshUserFromIdTokenResponse_FieldMask()
}

func (o *RefreshUserFromIdTokenResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRefreshUserFromIdTokenResponse_FieldMask()
}

func (o *RefreshUserFromIdTokenResponse) MakeDiffFieldMask(other *RefreshUserFromIdTokenResponse) *RefreshUserFromIdTokenResponse_FieldMask {
	if o == nil && other == nil {
		return &RefreshUserFromIdTokenResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRefreshUserFromIdTokenResponse_FieldMask()
	}

	res := &RefreshUserFromIdTokenResponse_FieldMask{}
	return res
}

func (o *RefreshUserFromIdTokenResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*RefreshUserFromIdTokenResponse))
}

func (o *RefreshUserFromIdTokenResponse) Clone() *RefreshUserFromIdTokenResponse {
	if o == nil {
		return nil
	}
	result := &RefreshUserFromIdTokenResponse{}
	return result
}

func (o *RefreshUserFromIdTokenResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *RefreshUserFromIdTokenResponse) Merge(source *RefreshUserFromIdTokenResponse) {
}

func (o *RefreshUserFromIdTokenResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*RefreshUserFromIdTokenResponse))
}

func (o *GetMySettingsRequest) GotenObjectExt() {}

func (o *GetMySettingsRequest) MakeFullFieldMask() *GetMySettingsRequest_FieldMask {
	return FullGetMySettingsRequest_FieldMask()
}

func (o *GetMySettingsRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullGetMySettingsRequest_FieldMask()
}

func (o *GetMySettingsRequest) MakeDiffFieldMask(other *GetMySettingsRequest) *GetMySettingsRequest_FieldMask {
	if o == nil && other == nil {
		return &GetMySettingsRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullGetMySettingsRequest_FieldMask()
	}

	res := &GetMySettingsRequest_FieldMask{}

	if len(o.GetKeys()) == len(other.GetKeys()) {
		for i, lValue := range o.GetKeys() {
			rValue := other.GetKeys()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &GetMySettingsRequest_FieldTerminalPath{selector: GetMySettingsRequest_FieldPathSelectorKeys})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &GetMySettingsRequest_FieldTerminalPath{selector: GetMySettingsRequest_FieldPathSelectorKeys})
	}
	return res
}

func (o *GetMySettingsRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*GetMySettingsRequest))
}

func (o *GetMySettingsRequest) Clone() *GetMySettingsRequest {
	if o == nil {
		return nil
	}
	result := &GetMySettingsRequest{}
	result.Keys = make([]string, len(o.Keys))
	for i, sourceValue := range o.Keys {
		result.Keys[i] = sourceValue
	}
	return result
}

func (o *GetMySettingsRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *GetMySettingsRequest) Merge(source *GetMySettingsRequest) {
	for _, sourceValue := range source.GetKeys() {
		exists := false
		for _, currentValue := range o.Keys {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.Keys = append(o.Keys, newDstElement)
		}
	}

}

func (o *GetMySettingsRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*GetMySettingsRequest))
}

func (o *GetMySettingsResponse) GotenObjectExt() {}

func (o *GetMySettingsResponse) MakeFullFieldMask() *GetMySettingsResponse_FieldMask {
	return FullGetMySettingsResponse_FieldMask()
}

func (o *GetMySettingsResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullGetMySettingsResponse_FieldMask()
}

func (o *GetMySettingsResponse) MakeDiffFieldMask(other *GetMySettingsResponse) *GetMySettingsResponse_FieldMask {
	if o == nil && other == nil {
		return &GetMySettingsResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullGetMySettingsResponse_FieldMask()
	}

	res := &GetMySettingsResponse_FieldMask{}

	if len(o.GetSettings()) == len(other.GetSettings()) {
		for i, lValue := range o.GetSettings() {
			rValue := other.GetSettings()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &GetMySettingsResponse_FieldTerminalPath{selector: GetMySettingsResponse_FieldPathSelectorSettings})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &GetMySettingsResponse_FieldTerminalPath{selector: GetMySettingsResponse_FieldPathSelectorSettings})
	}
	return res
}

func (o *GetMySettingsResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*GetMySettingsResponse))
}

func (o *GetMySettingsResponse) Clone() *GetMySettingsResponse {
	if o == nil {
		return nil
	}
	result := &GetMySettingsResponse{}
	result.Settings = map[string]string{}
	for key, sourceValue := range o.Settings {
		result.Settings[key] = sourceValue
	}
	return result
}

func (o *GetMySettingsResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *GetMySettingsResponse) Merge(source *GetMySettingsResponse) {
	if source.GetSettings() != nil {
		if o.Settings == nil {
			o.Settings = make(map[string]string, len(source.GetSettings()))
		}
		for key, sourceValue := range source.GetSettings() {
			o.Settings[key] = sourceValue
		}
	}
}

func (o *GetMySettingsResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*GetMySettingsResponse))
}

func (o *SetMySettingsRequest) GotenObjectExt() {}

func (o *SetMySettingsRequest) MakeFullFieldMask() *SetMySettingsRequest_FieldMask {
	return FullSetMySettingsRequest_FieldMask()
}

func (o *SetMySettingsRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullSetMySettingsRequest_FieldMask()
}

func (o *SetMySettingsRequest) MakeDiffFieldMask(other *SetMySettingsRequest) *SetMySettingsRequest_FieldMask {
	if o == nil && other == nil {
		return &SetMySettingsRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullSetMySettingsRequest_FieldMask()
	}

	res := &SetMySettingsRequest_FieldMask{}

	if len(o.GetSettings()) == len(other.GetSettings()) {
		for i, lValue := range o.GetSettings() {
			rValue := other.GetSettings()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &SetMySettingsRequest_FieldTerminalPath{selector: SetMySettingsRequest_FieldPathSelectorSettings})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &SetMySettingsRequest_FieldTerminalPath{selector: SetMySettingsRequest_FieldPathSelectorSettings})
	}
	return res
}

func (o *SetMySettingsRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*SetMySettingsRequest))
}

func (o *SetMySettingsRequest) Clone() *SetMySettingsRequest {
	if o == nil {
		return nil
	}
	result := &SetMySettingsRequest{}
	result.Settings = map[string]string{}
	for key, sourceValue := range o.Settings {
		result.Settings[key] = sourceValue
	}
	return result
}

func (o *SetMySettingsRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *SetMySettingsRequest) Merge(source *SetMySettingsRequest) {
	if source.GetSettings() != nil {
		if o.Settings == nil {
			o.Settings = make(map[string]string, len(source.GetSettings()))
		}
		for key, sourceValue := range source.GetSettings() {
			o.Settings[key] = sourceValue
		}
	}
}

func (o *SetMySettingsRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*SetMySettingsRequest))
}
