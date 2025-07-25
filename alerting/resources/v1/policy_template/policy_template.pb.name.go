// Code generated by protoc-gen-goten-resource
// Resource: PolicyTemplate
// DO NOT EDIT!!!

package policy_template

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/cloudwan/goten-sdk/runtime/goten"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	rcommon "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/common"
	document "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/document"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
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
	_ = &document.Document{}
	_ = &rcommon.LogCndSpec{}
	_ = &iam_project.Project{}
	_ = &meta.Meta{}
)

var policyTemplate_RegexpId = regexp.MustCompile("^(?P<policy_template_id>[\\w][\\w.-]{0,127})$")
var regexPath_Project = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})/policyTemplates/(?P<policy_template_id>-|[\\w][\\w.-]{0,127})$")

func (r *PolicyTemplate) MaybePopulateDefaults() error {
	policyTemplateInterface := interface{}(r)
	if defaulter, ok := policyTemplateInterface.(goten.Defaulter); ok {
		return defaulter.PopulateDefaults()
	}
	return nil
}

type Name struct {
	ParentName
	PolicyTemplateId string `firestore:"policyTemplateId"`
}

func ParseName(name string) (*Name, error) {
	var matches []string
	if matches = regexPath_Project.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			SetId(matches[2]).
			Name(), nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "unable to parse '%s' as PolicyTemplate name", name)
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
	if policyTemplate_RegexpId.MatchString(nameOrId) {
		return &Name{PolicyTemplateId: nameOrId}, nil
	} else {
		return nil, fmt.Errorf("unable to parse '%s' as PolicyTemplate name or id", nameOrId)
	}
}

func (name *Name) SetFromSegments(segments gotenresource.NameSegments) error {
	if len(segments) == 0 {
		return status.Errorf(codes.InvalidArgument, "No segments given for PolicyTemplate name")
	}
	if err := name.ParentName.SetFromSegments(segments[:len(segments)-1]); err != nil {
		return err
	}
	if segments[len(segments)-1].CollectionLowerJson != "policyTemplates" {
		return status.Errorf(codes.InvalidArgument, "unable to use segments %s to form PolicyTemplate name", segments)
	}
	name.PolicyTemplateId = segments[len(segments)-1].Id
	return nil
}

func (name *Name) GetProjectName() *iam_project.Name {
	if name == nil {
		return nil
	}
	return name.ParentName.GetProjectName()
}

func (name *Name) IsSpecified() bool {
	if name == nil || name.Pattern == "" || name.PolicyTemplateId == "" {
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
	if name.PolicyTemplateId == "" || name.PolicyTemplateId == gotenresource.WildcardId {
		return false
	}
	return true
}

func (name *Name) FullyQualifiedName() (string, error) {
	if !name.IsFullyQualified() {
		return "", status.Errorf(codes.InvalidArgument, "Name for PolicyTemplate is not fully qualified")
	}
	return fmt.Sprintf("//alerting.edgelq.com/%s", name.String()), nil
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
			"projectId":        name.ProjectId,
			"policyTemplateId": name.PolicyTemplateId,
		}
	}
	return map[string]string{
		"projectId":        "",
		"policyTemplateId": "",
	}
}

func (name *Name) GetSegments() gotenresource.NameSegments {
	if name == nil || name.Pattern == "" {
		return nil
	}
	segments := name.ParentName.GetSegments()
	return append(segments, gotenresource.NameSegment{
		CollectionLowerJson: "policyTemplates",
		Id:                  name.PolicyTemplateId,
	})
}

func (name *Name) GetIParentName() gotenresource.Name {
	if name == nil {
		return (*ParentName)(nil)
	}
	return &name.ParentName
}

func (name *Name) GetIUnderlyingParentName() gotenresource.Name {
	if parentName := name.GetProjectName(); parentName != nil {
		return parentName
	}
	return nil
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
	result += "policyTemplates/" + name.PolicyTemplateId
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
	if name.PolicyTemplateId != other1.PolicyTemplateId {
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
	if name.PolicyTemplateId != other1.PolicyTemplateId {
		return name.PolicyTemplateId == gotenresource.WildcardId
	}

	return true
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
	policyTemplate *PolicyTemplate
}

func MakeReference(name *Name, policyTemplate *PolicyTemplate) (*Reference, error) {
	return &Reference{
		Name:           *name,
		policyTemplate: policyTemplate,
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

func (ref *Reference) Resolve(resolved *PolicyTemplate) {
	ref.policyTemplate = resolved
}

func (ref *Reference) ResolveRaw(res gotenresource.Resource) error {
	if res == nil {
		ref.Resolve(nil)
		return nil
	}
	if typedRes, ok := res.(*PolicyTemplate); ok {
		ref.Resolve(typedRes)
		return nil
	}
	return status.Errorf(codes.Internal, "Invalid resource type for PolicyTemplate: %s", reflect.TypeOf(res).Elem().Name())
}

func (ref *Reference) Resolved() bool {
	return ref != nil && ref.policyTemplate != nil
}

func (ref *Reference) ClearCached() {
	ref.policyTemplate = nil
}

func (ref *Reference) GetPolicyTemplate() *PolicyTemplate {
	if ref == nil {
		return nil
	}
	return ref.policyTemplate
}

func (ref *Reference) GetRawResource() gotenresource.Resource {
	if ref == nil {
		return (*PolicyTemplate)(nil)
	}
	return ref.policyTemplate
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
		return "", status.Errorf(codes.InvalidArgument, "Name for PolicyTemplate is not fully qualified")
	}
	return fmt.Sprintf("//alerting.edgelq.com/%s", ref.String()), nil
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
		"projectId":        "",
		"policyTemplateId": "",
	}
}

func (ref *Reference) GetSegments() gotenresource.NameSegments {
	if ref != nil {
		return ref.Name.GetSegments()
	}
	return nil
}

func (ref *Reference) GetIParentName() gotenresource.Name {
	if ref == nil {
		return (*ParentName)(nil)
	}
	return ref.Name.GetIParentName()
}

func (ref *Reference) GetIUnderlyingParentName() gotenresource.Name {
	if ref != nil {
		return ref.Name.GetIUnderlyingParentName()
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
