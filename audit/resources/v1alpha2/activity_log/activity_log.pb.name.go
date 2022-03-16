// Code generated by protoc-gen-goten-resource
// Resource: ActivityLog
// DO NOT EDIT!!!

package activity_log

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strings"

	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/common/types/traits"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/cloudwan/goten-sdk/runtime/goten"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	audit_common "github.com/cloudwan/edgelq-sdk/audit/common/v1alpha2"
	rpc "github.com/cloudwan/edgelq-sdk/common/rpc"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// ensure the imports are used
var (
	_ = codes.NotFound
	_ = fmt.Stringer(nil)
	_ = proto.Message(nil)
	_ = status.Status{}
	_ = url.URL{}
	_ = strings.Builder{}

	_ = goten.GotenMessage(nil)
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &audit_common.Authentication{}
	_ = &rpc.Status{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &any.Any{}
	_ = &field_mask.FieldMask{}
	_ = &timestamp.Timestamp{}
)

var activityLog_RegexpId = regexp.MustCompile("^(?P<activity_log_id>[\\w./-=+]{1,128})$")
var regexPath = regexp.MustCompile("^activityLogs/(?P<activity_log_id>-|[\\w./-=+]{1,128})$")
var regexPath_Project = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})/activityLogs/(?P<activity_log_id>-|[\\w./-=+]{1,128})$")
var regexPath_Organization = regexp.MustCompile("^organizations/(?P<organization_id>-|[\\w][\\w.-]{0,127})/activityLogs/(?P<activity_log_id>-|[\\w./-=+]{1,128})$")

func (r *ActivityLog) MaybePopulateDefaults() error {
	activityLogInterface := interface{}(r)
	if defaulter, ok := activityLogInterface.(goten.Defaulter); ok {
		return defaulter.PopulateDefaults()
	}
	return nil
}

func (r *ActivityLog) GetRawName() gotenresource.Name {
	return r.GetName()
}

func (r *ActivityLog) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

type Name struct {
	ParentName
	ActivityLogId string `firestore:"activityLogId"`
}

func ParseName(name string) (*Name, error) {
	var matches []string
	if matches = regexPath.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetId(matches[1]).
			Name(), nil
	}
	if matches = regexPath_Project.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			SetId(matches[2]).
			Name(), nil
	}
	if matches = regexPath_Organization.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetOrganizationId(matches[1]).
			SetId(matches[2]).
			Name(), nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "unable to parse '%s' as ActivityLog name", name)
}

func MustParseName(name string) *Name {
	result, err := ParseName(name)
	if err != nil {
		panic(err)
	}
	return result
}

func ParseNameOrId(nameOrId string) (*Name, error) {
	name, err := ParseName(nameOrId)
	if err == nil {
		return name, err
	}
	if activityLog_RegexpId.MatchString(nameOrId) {
		return NewNameBuilder().SetId(nameOrId).Name(), nil
	} else {
		return nil, fmt.Errorf("unable to parse '%s' as ActivityLog name or id", name)
	}
}

func (name *Name) SetFromSegments(segments gotenresource.NameSegments) error {
	if len(segments) == 0 {
		return status.Errorf(codes.InvalidArgument, "No segments given for ActivityLog name")
	}
	if err := name.ParentName.SetFromSegments(segments[:len(segments)-1]); err != nil {
		return err
	}
	if segments[len(segments)-1].CollectionLowerJson != "activityLogs" {
		return status.Errorf(codes.InvalidArgument, "unable to use segments %s to form ActivityLog name", segments)
	}
	name.ActivityLogId = segments[len(segments)-1].Id
	return nil
}

func (name *Name) GetProjectName() *iam_project.Name {
	if name == nil {
		return nil
	}
	return name.ParentName.GetProjectName()
}
func (name *Name) GetOrganizationName() *iam_organization.Name {
	if name == nil {
		return nil
	}
	return name.ParentName.GetOrganizationName()
}

func (name *Name) IsSpecified() bool {
	if name == nil || name.Pattern == "" || name.ActivityLogId == "" {
		return false
	}
	return name.ParentName.IsSpecified()
}

func (name *Name) IsFullyQualified() bool {
	if name == nil {
		return false
	}
	if name.ParentName.IsFullyQualified() == false {
		return false
	}
	if name.ActivityLogId == "" || name.ActivityLogId == gotenresource.WildcardId {
		return false
	}
	return true
}

func (name *Name) FullyQualifiedName() (string, error) {
	if !name.IsFullyQualified() {
		return "", status.Errorf(codes.InvalidArgument, "Name for ActivityLog is not fully qualified")
	}
	return fmt.Sprintf("//audit.edgelq.com/%s", name.String()), nil
}

func (name *Name) String() string {
	if name == nil {
		return "<nil>"
	}
	if valueStr, err := name.ProtoString(); err != nil {
		panic(err)
	} else {
		return valueStr
	}
}

func (name *Name) AsReference() *Reference {
	return &Reference{Name: *name}
}

func (name *Name) AsRawReference() gotenresource.Reference {
	return name.AsReference()
}

func (name *Name) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

func (name *Name) GetPattern() gotenresource.NamePattern {
	if name == nil {
		return ""
	}
	return name.Pattern
}

func (name *Name) GetIdParts() map[string]string {
	if name != nil {
		return map[string]string{
			"activityLogId":  name.ActivityLogId,
			"projectId":      name.ProjectId,
			"organizationId": name.OrganizationId,
		}
	}
	return map[string]string{
		"activityLogId":  "",
		"projectId":      "",
		"organizationId": "",
	}
}

func (name *Name) GetSegments() gotenresource.NameSegments {
	if name == nil || name.Pattern == "" {
		return nil
	}
	segments := name.ParentName.GetSegments()
	return append(segments, gotenresource.NameSegment{
		CollectionLowerJson: "activityLogs",
		Id:                  name.ActivityLogId,
	})
}

// implement methods required by protobuf-go library for string-struct conversion

func (name *Name) ProtoString() (string, error) {
	if name == nil {
		return "", nil
	}
	result := ""
	parentPrefix, err := name.ParentName.ProtoString()
	if err != nil {
		return "", err
	}
	if parentPrefix != "" {
		result += parentPrefix + "/"
	}
	result += "activityLogs/" + name.ActivityLogId
	return result, nil
}

func (name *Name) ParseProtoString(data string) error {
	parsed, err := ParseName(data)
	if err != nil {
		return err
	}
	*name = *parsed
	return nil
}

// GotenEqual returns true if other is of same type and paths are equal (implements goten.Equaler interface)
func (name *Name) GotenEqual(other interface{}) bool {
	if other == nil {
		return name == nil
	}
	other1, ok := other.(*Name)
	if !ok {
		other2, ok := other.(Name)
		if ok {
			other1 = &other2
		} else {
			return false
		}
	}
	if other1 == nil {
		return name == nil
	} else if name == nil {
		return false
	}
	if name.ParentName.GotenEqual(other1.ParentName) == false {
		return false
	}
	if name.ActivityLogId != other1.ActivityLogId {
		return false
	}

	return true
}

// Matches is same as GotenEqual, but also will accept "other" if name is wildcard.
func (name *Name) Matches(other interface{}) bool {
	if other == nil {
		return name == nil
	}
	other1, ok := other.(*Name)
	if !ok {
		other2, ok := other.(Name)
		if ok {
			other1 = &other2
		} else {
			return false
		}
	}
	if other1 == nil {
		return name == nil
	} else if name == nil {
		return false
	}
	if name.ParentName.Matches(other1.ParentName) == false {
		return false
	}
	if name.ActivityLogId != other1.ActivityLogId {
		return name.ActivityLogId == gotenresource.WildcardId
	}

	return true
}

// Google CEL integration Type
var celName = types.NewTypeValue("Name", traits.ReceiverType)

func (name *Name) TypeName() string {
	return ".ntt.audit.v1alpha2.ActivityLog.Name"
}

func (name *Name) HasTrait(trait int) bool {
	return trait == traits.ReceiverType
}

func (name *Name) Equal(other ref.Val) ref.Val {
	return types.Bool(false)
}

func (name *Name) Value() interface{} {
	return name
}

func (name *Name) Match(pattern ref.Val) ref.Val {
	return types.Bool(true)
}

func (name *Name) Receive(function string, overload string, args []ref.Val) ref.Val {
	switch function {
	case "getId":
		return types.String(name.ActivityLogId)
	case "satisfies":
		rhsName, err := ParseName(string(args[0].(types.String)))
		if err != nil {
			return types.ValOrErr(celName, "function %s name parse error: %s", function, err)
		}
		return types.Bool(rhsName.Matches(name))
	default:
		return types.ValOrErr(celName, "no such function - %s", function)
	}
}

func (name *Name) ConvertToNative(typeDesc reflect.Type) (interface{}, error) {
	panic("not required")
}

func (name *Name) ConvertToType(typeVal ref.Type) ref.Val {
	switch typeVal.TypeName() {
	case types.StringType.TypeName():
		return types.String(name.String())
	default:
		panic(fmt.Errorf("unable to convert %s to CEL type %s", "Name", typeVal.TypeName()))
	}
}

func (name *Name) Type() ref.Type {
	return celName
}

// implement CustomTypeCliValue method
func (name *Name) SetFromCliFlag(raw string) error {
	parsedName, err := ParseName(raw)
	if err != nil {
		return err
	}
	*name = *parsedName
	return nil
}

type Reference struct {
	Name
	activityLog *ActivityLog
}

func MakeReference(name *Name, activityLog *ActivityLog) (*Reference, error) {
	return &Reference{
		Name:        *name,
		activityLog: activityLog,
	}, nil
}

func ParseReference(name string) (*Reference, error) {
	parsedName, err := ParseName(name)
	if err != nil {
		return nil, err
	}
	return MakeReference(parsedName, nil)
}

func MustParseReference(name string) *Reference {
	result, err := ParseReference(name)
	if err != nil {
		panic(err)
	}
	return result
}

func (ref *Reference) Resolve(resolved *ActivityLog) {
	ref.activityLog = resolved
}

func (ref *Reference) ResolveRaw(res gotenresource.Resource) error {
	if typedRes, ok := res.(*ActivityLog); ok && typedRes != nil {
		ref.Resolve(typedRes)
		return nil
	}
	return status.Errorf(codes.Internal, "Invalid resource type for ActivityLog: %s", reflect.TypeOf(res).Elem().Name())
}

func (ref *Reference) Resolved() bool {
	return ref.activityLog != nil
}

func (ref *Reference) ClearCached() {
	ref.activityLog = nil
}

func (ref *Reference) GetActivityLog() *ActivityLog {
	return ref.activityLog
}

func (ref *Reference) GetRawResource() gotenresource.Resource {
	return ref.activityLog
}

func (ref *Reference) IsFullyQualified() bool {
	if ref == nil {
		return false
	}
	return ref.Name.IsFullyQualified()
}

func (ref *Reference) IsSpecified() bool {
	if ref == nil {
		return false
	}
	return ref.Name.IsSpecified()
}

func (ref *Reference) FullyQualifiedName() (string, error) {
	if !ref.IsFullyQualified() {
		return "", status.Errorf(codes.InvalidArgument, "Name for ActivityLog is not fully qualified")
	}
	return fmt.Sprintf("//audit.edgelq.com/%s", ref.String()), nil
}

func (ref *Reference) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

func (ref *Reference) GetPattern() gotenresource.NamePattern {
	if ref == nil {
		return ""
	}
	return ref.Pattern
}

func (ref *Reference) GetIdParts() map[string]string {
	if ref != nil {
		return ref.Name.GetIdParts()
	}
	return map[string]string{
		"activityLogId":  "",
		"projectId":      "",
		"organizationId": "",
	}
}

func (ref *Reference) GetSegments() gotenresource.NameSegments {
	if ref != nil {
		return ref.Name.GetSegments()
	}
	return nil
}

func (ref *Reference) String() string {
	if ref == nil {
		return "<nil>"
	}
	return ref.Name.String()
}

// implement methods required by protobuf-go library for string-struct conversion

func (ref *Reference) ProtoString() (string, error) {
	if ref == nil {
		return "", nil
	}
	return ref.Name.ProtoString()
}

func (ref *Reference) ParseProtoString(data string) error {
	parsed, err := ParseReference(data)
	if err != nil {
		return err
	}
	*ref = *parsed
	return nil
}

// GotenEqual returns true if other is of same type and paths are equal (implements goten.Equaler interface)
func (ref *Reference) GotenEqual(other interface{}) bool {
	if other == nil {
		return ref == nil
	}
	other1, ok := other.(*Reference)
	if !ok {
		other2, ok := other.(Reference)
		if ok {
			other1 = &other2
		} else {
			return false
		}
	}
	if other1 == nil {
		return ref == nil
	} else if ref == nil {
		return false
	}

	return ref.Name.GotenEqual(other1.Name)
}

// Matches is same as GotenEqual, but also will accept "other" if name is wildcard.
func (name *Reference) Matches(other interface{}) bool {
	if other == nil {
		return name == nil
	}
	other1, ok := other.(*Reference)
	if !ok {
		other2, ok := other.(Reference)
		if ok {
			other1 = &other2
		} else {
			return false
		}
	}
	if other1 == nil {
		return name == nil
	} else if name == nil {
		return false
	}
	return name.Name.Matches(&other1.Name)
}

// implement CustomTypeCliValue method
func (ref *Reference) SetFromCliFlag(raw string) error {
	parsedRef, err := ParseReference(raw)
	if err != nil {
		return err
	}
	*ref = *parsedRef
	return nil
}
