// Code generated by protoc-gen-goten-resource
// Resource: RecoveryStoreShardingInfo
// DO NOT EDIT!!!

package recovery_store_sharding_info

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
	meta "github.com/cloudwan/goten-sdk/types/meta"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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
	_ = &durationpb.Duration{}
	_ = &timestamppb.Timestamp{}
	_ = &meta.Meta{}
)

var recoveryStoreShardingInfo_RegexpId = regexp.MustCompile("^(?P<recovery_store_sharding_info_id>[a-zA-Z0-9_.-]{1,128})$")
var regexPath_Region = regexp.MustCompile("^regions/(?P<region_id>-|[a-z][a-z0-9\\-]{0,28}[a-z0-9])/recoveryStoreShardingInfos/(?P<recovery_store_sharding_info_id>-|[a-zA-Z0-9_.-]{1,128})$")

func (r *RecoveryStoreShardingInfo) MaybePopulateDefaults() error {
	recoveryStoreShardingInfoInterface := interface{}(r)
	if defaulter, ok := recoveryStoreShardingInfoInterface.(goten.Defaulter); ok {
		return defaulter.PopulateDefaults()
	}
	return nil
}

type Name struct {
	ParentName
	RecoveryStoreShardingInfoId string `firestore:"recoveryStoreShardingInfoId"`
}

func ParseName(name string) (*Name, error) {
	var matches []string
	if matches = regexPath_Region.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetRegionId(matches[1]).
			SetId(matches[2]).
			Name(), nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "unable to parse '%s' as RecoveryStoreShardingInfo name", name)
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
	if recoveryStoreShardingInfo_RegexpId.MatchString(nameOrId) {
		return &Name{RecoveryStoreShardingInfoId: nameOrId}, nil
	} else {
		return nil, fmt.Errorf("unable to parse '%s' as RecoveryStoreShardingInfo name or id", nameOrId)
	}
}

func (name *Name) SetFromSegments(segments gotenresource.NameSegments) error {
	if len(segments) == 0 {
		return status.Errorf(codes.InvalidArgument, "No segments given for RecoveryStoreShardingInfo name")
	}
	if err := name.ParentName.SetFromSegments(segments[:len(segments)-1]); err != nil {
		return err
	}
	if segments[len(segments)-1].CollectionLowerJson != "recoveryStoreShardingInfos" {
		return status.Errorf(codes.InvalidArgument, "unable to use segments %s to form RecoveryStoreShardingInfo name", segments)
	}
	name.RecoveryStoreShardingInfoId = segments[len(segments)-1].Id
	return nil
}

func (name *Name) IsSpecified() bool {
	if name == nil || name.Pattern == "" || name.RecoveryStoreShardingInfoId == "" {
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
	if name.RecoveryStoreShardingInfoId == "" || name.RecoveryStoreShardingInfoId == gotenresource.WildcardId {
		return false
	}
	return true
}

func (name *Name) FullyQualifiedName() (string, error) {
	if !name.IsFullyQualified() {
		return "", status.Errorf(codes.InvalidArgument, "Name for RecoveryStoreShardingInfo is not fully qualified")
	}
	return fmt.Sprintf("//monitoring.edgelq.com/%s", name.String()), nil
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
			"regionId":                    name.RegionId,
			"recoveryStoreShardingInfoId": name.RecoveryStoreShardingInfoId,
		}
	}
	return map[string]string{
		"regionId":                    "",
		"recoveryStoreShardingInfoId": "",
	}
}

func (name *Name) GetSegments() gotenresource.NameSegments {
	if name == nil || name.Pattern == "" {
		return nil
	}
	segments := name.ParentName.GetSegments()
	return append(segments, gotenresource.NameSegment{
		CollectionLowerJson: "recoveryStoreShardingInfos",
		Id:                  name.RecoveryStoreShardingInfoId,
	})
}

func (name *Name) GetIParentName() gotenresource.Name {
	if name == nil {
		return (*ParentName)(nil)
	}
	return &name.ParentName
}

func (name *Name) GetIUnderlyingParentName() gotenresource.Name {
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
	result += "recoveryStoreShardingInfos/" + name.RecoveryStoreShardingInfoId
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
	if name.RecoveryStoreShardingInfoId != other1.RecoveryStoreShardingInfoId {
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
	if name.RecoveryStoreShardingInfoId != other1.RecoveryStoreShardingInfoId {
		return name.RecoveryStoreShardingInfoId == gotenresource.WildcardId
	}

	return true
}

// Google CEL integration Type
var celName = types.NewTypeValue("Name", traits.ReceiverType)

func (name *Name) TypeName() string {
	return ".ntt.monitoring.v4.RecoveryStoreShardingInfo.Name"
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
		return types.String(name.RecoveryStoreShardingInfoId)
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
	recoveryStoreShardingInfo *RecoveryStoreShardingInfo
}

func MakeReference(name *Name, recoveryStoreShardingInfo *RecoveryStoreShardingInfo) (*Reference, error) {
	return &Reference{
		Name:                      *name,
		recoveryStoreShardingInfo: recoveryStoreShardingInfo,
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

func (ref *Reference) Resolve(resolved *RecoveryStoreShardingInfo) {
	ref.recoveryStoreShardingInfo = resolved
}

func (ref *Reference) ResolveRaw(res gotenresource.Resource) error {
	if res == nil {
		ref.Resolve(nil)
		return nil
	}
	if typedRes, ok := res.(*RecoveryStoreShardingInfo); ok {
		ref.Resolve(typedRes)
		return nil
	}
	return status.Errorf(codes.Internal, "Invalid resource type for RecoveryStoreShardingInfo: %s", reflect.TypeOf(res).Elem().Name())
}

func (ref *Reference) Resolved() bool {
	return ref != nil && ref.recoveryStoreShardingInfo != nil
}

func (ref *Reference) ClearCached() {
	ref.recoveryStoreShardingInfo = nil
}

func (ref *Reference) GetRecoveryStoreShardingInfo() *RecoveryStoreShardingInfo {
	if ref == nil {
		return nil
	}
	return ref.recoveryStoreShardingInfo
}

func (ref *Reference) GetRawResource() gotenresource.Resource {
	if ref == nil {
		return (*RecoveryStoreShardingInfo)(nil)
	}
	return ref.recoveryStoreShardingInfo
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
		return "", status.Errorf(codes.InvalidArgument, "Name for RecoveryStoreShardingInfo is not fully qualified")
	}
	return fmt.Sprintf("//monitoring.edgelq.com/%s", ref.String()), nil
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
		"regionId":                    "",
		"recoveryStoreShardingInfoId": "",
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
