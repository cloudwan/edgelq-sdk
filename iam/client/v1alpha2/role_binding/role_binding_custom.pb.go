// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha2/role_binding_custom.proto
// DO NOT EDIT!!!

package role_binding_client

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
	multi_region_policy "github.com/cloudwan/edgelq-sdk/common/types/multi_region_policy"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	role_binding "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/role_binding"
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
	_ = &multi_region_policy.MultiRegionPolicy{}
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &role_binding.RoleBinding{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for method
// [SetupCreateOwnerRole][ntt.iam.v1alpha2.SetupCreateOwnerRole] This is almost
// same as CreateRoleBinding, but used when central region creates initial
// RoleBinding for newly created project or organization. This skips checking if
// Project/Organization exist in store because they do not yet! Instead it
// checks if create owner role binding already exists.
type SetupCreateOwnerRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Parent reference of ntt.iam.v1alpha2.RoleBinding
	Parent *role_binding.ParentReference `protobuf:"bytes,1,opt,customtype=ParentReference,name=parent,proto3" json:"parent,omitempty" firestore:"parent"`
	// RoleBinding resource body
	RoleBinding *role_binding.RoleBinding `protobuf:"bytes,2,opt,name=role_binding,json=roleBinding,proto3" json:"role_binding,omitempty" firestore:"roleBinding"`
	// Policy created with this project
	MultiRegionPolicy *multi_region_policy.MultiRegionPolicy `protobuf:"bytes,3,opt,name=multi_region_policy,json=multiRegionPolicy,proto3" json:"multi_region_policy,omitempty" firestore:"multiRegionPolicy"`
}

func (m *SetupCreateOwnerRoleRequest) Reset() {
	*m = SetupCreateOwnerRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_role_binding_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *SetupCreateOwnerRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*SetupCreateOwnerRoleRequest) ProtoMessage() {}

func (m *SetupCreateOwnerRoleRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_role_binding_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*SetupCreateOwnerRoleRequest) GotenMessage() {}

// Deprecated, Use SetupCreateOwnerRoleRequest.ProtoReflect.Descriptor instead.
func (*SetupCreateOwnerRoleRequest) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_role_binding_custom_proto_rawDescGZIP(), []int{0}
}

func (m *SetupCreateOwnerRoleRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *SetupCreateOwnerRoleRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *SetupCreateOwnerRoleRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *SetupCreateOwnerRoleRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *SetupCreateOwnerRoleRequest) GetParent() *role_binding.ParentReference {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *SetupCreateOwnerRoleRequest) GetRoleBinding() *role_binding.RoleBinding {
	if m != nil {
		return m.RoleBinding
	}
	return nil
}

func (m *SetupCreateOwnerRoleRequest) GetMultiRegionPolicy() *multi_region_policy.MultiRegionPolicy {
	if m != nil {
		return m.MultiRegionPolicy
	}
	return nil
}

func (m *SetupCreateOwnerRoleRequest) SetParent(fv *role_binding.ParentReference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Parent", "SetupCreateOwnerRoleRequest"))
	}
	m.Parent = fv
}

func (m *SetupCreateOwnerRoleRequest) SetRoleBinding(fv *role_binding.RoleBinding) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RoleBinding", "SetupCreateOwnerRoleRequest"))
	}
	m.RoleBinding = fv
}

func (m *SetupCreateOwnerRoleRequest) SetMultiRegionPolicy(fv *multi_region_policy.MultiRegionPolicy) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MultiRegionPolicy", "SetupCreateOwnerRoleRequest"))
	}
	m.MultiRegionPolicy = fv
}

var edgelq_iam_proto_v1alpha2_role_binding_custom_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha2_role_binding_custom_proto_rawDesc = []byte{
	0x0a, 0x33, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x34, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x02, 0x0a, 0x1b, 0x53, 0x65, 0x74, 0x75, 0x70, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xb2, 0xda, 0x21, 0x0f, 0x42, 0x0d, 0x0a, 0x0b, 0x52,
	0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0xba, 0x9d, 0x22, 0x00, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x0c, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x0c, 0xca, 0xc6, 0x27,
	0x04, 0x62, 0x02, 0x08, 0x01, 0xc8, 0xd5, 0x22, 0x01, 0x52, 0x0b, 0x72, 0x6f, 0x6c, 0x65, 0x42,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x4c, 0x0a, 0x13, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x11, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x3a, 0x1f, 0xc2, 0x85, 0x2c, 0x1b, 0x32, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x3a, 0x11, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0xd3, 0x01, 0xe8, 0xde, 0x21, 0x01, 0x0a, 0x17, 0x63, 0x6f,
	0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x16, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a,
	0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3b, 0x72, 0x6f, 0x6c,
	0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0xd2, 0x84, 0xd1, 0x02, 0x46, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x73, 0x12, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha2_role_binding_custom_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha2_role_binding_custom_proto_rawDescData = edgelq_iam_proto_v1alpha2_role_binding_custom_proto_rawDesc
)

func edgelq_iam_proto_v1alpha2_role_binding_custom_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha2_role_binding_custom_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha2_role_binding_custom_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha2_role_binding_custom_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha2_role_binding_custom_proto_rawDescData
}

var edgelq_iam_proto_v1alpha2_role_binding_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_iam_proto_v1alpha2_role_binding_custom_proto_goTypes = []interface{}{
	(*SetupCreateOwnerRoleRequest)(nil),           // 0: ntt.iam.v1alpha2.SetupCreateOwnerRoleRequest
	(*role_binding.RoleBinding)(nil),              // 1: ntt.iam.v1alpha2.RoleBinding
	(*multi_region_policy.MultiRegionPolicy)(nil), // 2: ntt.types.MultiRegionPolicy
}
var edgelq_iam_proto_v1alpha2_role_binding_custom_proto_depIdxs = []int32{
	1, // 0: ntt.iam.v1alpha2.SetupCreateOwnerRoleRequest.role_binding:type_name -> ntt.iam.v1alpha2.RoleBinding
	2, // 1: ntt.iam.v1alpha2.SetupCreateOwnerRoleRequest.multi_region_policy:type_name -> ntt.types.MultiRegionPolicy
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha2_role_binding_custom_proto_init() }
func edgelq_iam_proto_v1alpha2_role_binding_custom_proto_init() {
	if edgelq_iam_proto_v1alpha2_role_binding_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha2_role_binding_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetupCreateOwnerRoleRequest); i {
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
			RawDescriptor: edgelq_iam_proto_v1alpha2_role_binding_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha2_role_binding_custom_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha2_role_binding_custom_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1alpha2_role_binding_custom_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha2_role_binding_custom_proto = out.File
	edgelq_iam_proto_v1alpha2_role_binding_custom_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha2_role_binding_custom_proto_goTypes = nil
	edgelq_iam_proto_v1alpha2_role_binding_custom_proto_depIdxs = nil
}
