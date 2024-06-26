// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1/attestation_domain.proto
// DO NOT EDIT!!!

package attestation_domain

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1/common"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
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
	_ = &iam_common.PCR{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

func (o *AttestationDomain) GotenObjectExt() {}

func (o *AttestationDomain) MakeFullFieldMask() *AttestationDomain_FieldMask {
	return FullAttestationDomain_FieldMask()
}

func (o *AttestationDomain) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAttestationDomain_FieldMask()
}

func (o *AttestationDomain) MakeDiffFieldMask(other *AttestationDomain) *AttestationDomain_FieldMask {
	if o == nil && other == nil {
		return &AttestationDomain_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAttestationDomain_FieldMask()
	}

	res := &AttestationDomain_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &AttestationDomain_FieldTerminalPath{selector: AttestationDomain_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AttestationDomain_FieldTerminalPath{selector: AttestationDomain_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AttestationDomain_FieldSubPath{selector: AttestationDomain_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &AttestationDomain_FieldTerminalPath{selector: AttestationDomain_FieldPathSelectorDisplayName})
	}
	if o.GetInsecureSkipManufacturerEkcertVerification() != other.GetInsecureSkipManufacturerEkcertVerification() {
		res.Paths = append(res.Paths, &AttestationDomain_FieldTerminalPath{selector: AttestationDomain_FieldPathSelectorInsecureSkipManufacturerEkcertVerification})
	}

	if len(o.GetPolicies()) == len(other.GetPolicies()) {
		for i, lValue := range o.GetPolicies() {
			rValue := other.GetPolicies()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &AttestationDomain_FieldTerminalPath{selector: AttestationDomain_FieldPathSelectorPolicies})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &AttestationDomain_FieldTerminalPath{selector: AttestationDomain_FieldPathSelectorPolicies})
	}

	if len(o.GetEnrollmentList()) == len(other.GetEnrollmentList()) {
		for i, lValue := range o.GetEnrollmentList() {
			rValue := other.GetEnrollmentList()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &AttestationDomain_FieldTerminalPath{selector: AttestationDomain_FieldPathSelectorEnrollmentList})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &AttestationDomain_FieldTerminalPath{selector: AttestationDomain_FieldPathSelectorEnrollmentList})
	}
	return res
}

func (o *AttestationDomain) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*AttestationDomain))
}

func (o *AttestationDomain) Clone() *AttestationDomain {
	if o == nil {
		return nil
	}
	result := &AttestationDomain{}
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
	result.DisplayName = o.DisplayName
	result.InsecureSkipManufacturerEkcertVerification = o.InsecureSkipManufacturerEkcertVerification
	result.Policies = make([]*AttestationDomain_Policy, len(o.Policies))
	for i, sourceValue := range o.Policies {
		result.Policies[i] = sourceValue.Clone()
	}
	result.EnrollmentList = make([]*AttestationDomain_EnrolledKey, len(o.EnrollmentList))
	for i, sourceValue := range o.EnrollmentList {
		result.EnrollmentList[i] = sourceValue.Clone()
	}
	return result
}

func (o *AttestationDomain) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *AttestationDomain) Merge(source *AttestationDomain) {
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
	o.DisplayName = source.GetDisplayName()
	o.InsecureSkipManufacturerEkcertVerification = source.GetInsecureSkipManufacturerEkcertVerification()
	for _, sourceValue := range source.GetPolicies() {
		exists := false
		for _, currentValue := range o.Policies {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *AttestationDomain_Policy
			if sourceValue != nil {
				newDstElement = new(AttestationDomain_Policy)
				newDstElement.Merge(sourceValue)
			}
			o.Policies = append(o.Policies, newDstElement)
		}
	}

	for _, sourceValue := range source.GetEnrollmentList() {
		exists := false
		for _, currentValue := range o.EnrollmentList {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *AttestationDomain_EnrolledKey
			if sourceValue != nil {
				newDstElement = new(AttestationDomain_EnrolledKey)
				newDstElement.Merge(sourceValue)
			}
			o.EnrollmentList = append(o.EnrollmentList, newDstElement)
		}
	}

}

func (o *AttestationDomain) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*AttestationDomain))
}

func (o *AttestationDomain_Policy) GotenObjectExt() {}

func (o *AttestationDomain_Policy) MakeFullFieldMask() *AttestationDomain_Policy_FieldMask {
	return FullAttestationDomain_Policy_FieldMask()
}

func (o *AttestationDomain_Policy) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAttestationDomain_Policy_FieldMask()
}

func (o *AttestationDomain_Policy) MakeDiffFieldMask(other *AttestationDomain_Policy) *AttestationDomain_Policy_FieldMask {
	if o == nil && other == nil {
		return &AttestationDomain_Policy_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAttestationDomain_Policy_FieldMask()
	}

	res := &AttestationDomain_Policy_FieldMask{}
	if o.GetManufacturerRootCaCertsPem() != other.GetManufacturerRootCaCertsPem() {
		res.Paths = append(res.Paths, &AttestationDomainPolicy_FieldTerminalPath{selector: AttestationDomainPolicy_FieldPathSelectorManufacturerRootCaCertsPem})
	}
	if o.GetRequireEnrollment() != other.GetRequireEnrollment() {
		res.Paths = append(res.Paths, &AttestationDomainPolicy_FieldTerminalPath{selector: AttestationDomainPolicy_FieldPathSelectorRequireEnrollment})
	}
	if o.GetVerifyEventLog() != other.GetVerifyEventLog() {
		res.Paths = append(res.Paths, &AttestationDomainPolicy_FieldTerminalPath{selector: AttestationDomainPolicy_FieldPathSelectorVerifyEventLog})
	}

	if len(o.GetExpectedPcrs()) == len(other.GetExpectedPcrs()) {
		for i, lValue := range o.GetExpectedPcrs() {
			rValue := other.GetExpectedPcrs()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &AttestationDomainPolicy_FieldTerminalPath{selector: AttestationDomainPolicy_FieldPathSelectorExpectedPcrs})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &AttestationDomainPolicy_FieldTerminalPath{selector: AttestationDomainPolicy_FieldPathSelectorExpectedPcrs})
	}
	return res
}

func (o *AttestationDomain_Policy) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*AttestationDomain_Policy))
}

func (o *AttestationDomain_Policy) Clone() *AttestationDomain_Policy {
	if o == nil {
		return nil
	}
	result := &AttestationDomain_Policy{}
	result.ManufacturerRootCaCertsPem = o.ManufacturerRootCaCertsPem
	result.RequireEnrollment = o.RequireEnrollment
	result.VerifyEventLog = o.VerifyEventLog
	result.ExpectedPcrs = make([]*iam_common.PCR, len(o.ExpectedPcrs))
	for i, sourceValue := range o.ExpectedPcrs {
		result.ExpectedPcrs[i] = sourceValue.Clone()
	}
	return result
}

func (o *AttestationDomain_Policy) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *AttestationDomain_Policy) Merge(source *AttestationDomain_Policy) {
	o.ManufacturerRootCaCertsPem = source.GetManufacturerRootCaCertsPem()
	o.RequireEnrollment = source.GetRequireEnrollment()
	o.VerifyEventLog = source.GetVerifyEventLog()
	for _, sourceValue := range source.GetExpectedPcrs() {
		exists := false
		for _, currentValue := range o.ExpectedPcrs {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *iam_common.PCR
			if sourceValue != nil {
				newDstElement = new(iam_common.PCR)
				newDstElement.Merge(sourceValue)
			}
			o.ExpectedPcrs = append(o.ExpectedPcrs, newDstElement)
		}
	}

}

func (o *AttestationDomain_Policy) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*AttestationDomain_Policy))
}

func (o *AttestationDomain_EnrolledKey) GotenObjectExt() {}

func (o *AttestationDomain_EnrolledKey) MakeFullFieldMask() *AttestationDomain_EnrolledKey_FieldMask {
	return FullAttestationDomain_EnrolledKey_FieldMask()
}

func (o *AttestationDomain_EnrolledKey) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAttestationDomain_EnrolledKey_FieldMask()
}

func (o *AttestationDomain_EnrolledKey) MakeDiffFieldMask(other *AttestationDomain_EnrolledKey) *AttestationDomain_EnrolledKey_FieldMask {
	if o == nil && other == nil {
		return &AttestationDomain_EnrolledKey_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAttestationDomain_EnrolledKey_FieldMask()
	}

	res := &AttestationDomain_EnrolledKey_FieldMask{}
	if o.GetPubkeyPem() != other.GetPubkeyPem() {
		res.Paths = append(res.Paths, &AttestationDomainEnrolledKey_FieldTerminalPath{selector: AttestationDomainEnrolledKey_FieldPathSelectorPubkeyPem})
	}
	if o.GetComment() != other.GetComment() {
		res.Paths = append(res.Paths, &AttestationDomainEnrolledKey_FieldTerminalPath{selector: AttestationDomainEnrolledKey_FieldPathSelectorComment})
	}
	return res
}

func (o *AttestationDomain_EnrolledKey) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*AttestationDomain_EnrolledKey))
}

func (o *AttestationDomain_EnrolledKey) Clone() *AttestationDomain_EnrolledKey {
	if o == nil {
		return nil
	}
	result := &AttestationDomain_EnrolledKey{}
	result.PubkeyPem = o.PubkeyPem
	result.Comment = o.Comment
	return result
}

func (o *AttestationDomain_EnrolledKey) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *AttestationDomain_EnrolledKey) Merge(source *AttestationDomain_EnrolledKey) {
	o.PubkeyPem = source.GetPubkeyPem()
	o.Comment = source.GetComment()
}

func (o *AttestationDomain_EnrolledKey) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*AttestationDomain_EnrolledKey))
}
