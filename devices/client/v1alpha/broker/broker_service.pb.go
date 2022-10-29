// Code generated by protoc-gen-goten-go
// File: edgelq/devices/proto/v1alpha/broker_service.proto
// DO NOT EDIT!!!

package broker_client

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

var edgelq_devices_proto_v1alpha_broker_service_proto preflect.FileDescriptor

var edgelq_devices_proto_v1alpha_broker_service_proto_rawDesc = []byte{
	0x0a, 0x31, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x29, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xdc, 0x09, 0x0a,
	0x0d, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8a,
	0x03, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x30, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x88, 0x02, 0x82,
	0xdb, 0x21, 0x3a, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x14, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2a, 0x1a, 0x0a, 0x18, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x8a, 0xdb, 0x21,
	0x22, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x4c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0xa2, 0xdc, 0x21, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4f, 0x22, 0x4d, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x46, 0x6f,
	0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x92, 0x97, 0x22, 0x3f,
	0x0a, 0x17, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xc8,
	0xd5, 0x22, 0x01, 0xd0, 0xd5, 0x22, 0x02, 0x28, 0x01, 0x30, 0x01, 0x12, 0xb1, 0x03, 0x0a, 0x1b,
	0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x37, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9a,
	0x02, 0x82, 0xdb, 0x21, 0x3f, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x6f,
	0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2a, 0x18, 0x0a, 0x16, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x8a, 0xdb, 0x21, 0x29, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0xa2, 0xdc, 0x21, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x54, 0x22, 0x52, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x3d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x2a, 0x7d, 0x3a, 0x6f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x92,
	0x97, 0x22, 0x44, 0x0a, 0x15, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0xc8, 0xd5, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12,
	0xd8, 0x02, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2b, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x54, 0x6f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2c, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x6f,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xe5,
	0x01, 0x82, 0xdb, 0x21, 0x30, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2a, 0x15,
	0x0a, 0x13, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x8a, 0xdb, 0x21, 0x1d, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x6f,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0xa2, 0xdc, 0x21, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x45,
	0x22, 0x43, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6f, 0x70, 0x65, 0x6e,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x3d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x92, 0x97, 0x22, 0x35, 0x0a, 0x12, 0x6f, 0x70, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0xc8, 0xd5,
	0x22, 0x01, 0xd0, 0xd5, 0x22, 0x02, 0x28, 0x01, 0x30, 0x01, 0x1a, 0x2f, 0xca, 0x41, 0x12, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f,
	0x6d, 0xd2, 0x41, 0x17, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0xbe, 0x03, 0xe8, 0xde,
	0x21, 0x01, 0x82, 0xff, 0xd0, 0x02, 0x49, 0x0a, 0x0d, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x8a, 0xff, 0xd0, 0x02, 0x49, 0x0a, 0x0d, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x92, 0x8c,
	0xd1, 0x02, 0x51, 0x0a, 0x11, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x62, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x42, 0x12, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x3b, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0xd2, 0x84,
	0xd1, 0x02, 0x49, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x73, 0x12, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_devices_proto_v1alpha_broker_service_proto_rawDescOnce sync.Once
	edgelq_devices_proto_v1alpha_broker_service_proto_rawDescData = edgelq_devices_proto_v1alpha_broker_service_proto_rawDesc
)

func edgelq_devices_proto_v1alpha_broker_service_proto_rawDescGZIP() []byte {
	edgelq_devices_proto_v1alpha_broker_service_proto_rawDescOnce.Do(func() {
		edgelq_devices_proto_v1alpha_broker_service_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_devices_proto_v1alpha_broker_service_proto_rawDescData)
	})
	return edgelq_devices_proto_v1alpha_broker_service_proto_rawDescData
}

var edgelq_devices_proto_v1alpha_broker_service_proto_goTypes = []interface{}{
	(*ListenForConnectionsRequest)(nil),         // 0: ntt.devices.v1alpha.ListenForConnectionsRequest
	(*OpenConnectionChannelSocketRequest)(nil),  // 1: ntt.devices.v1alpha.OpenConnectionChannelSocketRequest
	(*ConnectToDeviceRequest)(nil),              // 2: ntt.devices.v1alpha.ConnectToDeviceRequest
	(*ListenForConnectionsResponse)(nil),        // 3: ntt.devices.v1alpha.ListenForConnectionsResponse
	(*OpenConnectionChannelSocketResponse)(nil), // 4: ntt.devices.v1alpha.OpenConnectionChannelSocketResponse
	(*ConnectToDeviceResponse)(nil),             // 5: ntt.devices.v1alpha.ConnectToDeviceResponse
}
var edgelq_devices_proto_v1alpha_broker_service_proto_depIdxs = []int32{
	0, // 0: ntt.devices.v1alpha.BrokerService.ListenForConnections:input_type -> ntt.devices.v1alpha.ListenForConnectionsRequest
	1, // 1: ntt.devices.v1alpha.BrokerService.OpenConnectionChannelSocket:input_type -> ntt.devices.v1alpha.OpenConnectionChannelSocketRequest
	2, // 2: ntt.devices.v1alpha.BrokerService.ConnectToDevice:input_type -> ntt.devices.v1alpha.ConnectToDeviceRequest
	3, // 3: ntt.devices.v1alpha.BrokerService.ListenForConnections:output_type -> ntt.devices.v1alpha.ListenForConnectionsResponse
	4, // 4: ntt.devices.v1alpha.BrokerService.OpenConnectionChannelSocket:output_type -> ntt.devices.v1alpha.OpenConnectionChannelSocketResponse
	5, // 5: ntt.devices.v1alpha.BrokerService.ConnectToDevice:output_type -> ntt.devices.v1alpha.ConnectToDeviceResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { edgelq_devices_proto_v1alpha_broker_service_proto_init() }
func edgelq_devices_proto_v1alpha_broker_service_proto_init() {
	if edgelq_devices_proto_v1alpha_broker_service_proto != nil {
		return
	}

	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_devices_proto_v1alpha_broker_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           edgelq_devices_proto_v1alpha_broker_service_proto_goTypes,
		DependencyIndexes: edgelq_devices_proto_v1alpha_broker_service_proto_depIdxs,
	}.Build()
	edgelq_devices_proto_v1alpha_broker_service_proto = out.File
	edgelq_devices_proto_v1alpha_broker_service_proto_rawDesc = nil
	edgelq_devices_proto_v1alpha_broker_service_proto_goTypes = nil
	edgelq_devices_proto_v1alpha_broker_service_proto_depIdxs = nil
}
