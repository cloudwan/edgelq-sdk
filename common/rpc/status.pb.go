// Code generated by protoc-gen-goten-go
// File: edgelq/common/rpc/status.proto
// DO NOT EDIT!!!

package rpc

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
import (
	any "github.com/golang/protobuf/ptypes/any"
)

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
var (
	_ = &any.Any{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The `Status` type defines a logical error model that is suitable for
// different programming environments, including REST APIs and RPC APIs. It is
// used by [gRPC](https://github.com/grpc). The error model is designed to be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message, and error details. The error code should be an enum value of
// [google.rpc.Code][google.rpc.Code], but it may accept additional error codes
// if needed.  The error message should be a developer-facing English message
// that helps developers *understand* and *resolve* the error. If a localized
// user-facing error message is needed, put the localized message in the error
// details or localize it in the client. The optional error details may contain
// arbitrary information about the error. There is a predefined set of error
// detail types in the package `google.rpc` that can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error model, but it
// is not necessarily the actual wire format. When the `Status` message is
// exposed in different client libraries and different wire protocols, it can be
// mapped differently. For example, it will likely be mapped to some exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety of
// environments, either with or without APIs, to provide a
// consistent developer experience across different environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the client,
//     it may embed the `Status` in the normal response to indicate the partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step may
//     have a `Status` message for error reporting.
//
// - Batch operations. If a client uses batch request and batch response, the
//     `Status` message should be used directly inside batch response, one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous operation
//     results in its response, the status of those operations should be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message `Status` could
//     be used directly after any stripping needed for security/privacy reasons.
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// The status code, which should be an enum value of
	// [google.rpc.Code][google.rpc.Code].
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty" firestore:"code"`
	// A developer-facing error message, which should be in English. Any
	// user-facing error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	// by the client.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty" firestore:"message"`
	// A list of messages that carry the error details.  There is a common set of
	// message types for APIs to use.
	Details []*any.Any `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty" firestore:"details"`
}

func (m *Status) Reset() {
	*m = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_common_rpc_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Status) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Status) ProtoMessage() {}

func (m *Status) ProtoReflect() preflect.Message {
	mi := &edgelq_common_rpc_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Status) GotenMessage() {}

// Deprecated, Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return edgelq_common_rpc_status_proto_rawDescGZIP(), []int{0}
}

func (m *Status) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Status) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Status) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Status) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return int32(0)
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Status) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Status) SetCode(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Code", "Status"))
	}
	m.Code = fv
}

func (m *Status) SetMessage(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Message", "Status"))
	}
	m.Message = fv
}

func (m *Status) SetDetails(fv []*any.Any) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Details", "Status"))
	}
	m.Details = fv
}

var edgelq_common_rpc_status_proto preflect.FileDescriptor

var edgelq_common_rpc_status_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x6e, 0x74, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x47, 0xe8, 0xde, 0x21,
	0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x62, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x3b, 0x72, 0x70, 0x63, 0xa2, 0x02,
	0x03, 0x52, 0x50, 0x43, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_common_rpc_status_proto_rawDescOnce sync.Once
	edgelq_common_rpc_status_proto_rawDescData = edgelq_common_rpc_status_proto_rawDesc
)

func edgelq_common_rpc_status_proto_rawDescGZIP() []byte {
	edgelq_common_rpc_status_proto_rawDescOnce.Do(func() {
		edgelq_common_rpc_status_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_common_rpc_status_proto_rawDescData)
	})
	return edgelq_common_rpc_status_proto_rawDescData
}

var edgelq_common_rpc_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_common_rpc_status_proto_goTypes = []interface{}{
	(*Status)(nil),  // 0: ntt.rpc.Status
	(*any.Any)(nil), // 1: google.protobuf.Any
}
var edgelq_common_rpc_status_proto_depIdxs = []int32{
	1, // 0: ntt.rpc.Status.details:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { edgelq_common_rpc_status_proto_init() }
func edgelq_common_rpc_status_proto_init() {
	if edgelq_common_rpc_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_common_rpc_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}

	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_common_rpc_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_common_rpc_status_proto_goTypes,
		DependencyIndexes: edgelq_common_rpc_status_proto_depIdxs,
		MessageInfos:      edgelq_common_rpc_status_proto_msgTypes,
	}.Build()
	edgelq_common_rpc_status_proto = out.File
	edgelq_common_rpc_status_proto_rawDesc = nil
	edgelq_common_rpc_status_proto_goTypes = nil
	edgelq_common_rpc_status_proto_depIdxs = nil
}
