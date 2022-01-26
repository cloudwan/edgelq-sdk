// Code generated by protoc-gen-goten-go
// File: edgelq/common/api/launch_stage.proto
// DO NOT EDIT!!!

package api

import (
	"fmt"
	"reflect"
	"sync"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import ()

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = reflect.Method{}
	_ = sync.Once{}

	_ = protojson.MarshalOptions{}
	_ = proto.MarshalOptions{}
	_ = preflect.Value{}
	_ = protoimpl.DescBuilder{}
)

// make sure we're using proto imports
var ()

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The launch stage as defined by [Google Cloud Platform
// Launch Stages](http://cloud.google.com/terms/launch-stages).
type LaunchStage int32

const (
	// Do not use this default value.
	LaunchStage_LAUNCH_STAGE_UNSPECIFIED LaunchStage = 0
	// Early Access features are limited to a closed group of testers. To use
	// these features, you must sign up in advance and sign a Trusted Tester
	// agreement (which includes confidentiality provisions). These features may
	// be unstable, changed in backward-incompatible ways, and are not
	// guaranteed to be released.
	LaunchStage_EARLY_ACCESS LaunchStage = 1
	// Alpha is a limited availability test for releases before they are cleared
	// for widespread use. By Alpha, all significant design issues are resolved
	// and we are in the process of verifying functionality. Alpha customers
	// need to apply for access, agree to applicable terms, and have their
	// projects whitelisted. Alpha releases don’t have to be feature complete,
	// no SLAs are provided, and there are no technical support obligations, but
	// they will be far enough along that customers can actually use them in
	// test environments or for limited-use tests -- just like they would in
	// normal production cases.
	LaunchStage_ALPHA LaunchStage = 2
	// Beta is the point at which we are ready to open a release for any
	// customer to use. There are no SLA or technical support obligations in a
	// Beta release. Products will be complete from a feature perspective, but
	// may have some open outstanding issues. Beta releases are suitable for
	// limited production use cases.
	LaunchStage_BETA LaunchStage = 3
	// GA features are open to all developers and are considered stable and
	// fully qualified for production use.
	LaunchStage_GA LaunchStage = 4
	// Deprecated features are scheduled to be shut down and removed. For more
	// information, see the “Deprecation Policy” section of our [Terms of
	// Service](https://cloud.google.com/terms/)
	// and the [Google Cloud Platform Subject to the Deprecation
	// Policy](https://cloud.google.com/terms/deprecation) documentation.
	LaunchStage_DEPRECATED LaunchStage = 5
)

var (
	LaunchStage_name = map[int32]string{
		0: "LAUNCH_STAGE_UNSPECIFIED",
		1: "EARLY_ACCESS",
		2: "ALPHA",
		3: "BETA",
		4: "GA",
		5: "DEPRECATED",
	}

	LaunchStage_value = map[string]int32{
		"LAUNCH_STAGE_UNSPECIFIED": 0,
		"EARLY_ACCESS":             1,
		"ALPHA":                    2,
		"BETA":                     3,
		"GA":                       4,
		"DEPRECATED":               5,
	}
)

func (x LaunchStage) Enum() *LaunchStage {
	p := new(LaunchStage)
	*p = x
	return p
}

func (x LaunchStage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), preflect.EnumNumber(x))
}

func (LaunchStage) Descriptor() preflect.EnumDescriptor {
	return edgelq_common_api_launch_stage_proto_enumTypes[0].Descriptor()
}

func (LaunchStage) Type() preflect.EnumType {
	return &edgelq_common_api_launch_stage_proto_enumTypes[0]
}

func (x LaunchStage) Number() preflect.EnumNumber {
	return preflect.EnumNumber(x)
}

// Deprecated, Use LaunchStage.ProtoReflect.Descriptor instead.
func (LaunchStage) EnumDescriptor() ([]byte, []int) {
	return edgelq_common_api_launch_stage_proto_rawDescGZIP(), []int{0}
}

var edgelq_common_api_launch_stage_proto preflect.FileDescriptor

var edgelq_common_api_launch_stage_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2a,
	0x6a, 0x0a, 0x0b, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1c,
	0x0a, 0x18, 0x4c, 0x41, 0x55, 0x4e, 0x43, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x45, 0x41, 0x52, 0x4c, 0x59, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x41, 0x4c, 0x50, 0x48, 0x41, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x45, 0x54,
	0x41, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x41, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x44,
	0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x05, 0x42, 0x44, 0x0a, 0x0e, 0x63,
	0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x50, 0x01, 0x5a,
	0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0xa2, 0x02, 0x04, 0x47, 0x41, 0x50,
	0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_common_api_launch_stage_proto_rawDescOnce sync.Once
	edgelq_common_api_launch_stage_proto_rawDescData = edgelq_common_api_launch_stage_proto_rawDesc
)

func edgelq_common_api_launch_stage_proto_rawDescGZIP() []byte {
	edgelq_common_api_launch_stage_proto_rawDescOnce.Do(func() {
		edgelq_common_api_launch_stage_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_common_api_launch_stage_proto_rawDescData)
	})
	return edgelq_common_api_launch_stage_proto_rawDescData
}

var edgelq_common_api_launch_stage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var edgelq_common_api_launch_stage_proto_goTypes = []interface{}{
	(LaunchStage)(0), // 0: ntt.api.LaunchStage
}
var edgelq_common_api_launch_stage_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { edgelq_common_api_launch_stage_proto_init() }
func edgelq_common_api_launch_stage_proto_init() {
	if edgelq_common_api_launch_stage_proto != nil {
		return
	}

	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_common_api_launch_stage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_common_api_launch_stage_proto_goTypes,
		DependencyIndexes: edgelq_common_api_launch_stage_proto_depIdxs,
		EnumInfos:         edgelq_common_api_launch_stage_proto_enumTypes,
	}.Build()
	edgelq_common_api_launch_stage_proto = out.File
	edgelq_common_api_launch_stage_proto_rawDesc = nil
	edgelq_common_api_launch_stage_proto_goTypes = nil
	edgelq_common_api_launch_stage_proto_depIdxs = nil
}
