// Code generated by protoc-gen-goten-validate
// File: edgelq/devices/proto/v1alpha2/device.proto
// DO NOT EDIT!!!

package device

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	gotenvalidate "github.com/cloudwan/goten-sdk/runtime/validate"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/project"
	iam_attestation_domain "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/attestation_domain"
	iam_iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/common"
	iam_service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/service_account"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

var (
	_ = bytes.Equal
	_ = errors.New
	_ = fmt.Errorf
	_ = net.ParseIP
	_ = regexp.Match
	_ = strings.Split
	_ = time.Now
	_ = utf8.RuneCountInString
	_ = url.Parse
	_ = durationpb.Duration{}
	_ = timestamppb.Timestamp{}
	_ = gotenvalidate.NewValidationError
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &project.Project{}
	_ = &iam_attestation_domain.AttestationDomain{}
	_ = &iam_iam_common.Actor{}
	_ = &iam_service_account.ServiceAccount{}
	_ = &duration.Duration{}
	_ = &field_mask.FieldMask{}
	_ = &timestamp.Timestamp{}
)

func (obj *Device) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Metadata).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Device", "metadata", obj.Metadata, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Spec).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Device", "spec", obj.Spec, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Status).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Device", "status", obj.Status, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.PublicListingSpec).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Device", "publicListingSpec", obj.PublicListingSpec, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.NetConfig).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Spec", "netConfig", obj.NetConfig, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.SshConfig).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Spec", "sshConfig", obj.SshConfig, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.AttestationConfig).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Spec", "attestationConfig", obj.AttestationConfig, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Addresses {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("Status", "addresses", obj.Addresses[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.Conditions {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("Status", "conditions", obj.Conditions[idx], "nested object validation failed", err)
			}
		}
	}
	if subobj, ok := interface{}(obj.DeviceInfo).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Status", "deviceInfo", obj.DeviceInfo, "nested object validation failed", err)
		}
	}
	for idx, elem := range obj.AttestationStatus {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("Status", "attestationStatus", obj.AttestationStatus[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_PublicListingSpec) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Ethernets {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("NetworkConfig", "ethernets", obj.Ethernets[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.Wifis {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("NetworkConfig", "wifis", obj.Wifis[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.Bridges {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("NetworkConfig", "bridges", obj.Bridges[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.Bonds {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("NetworkConfig", "bonds", obj.Bonds[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.Tunnels {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("NetworkConfig", "tunnels", obj.Tunnels[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.Vlans {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("NetworkConfig", "vlans", obj.Vlans[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_SSHConfig) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.SshAuthorized {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("SSHConfig", "sshAuthorized", obj.SshAuthorized[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_AttestationConfig) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig_CommonOpts) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Dhcp4Overrides).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CommonOpts", "dhcp4Overrides", obj.Dhcp4Overrides, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Dhcp6Overrides).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CommonOpts", "dhcp6Overrides", obj.Dhcp6Overrides, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Nameservers).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CommonOpts", "nameservers", obj.Nameservers, "nested object validation failed", err)
		}
	}
	for idx, elem := range obj.Routes {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("CommonOpts", "routes", obj.Routes[idx], "nested object validation failed", err)
			}
		}
	}
	if subobj, ok := interface{}(obj.RoutingPolicy).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CommonOpts", "routingPolicy", obj.RoutingPolicy, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Auth).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CommonOpts", "auth", obj.Auth, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig_EthOpts) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Match).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("EthOpts", "match", obj.Match, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Opts).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("EthOpts", "opts", obj.Opts, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig_WifiOpts) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Match).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("WifiOpts", "match", obj.Match, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Opts).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("WifiOpts", "opts", obj.Opts, "nested object validation failed", err)
		}
	}
	for idx, elem := range obj.AccessPoints {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("WifiOpts", "accessPoints", obj.AccessPoints[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig_BridgesOpts) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Opts).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("BridgesOpts", "opts", obj.Opts, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Parameters).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("BridgesOpts", "parameters", obj.Parameters, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig_BondsOpts) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Opts).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("BondsOpts", "opts", obj.Opts, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Parameters).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("BondsOpts", "parameters", obj.Parameters, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig_TunnelsOpts) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Opts).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("TunnelsOpts", "opts", obj.Opts, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig_VlansOpts) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Opts).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("VlansOpts", "opts", obj.Opts, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig_CommonOpts_DHCPOverrides) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig_CommonOpts_Nameservers) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig_CommonOpts_Routes) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig_CommonOpts_RoutingPolicy) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig_CommonOpts_Auth) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig_EthOpts_Match) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig_WifiOpts_Match) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig_WifiOpts_AccessPoint) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig_BridgesOpts_Parameters) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig_BondsOpts_Parameters) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_NetworkConfig_TunnelsOpts_Key) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Spec_SSHConfig_AuthKey) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_Address) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_Condition) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.HardwareInformation).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("DeviceInfo", "hardwareInformation", obj.HardwareInformation, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Os).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("HardwareInformation", "os", obj.Os, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Bios).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("HardwareInformation", "bios", obj.Bios, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.System).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("HardwareInformation", "system", obj.System, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Cpu).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("HardwareInformation", "cpu", obj.Cpu, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Block).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("HardwareInformation", "block", obj.Block, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Network).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("HardwareInformation", "network", obj.Network, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Gpu).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("HardwareInformation", "gpu", obj.Gpu, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.MemoryInfo).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("HardwareInformation", "memoryInfo", obj.MemoryInfo, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_Capability) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_OS) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_BIOS) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_System) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Configuration).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("System", "configuration", obj.Configuration, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_CPU) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Processors {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("CPU", "processors", obj.Processors[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_Block) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Disks {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("Block", "disks", obj.Disks[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_Network) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Nics {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("Network", "nics", obj.Nics[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_GPU) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.GraphicCards {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("GPU", "graphicCards", obj.GraphicCards[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_PCIDevice) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_MemoryInfo) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Memory {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("MemoryInfo", "memory", obj.Memory[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_System_Configuration) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_CPU_Processor) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Capabilities {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("Processor", "capabilities", obj.Capabilities[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.CacheInfo {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("Processor", "cacheInfo", obj.CacheInfo[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_CPU_Processor_Cache) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_Block_Disk) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Partitions {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("Disk", "partitions", obj.Partitions[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_Block_Disk_Partition) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_Network_NIC) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_GPU_GraphicCard) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Device).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("GraphicCard", "device", obj.Device, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_MemoryInfo_Memory) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.MemoryBanks {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("Memory", "memoryBanks", obj.MemoryBanks[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Device_Status_DeviceInfo_HardwareInformation_MemoryInfo_Memory_MemoryBank) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
