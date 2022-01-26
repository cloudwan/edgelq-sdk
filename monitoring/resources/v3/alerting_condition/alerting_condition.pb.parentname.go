// Code generated by protoc-gen-goten-resource
// Resource: AlertingCondition
// DO NOT EDIT!!!

package alerting_condition

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
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	monitoring_common "github.com/cloudwan/edgelq-sdk/monitoring/common/v3"
	alerting_policy "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_policy"
	duration "github.com/golang/protobuf/ptypes/duration"
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
	_ = &ntt_meta.Meta{}
	_ = &alerting_policy.AlertingPolicy{}
	_ = &monitoring_common.LabelDescriptor{}
	_ = &duration.Duration{}
)

var parentRegexPath_Project_Region_AlertingPolicy = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})/regions/(?P<region_id>-|[a-zA-Z0-9-]{1,128})/alertingPolicies/(?P<alerting_policy_id>-|[\\w][\\w.-]{0,127})$")

type ParentName struct {
	NamePattern
	ProjectId        string `firestore:"projectId"`
	RegionId         string `firestore:"regionId"`
	AlertingPolicyId string `firestore:"alertingPolicyId"`
}

func ParseParentName(name string) (*ParentName, error) {
	var matches []string
	if matches = parentRegexPath_Project_Region_AlertingPolicy.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			SetRegionId(matches[2]).
			SetAlertingPolicyId(matches[3]).
			Parent(), nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "unable to parse '%s' as AlertingCondition parent name", name)
}

func MustParseParentName(name string) *ParentName {
	result, err := ParseParentName(name)
	if err != nil {
		panic(err)
	}
	return result
}

func (name *ParentName) SetFromSegments(segments gotenresource.NameSegments) error {
	if len(segments) == 3 && segments[0].CollectionLowerJson == "projects" && segments[1].CollectionLowerJson == "regions" && segments[2].CollectionLowerJson == "alertingPolicies" {
		name.Pattern = NamePattern_Project_Region_AlertingPolicy
		name.ProjectId = segments[0].Id
		name.RegionId = segments[1].Id
		name.AlertingPolicyId = segments[2].Id
		return nil
	}
	return status.Errorf(codes.InvalidArgument, "unable to use segments %s to form AlertingCondition parent name", segments)
}

func (name *ParentName) GetAlertingPolicyName() *alerting_policy.Name {
	if name == nil {
		return nil
	}

	switch name.Pattern {
	case NamePattern_Project_Region_AlertingPolicy:
		return alerting_policy.NewNameBuilder().
			SetId(name.AlertingPolicyId).
			SetProjectId(name.ProjectId).
			SetRegionId(name.RegionId).
			Name()
	default:
		return nil
	}
}

func (name *ParentName) IsSpecified() bool {
	if name == nil {
		return false
	}
	return name.Pattern == NamePattern_Project_Region_AlertingPolicy
}

func (name *ParentName) IsFullyQualified() bool {
	if name == nil || name.Pattern == "" {
		return false
	}

	switch name.Pattern {
	case NamePattern_Project_Region_AlertingPolicy:
		return name.ProjectId != "" && name.ProjectId != gotenresource.WildcardId && name.RegionId != "" && name.RegionId != gotenresource.WildcardId && name.AlertingPolicyId != "" && name.AlertingPolicyId != gotenresource.WildcardId
	}

	return false
}

func (name *ParentName) FullyQualifiedName() (string, error) {
	if !name.IsFullyQualified() {
		return "", status.Errorf(codes.InvalidArgument, "Parent name for AlertingCondition is not fully qualified")
	}
	return fmt.Sprintf("//monitoring.edgelq.com/%s", name.String()), nil
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
			"projectId":        name.ProjectId,
			"regionId":         name.RegionId,
			"alertingPolicyId": name.AlertingPolicyId,
		}
	}
	return map[string]string{
		"projectId":        "",
		"regionId":         "",
		"alertingPolicyId": "",
	}
}

func (name *ParentName) GetSegments() gotenresource.NameSegments {
	if name == nil {
		return nil
	}

	switch name.Pattern {
	case NamePattern_Project_Region_AlertingPolicy:
		return gotenresource.NameSegments{
			gotenresource.NameSegment{
				CollectionLowerJson: "projects",
				Id:                  name.ProjectId,
			},
			gotenresource.NameSegment{
				CollectionLowerJson: "regions",
				Id:                  name.RegionId,
			},
			gotenresource.NameSegment{
				CollectionLowerJson: "alertingPolicies",
				Id:                  name.AlertingPolicyId,
			},
		}
	}
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
	case NamePattern_Project_Region_AlertingPolicy:
		return ancestor == "projects" || ancestor == "regions" || ancestor == "alertingPolicies"
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
	case NamePattern_Project_Region_AlertingPolicy:
		return "projects/" + name.ProjectId + "/regions/" + name.RegionId + "/alertingPolicies/" + name.AlertingPolicyId, nil
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
	if name.RegionId != other1.RegionId {
		return false
	}
	if name.AlertingPolicyId != other1.AlertingPolicyId {
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
	case NamePattern_Project_Region_AlertingPolicy:
		if name.ProjectId != other1.ProjectId &&
			name.ProjectId != gotenresource.WildcardId {
			return false
		}
		if name.RegionId != other1.RegionId &&
			name.RegionId != gotenresource.WildcardId {
			return false
		}
		if name.AlertingPolicyId != other1.AlertingPolicyId &&
			name.AlertingPolicyId != gotenresource.WildcardId {
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
	alertingPolicy *alerting_policy.AlertingPolicy
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
func (ref *ParentReference) GetAlertingPolicyReference() *alerting_policy.Reference {
	if ref == nil {
		return nil
	}

	switch ref.Pattern {
	case NamePattern_Project_Region_AlertingPolicy:
		return alerting_policy.NewNameBuilder().
			SetId(ref.AlertingPolicyId).
			SetProjectId(ref.ProjectId).
			SetRegionId(ref.RegionId).
			Reference()
	default:
		return nil
	}
}

func (ref *ParentReference) GetUnderlyingReference() gotenresource.Reference {
	if ref == nil {
		return nil
	}
	alertingPolicyRef := ref.GetAlertingPolicyReference()
	if alertingPolicyRef != nil {
		return alertingPolicyRef
	}

	return nil
}

func (ref *ParentReference) ResolveRaw(res gotenresource.Resource) error {
	switch typedRes := res.(type) {
	case *alerting_policy.AlertingPolicy:
		if name := ref.GetAlertingPolicyName(); name == nil {
			return status.Errorf(codes.InvalidArgument, "cannot set AlertingPolicy as parent of AlertingCondition, because pattern does not match")
		}
		ref.alertingPolicy = typedRes
		return nil
	default:
		return status.Errorf(codes.Internal, "Invalid parent type for AlertingCondition, got %s", reflect.TypeOf(res).Elem().Name())
	}
}

func (ref *ParentReference) Resolved() bool {
	if name := ref.GetAlertingPolicyName(); name != nil {
		return ref.alertingPolicy != nil
	}
	return true
}

func (ref *ParentReference) ClearCached() {
	ref.alertingPolicy = nil
}

func (ref *ParentReference) GetAlertingPolicy() *alerting_policy.AlertingPolicy {
	if ref == nil {
		return nil
	}
	return ref.alertingPolicy
}

func (ref *ParentReference) GetRawResource() gotenresource.Resource {
	if name := ref.ParentName.GetAlertingPolicyName(); name != nil {
		return ref.alertingPolicy
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
		"projectId":        "",
		"regionId":         "",
		"alertingPolicyId": "",
	}
}

func (ref *ParentReference) GetSegments() gotenresource.NameSegments {
	if ref != nil {
		return ref.ParentName.GetSegments()
	}
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
	if ref.alertingPolicy != other1.alertingPolicy {
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

// implement CustomTypeCliValue method
func (ref *ParentReference) SetFromCliFlag(raw string) error {
	parsedRef, err := ParseParentReference(raw)
	if err != nil {
		return err
	}
	*ref = *parsedRef
	return nil
}
