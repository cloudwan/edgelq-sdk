// Code generated by protoc-gen-goten-object
// File: edgelq/devices/proto/v1/customized_image.proto
// DO NOT EDIT!!!

package customized_image

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	os_version "github.com/cloudwan/edgelq-sdk/devices/resources/v1/os_version"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1/project"
	iam_service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1/service_account"
	iam_service_account_key "github.com/cloudwan/edgelq-sdk/iam/resources/v1/service_account_key"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(fmt.Stringer)
	_ = new(sort.Interface)

	_ = new(proto.Message)
	_ = googlefieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &os_version.OsVersion{}
	_ = &project.Project{}
	_ = &iam_service_account.ServiceAccount{}
	_ = &iam_service_account_key.ServiceAccountKey{}
	_ = &meta.Meta{}
)

func (o *CustomizedImage) GotenObjectExt() {}

func (o *CustomizedImage) MakeFullFieldMask() *CustomizedImage_FieldMask {
	return FullCustomizedImage_FieldMask()
}

func (o *CustomizedImage) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullCustomizedImage_FieldMask()
}

func (o *CustomizedImage) MakeDiffFieldMask(other *CustomizedImage) *CustomizedImage_FieldMask {
	if o == nil && other == nil {
		return &CustomizedImage_FieldMask{}
	}
	if o == nil || other == nil {
		return FullCustomizedImage_FieldMask()
	}

	res := &CustomizedImage_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &CustomizedImage_FieldTerminalPath{selector: CustomizedImage_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &CustomizedImage_FieldTerminalPath{selector: CustomizedImage_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &CustomizedImage_FieldSubPath{selector: CustomizedImage_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetSpec().MakeDiffFieldMask(other.GetSpec())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &CustomizedImage_FieldTerminalPath{selector: CustomizedImage_FieldPathSelectorSpec})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &CustomizedImage_FieldSubPath{selector: CustomizedImage_FieldPathSelectorSpec, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetStatus().MakeDiffFieldMask(other.GetStatus())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &CustomizedImage_FieldTerminalPath{selector: CustomizedImage_FieldPathSelectorStatus})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &CustomizedImage_FieldSubPath{selector: CustomizedImage_FieldPathSelectorStatus, subPath: subpath})
			}
		}
	}
	return res
}

func (o *CustomizedImage) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*CustomizedImage))
}

func (o *CustomizedImage) Clone() *CustomizedImage {
	if o == nil {
		return nil
	}
	result := &CustomizedImage{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &Name{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Metadata = o.Metadata.Clone()
	result.Spec = o.Spec.Clone()
	result.Status = o.Status.Clone()
	return result
}

func (o *CustomizedImage) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *CustomizedImage) Merge(source *CustomizedImage) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &Name{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
	if source.GetSpec() != nil {
		if o.Spec == nil {
			o.Spec = new(CustomizedImage_Spec)
		}
		o.Spec.Merge(source.GetSpec())
	}
	if source.GetStatus() != nil {
		if o.Status == nil {
			o.Status = new(CustomizedImage_Status)
		}
		o.Status.Merge(source.GetStatus())
	}
}

func (o *CustomizedImage) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*CustomizedImage))
}

func (o *CustomizedImage_Spec) GotenObjectExt() {}

func (o *CustomizedImage_Spec) MakeFullFieldMask() *CustomizedImage_Spec_FieldMask {
	return FullCustomizedImage_Spec_FieldMask()
}

func (o *CustomizedImage_Spec) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullCustomizedImage_Spec_FieldMask()
}

func (o *CustomizedImage_Spec) MakeDiffFieldMask(other *CustomizedImage_Spec) *CustomizedImage_Spec_FieldMask {
	if o == nil && other == nil {
		return &CustomizedImage_Spec_FieldMask{}
	}
	if o == nil || other == nil {
		return FullCustomizedImage_Spec_FieldMask()
	}

	res := &CustomizedImage_Spec_FieldMask{}
	if o.GetVersion() != other.GetVersion() {
		res.Paths = append(res.Paths, &CustomizedImageSpec_FieldTerminalPath{selector: CustomizedImageSpec_FieldPathSelectorVersion})
	}
	if o.GetDeviceType() != other.GetDeviceType() {
		res.Paths = append(res.Paths, &CustomizedImageSpec_FieldTerminalPath{selector: CustomizedImageSpec_FieldPathSelectorDeviceType})
	}
	if o.GetOsVersion().String() != other.GetOsVersion().String() {
		res.Paths = append(res.Paths, &CustomizedImageSpec_FieldTerminalPath{selector: CustomizedImageSpec_FieldPathSelectorOsVersion})
	}
	if o.GetProvisioningPolicy() != other.GetProvisioningPolicy() {
		res.Paths = append(res.Paths, &CustomizedImageSpec_FieldTerminalPath{selector: CustomizedImageSpec_FieldPathSelectorProvisioningPolicy})
	}
	if o.GetServiceAccount().String() != other.GetServiceAccount().String() {
		res.Paths = append(res.Paths, &CustomizedImageSpec_FieldTerminalPath{selector: CustomizedImageSpec_FieldPathSelectorServiceAccount})
	}
	if o.GetServiceAccountKey().String() != other.GetServiceAccountKey().String() {
		res.Paths = append(res.Paths, &CustomizedImageSpec_FieldTerminalPath{selector: CustomizedImageSpec_FieldPathSelectorServiceAccountKey})
	}
	if o.GetPassword() != other.GetPassword() {
		res.Paths = append(res.Paths, &CustomizedImageSpec_FieldTerminalPath{selector: CustomizedImageSpec_FieldPathSelectorPassword})
	}
	if o.GetEncryption() != other.GetEncryption() {
		res.Paths = append(res.Paths, &CustomizedImageSpec_FieldTerminalPath{selector: CustomizedImageSpec_FieldPathSelectorEncryption})
	}
	if o.GetEncryptionPassword() != other.GetEncryptionPassword() {
		res.Paths = append(res.Paths, &CustomizedImageSpec_FieldTerminalPath{selector: CustomizedImageSpec_FieldPathSelectorEncryptionPassword})
	}
	if o.GetDiskMapping() != other.GetDiskMapping() {
		res.Paths = append(res.Paths, &CustomizedImageSpec_FieldTerminalPath{selector: CustomizedImageSpec_FieldPathSelectorDiskMapping})
	}
	if o.GetNetworkAgent() != other.GetNetworkAgent() {
		res.Paths = append(res.Paths, &CustomizedImageSpec_FieldTerminalPath{selector: CustomizedImageSpec_FieldPathSelectorNetworkAgent})
	}
	if o.GetNtp() != other.GetNtp() {
		res.Paths = append(res.Paths, &CustomizedImageSpec_FieldTerminalPath{selector: CustomizedImageSpec_FieldPathSelectorNtp})
	}
	if o.GetHttpProxy() != other.GetHttpProxy() {
		res.Paths = append(res.Paths, &CustomizedImageSpec_FieldTerminalPath{selector: CustomizedImageSpec_FieldPathSelectorHttpProxy})
	}
	if o.GetHttpsProxy() != other.GetHttpsProxy() {
		res.Paths = append(res.Paths, &CustomizedImageSpec_FieldTerminalPath{selector: CustomizedImageSpec_FieldPathSelectorHttpsProxy})
	}
	if o.GetNoProxy() != other.GetNoProxy() {
		res.Paths = append(res.Paths, &CustomizedImageSpec_FieldTerminalPath{selector: CustomizedImageSpec_FieldPathSelectorNoProxy})
	}
	return res
}

func (o *CustomizedImage_Spec) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*CustomizedImage_Spec))
}

func (o *CustomizedImage_Spec) Clone() *CustomizedImage_Spec {
	if o == nil {
		return nil
	}
	result := &CustomizedImage_Spec{}
	result.Version = o.Version
	result.DeviceType = o.DeviceType
	if o.OsVersion == nil {
		result.OsVersion = nil
	} else if data, err := o.OsVersion.ProtoString(); err != nil {
		panic(err)
	} else {
		result.OsVersion = &os_version.Reference{}
		if err := result.OsVersion.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.ProvisioningPolicy = o.ProvisioningPolicy
	if o.ServiceAccount == nil {
		result.ServiceAccount = nil
	} else if data, err := o.ServiceAccount.ProtoString(); err != nil {
		panic(err)
	} else {
		result.ServiceAccount = &iam_service_account.Reference{}
		if err := result.ServiceAccount.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	if o.ServiceAccountKey == nil {
		result.ServiceAccountKey = nil
	} else if data, err := o.ServiceAccountKey.ProtoString(); err != nil {
		panic(err)
	} else {
		result.ServiceAccountKey = &iam_service_account_key.Reference{}
		if err := result.ServiceAccountKey.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Password = o.Password
	result.Encryption = o.Encryption
	result.EncryptionPassword = o.EncryptionPassword
	result.DiskMapping = o.DiskMapping
	result.NetworkAgent = o.NetworkAgent
	result.Ntp = o.Ntp
	result.HttpProxy = o.HttpProxy
	result.HttpsProxy = o.HttpsProxy
	result.NoProxy = o.NoProxy
	return result
}

func (o *CustomizedImage_Spec) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *CustomizedImage_Spec) Merge(source *CustomizedImage_Spec) {
	o.Version = source.GetVersion()
	o.DeviceType = source.GetDeviceType()
	if source.GetOsVersion() != nil {
		if data, err := source.GetOsVersion().ProtoString(); err != nil {
			panic(err)
		} else {
			o.OsVersion = &os_version.Reference{}
			if err := o.OsVersion.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.OsVersion = nil
	}
	o.ProvisioningPolicy = source.GetProvisioningPolicy()
	if source.GetServiceAccount() != nil {
		if data, err := source.GetServiceAccount().ProtoString(); err != nil {
			panic(err)
		} else {
			o.ServiceAccount = &iam_service_account.Reference{}
			if err := o.ServiceAccount.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.ServiceAccount = nil
	}
	if source.GetServiceAccountKey() != nil {
		if data, err := source.GetServiceAccountKey().ProtoString(); err != nil {
			panic(err)
		} else {
			o.ServiceAccountKey = &iam_service_account_key.Reference{}
			if err := o.ServiceAccountKey.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.ServiceAccountKey = nil
	}
	o.Password = source.GetPassword()
	o.Encryption = source.GetEncryption()
	o.EncryptionPassword = source.GetEncryptionPassword()
	o.DiskMapping = source.GetDiskMapping()
	o.NetworkAgent = source.GetNetworkAgent()
	o.Ntp = source.GetNtp()
	o.HttpProxy = source.GetHttpProxy()
	o.HttpsProxy = source.GetHttpsProxy()
	o.NoProxy = source.GetNoProxy()
}

func (o *CustomizedImage_Spec) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*CustomizedImage_Spec))
}

func (o *CustomizedImage_Status) GotenObjectExt() {}

func (o *CustomizedImage_Status) MakeFullFieldMask() *CustomizedImage_Status_FieldMask {
	return FullCustomizedImage_Status_FieldMask()
}

func (o *CustomizedImage_Status) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullCustomizedImage_Status_FieldMask()
}

func (o *CustomizedImage_Status) MakeDiffFieldMask(other *CustomizedImage_Status) *CustomizedImage_Status_FieldMask {
	if o == nil && other == nil {
		return &CustomizedImage_Status_FieldMask{}
	}
	if o == nil || other == nil {
		return FullCustomizedImage_Status_FieldMask()
	}

	res := &CustomizedImage_Status_FieldMask{}
	if o.GetState() != other.GetState() {
		res.Paths = append(res.Paths, &CustomizedImageStatus_FieldTerminalPath{selector: CustomizedImageStatus_FieldPathSelectorState})
	}
	if o.GetLog() != other.GetLog() {
		res.Paths = append(res.Paths, &CustomizedImageStatus_FieldTerminalPath{selector: CustomizedImageStatus_FieldPathSelectorLog})
	}
	if o.GetFile() != other.GetFile() {
		res.Paths = append(res.Paths, &CustomizedImageStatus_FieldTerminalPath{selector: CustomizedImageStatus_FieldPathSelectorFile})
	}
	if o.GetMd5Sum() != other.GetMd5Sum() {
		res.Paths = append(res.Paths, &CustomizedImageStatus_FieldTerminalPath{selector: CustomizedImageStatus_FieldPathSelectorMd5Sum})
	}
	return res
}

func (o *CustomizedImage_Status) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*CustomizedImage_Status))
}

func (o *CustomizedImage_Status) Clone() *CustomizedImage_Status {
	if o == nil {
		return nil
	}
	result := &CustomizedImage_Status{}
	result.State = o.State
	result.Log = o.Log
	result.File = o.File
	result.Md5Sum = o.Md5Sum
	return result
}

func (o *CustomizedImage_Status) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *CustomizedImage_Status) Merge(source *CustomizedImage_Status) {
	o.State = source.GetState()
	o.Log = source.GetLog()
	o.File = source.GetFile()
	o.Md5Sum = source.GetMd5Sum()
}

func (o *CustomizedImage_Status) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*CustomizedImage_Status))
}
