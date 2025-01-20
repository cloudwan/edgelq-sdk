// Code generated by protoc-gen-goten-object
// File: edgelq/logging/proto/v1/log.proto
// DO NOT EDIT!!!

package log

// proto imports
import (
	iam_iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1/common"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	bucket "github.com/cloudwan/edgelq-sdk/logging/resources/v1/bucket"
	common "github.com/cloudwan/edgelq-sdk/logging/resources/v1/common"
	log_descriptor "github.com/cloudwan/edgelq-sdk/logging/resources/v1/log_descriptor"
	meta_common "github.com/cloudwan/goten-sdk/meta-service/resources/v1/common"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	multi_region_policy "github.com/cloudwan/goten-sdk/types/multi_region_policy"
	anypb "google.golang.org/protobuf/types/known/anypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// make sure we're using proto imports
var (
	_ = &iam_iam_common.PCR{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &bucket.Bucket{}
	_ = &common.LabelDescriptor{}
	_ = &log_descriptor.LogDescriptor{}
	_ = &anypb.Any{}
	_ = &structpb.Struct{}
	_ = &timestamppb.Timestamp{}
	_ = &meta_common.LabelledDomain{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
)

type LogFieldPathBuilder struct{}

func NewLogFieldPathBuilder() LogFieldPathBuilder {
	return LogFieldPathBuilder{}
}
func (LogFieldPathBuilder) Name() LogPathSelectorName {
	return LogPathSelectorName{}
}
func (LogFieldPathBuilder) Scope() LogPathSelectorScope {
	return LogPathSelectorScope{}
}
func (LogFieldPathBuilder) Service() LogPathSelectorService {
	return LogPathSelectorService{}
}
func (LogFieldPathBuilder) Region() LogPathSelectorRegion {
	return LogPathSelectorRegion{}
}
func (LogFieldPathBuilder) Version() LogPathSelectorVersion {
	return LogPathSelectorVersion{}
}
func (LogFieldPathBuilder) LogDescriptor() LogPathSelectorLogDescriptor {
	return LogPathSelectorLogDescriptor{}
}
func (LogFieldPathBuilder) Labels() LogPathSelectorLabels {
	return LogPathSelectorLabels{}
}
func (LogFieldPathBuilder) Time() LogPathSelectorTime {
	return LogPathSelectorTime{}
}
func (LogFieldPathBuilder) JsonPayload() LogPathSelectorJsonPayload {
	return LogPathSelectorJsonPayload{}
}
func (LogFieldPathBuilder) PbPayload() LogPathSelectorPbPayload {
	return LogPathSelectorPbPayload{}
}
func (LogFieldPathBuilder) StringPayload() LogPathSelectorStringPayload {
	return LogPathSelectorStringPayload{}
}
func (LogFieldPathBuilder) BytesPayload() LogPathSelectorBytesPayload {
	return LogPathSelectorBytesPayload{}
}
func (LogFieldPathBuilder) BinKey() LogPathSelectorBinKey {
	return LogPathSelectorBinKey{}
}

type LogPathSelectorName struct{}

func (LogPathSelectorName) FieldPath() *Log_FieldTerminalPath {
	return &Log_FieldTerminalPath{selector: Log_FieldPathSelectorName}
}

func (s LogPathSelectorName) WithValue(value *Name) *Log_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Log_FieldTerminalPathValue)
}

func (s LogPathSelectorName) WithArrayOfValues(values []*Name) *Log_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Log_FieldTerminalPathArrayOfValues)
}

type LogPathSelectorScope struct{}

func (LogPathSelectorScope) FieldPath() *Log_FieldTerminalPath {
	return &Log_FieldTerminalPath{selector: Log_FieldPathSelectorScope}
}

func (s LogPathSelectorScope) WithValue(value string) *Log_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Log_FieldTerminalPathValue)
}

func (s LogPathSelectorScope) WithArrayOfValues(values []string) *Log_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Log_FieldTerminalPathArrayOfValues)
}

type LogPathSelectorService struct{}

func (LogPathSelectorService) FieldPath() *Log_FieldTerminalPath {
	return &Log_FieldTerminalPath{selector: Log_FieldPathSelectorService}
}

func (s LogPathSelectorService) WithValue(value string) *Log_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Log_FieldTerminalPathValue)
}

func (s LogPathSelectorService) WithArrayOfValues(values []string) *Log_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Log_FieldTerminalPathArrayOfValues)
}

type LogPathSelectorRegion struct{}

func (LogPathSelectorRegion) FieldPath() *Log_FieldTerminalPath {
	return &Log_FieldTerminalPath{selector: Log_FieldPathSelectorRegion}
}

func (s LogPathSelectorRegion) WithValue(value string) *Log_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Log_FieldTerminalPathValue)
}

func (s LogPathSelectorRegion) WithArrayOfValues(values []string) *Log_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Log_FieldTerminalPathArrayOfValues)
}

type LogPathSelectorVersion struct{}

func (LogPathSelectorVersion) FieldPath() *Log_FieldTerminalPath {
	return &Log_FieldTerminalPath{selector: Log_FieldPathSelectorVersion}
}

func (s LogPathSelectorVersion) WithValue(value string) *Log_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Log_FieldTerminalPathValue)
}

func (s LogPathSelectorVersion) WithArrayOfValues(values []string) *Log_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Log_FieldTerminalPathArrayOfValues)
}

type LogPathSelectorLogDescriptor struct{}

func (LogPathSelectorLogDescriptor) FieldPath() *Log_FieldTerminalPath {
	return &Log_FieldTerminalPath{selector: Log_FieldPathSelectorLogDescriptor}
}

func (s LogPathSelectorLogDescriptor) WithValue(value *log_descriptor.Reference) *Log_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Log_FieldTerminalPathValue)
}

func (s LogPathSelectorLogDescriptor) WithArrayOfValues(values []*log_descriptor.Reference) *Log_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Log_FieldTerminalPathArrayOfValues)
}

type LogPathSelectorLabels struct{}

func (LogPathSelectorLabels) FieldPath() *Log_FieldTerminalPath {
	return &Log_FieldTerminalPath{selector: Log_FieldPathSelectorLabels}
}

func (s LogPathSelectorLabels) WithValue(value map[string]string) *Log_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Log_FieldTerminalPathValue)
}

func (s LogPathSelectorLabels) WithArrayOfValues(values []map[string]string) *Log_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Log_FieldTerminalPathArrayOfValues)
}

func (LogPathSelectorLabels) WithKey(key string) LogMapPathSelectorLabels {
	return LogMapPathSelectorLabels{key: key}
}

type LogMapPathSelectorLabels struct {
	key string
}

func (s LogMapPathSelectorLabels) FieldPath() *Log_FieldPathMap {
	return &Log_FieldPathMap{selector: Log_FieldPathSelectorLabels, key: s.key}
}

func (s LogMapPathSelectorLabels) WithValue(value string) *Log_FieldPathMapValue {
	return s.FieldPath().WithIValue(value).(*Log_FieldPathMapValue)
}

func (s LogMapPathSelectorLabels) WithArrayOfValues(values []string) *Log_FieldPathMapArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Log_FieldPathMapArrayOfValues)
}

type LogPathSelectorTime struct{}

func (LogPathSelectorTime) FieldPath() *Log_FieldTerminalPath {
	return &Log_FieldTerminalPath{selector: Log_FieldPathSelectorTime}
}

func (s LogPathSelectorTime) WithValue(value *timestamppb.Timestamp) *Log_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Log_FieldTerminalPathValue)
}

func (s LogPathSelectorTime) WithArrayOfValues(values []*timestamppb.Timestamp) *Log_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Log_FieldTerminalPathArrayOfValues)
}

type LogPathSelectorJsonPayload struct{}

func (LogPathSelectorJsonPayload) FieldPath() *Log_FieldTerminalPath {
	return &Log_FieldTerminalPath{selector: Log_FieldPathSelectorJsonPayload}
}

func (s LogPathSelectorJsonPayload) WithValue(value *structpb.Struct) *Log_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Log_FieldTerminalPathValue)
}

func (s LogPathSelectorJsonPayload) WithArrayOfValues(values []*structpb.Struct) *Log_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Log_FieldTerminalPathArrayOfValues)
}

type LogPathSelectorPbPayload struct{}

func (LogPathSelectorPbPayload) FieldPath() *Log_FieldTerminalPath {
	return &Log_FieldTerminalPath{selector: Log_FieldPathSelectorPbPayload}
}

func (s LogPathSelectorPbPayload) WithValue(value *anypb.Any) *Log_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Log_FieldTerminalPathValue)
}

func (s LogPathSelectorPbPayload) WithArrayOfValues(values []*anypb.Any) *Log_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Log_FieldTerminalPathArrayOfValues)
}

type LogPathSelectorStringPayload struct{}

func (LogPathSelectorStringPayload) FieldPath() *Log_FieldTerminalPath {
	return &Log_FieldTerminalPath{selector: Log_FieldPathSelectorStringPayload}
}

func (s LogPathSelectorStringPayload) WithValue(value string) *Log_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Log_FieldTerminalPathValue)
}

func (s LogPathSelectorStringPayload) WithArrayOfValues(values []string) *Log_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Log_FieldTerminalPathArrayOfValues)
}

type LogPathSelectorBytesPayload struct{}

func (LogPathSelectorBytesPayload) FieldPath() *Log_FieldTerminalPath {
	return &Log_FieldTerminalPath{selector: Log_FieldPathSelectorBytesPayload}
}

func (s LogPathSelectorBytesPayload) WithValue(value []byte) *Log_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Log_FieldTerminalPathValue)
}

func (s LogPathSelectorBytesPayload) WithArrayOfValues(values [][]byte) *Log_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Log_FieldTerminalPathArrayOfValues)
}

type LogPathSelectorBinKey struct{}

func (LogPathSelectorBinKey) FieldPath() *Log_FieldTerminalPath {
	return &Log_FieldTerminalPath{selector: Log_FieldPathSelectorBinKey}
}

func (s LogPathSelectorBinKey) WithValue(value string) *Log_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Log_FieldTerminalPathValue)
}

func (s LogPathSelectorBinKey) WithArrayOfValues(values []string) *Log_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Log_FieldTerminalPathArrayOfValues)
}
