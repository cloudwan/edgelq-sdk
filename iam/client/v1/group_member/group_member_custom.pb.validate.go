// Code generated by protoc-gen-goten-validate
// File: edgelq/iam/proto/v1/group_member_custom.proto
// DO NOT EDIT!!!

package group_member_client

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	gotenvalidate "github.com/cloudwan/goten-sdk/runtime/validate"
)

// proto imports
import (
	iam_common "github.com/cloudwan/edgelq-sdk/iam/client/v1/common"
	group_member "github.com/cloudwan/edgelq-sdk/iam/resources/v1/group_member"
	view "github.com/cloudwan/goten-sdk/types/view"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

var (
	_ = bytes.Equal
	_ = errors.New
	_ = fmt.Errorf
	_ = net.ParseIP
	_ = regexp.Match
	_ = strings.Split
	_ = time.Now
	_ = utf8.RuneCountInString
	_ = url.Parse
	_ = gotenvalidate.NewValidationError
)

// make sure we're using proto imports
var (
	_ = &iam_common.MembersMasks{}
	_ = &group_member.GroupMember{}
	_ = &fieldmaskpb.FieldMask{}
	_ = view.View(0)
)

func (obj *ListGroupMembersWithMembersRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if !(obj.PageSize >= 0 && obj.PageSize < 1000) {
		return gotenvalidate.NewValidationError("ListGroupMembersWithMembersRequest", "pageSize", obj.PageSize, "field must be in range [0, 1000)", nil)
	}
	if subobj, ok := interface{}(obj.MembersMasks).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ListGroupMembersWithMembersRequest", "membersMasks", obj.MembersMasks, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListGroupMembersWithMembersResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.GroupMembers {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ListGroupMembersWithMembersResponse", "groupMembers", obj.GroupMembers[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.MatchingMembers {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ListGroupMembersWithMembersResponse", "matchingMembers", obj.MatchingMembers[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
