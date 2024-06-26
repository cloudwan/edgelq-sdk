// Code generated by protoc-gen-goten-object
// File: edgelq/devices/proto/v1alpha2/provisioning_approval_request.proto
// DO NOT EDIT!!!

package provisioning_approval_request

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	provisioning_policy "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/provisioning_policy"
	iam_service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/service_account"
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
	_ = &provisioning_policy.ProvisioningPolicy{}
	_ = &iam_service_account.ServiceAccount{}
	_ = &meta.Meta{}
)

func (o *ProvisioningApprovalRequest) GotenObjectExt() {}

func (o *ProvisioningApprovalRequest) MakeFullFieldMask() *ProvisioningApprovalRequest_FieldMask {
	return FullProvisioningApprovalRequest_FieldMask()
}

func (o *ProvisioningApprovalRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullProvisioningApprovalRequest_FieldMask()
}

func (o *ProvisioningApprovalRequest) MakeDiffFieldMask(other *ProvisioningApprovalRequest) *ProvisioningApprovalRequest_FieldMask {
	if o == nil && other == nil {
		return &ProvisioningApprovalRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullProvisioningApprovalRequest_FieldMask()
	}

	res := &ProvisioningApprovalRequest_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &ProvisioningApprovalRequest_FieldTerminalPath{selector: ProvisioningApprovalRequest_FieldPathSelectorName})
	}
	{
		subMask := o.GetSpec().MakeDiffFieldMask(other.GetSpec())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProvisioningApprovalRequest_FieldTerminalPath{selector: ProvisioningApprovalRequest_FieldPathSelectorSpec})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProvisioningApprovalRequest_FieldSubPath{selector: ProvisioningApprovalRequest_FieldPathSelectorSpec, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetStatus().MakeDiffFieldMask(other.GetStatus())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProvisioningApprovalRequest_FieldTerminalPath{selector: ProvisioningApprovalRequest_FieldPathSelectorStatus})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProvisioningApprovalRequest_FieldSubPath{selector: ProvisioningApprovalRequest_FieldPathSelectorStatus, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProvisioningApprovalRequest_FieldTerminalPath{selector: ProvisioningApprovalRequest_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProvisioningApprovalRequest_FieldSubPath{selector: ProvisioningApprovalRequest_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	return res
}

func (o *ProvisioningApprovalRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ProvisioningApprovalRequest))
}

func (o *ProvisioningApprovalRequest) Clone() *ProvisioningApprovalRequest {
	if o == nil {
		return nil
	}
	result := &ProvisioningApprovalRequest{}
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
	result.Spec = o.Spec.Clone()
	result.Status = o.Status.Clone()
	result.Metadata = o.Metadata.Clone()
	return result
}

func (o *ProvisioningApprovalRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ProvisioningApprovalRequest) Merge(source *ProvisioningApprovalRequest) {
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
	if source.GetSpec() != nil {
		if o.Spec == nil {
			o.Spec = new(ProvisioningApprovalRequest_Spec)
		}
		o.Spec.Merge(source.GetSpec())
	}
	if source.GetStatus() != nil {
		if o.Status == nil {
			o.Status = new(ProvisioningApprovalRequest_Status)
		}
		o.Status.Merge(source.GetStatus())
	}
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
}

func (o *ProvisioningApprovalRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ProvisioningApprovalRequest))
}

func (o *ProvisioningApprovalRequest_Spec) GotenObjectExt() {}

func (o *ProvisioningApprovalRequest_Spec) MakeFullFieldMask() *ProvisioningApprovalRequest_Spec_FieldMask {
	return FullProvisioningApprovalRequest_Spec_FieldMask()
}

func (o *ProvisioningApprovalRequest_Spec) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullProvisioningApprovalRequest_Spec_FieldMask()
}

func (o *ProvisioningApprovalRequest_Spec) MakeDiffFieldMask(other *ProvisioningApprovalRequest_Spec) *ProvisioningApprovalRequest_Spec_FieldMask {
	if o == nil && other == nil {
		return &ProvisioningApprovalRequest_Spec_FieldMask{}
	}
	if o == nil || other == nil {
		return FullProvisioningApprovalRequest_Spec_FieldMask()
	}

	res := &ProvisioningApprovalRequest_Spec_FieldMask{}
	if o.GetConclusion() != other.GetConclusion() {
		res.Paths = append(res.Paths, &ProvisioningApprovalRequestSpec_FieldTerminalPath{selector: ProvisioningApprovalRequestSpec_FieldPathSelectorConclusion})
	}
	if o.GetServiceAccount().String() != other.GetServiceAccount().String() {
		res.Paths = append(res.Paths, &ProvisioningApprovalRequestSpec_FieldTerminalPath{selector: ProvisioningApprovalRequestSpec_FieldPathSelectorServiceAccount})
	}
	return res
}

func (o *ProvisioningApprovalRequest_Spec) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ProvisioningApprovalRequest_Spec))
}

func (o *ProvisioningApprovalRequest_Spec) Clone() *ProvisioningApprovalRequest_Spec {
	if o == nil {
		return nil
	}
	result := &ProvisioningApprovalRequest_Spec{}
	result.Conclusion = o.Conclusion
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
	return result
}

func (o *ProvisioningApprovalRequest_Spec) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ProvisioningApprovalRequest_Spec) Merge(source *ProvisioningApprovalRequest_Spec) {
	o.Conclusion = source.GetConclusion()
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
}

func (o *ProvisioningApprovalRequest_Spec) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ProvisioningApprovalRequest_Spec))
}

func (o *ProvisioningApprovalRequest_Status) GotenObjectExt() {}

func (o *ProvisioningApprovalRequest_Status) MakeFullFieldMask() *ProvisioningApprovalRequest_Status_FieldMask {
	return FullProvisioningApprovalRequest_Status_FieldMask()
}

func (o *ProvisioningApprovalRequest_Status) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullProvisioningApprovalRequest_Status_FieldMask()
}

func (o *ProvisioningApprovalRequest_Status) MakeDiffFieldMask(other *ProvisioningApprovalRequest_Status) *ProvisioningApprovalRequest_Status_FieldMask {
	if o == nil && other == nil {
		return &ProvisioningApprovalRequest_Status_FieldMask{}
	}
	if o == nil || other == nil {
		return FullProvisioningApprovalRequest_Status_FieldMask()
	}

	res := &ProvisioningApprovalRequest_Status_FieldMask{}
	return res
}

func (o *ProvisioningApprovalRequest_Status) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ProvisioningApprovalRequest_Status))
}

func (o *ProvisioningApprovalRequest_Status) Clone() *ProvisioningApprovalRequest_Status {
	if o == nil {
		return nil
	}
	result := &ProvisioningApprovalRequest_Status{}
	return result
}

func (o *ProvisioningApprovalRequest_Status) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ProvisioningApprovalRequest_Status) Merge(source *ProvisioningApprovalRequest_Status) {
}

func (o *ProvisioningApprovalRequest_Status) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ProvisioningApprovalRequest_Status))
}
