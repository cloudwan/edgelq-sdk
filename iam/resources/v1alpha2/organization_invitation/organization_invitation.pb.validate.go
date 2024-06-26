// Code generated by protoc-gen-goten-validate
// File: edgelq/iam/proto/v1alpha2/organization_invitation.proto
// DO NOT EDIT!!!

package organization_invitation

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
	iam_invitation "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/invitation"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	meta "github.com/cloudwan/goten-sdk/types/meta"
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
	_ = &iam_invitation.Actor{}
	_ = &organization.Organization{}
	_ = &meta.Meta{}
)

func (obj *OrganizationInvitation) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.Invitation == nil {
		return gotenvalidate.NewValidationError("OrganizationInvitation", "invitation", obj.Invitation, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.Invitation).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("OrganizationInvitation", "invitation", obj.Invitation, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Metadata).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("OrganizationInvitation", "metadata", obj.Metadata, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
