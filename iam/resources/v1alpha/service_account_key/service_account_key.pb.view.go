// Code generated by protoc-gen-goten-resource
// Resource: ServiceAccountKey
// DO NOT EDIT!!!

package service_account_key

import (
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/runtime/api/view"
)

// proto imports
import (
	service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/service_account"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// ensure the imports are used
var (
	_ = fieldmaskpb.FieldMask{}

	_ = view.View_UNSPECIFIED
)

// make sure we're using proto imports
var (
	_ = &service_account.ServiceAccount{}
	_ = &timestamp.Timestamp{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *ServiceAccountKey_FieldMask) *ServiceAccountKey_FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}

	switch viewName {
	case view.View_UNSPECIFIED:
		return extraMask
	case view.View_FULL:
		return nil
	case view.View_BASIC:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "display_name", "algorithm", "valid_not_before", "valid_not_after")
		break
	case view.View_NAME:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "display_name")
		break
	default:
		return extraMask
	}
	if extraMask != nil {
		protoFieldMask.Paths = append(protoFieldMask.Paths, extraMask.ToProtoFieldMask().Paths...)
	}
	res := &ServiceAccountKey_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
