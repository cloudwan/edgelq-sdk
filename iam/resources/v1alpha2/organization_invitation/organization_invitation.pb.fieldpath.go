// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha2/organization_invitation.proto
// DO NOT EDIT!!!

package organization_invitation

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	"github.com/cloudwan/goten-sdk/runtime/strcase"
)

// proto imports
import (
	iam_invitation "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/invitation"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(json.Marshaler)
	_ = new(fmt.Stringer)
	_ = reflect.DeepEqual
	_ = strings.Builder{}
	_ = time.Second

	_ = strcase.ToLowerCamel
	_ = codes.NotFound
	_ = status.Status{}
	_ = protojson.UnmarshalOptions{}
	_ = new(proto.Message)
	_ = protoregistry.GlobalTypes

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &iam_invitation.Actor{}
	_ = &organization.Organization{}
	_ = &meta.Meta{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type OrganizationInvitation_FieldPath interface {
	gotenobject.FieldPath
	Selector() OrganizationInvitation_FieldPathSelector
	Get(source *OrganizationInvitation) []interface{}
	GetSingle(source *OrganizationInvitation) (interface{}, bool)
	ClearValue(item *OrganizationInvitation)

	// Those methods build corresponding OrganizationInvitation_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) OrganizationInvitation_FieldPathValue
	WithIArrayOfValues(values interface{}) OrganizationInvitation_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) OrganizationInvitation_FieldPathArrayItemValue
}

type OrganizationInvitation_FieldPathSelector int32

const (
	OrganizationInvitation_FieldPathSelectorName       OrganizationInvitation_FieldPathSelector = 0
	OrganizationInvitation_FieldPathSelectorInvitation OrganizationInvitation_FieldPathSelector = 1
	OrganizationInvitation_FieldPathSelectorMetadata   OrganizationInvitation_FieldPathSelector = 2
)

func (s OrganizationInvitation_FieldPathSelector) String() string {
	switch s {
	case OrganizationInvitation_FieldPathSelectorName:
		return "name"
	case OrganizationInvitation_FieldPathSelectorInvitation:
		return "invitation"
	case OrganizationInvitation_FieldPathSelectorMetadata:
		return "metadata"
	default:
		panic(fmt.Sprintf("Invalid selector for OrganizationInvitation: %d", s))
	}
}

func BuildOrganizationInvitation_FieldPath(fp gotenobject.RawFieldPath) (OrganizationInvitation_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object OrganizationInvitation")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &OrganizationInvitation_FieldTerminalPath{selector: OrganizationInvitation_FieldPathSelectorName}, nil
		case "invitation":
			return &OrganizationInvitation_FieldTerminalPath{selector: OrganizationInvitation_FieldPathSelectorInvitation}, nil
		case "metadata":
			return &OrganizationInvitation_FieldTerminalPath{selector: OrganizationInvitation_FieldPathSelectorMetadata}, nil
		}
	} else {
		switch fp[0] {
		case "invitation":
			if subpath, err := iam_invitation.BuildInvitation_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &OrganizationInvitation_FieldSubPath{selector: OrganizationInvitation_FieldPathSelectorInvitation, subPath: subpath}, nil
			}
		case "metadata":
			if subpath, err := meta.BuildMeta_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &OrganizationInvitation_FieldSubPath{selector: OrganizationInvitation_FieldPathSelectorMetadata, subPath: subpath}, nil
			}
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object OrganizationInvitation", fp)
}

func ParseOrganizationInvitation_FieldPath(rawField string) (OrganizationInvitation_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildOrganizationInvitation_FieldPath(fp)
}

func MustParseOrganizationInvitation_FieldPath(rawField string) OrganizationInvitation_FieldPath {
	fp, err := ParseOrganizationInvitation_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type OrganizationInvitation_FieldTerminalPath struct {
	selector OrganizationInvitation_FieldPathSelector
}

var _ OrganizationInvitation_FieldPath = (*OrganizationInvitation_FieldTerminalPath)(nil)

func (fp *OrganizationInvitation_FieldTerminalPath) Selector() OrganizationInvitation_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *OrganizationInvitation_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *OrganizationInvitation_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source OrganizationInvitation
func (fp *OrganizationInvitation_FieldTerminalPath) Get(source *OrganizationInvitation) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case OrganizationInvitation_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		case OrganizationInvitation_FieldPathSelectorInvitation:
			if source.Invitation != nil {
				values = append(values, source.Invitation)
			}
		case OrganizationInvitation_FieldPathSelectorMetadata:
			if source.Metadata != nil {
				values = append(values, source.Metadata)
			}
		default:
			panic(fmt.Sprintf("Invalid selector for OrganizationInvitation: %d", fp.selector))
		}
	}
	return
}

func (fp *OrganizationInvitation_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*OrganizationInvitation))
}

// GetSingle returns value pointed by specific field of from source OrganizationInvitation
func (fp *OrganizationInvitation_FieldTerminalPath) GetSingle(source *OrganizationInvitation) (interface{}, bool) {
	switch fp.selector {
	case OrganizationInvitation_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	case OrganizationInvitation_FieldPathSelectorInvitation:
		res := source.GetInvitation()
		return res, res != nil
	case OrganizationInvitation_FieldPathSelectorMetadata:
		res := source.GetMetadata()
		return res, res != nil
	default:
		panic(fmt.Sprintf("Invalid selector for OrganizationInvitation: %d", fp.selector))
	}
}

func (fp *OrganizationInvitation_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*OrganizationInvitation))
}

// GetDefault returns a default value of the field type
func (fp *OrganizationInvitation_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case OrganizationInvitation_FieldPathSelectorName:
		return (*Name)(nil)
	case OrganizationInvitation_FieldPathSelectorInvitation:
		return (*iam_invitation.Invitation)(nil)
	case OrganizationInvitation_FieldPathSelectorMetadata:
		return (*meta.Meta)(nil)
	default:
		panic(fmt.Sprintf("Invalid selector for OrganizationInvitation: %d", fp.selector))
	}
}

func (fp *OrganizationInvitation_FieldTerminalPath) ClearValue(item *OrganizationInvitation) {
	if item != nil {
		switch fp.selector {
		case OrganizationInvitation_FieldPathSelectorName:
			item.Name = nil
		case OrganizationInvitation_FieldPathSelectorInvitation:
			item.Invitation = nil
		case OrganizationInvitation_FieldPathSelectorMetadata:
			item.Metadata = nil
		default:
			panic(fmt.Sprintf("Invalid selector for OrganizationInvitation: %d", fp.selector))
		}
	}
}

func (fp *OrganizationInvitation_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*OrganizationInvitation))
}

// IsLeaf - whether field path is holds simple value
func (fp *OrganizationInvitation_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == OrganizationInvitation_FieldPathSelectorName
}

func (fp *OrganizationInvitation_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *OrganizationInvitation_FieldTerminalPath) WithIValue(value interface{}) OrganizationInvitation_FieldPathValue {
	switch fp.selector {
	case OrganizationInvitation_FieldPathSelectorName:
		return &OrganizationInvitation_FieldTerminalPathValue{OrganizationInvitation_FieldTerminalPath: *fp, value: value.(*Name)}
	case OrganizationInvitation_FieldPathSelectorInvitation:
		return &OrganizationInvitation_FieldTerminalPathValue{OrganizationInvitation_FieldTerminalPath: *fp, value: value.(*iam_invitation.Invitation)}
	case OrganizationInvitation_FieldPathSelectorMetadata:
		return &OrganizationInvitation_FieldTerminalPathValue{OrganizationInvitation_FieldTerminalPath: *fp, value: value.(*meta.Meta)}
	default:
		panic(fmt.Sprintf("Invalid selector for OrganizationInvitation: %d", fp.selector))
	}
}

func (fp *OrganizationInvitation_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *OrganizationInvitation_FieldTerminalPath) WithIArrayOfValues(values interface{}) OrganizationInvitation_FieldPathArrayOfValues {
	fpaov := &OrganizationInvitation_FieldTerminalPathArrayOfValues{OrganizationInvitation_FieldTerminalPath: *fp}
	switch fp.selector {
	case OrganizationInvitation_FieldPathSelectorName:
		return &OrganizationInvitation_FieldTerminalPathArrayOfValues{OrganizationInvitation_FieldTerminalPath: *fp, values: values.([]*Name)}
	case OrganizationInvitation_FieldPathSelectorInvitation:
		return &OrganizationInvitation_FieldTerminalPathArrayOfValues{OrganizationInvitation_FieldTerminalPath: *fp, values: values.([]*iam_invitation.Invitation)}
	case OrganizationInvitation_FieldPathSelectorMetadata:
		return &OrganizationInvitation_FieldTerminalPathArrayOfValues{OrganizationInvitation_FieldTerminalPath: *fp, values: values.([]*meta.Meta)}
	default:
		panic(fmt.Sprintf("Invalid selector for OrganizationInvitation: %d", fp.selector))
	}
	return fpaov
}

func (fp *OrganizationInvitation_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *OrganizationInvitation_FieldTerminalPath) WithIArrayItemValue(value interface{}) OrganizationInvitation_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for OrganizationInvitation: %d", fp.selector))
	}
}

func (fp *OrganizationInvitation_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

type OrganizationInvitation_FieldSubPath struct {
	selector OrganizationInvitation_FieldPathSelector
	subPath  gotenobject.FieldPath
}

var _ OrganizationInvitation_FieldPath = (*OrganizationInvitation_FieldSubPath)(nil)

func (fps *OrganizationInvitation_FieldSubPath) Selector() OrganizationInvitation_FieldPathSelector {
	return fps.selector
}
func (fps *OrganizationInvitation_FieldSubPath) AsInvitationSubPath() (iam_invitation.Invitation_FieldPath, bool) {
	res, ok := fps.subPath.(iam_invitation.Invitation_FieldPath)
	return res, ok
}
func (fps *OrganizationInvitation_FieldSubPath) AsMetadataSubPath() (meta.Meta_FieldPath, bool) {
	res, ok := fps.subPath.(meta.Meta_FieldPath)
	return res, ok
}

// String returns path representation in proto convention
func (fps *OrganizationInvitation_FieldSubPath) String() string {
	return fps.selector.String() + "." + fps.subPath.String()
}

// JSONString returns path representation is JSON convention
func (fps *OrganizationInvitation_FieldSubPath) JSONString() string {
	return strcase.ToLowerCamel(fps.selector.String()) + "." + fps.subPath.JSONString()
}

// Get returns all values pointed by selected field from source OrganizationInvitation
func (fps *OrganizationInvitation_FieldSubPath) Get(source *OrganizationInvitation) (values []interface{}) {
	switch fps.selector {
	case OrganizationInvitation_FieldPathSelectorInvitation:
		values = append(values, fps.subPath.GetRaw(source.GetInvitation())...)
	case OrganizationInvitation_FieldPathSelectorMetadata:
		values = append(values, fps.subPath.GetRaw(source.GetMetadata())...)
	default:
		panic(fmt.Sprintf("Invalid selector for OrganizationInvitation: %d", fps.selector))
	}
	return
}

func (fps *OrganizationInvitation_FieldSubPath) GetRaw(source proto.Message) []interface{} {
	return fps.Get(source.(*OrganizationInvitation))
}

// GetSingle returns value of selected field from source OrganizationInvitation
func (fps *OrganizationInvitation_FieldSubPath) GetSingle(source *OrganizationInvitation) (interface{}, bool) {
	switch fps.selector {
	case OrganizationInvitation_FieldPathSelectorInvitation:
		if source.GetInvitation() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetInvitation())
	case OrganizationInvitation_FieldPathSelectorMetadata:
		if source.GetMetadata() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for OrganizationInvitation: %d", fps.selector))
	}
}

func (fps *OrganizationInvitation_FieldSubPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fps.GetSingle(source.(*OrganizationInvitation))
}

// GetDefault returns a default value of the field type
func (fps *OrganizationInvitation_FieldSubPath) GetDefault() interface{} {
	return fps.subPath.GetDefault()
}

func (fps *OrganizationInvitation_FieldSubPath) ClearValue(item *OrganizationInvitation) {
	if item != nil {
		switch fps.selector {
		case OrganizationInvitation_FieldPathSelectorInvitation:
			fps.subPath.ClearValueRaw(item.Invitation)
		case OrganizationInvitation_FieldPathSelectorMetadata:
			fps.subPath.ClearValueRaw(item.Metadata)
		default:
			panic(fmt.Sprintf("Invalid selector for OrganizationInvitation: %d", fps.selector))
		}
	}
}

func (fps *OrganizationInvitation_FieldSubPath) ClearValueRaw(item proto.Message) {
	fps.ClearValue(item.(*OrganizationInvitation))
}

// IsLeaf - whether field path is holds simple value
func (fps *OrganizationInvitation_FieldSubPath) IsLeaf() bool {
	return fps.subPath.IsLeaf()
}

func (fps *OrganizationInvitation_FieldSubPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	iPaths := []gotenobject.FieldPath{&OrganizationInvitation_FieldTerminalPath{selector: fps.selector}}
	iPaths = append(iPaths, fps.subPath.SplitIntoTerminalIPaths()...)
	return iPaths
}

func (fps *OrganizationInvitation_FieldSubPath) WithIValue(value interface{}) OrganizationInvitation_FieldPathValue {
	return &OrganizationInvitation_FieldSubPathValue{fps, fps.subPath.WithRawIValue(value)}
}

func (fps *OrganizationInvitation_FieldSubPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fps.WithIValue(value)
}

func (fps *OrganizationInvitation_FieldSubPath) WithIArrayOfValues(values interface{}) OrganizationInvitation_FieldPathArrayOfValues {
	return &OrganizationInvitation_FieldSubPathArrayOfValues{fps, fps.subPath.WithRawIArrayOfValues(values)}
}

func (fps *OrganizationInvitation_FieldSubPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fps.WithIArrayOfValues(values)
}

func (fps *OrganizationInvitation_FieldSubPath) WithIArrayItemValue(value interface{}) OrganizationInvitation_FieldPathArrayItemValue {
	return &OrganizationInvitation_FieldSubPathArrayItemValue{fps, fps.subPath.WithRawIArrayItemValue(value)}
}

func (fps *OrganizationInvitation_FieldSubPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fps.WithIArrayItemValue(value)
}

// OrganizationInvitation_FieldPathValue allows storing values for OrganizationInvitation fields according to their type
type OrganizationInvitation_FieldPathValue interface {
	OrganizationInvitation_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **OrganizationInvitation)
	CompareWith(*OrganizationInvitation) (cmp int, comparable bool)
}

func ParseOrganizationInvitation_FieldPathValue(pathStr, valueStr string) (OrganizationInvitation_FieldPathValue, error) {
	fp, err := ParseOrganizationInvitation_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing OrganizationInvitation field path value from %s: %v", valueStr, err)
	}
	return fpv.(OrganizationInvitation_FieldPathValue), nil
}

func MustParseOrganizationInvitation_FieldPathValue(pathStr, valueStr string) OrganizationInvitation_FieldPathValue {
	fpv, err := ParseOrganizationInvitation_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type OrganizationInvitation_FieldTerminalPathValue struct {
	OrganizationInvitation_FieldTerminalPath
	value interface{}
}

var _ OrganizationInvitation_FieldPathValue = (*OrganizationInvitation_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'OrganizationInvitation' as interface{}
func (fpv *OrganizationInvitation_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *OrganizationInvitation_FieldTerminalPathValue) AsNameValue() (*Name, bool) {
	res, ok := fpv.value.(*Name)
	return res, ok
}
func (fpv *OrganizationInvitation_FieldTerminalPathValue) AsInvitationValue() (*iam_invitation.Invitation, bool) {
	res, ok := fpv.value.(*iam_invitation.Invitation)
	return res, ok
}
func (fpv *OrganizationInvitation_FieldTerminalPathValue) AsMetadataValue() (*meta.Meta, bool) {
	res, ok := fpv.value.(*meta.Meta)
	return res, ok
}

// SetTo stores value for selected field for object OrganizationInvitation
func (fpv *OrganizationInvitation_FieldTerminalPathValue) SetTo(target **OrganizationInvitation) {
	if *target == nil {
		*target = new(OrganizationInvitation)
	}
	switch fpv.selector {
	case OrganizationInvitation_FieldPathSelectorName:
		(*target).Name = fpv.value.(*Name)
	case OrganizationInvitation_FieldPathSelectorInvitation:
		(*target).Invitation = fpv.value.(*iam_invitation.Invitation)
	case OrganizationInvitation_FieldPathSelectorMetadata:
		(*target).Metadata = fpv.value.(*meta.Meta)
	default:
		panic(fmt.Sprintf("Invalid selector for OrganizationInvitation: %d", fpv.selector))
	}
}

func (fpv *OrganizationInvitation_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*OrganizationInvitation)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'OrganizationInvitation_FieldTerminalPathValue' with the value under path in 'OrganizationInvitation'.
func (fpv *OrganizationInvitation_FieldTerminalPathValue) CompareWith(source *OrganizationInvitation) (int, bool) {
	switch fpv.selector {
	case OrganizationInvitation_FieldPathSelectorName:
		leftValue := fpv.value.(*Name)
		rightValue := source.GetName()
		if leftValue == nil {
			if rightValue != nil {
				return -1, true
			}
			return 0, true
		}
		if rightValue == nil {
			return 1, true
		}
		if leftValue.String() == rightValue.String() {
			return 0, true
		} else if leftValue.String() < rightValue.String() {
			return -1, true
		} else {
			return 1, true
		}
	case OrganizationInvitation_FieldPathSelectorInvitation:
		return 0, false
	case OrganizationInvitation_FieldPathSelectorMetadata:
		return 0, false
	default:
		panic(fmt.Sprintf("Invalid selector for OrganizationInvitation: %d", fpv.selector))
	}
}

func (fpv *OrganizationInvitation_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*OrganizationInvitation))
}

type OrganizationInvitation_FieldSubPathValue struct {
	OrganizationInvitation_FieldPath
	subPathValue gotenobject.FieldPathValue
}

var _ OrganizationInvitation_FieldPathValue = (*OrganizationInvitation_FieldSubPathValue)(nil)

func (fpvs *OrganizationInvitation_FieldSubPathValue) AsInvitationPathValue() (iam_invitation.Invitation_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(iam_invitation.Invitation_FieldPathValue)
	return res, ok
}
func (fpvs *OrganizationInvitation_FieldSubPathValue) AsMetadataPathValue() (meta.Meta_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(meta.Meta_FieldPathValue)
	return res, ok
}

func (fpvs *OrganizationInvitation_FieldSubPathValue) SetTo(target **OrganizationInvitation) {
	if *target == nil {
		*target = new(OrganizationInvitation)
	}
	switch fpvs.Selector() {
	case OrganizationInvitation_FieldPathSelectorInvitation:
		fpvs.subPathValue.(iam_invitation.Invitation_FieldPathValue).SetTo(&(*target).Invitation)
	case OrganizationInvitation_FieldPathSelectorMetadata:
		fpvs.subPathValue.(meta.Meta_FieldPathValue).SetTo(&(*target).Metadata)
	default:
		panic(fmt.Sprintf("Invalid selector for OrganizationInvitation: %d", fpvs.Selector()))
	}
}

func (fpvs *OrganizationInvitation_FieldSubPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*OrganizationInvitation)
	fpvs.SetTo(&typedObject)
}

func (fpvs *OrganizationInvitation_FieldSubPathValue) GetRawValue() interface{} {
	return fpvs.subPathValue.GetRawValue()
}

func (fpvs *OrganizationInvitation_FieldSubPathValue) CompareWith(source *OrganizationInvitation) (int, bool) {
	switch fpvs.Selector() {
	case OrganizationInvitation_FieldPathSelectorInvitation:
		return fpvs.subPathValue.(iam_invitation.Invitation_FieldPathValue).CompareWith(source.GetInvitation())
	case OrganizationInvitation_FieldPathSelectorMetadata:
		return fpvs.subPathValue.(meta.Meta_FieldPathValue).CompareWith(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for OrganizationInvitation: %d", fpvs.Selector()))
	}
}

func (fpvs *OrganizationInvitation_FieldSubPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpvs.CompareWith(source.(*OrganizationInvitation))
}

// OrganizationInvitation_FieldPathArrayItemValue allows storing single item in Path-specific values for OrganizationInvitation according to their type
// Present only for array (repeated) types.
type OrganizationInvitation_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	OrganizationInvitation_FieldPath
	ContainsValue(*OrganizationInvitation) bool
}

// ParseOrganizationInvitation_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseOrganizationInvitation_FieldPathArrayItemValue(pathStr, valueStr string) (OrganizationInvitation_FieldPathArrayItemValue, error) {
	fp, err := ParseOrganizationInvitation_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing OrganizationInvitation field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(OrganizationInvitation_FieldPathArrayItemValue), nil
}

func MustParseOrganizationInvitation_FieldPathArrayItemValue(pathStr, valueStr string) OrganizationInvitation_FieldPathArrayItemValue {
	fpaiv, err := ParseOrganizationInvitation_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type OrganizationInvitation_FieldTerminalPathArrayItemValue struct {
	OrganizationInvitation_FieldTerminalPath
	value interface{}
}

var _ OrganizationInvitation_FieldPathArrayItemValue = (*OrganizationInvitation_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object OrganizationInvitation as interface{}
func (fpaiv *OrganizationInvitation_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *OrganizationInvitation_FieldTerminalPathArrayItemValue) GetSingle(source *OrganizationInvitation) (interface{}, bool) {
	return nil, false
}

func (fpaiv *OrganizationInvitation_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*OrganizationInvitation))
}

// Contains returns a boolean indicating if value that is being held is present in given 'OrganizationInvitation'
func (fpaiv *OrganizationInvitation_FieldTerminalPathArrayItemValue) ContainsValue(source *OrganizationInvitation) bool {
	slice := fpaiv.OrganizationInvitation_FieldTerminalPath.Get(source)
	for _, v := range slice {
		if asProtoMsg, ok := fpaiv.value.(proto.Message); ok {
			if proto.Equal(asProtoMsg, v.(proto.Message)) {
				return true
			}
		} else if reflect.DeepEqual(v, fpaiv.value) {
			return true
		}
	}
	return false
}

type OrganizationInvitation_FieldSubPathArrayItemValue struct {
	OrganizationInvitation_FieldPath
	subPathItemValue gotenobject.FieldPathArrayItemValue
}

// GetRawValue returns stored array item value
func (fpaivs *OrganizationInvitation_FieldSubPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaivs.subPathItemValue.GetRawItemValue()
}
func (fpaivs *OrganizationInvitation_FieldSubPathArrayItemValue) AsInvitationPathItemValue() (iam_invitation.Invitation_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(iam_invitation.Invitation_FieldPathArrayItemValue)
	return res, ok
}
func (fpaivs *OrganizationInvitation_FieldSubPathArrayItemValue) AsMetadataPathItemValue() (meta.Meta_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue)
	return res, ok
}

// Contains returns a boolean indicating if value that is being held is present in given 'OrganizationInvitation'
func (fpaivs *OrganizationInvitation_FieldSubPathArrayItemValue) ContainsValue(source *OrganizationInvitation) bool {
	switch fpaivs.Selector() {
	case OrganizationInvitation_FieldPathSelectorInvitation:
		return fpaivs.subPathItemValue.(iam_invitation.Invitation_FieldPathArrayItemValue).ContainsValue(source.GetInvitation())
	case OrganizationInvitation_FieldPathSelectorMetadata:
		return fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue).ContainsValue(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for OrganizationInvitation: %d", fpaivs.Selector()))
	}
}

// OrganizationInvitation_FieldPathArrayOfValues allows storing slice of values for OrganizationInvitation fields according to their type
type OrganizationInvitation_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	OrganizationInvitation_FieldPath
}

func ParseOrganizationInvitation_FieldPathArrayOfValues(pathStr, valuesStr string) (OrganizationInvitation_FieldPathArrayOfValues, error) {
	fp, err := ParseOrganizationInvitation_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing OrganizationInvitation field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(OrganizationInvitation_FieldPathArrayOfValues), nil
}

func MustParseOrganizationInvitation_FieldPathArrayOfValues(pathStr, valuesStr string) OrganizationInvitation_FieldPathArrayOfValues {
	fpaov, err := ParseOrganizationInvitation_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type OrganizationInvitation_FieldTerminalPathArrayOfValues struct {
	OrganizationInvitation_FieldTerminalPath
	values interface{}
}

var _ OrganizationInvitation_FieldPathArrayOfValues = (*OrganizationInvitation_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *OrganizationInvitation_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case OrganizationInvitation_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*Name) {
			values = append(values, v)
		}
	case OrganizationInvitation_FieldPathSelectorInvitation:
		for _, v := range fpaov.values.([]*iam_invitation.Invitation) {
			values = append(values, v)
		}
	case OrganizationInvitation_FieldPathSelectorMetadata:
		for _, v := range fpaov.values.([]*meta.Meta) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *OrganizationInvitation_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*Name, bool) {
	res, ok := fpaov.values.([]*Name)
	return res, ok
}
func (fpaov *OrganizationInvitation_FieldTerminalPathArrayOfValues) AsInvitationArrayOfValues() ([]*iam_invitation.Invitation, bool) {
	res, ok := fpaov.values.([]*iam_invitation.Invitation)
	return res, ok
}
func (fpaov *OrganizationInvitation_FieldTerminalPathArrayOfValues) AsMetadataArrayOfValues() ([]*meta.Meta, bool) {
	res, ok := fpaov.values.([]*meta.Meta)
	return res, ok
}

type OrganizationInvitation_FieldSubPathArrayOfValues struct {
	OrganizationInvitation_FieldPath
	subPathArrayOfValues gotenobject.FieldPathArrayOfValues
}

var _ OrganizationInvitation_FieldPathArrayOfValues = (*OrganizationInvitation_FieldSubPathArrayOfValues)(nil)

func (fpsaov *OrganizationInvitation_FieldSubPathArrayOfValues) GetRawValues() []interface{} {
	return fpsaov.subPathArrayOfValues.GetRawValues()
}
func (fpsaov *OrganizationInvitation_FieldSubPathArrayOfValues) AsInvitationPathArrayOfValues() (iam_invitation.Invitation_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(iam_invitation.Invitation_FieldPathArrayOfValues)
	return res, ok
}
func (fpsaov *OrganizationInvitation_FieldSubPathArrayOfValues) AsMetadataPathArrayOfValues() (meta.Meta_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(meta.Meta_FieldPathArrayOfValues)
	return res, ok
}
