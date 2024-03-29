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
	common "github.com/cloudwan/edgelq-sdk/audit/resources/v1alpha2/common"
	rpc "github.com/cloudwan/edgelq-sdk/common/rpc"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	anypb "google.golang.org/protobuf/types/known/anypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// ensure the imports are used
var (
	_ = codes.NotFound
	_ = new(fmt.Stringer)
	_ = new(proto.Message)
	_ = status.Status{}
	_ = url.URL{}
	_ = strings.Builder{}

	_ = new(goten.GotenMessage)
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &common.Authentication{}
	_ = &rpc.Status{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &anypb.Any{}
	_ = &fieldmaskpb.FieldMask{}
	_ = &timestamppb.Timestamp{}
	_ = &meta.Meta{}
)

var parentRegexPath = regexp.MustCompile("^$")
var parentRegexPath_Project = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})$")
var parentRegexPath_Organization = regexp.MustCompile("^organizations/(?P<organization_id>-|[\\w][\\w.-]{0,127})$")

type ParentName struct {
	NamePattern
	ProjectId      string `firestore:"projectId"`
	OrganizationId string `firestore:"organizationId"`
}

func ParseParentName(name string) (*ParentName, error) {
	var matches []string
	if matches = parentRegexPath.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			Parent(), nil
	}
	if matches = parentRegexPath_Project.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			Parent(), nil
	}
	if matches = parentRegexPath_Organization.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetOrganizationId(matches[1]).
			Parent(), nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "unable to parse '%s' as ActivityLog parent name", name)
}

func MustParseParentName(name string) *ParentName {
	result, err := ParseParentName(name)
	if err != nil {
		panic(err)
	}
	return result
}

func (name *ParentName) SetFromSegments(segments gotenresource.NameSegments) error {
	if len(segments) == 0 {
		name.Pattern = NamePattern_Nil
		return nil
	} else if len(segments) == 1 && segments[0].CollectionLowerJson == "projects" {
		name.Pattern = NamePattern_Project
		name.ProjectId = segments[0].Id
		return nil
	} else if len(segments) == 1 && segments[0].CollectionLowerJson == "organizations" {
		name.Pattern = NamePattern_Organization
		name.OrganizationId = segments[0].Id
		return nil
	}
	return status.Errorf(codes.InvalidArgument, "unable to use segments %s to form ActivityLog parent name", segments)
}

func (name *ParentName) GetProjectName() *iam_project.Name {
	if name == nil {
		return nil
	}

	switch name.Pattern {
	case NamePattern_Project:
		return iam_project.NewNameBuilder().
			SetId(name.ProjectId).
			Name()
	default:
		return nil
	}
}

func (name *ParentName) GetOrganizationName() *iam_organization.Name {
	if name == nil {
		return nil
	}

	switch name.Pattern {
	case NamePattern_Organization:
		return iam_organization.NewNameBuilder().
			SetId(name.OrganizationId).
			Name()
	default:
		return nil
	}
}

func (name *ParentName) IsSpecified() bool {
	if name == nil || name.Pattern == "" {
		return false
	}
	switch name.Pattern {
	case NamePattern_Nil:
		return true
	case NamePattern_Project:
		return name.ProjectId != ""
	case NamePattern_Organization:
		return name.OrganizationId != ""
	}
	return false
}

func (name *ParentName) IsFullyQualified() bool {
	if name == nil || name.Pattern == "" {
		return false
	}

	switch name.Pattern {
	case NamePattern_Nil:
		return true
	case NamePattern_Project:
		return name.ProjectId != "" && name.ProjectId != gotenresource.WildcardId
	case NamePattern_Organization:
		return name.OrganizationId != "" && name.OrganizationId != gotenresource.WildcardId
	}

	return false
}

func (name *ParentName) FullyQualifiedName() (string, error) {
	if !name.IsFullyQualified() {
		return "", status.Errorf(codes.InvalidArgument, "Parent name for ActivityLog is not fully qualified")
	}
	return fmt.Sprintf("//audit.edgelq.com/%s", name.String()), nil
}

func (name *ParentName) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

func (name *ParentName) GetPattern() gotenresource.NamePattern {
	if name == nil {
		return ""
	}
	return name.Pattern
}

func (name *ParentName) GetIdParts() map[string]string {
	if name != nil {
		return map[string]string{
			"projectId":      name.ProjectId,
			"organizationId": name.OrganizationId,
		}
	}
	return map[string]string{
		"projectId":      "",
		"organizationId": "",
	}
}

func (name *ParentName) GetSegments() gotenresource.NameSegments {
	if name == nil {
		return nil
	}

	switch name.Pattern {
	case NamePattern_Nil:
		return gotenresource.NameSegments{}
	case NamePattern_Project:
		return gotenresource.NameSegments{
			gotenresource.NameSegment{
				CollectionLowerJson: "projects",
				Id:                  name.ProjectId,
			},
		}
	case NamePattern_Organization:
		return gotenresource.NameSegments{
			gotenresource.NameSegment{
				CollectionLowerJson: "organizations",
				Id:                  name.OrganizationId,
			},
		}
	}
	return nil
}

func (name *ParentName) GetIParentName() gotenresource.Name {
	return nil
}

func (name *ParentName) GetIUnderlyingParentName() gotenresource.Name {
	return nil
}

func (name *ParentName) String() string {
	if name == nil {
		return "<nil>"
	}

	if valueStr, err := name.ProtoString(); err != nil {
		panic(err)
	} else {
		return valueStr
	}
}

func (name *ParentName) DescendsFrom(ancestor string) bool {
	if name == nil {
		return false
	}

	switch name.Pattern {
	case NamePattern_Project:
		return ancestor == "projects"
	case NamePattern_Organization:
		return ancestor == "organizations"
	}

	return false
}

func (name *ParentName) AsReference() *ParentReference {
	return &ParentReference{ParentName: *name}
}

func (name *ParentName) AsRawReference() gotenresource.Reference {
	return name.AsReference()
}

// implement methods required by protobuf-go library for string-struct conversion

func (name *ParentName) ProtoString() (string, error) {
	if name == nil {
		return "", nil
	}
	switch name.Pattern {
	case NamePattern_Project:
		return "projects/" + name.ProjectId, nil
	case NamePattern_Organization:
		return "organizations/" + name.OrganizationId, nil
	}
	return "", nil
}

func (name *ParentName) ParseProtoString(data string) error {
	parsed, err := ParseParentName(data)
	if err != nil {
		return err
	}
	*name = *parsed
	return nil
}

// GotenEqual returns true if other is of same type and paths are equal (implements goten.Equaler interface)
func (name *ParentName) GotenEqual(other interface{}) bool {
	if other == nil {
		return name == nil
	}
	other1, ok := other.(*ParentName)
	if !ok {
		other2, ok := other.(ParentName)
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
	if name.ProjectId != other1.ProjectId {
		return false
	}
	if name.OrganizationId != other1.OrganizationId {
		return false
	}
	if name.Pattern != other1.Pattern {
		return false
	}

	return true
}

// Matches is same as GotenEqual, but also will accept "other" if name is wildcard.
func (name *ParentName) Matches(other interface{}) bool {
	if other == nil {
		return name == nil
	}
	other1, ok := other.(*ParentName)
	if !ok {
		other2, ok := other.(ParentName)
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

	if name.Pattern != other1.Pattern {
		return false
	}
	switch name.Pattern {
	case NamePattern_Project:
		if name.ProjectId != other1.ProjectId &&
			name.ProjectId != gotenresource.WildcardId {
			return false
		}
	case NamePattern_Organization:
		if name.OrganizationId != other1.OrganizationId &&
			name.OrganizationId != gotenresource.WildcardId {
			return false
		}
	}

	return true
}

// implement CustomTypeCliValue method
func (name *ParentName) SetFromCliFlag(raw string) error {
	parsedName, err := ParseParentName(raw)
	if err != nil {
		return err
	}
	*name = *parsedName
	return nil
}

type ParentReference struct {
	ParentName
	project      *iam_project.Project
	organization *iam_organization.Organization
}

func MakeParentReference(name *ParentName) (*ParentReference, error) {
	return &ParentReference{
		ParentName: *name,
	}, nil
}

func ParseParentReference(name string) (*ParentReference, error) {
	parsedName, err := ParseParentName(name)
	if err != nil {
		return nil, err
	}
	return MakeParentReference(parsedName)
}

func MustParseParentReference(name string) *ParentReference {
	result, err := ParseParentReference(name)
	if err != nil {
		panic(err)
	}
	return result
}
func (ref *ParentReference) GetProjectReference() *iam_project.Reference {
	if ref == nil {
		return nil
	}

	switch ref.Pattern {
	case NamePattern_Project:
		return iam_project.NewNameBuilder().
			SetId(ref.ProjectId).
			Reference()
	default:
		return nil
	}
}
func (ref *ParentReference) GetOrganizationReference() *iam_organization.Reference {
	if ref == nil {
		return nil
	}

	switch ref.Pattern {
	case NamePattern_Organization:
		return iam_organization.NewNameBuilder().
			SetId(ref.OrganizationId).
			Reference()
	default:
		return nil
	}
}

func (ref *ParentReference) GetUnderlyingReference() gotenresource.Reference {
	if ref == nil {
		return nil
	}
	projectRef := ref.GetProjectReference()
	if projectRef != nil {
		return projectRef
	}
	organizationRef := ref.GetOrganizationReference()
	if organizationRef != nil {
		return organizationRef
	}

	return nil
}

func (ref *ParentReference) ResolveRaw(res gotenresource.Resource) error {
	switch typedRes := res.(type) {
	case *iam_project.Project:
		if name := ref.GetProjectName(); name == nil {
			return status.Errorf(codes.InvalidArgument, "cannot set Project as parent of ActivityLog, because pattern does not match")
		}
		ref.project = typedRes
		return nil
	case *iam_organization.Organization:
		if name := ref.GetOrganizationName(); name == nil {
			return status.Errorf(codes.InvalidArgument, "cannot set Organization as parent of ActivityLog, because pattern does not match")
		}
		ref.organization = typedRes
		return nil
	default:
		return status.Errorf(codes.Internal, "Invalid parent type for ActivityLog, got %s", reflect.TypeOf(res).Elem().Name())
	}
}

func (ref *ParentReference) Resolved() bool {
	if name := ref.GetProjectName(); name != nil {
		return ref.project != nil
	}
	if name := ref.GetOrganizationName(); name != nil {
		return ref.organization != nil
	}
	return true
}

func (ref *ParentReference) ClearCached() {
	ref.project = nil
	ref.organization = nil
}

func (ref *ParentReference) GetProject() *iam_project.Project {
	if ref == nil {
		return nil
	}
	return ref.project
}
func (ref *ParentReference) GetOrganization() *iam_organization.Organization {
	if ref == nil {
		return nil
	}
	return ref.organization
}

func (ref *ParentReference) GetRawResource() gotenresource.Resource {
	if name := ref.ParentName.GetProjectName(); name != nil {
		return ref.project
	}
	if name := ref.ParentName.GetOrganizationName(); name != nil {
		return ref.organization
	}
	return nil
}

func (ref *ParentReference) IsFullyQualified() bool {
	if ref == nil {
		return false
	}
	return ref.ParentName.IsFullyQualified()
}

func (ref *ParentReference) IsSpecified() bool {
	if ref == nil {
		return false
	}
	return ref.ParentName.IsSpecified()
}

func (ref *ParentReference) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

func (ref *ParentReference) GetPattern() gotenresource.NamePattern {
	if ref == nil {
		return ""
	}
	return ref.Pattern
}

func (ref *ParentReference) GetIdParts() map[string]string {
	if ref != nil {
		return ref.ParentName.GetIdParts()
	}
	return map[string]string{
		"projectId":      "",
		"organizationId": "",
	}
}

func (ref *ParentReference) GetSegments() gotenresource.NameSegments {
	if ref != nil {
		return ref.ParentName.GetSegments()
	}
	return nil
}

func (ref *ParentReference) GetIParentName() gotenresource.Name {
	return nil
}

func (ref *ParentReference) GetIUnderlyingParentName() gotenresource.Name {
	return nil
}

func (ref *ParentReference) String() string {
	if ref == nil {
		return "<nil>"
	}
	return ref.ParentName.String()
}

// implement methods required by protobuf-go library for string-struct conversion

func (ref *ParentReference) ProtoString() (string, error) {
	if ref == nil {
		return "", nil
	}
	return ref.ParentName.ProtoString()
}

func (ref *ParentReference) ParseProtoString(data string) error {
	parsed, err := ParseParentReference(data)
	if err != nil {
		return err
	}
	*ref = *parsed
	return nil
}

// GotenEqual returns true if other is of same type and paths are equal (implements goten.Equaler interface)
func (ref *ParentReference) GotenEqual(other interface{}) bool {
	if other == nil {
		return ref == nil
	}
	other1, ok := other.(*ParentReference)
	if !ok {
		other2, ok := other.(ParentReference)
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
	if ref.project != other1.project {
		return false
	}
	if ref.organization != other1.organization {
		return false
	}

	return ref.ParentName.GotenEqual(other1.ParentName)
}

// Matches is same as GotenEqual, but also will accept "other" if name is wildcard.
func (name *ParentReference) Matches(other interface{}) bool {
	if other == nil {
		return name == nil
	}
	other1, ok := other.(*ParentReference)
	if !ok {
		other2, ok := other.(ParentReference)
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
	return name.ParentName.Matches(&other1.ParentName)
}

// Google CEL integration Type
var celParentName = types.NewTypeValue("ParentName", traits.ReceiverType)

func (name *ParentName) TypeName() string {
	return ".ntt.audit.v1alpha2.ActivityLog.ParentName"
}

func (name *ParentName) HasTrait(trait int) bool {
	return trait == traits.ReceiverType
}

func (name *ParentName) Equal(other ref.Val) ref.Val {
	return types.Bool(false)
}

func (name *ParentName) Value() interface{} {
	return name
}

func (name *ParentName) Match(pattern ref.Val) ref.Val {
	return types.Bool(true)
}

func (name *ParentName) Receive(function string, overload string, args []ref.Val) ref.Val {
	switch function {
	case "satisfies":
		rhsName, err := ParseParentName(string(args[0].(types.String)))
		if err != nil {
			return types.ValOrErr(celParentName, "function %s name parse error: %s", function, err)
		}
		return types.Bool(rhsName.Matches(name))
	default:
		return types.ValOrErr(celParentName, "no such function - %s", function)
	}
}

func (name *ParentName) ConvertToNative(typeDesc reflect.Type) (interface{}, error) {
	panic("not required")
}

func (name *ParentName) ConvertToType(typeVal ref.Type) ref.Val {
	switch typeVal.TypeName() {
	case types.StringType.TypeName():
		return types.String(name.String())
	default:
		panic(fmt.Errorf("unable to convert %s to CEL type %s", "ParentName", typeVal.TypeName()))
	}
}

func (name *ParentName) Type() ref.Type {
	return celParentName
}

// implement CustomTypeCliValue method
func (ref *ParentReference) SetFromCliFlag(raw string) error {
	parsedRef, err := ParseParentReference(raw)
	if err != nil {
		return err
	}
	*ref = *parsedRef
	return nil
}
