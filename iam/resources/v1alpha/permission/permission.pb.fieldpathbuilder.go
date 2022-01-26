// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha/permission.proto
// DO NOT EDIT!!!

package permission

// proto imports
import ()

// make sure we're using proto imports
var ()

type PermissionFieldPathBuilder struct{}

func NewPermissionFieldPathBuilder() PermissionFieldPathBuilder {
	return PermissionFieldPathBuilder{}
}
func (PermissionFieldPathBuilder) Name() PermissionPathSelectorName {
	return PermissionPathSelectorName{}
}
func (PermissionFieldPathBuilder) Title() PermissionPathSelectorTitle {
	return PermissionPathSelectorTitle{}
}
func (PermissionFieldPathBuilder) Description() PermissionPathSelectorDescription {
	return PermissionPathSelectorDescription{}
}

type PermissionPathSelectorName struct{}

func (PermissionPathSelectorName) FieldPath() *Permission_FieldTerminalPath {
	return &Permission_FieldTerminalPath{selector: Permission_FieldPathSelectorName}
}

func (s PermissionPathSelectorName) WithValue(value *Name) *Permission_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldTerminalPathValue)
}

func (s PermissionPathSelectorName) WithArrayOfValues(values []*Name) *Permission_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldTerminalPathArrayOfValues)
}

type PermissionPathSelectorTitle struct{}

func (PermissionPathSelectorTitle) FieldPath() *Permission_FieldTerminalPath {
	return &Permission_FieldTerminalPath{selector: Permission_FieldPathSelectorTitle}
}

func (s PermissionPathSelectorTitle) WithValue(value string) *Permission_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldTerminalPathValue)
}

func (s PermissionPathSelectorTitle) WithArrayOfValues(values []string) *Permission_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldTerminalPathArrayOfValues)
}

type PermissionPathSelectorDescription struct{}

func (PermissionPathSelectorDescription) FieldPath() *Permission_FieldTerminalPath {
	return &Permission_FieldTerminalPath{selector: Permission_FieldPathSelectorDescription}
}

func (s PermissionPathSelectorDescription) WithValue(value string) *Permission_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldTerminalPathValue)
}

func (s PermissionPathSelectorDescription) WithArrayOfValues(values []string) *Permission_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldTerminalPathArrayOfValues)
}
