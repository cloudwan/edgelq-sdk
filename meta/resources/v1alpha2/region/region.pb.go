// Code generated by protoc-gen-goten-go
// File: edgelq/meta/proto/v1alpha2/region.proto
// DO NOT EDIT!!!

package region

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
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
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
	_ = &ntt_meta.Meta{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Region Resource
type Region struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of Region
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Region title
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty" firestore:"title"`
	// Region domain, for example us-west.edgelq.com
	Domain string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty" firestore:"domain"`
	// Region location
	Location *Region_RegionLocation `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty" firestore:"location"`
	// Whether is default. Only one record is allowed to have this attribute on.
	IsDefault bool `protobuf:"varint,5,opt,name=is_default,json=isDefault,proto3" json:"is_default,omitempty" firestore:"isDefault"`
	// Hardcoded scores regarding connectivity preference.
	ConnectivityScores []*Region_RegionConnectivityPreference `protobuf:"bytes,6,rep,name=connectivity_scores,json=connectivityScores,proto3" json:"connectivity_scores,omitempty" firestore:"connectivityScores"`
	// Metadata
	Metadata *ntt_meta.Meta `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
}

func (m *Region) Reset() {
	*m = Region{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_meta_proto_v1alpha2_region_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Region) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Region) ProtoMessage() {}

func (m *Region) ProtoReflect() preflect.Message {
	mi := &edgelq_meta_proto_v1alpha2_region_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Region) GotenMessage() {}

// Deprecated, Use Region.ProtoReflect.Descriptor instead.
func (*Region) Descriptor() ([]byte, []int) {
	return edgelq_meta_proto_v1alpha2_region_proto_rawDescGZIP(), []int{0}
}

func (m *Region) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Region) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Region) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Region) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Region) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Region) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Region) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Region) GetLocation() *Region_RegionLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Region) GetIsDefault() bool {
	if m != nil {
		return m.IsDefault
	}
	return false
}

func (m *Region) GetConnectivityScores() []*Region_RegionConnectivityPreference {
	if m != nil {
		return m.ConnectivityScores
	}
	return nil
}

func (m *Region) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Region) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "Region"))
	}
	m.Name = fv
}

func (m *Region) SetTitle(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Title", "Region"))
	}
	m.Title = fv
}

func (m *Region) SetDomain(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Domain", "Region"))
	}
	m.Domain = fv
}

func (m *Region) SetLocation(fv *Region_RegionLocation) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Location", "Region"))
	}
	m.Location = fv
}

func (m *Region) SetIsDefault(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "IsDefault", "Region"))
	}
	m.IsDefault = fv
}

func (m *Region) SetConnectivityScores(fv []*Region_RegionConnectivityPreference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ConnectivityScores", "Region"))
	}
	m.ConnectivityScores = fv
}

func (m *Region) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "Region"))
	}
	m.Metadata = fv
}

// RegionLocation describes deployment location of a region.
type Region_RegionLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Continent     string `protobuf:"bytes,1,opt,name=continent,proto3" json:"continent,omitempty" firestore:"continent"`
	Country       string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty" firestore:"country"`
	Agglomeration string `protobuf:"bytes,3,opt,name=agglomeration,proto3" json:"agglomeration,omitempty" firestore:"agglomeration"`
	City          string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty" firestore:"city"`
	Cloud         string `protobuf:"bytes,5,opt,name=cloud,proto3" json:"cloud,omitempty" firestore:"cloud"`
}

func (m *Region_RegionLocation) Reset() {
	*m = Region_RegionLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_meta_proto_v1alpha2_region_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Region_RegionLocation) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Region_RegionLocation) ProtoMessage() {}

func (m *Region_RegionLocation) ProtoReflect() preflect.Message {
	mi := &edgelq_meta_proto_v1alpha2_region_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Region_RegionLocation) GotenMessage() {}

// Deprecated, Use Region_RegionLocation.ProtoReflect.Descriptor instead.
func (*Region_RegionLocation) Descriptor() ([]byte, []int) {
	return edgelq_meta_proto_v1alpha2_region_proto_rawDescGZIP(), []int{0, 0}
}

func (m *Region_RegionLocation) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Region_RegionLocation) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Region_RegionLocation) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Region_RegionLocation) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Region_RegionLocation) GetContinent() string {
	if m != nil {
		return m.Continent
	}
	return ""
}

func (m *Region_RegionLocation) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Region_RegionLocation) GetAgglomeration() string {
	if m != nil {
		return m.Agglomeration
	}
	return ""
}

func (m *Region_RegionLocation) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Region_RegionLocation) GetCloud() string {
	if m != nil {
		return m.Cloud
	}
	return ""
}

func (m *Region_RegionLocation) SetContinent(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Continent", "Region_RegionLocation"))
	}
	m.Continent = fv
}

func (m *Region_RegionLocation) SetCountry(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Country", "Region_RegionLocation"))
	}
	m.Country = fv
}

func (m *Region_RegionLocation) SetAgglomeration(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Agglomeration", "Region_RegionLocation"))
	}
	m.Agglomeration = fv
}

func (m *Region_RegionLocation) SetCity(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "City", "Region_RegionLocation"))
	}
	m.City = fv
}

func (m *Region_RegionLocation) SetCloud(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Cloud", "Region_RegionLocation"))
	}
	m.Cloud = fv
}

// RegionConnectivityPreference is a hardcoded connectivity score to other
// region.
type Region_RegionConnectivityPreference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Destination region
	Dest *Reference `protobuf:"bytes,2,opt,customtype=Reference,name=dest,proto3" json:"dest,omitempty" firestore:"dest"`
	// Hardcoded score
	Score int32 `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty" firestore:"score"`
}

func (m *Region_RegionConnectivityPreference) Reset() {
	*m = Region_RegionConnectivityPreference{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_meta_proto_v1alpha2_region_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Region_RegionConnectivityPreference) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Region_RegionConnectivityPreference) ProtoMessage() {}

func (m *Region_RegionConnectivityPreference) ProtoReflect() preflect.Message {
	mi := &edgelq_meta_proto_v1alpha2_region_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Region_RegionConnectivityPreference) GotenMessage() {}

// Deprecated, Use Region_RegionConnectivityPreference.ProtoReflect.Descriptor instead.
func (*Region_RegionConnectivityPreference) Descriptor() ([]byte, []int) {
	return edgelq_meta_proto_v1alpha2_region_proto_rawDescGZIP(), []int{0, 1}
}

func (m *Region_RegionConnectivityPreference) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Region_RegionConnectivityPreference) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Region_RegionConnectivityPreference) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Region_RegionConnectivityPreference) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Region_RegionConnectivityPreference) GetDest() *Reference {
	if m != nil {
		return m.Dest
	}
	return nil
}

func (m *Region_RegionConnectivityPreference) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return int32(0)
}

func (m *Region_RegionConnectivityPreference) SetDest(fv *Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Dest", "Region_RegionConnectivityPreference"))
	}
	m.Dest = fv
}

func (m *Region_RegionConnectivityPreference) SetScore(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Score", "Region_RegionConnectivityPreference"))
	}
	m.Score = fv
}

var edgelq_meta_proto_v1alpha2_region_proto preflect.FileDescriptor

var edgelq_meta_proto_v1alpha2_region_proto_rawDesc = []byte{
	0x0a, 0x27, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x74, 0x74, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x34, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x05, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e,
	0xb2, 0xda, 0x21, 0x0a, 0x0a, 0x08, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x44, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x67, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x12, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x98, 0x01,
	0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x67, 0x67, 0x6c,
	0x6f, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x61, 0x67, 0x67, 0x6c, 0x6f, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x1a, 0x58, 0x0a, 0x1c, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xb2, 0xda, 0x21, 0x0a, 0x12, 0x08, 0x0a, 0x06,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x3a, 0xa3, 0x01, 0xea, 0x41, 0x2a, 0x0a, 0x16, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x10, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x7d, 0x92, 0xd9, 0x21, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xda, 0x94,
	0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xc2, 0x85, 0x2c, 0x4a, 0x22, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x22, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x0a, 0x69,
	0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22, 0x13, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x42, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xd1, 0x02, 0xe8, 0xde, 0x21, 0x01, 0xd2,
	0xff, 0xd0, 0x02, 0x45, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x42, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x3b, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0xd2, 0x84, 0xd1, 0x02, 0x47, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0xa2, 0x80, 0xd1, 0x02, 0x47, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_meta_proto_v1alpha2_region_proto_rawDescOnce sync.Once
	edgelq_meta_proto_v1alpha2_region_proto_rawDescData = edgelq_meta_proto_v1alpha2_region_proto_rawDesc
)

func edgelq_meta_proto_v1alpha2_region_proto_rawDescGZIP() []byte {
	edgelq_meta_proto_v1alpha2_region_proto_rawDescOnce.Do(func() {
		edgelq_meta_proto_v1alpha2_region_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_meta_proto_v1alpha2_region_proto_rawDescData)
	})
	return edgelq_meta_proto_v1alpha2_region_proto_rawDescData
}

var edgelq_meta_proto_v1alpha2_region_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var edgelq_meta_proto_v1alpha2_region_proto_goTypes = []interface{}{
	(*Region)(nil),                              // 0: ntt.meta.v1alpha2.Region
	(*Region_RegionLocation)(nil),               // 1: ntt.meta.v1alpha2.Region.RegionLocation
	(*Region_RegionConnectivityPreference)(nil), // 2: ntt.meta.v1alpha2.Region.RegionConnectivityPreference
	(*ntt_meta.Meta)(nil),                       // 3: ntt.types.Meta
}
var edgelq_meta_proto_v1alpha2_region_proto_depIdxs = []int32{
	1, // 0: ntt.meta.v1alpha2.Region.location:type_name -> ntt.meta.v1alpha2.Region.RegionLocation
	2, // 1: ntt.meta.v1alpha2.Region.connectivity_scores:type_name -> ntt.meta.v1alpha2.Region.RegionConnectivityPreference
	3, // 2: ntt.meta.v1alpha2.Region.metadata:type_name -> ntt.types.Meta
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { edgelq_meta_proto_v1alpha2_region_proto_init() }
func edgelq_meta_proto_v1alpha2_region_proto_init() {
	if edgelq_meta_proto_v1alpha2_region_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_meta_proto_v1alpha2_region_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Region); i {
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
		edgelq_meta_proto_v1alpha2_region_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Region_RegionLocation); i {
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
		edgelq_meta_proto_v1alpha2_region_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Region_RegionConnectivityPreference); i {
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
			RawDescriptor: edgelq_meta_proto_v1alpha2_region_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_meta_proto_v1alpha2_region_proto_goTypes,
		DependencyIndexes: edgelq_meta_proto_v1alpha2_region_proto_depIdxs,
		MessageInfos:      edgelq_meta_proto_v1alpha2_region_proto_msgTypes,
	}.Build()
	edgelq_meta_proto_v1alpha2_region_proto = out.File
	edgelq_meta_proto_v1alpha2_region_proto_rawDesc = nil
	edgelq_meta_proto_v1alpha2_region_proto_goTypes = nil
	edgelq_meta_proto_v1alpha2_region_proto_depIdxs = nil
}
