// Code generated by protoc-gen-goten-object
// File: edgelq/cellular-api/proto/v1/contract.proto
// DO NOT EDIT!!!

package contract

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	secrets_secret "github.com/cloudwan/edgelq-sdk/secrets/resources/v1/secret"
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
	_ = &secrets_secret.Secret{}
	_ = &meta.Meta{}
)

func (o *Contract) GotenObjectExt() {}

func (o *Contract) MakeFullFieldMask() *Contract_FieldMask {
	return FullContract_FieldMask()
}

func (o *Contract) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullContract_FieldMask()
}

func (o *Contract) MakeDiffFieldMask(other *Contract) *Contract_FieldMask {
	if o == nil && other == nil {
		return &Contract_FieldMask{}
	}
	if o == nil || other == nil {
		return FullContract_FieldMask()
	}

	res := &Contract_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &Contract_FieldTerminalPath{selector: Contract_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Contract_FieldTerminalPath{selector: Contract_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Contract_FieldSubPath{selector: Contract_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetRatePlan() != other.GetRatePlan() {
		res.Paths = append(res.Paths, &Contract_FieldTerminalPath{selector: Contract_FieldPathSelectorRatePlan})
	}
	{
		_, leftSelected := o.CarrierAccount.(*Contract_Transatel)
		_, rightSelected := other.CarrierAccount.(*Contract_Transatel)
		if leftSelected == rightSelected {
			if o.GetTransatel().String() != other.GetTransatel().String() {
				res.Paths = append(res.Paths, &Contract_FieldTerminalPath{selector: Contract_FieldPathSelectorTransatel})
			}
		} else {
			res.Paths = append(res.Paths, &Contract_FieldTerminalPath{selector: Contract_FieldPathSelectorTransatel})
		}
	}
	{
		_, leftSelected := o.CarrierAccount.(*Contract_Cisco)
		_, rightSelected := other.CarrierAccount.(*Contract_Cisco)
		if leftSelected == rightSelected {
			if o.GetCisco().String() != other.GetCisco().String() {
				res.Paths = append(res.Paths, &Contract_FieldTerminalPath{selector: Contract_FieldPathSelectorCisco})
			}
		} else {
			res.Paths = append(res.Paths, &Contract_FieldTerminalPath{selector: Contract_FieldPathSelectorCisco})
		}
	}
	{
		_, leftSelected := o.CarrierAccount.(*Contract_Celona)
		_, rightSelected := other.CarrierAccount.(*Contract_Celona)
		if leftSelected == rightSelected {
			if o.GetCelona().String() != other.GetCelona().String() {
				res.Paths = append(res.Paths, &Contract_FieldTerminalPath{selector: Contract_FieldPathSelectorCelona})
			}
		} else {
			res.Paths = append(res.Paths, &Contract_FieldTerminalPath{selector: Contract_FieldPathSelectorCelona})
		}
	}
	return res
}

func (o *Contract) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Contract))
}

func (o *Contract) Clone() *Contract {
	if o == nil {
		return nil
	}
	result := &Contract{}
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
	result.RatePlan = o.RatePlan
	if o, ok := o.CarrierAccount.(*Contract_Transatel); ok {
		result.CarrierAccount = (*Contract_Transatel)(nil)
		if o != nil {
			result.CarrierAccount = &Contract_Transatel{}
			result := result.CarrierAccount.(*Contract_Transatel)
			if o.Transatel == nil {
				result.Transatel = nil
			} else if data, err := o.Transatel.ProtoString(); err != nil {
				panic(err)
			} else {
				result.Transatel = &secrets_secret.Reference{}
				if err := result.Transatel.ParseProtoString(data); err != nil {
					panic(err)
				}
			}
		}
	}
	if o, ok := o.CarrierAccount.(*Contract_Cisco); ok {
		result.CarrierAccount = (*Contract_Cisco)(nil)
		if o != nil {
			result.CarrierAccount = &Contract_Cisco{}
			result := result.CarrierAccount.(*Contract_Cisco)
			if o.Cisco == nil {
				result.Cisco = nil
			} else if data, err := o.Cisco.ProtoString(); err != nil {
				panic(err)
			} else {
				result.Cisco = &secrets_secret.Reference{}
				if err := result.Cisco.ParseProtoString(data); err != nil {
					panic(err)
				}
			}
		}
	}
	if o, ok := o.CarrierAccount.(*Contract_Celona); ok {
		result.CarrierAccount = (*Contract_Celona)(nil)
		if o != nil {
			result.CarrierAccount = &Contract_Celona{}
			result := result.CarrierAccount.(*Contract_Celona)
			if o.Celona == nil {
				result.Celona = nil
			} else if data, err := o.Celona.ProtoString(); err != nil {
				panic(err)
			} else {
				result.Celona = &secrets_secret.Reference{}
				if err := result.Celona.ParseProtoString(data); err != nil {
					panic(err)
				}
			}
		}
	}
	return result
}

func (o *Contract) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Contract) Merge(source *Contract) {
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
	o.RatePlan = source.GetRatePlan()
	if source, ok := source.GetCarrierAccount().(*Contract_Transatel); ok {
		if dstOneOf, ok := o.CarrierAccount.(*Contract_Transatel); !ok || dstOneOf == nil {
			o.CarrierAccount = &Contract_Transatel{}
		}
		if source != nil {
			o := o.CarrierAccount.(*Contract_Transatel)
			if source.Transatel != nil {
				if data, err := source.Transatel.ProtoString(); err != nil {
					panic(err)
				} else {
					o.Transatel = &secrets_secret.Reference{}
					if err := o.Transatel.ParseProtoString(data); err != nil {
						panic(err)
					}
				}
			} else {
				o.Transatel = nil
			}
		}
	}
	if source, ok := source.GetCarrierAccount().(*Contract_Cisco); ok {
		if dstOneOf, ok := o.CarrierAccount.(*Contract_Cisco); !ok || dstOneOf == nil {
			o.CarrierAccount = &Contract_Cisco{}
		}
		if source != nil {
			o := o.CarrierAccount.(*Contract_Cisco)
			if source.Cisco != nil {
				if data, err := source.Cisco.ProtoString(); err != nil {
					panic(err)
				} else {
					o.Cisco = &secrets_secret.Reference{}
					if err := o.Cisco.ParseProtoString(data); err != nil {
						panic(err)
					}
				}
			} else {
				o.Cisco = nil
			}
		}
	}
	if source, ok := source.GetCarrierAccount().(*Contract_Celona); ok {
		if dstOneOf, ok := o.CarrierAccount.(*Contract_Celona); !ok || dstOneOf == nil {
			o.CarrierAccount = &Contract_Celona{}
		}
		if source != nil {
			o := o.CarrierAccount.(*Contract_Celona)
			if source.Celona != nil {
				if data, err := source.Celona.ProtoString(); err != nil {
					panic(err)
				} else {
					o.Celona = &secrets_secret.Reference{}
					if err := o.Celona.ParseProtoString(data); err != nil {
						panic(err)
					}
				}
			} else {
				o.Celona = nil
			}
		}
	}
}

func (o *Contract) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Contract))
}
