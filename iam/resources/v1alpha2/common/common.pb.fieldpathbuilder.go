// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha2/common.proto
// DO NOT EDIT!!!

package iam_common

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/condition"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/permission"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	role "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/role"
	service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/service_account"
	user "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/user"
	policy "github.com/cloudwan/edgelq-sdk/meta/multi_region/proto/policy"
	syncing_meta "github.com/cloudwan/edgelq-sdk/meta/multi_region/proto/syncing_meta"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &condition.Condition{}
	_ = &organization.Organization{}
	_ = &permission.Permission{}
	_ = &project.Project{}
	_ = &role.Role{}
	_ = &service_account.ServiceAccount{}
	_ = &user.User{}
	_ = &policy.Policy{}
	_ = &syncing_meta.SyncingMeta{}
	_ = &timestamp.Timestamp{}
)

type ActorFieldPathBuilder struct{}

func NewActorFieldPathBuilder() ActorFieldPathBuilder {
	return ActorFieldPathBuilder{}
}
func (ActorFieldPathBuilder) User() ActorPathSelectorUser {
	return ActorPathSelectorUser{}
}
func (ActorFieldPathBuilder) ServiceAccount() ActorPathSelectorServiceAccount {
	return ActorPathSelectorServiceAccount{}
}

type ActorPathSelectorUser struct{}

func (ActorPathSelectorUser) FieldPath() *Actor_FieldTerminalPath {
	return &Actor_FieldTerminalPath{selector: Actor_FieldPathSelectorUser}
}

func (s ActorPathSelectorUser) WithValue(value *user.Reference) *Actor_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Actor_FieldTerminalPathValue)
}

func (s ActorPathSelectorUser) WithArrayOfValues(values []*user.Reference) *Actor_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Actor_FieldTerminalPathArrayOfValues)
}

type ActorPathSelectorServiceAccount struct{}

func (ActorPathSelectorServiceAccount) FieldPath() *Actor_FieldTerminalPath {
	return &Actor_FieldTerminalPath{selector: Actor_FieldPathSelectorServiceAccount}
}

func (s ActorPathSelectorServiceAccount) WithValue(value *service_account.Reference) *Actor_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Actor_FieldTerminalPathValue)
}

func (s ActorPathSelectorServiceAccount) WithArrayOfValues(values []*service_account.Reference) *Actor_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Actor_FieldTerminalPathArrayOfValues)
}

type InvitationFieldPathBuilder struct{}

func NewInvitationFieldPathBuilder() InvitationFieldPathBuilder {
	return InvitationFieldPathBuilder{}
}
func (InvitationFieldPathBuilder) InviteeEmail() InvitationPathSelectorInviteeEmail {
	return InvitationPathSelectorInviteeEmail{}
}
func (InvitationFieldPathBuilder) InviterActor() InvitationPathSelectorInviterActor {
	return InvitationPathSelectorInviterActor{}
}
func (InvitationFieldPathBuilder) InviterFullName() InvitationPathSelectorInviterFullName {
	return InvitationPathSelectorInviterFullName{}
}
func (InvitationFieldPathBuilder) InviterEmail() InvitationPathSelectorInviterEmail {
	return InvitationPathSelectorInviterEmail{}
}
func (InvitationFieldPathBuilder) Roles() InvitationPathSelectorRoles {
	return InvitationPathSelectorRoles{}
}
func (InvitationFieldPathBuilder) ExpirationDate() InvitationPathSelectorExpirationDate {
	return InvitationPathSelectorExpirationDate{}
}
func (InvitationFieldPathBuilder) State() InvitationPathSelectorState {
	return InvitationPathSelectorState{}
}

type InvitationPathSelectorInviteeEmail struct{}

func (InvitationPathSelectorInviteeEmail) FieldPath() *Invitation_FieldTerminalPath {
	return &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorInviteeEmail}
}

func (s InvitationPathSelectorInviteeEmail) WithValue(value string) *Invitation_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Invitation_FieldTerminalPathValue)
}

func (s InvitationPathSelectorInviteeEmail) WithArrayOfValues(values []string) *Invitation_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Invitation_FieldTerminalPathArrayOfValues)
}

type InvitationPathSelectorInviterActor struct{}

func (InvitationPathSelectorInviterActor) FieldPath() *Invitation_FieldTerminalPath {
	return &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorInviterActor}
}

func (s InvitationPathSelectorInviterActor) WithValue(value *Actor) *Invitation_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Invitation_FieldTerminalPathValue)
}

func (s InvitationPathSelectorInviterActor) WithArrayOfValues(values []*Actor) *Invitation_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Invitation_FieldTerminalPathArrayOfValues)
}

func (InvitationPathSelectorInviterActor) WithSubPath(subPath Actor_FieldPath) *Invitation_FieldSubPath {
	return &Invitation_FieldSubPath{selector: Invitation_FieldPathSelectorInviterActor, subPath: subPath}
}

func (s InvitationPathSelectorInviterActor) WithSubValue(subPathValue Actor_FieldPathValue) *Invitation_FieldSubPathValue {
	return &Invitation_FieldSubPathValue{Invitation_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s InvitationPathSelectorInviterActor) WithSubArrayOfValues(subPathArrayOfValues Actor_FieldPathArrayOfValues) *Invitation_FieldSubPathArrayOfValues {
	return &Invitation_FieldSubPathArrayOfValues{Invitation_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s InvitationPathSelectorInviterActor) WithSubArrayItemValue(subPathArrayItemValue Actor_FieldPathArrayItemValue) *Invitation_FieldSubPathArrayItemValue {
	return &Invitation_FieldSubPathArrayItemValue{Invitation_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (InvitationPathSelectorInviterActor) User() InvitationPathSelectorInviterActorUser {
	return InvitationPathSelectorInviterActorUser{}
}

func (InvitationPathSelectorInviterActor) ServiceAccount() InvitationPathSelectorInviterActorServiceAccount {
	return InvitationPathSelectorInviterActorServiceAccount{}
}

type InvitationPathSelectorInviterActorUser struct{}

func (InvitationPathSelectorInviterActorUser) FieldPath() *Invitation_FieldSubPath {
	return &Invitation_FieldSubPath{
		selector: Invitation_FieldPathSelectorInviterActor,
		subPath:  NewActorFieldPathBuilder().User().FieldPath(),
	}
}

func (s InvitationPathSelectorInviterActorUser) WithValue(value *user.Reference) *Invitation_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Invitation_FieldSubPathValue)
}

func (s InvitationPathSelectorInviterActorUser) WithArrayOfValues(values []*user.Reference) *Invitation_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Invitation_FieldSubPathArrayOfValues)
}

type InvitationPathSelectorInviterActorServiceAccount struct{}

func (InvitationPathSelectorInviterActorServiceAccount) FieldPath() *Invitation_FieldSubPath {
	return &Invitation_FieldSubPath{
		selector: Invitation_FieldPathSelectorInviterActor,
		subPath:  NewActorFieldPathBuilder().ServiceAccount().FieldPath(),
	}
}

func (s InvitationPathSelectorInviterActorServiceAccount) WithValue(value *service_account.Reference) *Invitation_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Invitation_FieldSubPathValue)
}

func (s InvitationPathSelectorInviterActorServiceAccount) WithArrayOfValues(values []*service_account.Reference) *Invitation_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Invitation_FieldSubPathArrayOfValues)
}

type InvitationPathSelectorInviterFullName struct{}

func (InvitationPathSelectorInviterFullName) FieldPath() *Invitation_FieldTerminalPath {
	return &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorInviterFullName}
}

func (s InvitationPathSelectorInviterFullName) WithValue(value string) *Invitation_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Invitation_FieldTerminalPathValue)
}

func (s InvitationPathSelectorInviterFullName) WithArrayOfValues(values []string) *Invitation_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Invitation_FieldTerminalPathArrayOfValues)
}

type InvitationPathSelectorInviterEmail struct{}

func (InvitationPathSelectorInviterEmail) FieldPath() *Invitation_FieldTerminalPath {
	return &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorInviterEmail}
}

func (s InvitationPathSelectorInviterEmail) WithValue(value string) *Invitation_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Invitation_FieldTerminalPathValue)
}

func (s InvitationPathSelectorInviterEmail) WithArrayOfValues(values []string) *Invitation_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Invitation_FieldTerminalPathArrayOfValues)
}

type InvitationPathSelectorRoles struct{}

func (InvitationPathSelectorRoles) FieldPath() *Invitation_FieldTerminalPath {
	return &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorRoles}
}

func (s InvitationPathSelectorRoles) WithValue(value []*role.Reference) *Invitation_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Invitation_FieldTerminalPathValue)
}

func (s InvitationPathSelectorRoles) WithArrayOfValues(values [][]*role.Reference) *Invitation_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Invitation_FieldTerminalPathArrayOfValues)
}

func (s InvitationPathSelectorRoles) WithItemValue(value *role.Reference) *Invitation_FieldTerminalPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Invitation_FieldTerminalPathArrayItemValue)
}

type InvitationPathSelectorExpirationDate struct{}

func (InvitationPathSelectorExpirationDate) FieldPath() *Invitation_FieldTerminalPath {
	return &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorExpirationDate}
}

func (s InvitationPathSelectorExpirationDate) WithValue(value *timestamp.Timestamp) *Invitation_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Invitation_FieldTerminalPathValue)
}

func (s InvitationPathSelectorExpirationDate) WithArrayOfValues(values []*timestamp.Timestamp) *Invitation_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Invitation_FieldTerminalPathArrayOfValues)
}

type InvitationPathSelectorState struct{}

func (InvitationPathSelectorState) FieldPath() *Invitation_FieldTerminalPath {
	return &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorState}
}

func (s InvitationPathSelectorState) WithValue(value Invitation_State) *Invitation_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Invitation_FieldTerminalPathValue)
}

func (s InvitationPathSelectorState) WithArrayOfValues(values []Invitation_State) *Invitation_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Invitation_FieldTerminalPathArrayOfValues)
}

type PCRFieldPathBuilder struct{}

func NewPCRFieldPathBuilder() PCRFieldPathBuilder {
	return PCRFieldPathBuilder{}
}
func (PCRFieldPathBuilder) Index() PCRPathSelectorIndex {
	return PCRPathSelectorIndex{}
}
func (PCRFieldPathBuilder) DigestHex() PCRPathSelectorDigestHex {
	return PCRPathSelectorDigestHex{}
}
func (PCRFieldPathBuilder) DigestAlg() PCRPathSelectorDigestAlg {
	return PCRPathSelectorDigestAlg{}
}
func (PCRFieldPathBuilder) Comment() PCRPathSelectorComment {
	return PCRPathSelectorComment{}
}

type PCRPathSelectorIndex struct{}

func (PCRPathSelectorIndex) FieldPath() *PCR_FieldTerminalPath {
	return &PCR_FieldTerminalPath{selector: PCR_FieldPathSelectorIndex}
}

func (s PCRPathSelectorIndex) WithValue(value uint32) *PCR_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*PCR_FieldTerminalPathValue)
}

func (s PCRPathSelectorIndex) WithArrayOfValues(values []uint32) *PCR_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*PCR_FieldTerminalPathArrayOfValues)
}

type PCRPathSelectorDigestHex struct{}

func (PCRPathSelectorDigestHex) FieldPath() *PCR_FieldTerminalPath {
	return &PCR_FieldTerminalPath{selector: PCR_FieldPathSelectorDigestHex}
}

func (s PCRPathSelectorDigestHex) WithValue(value string) *PCR_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*PCR_FieldTerminalPathValue)
}

func (s PCRPathSelectorDigestHex) WithArrayOfValues(values []string) *PCR_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*PCR_FieldTerminalPathArrayOfValues)
}

type PCRPathSelectorDigestAlg struct{}

func (PCRPathSelectorDigestAlg) FieldPath() *PCR_FieldTerminalPath {
	return &PCR_FieldTerminalPath{selector: PCR_FieldPathSelectorDigestAlg}
}

func (s PCRPathSelectorDigestAlg) WithValue(value DigestAlg) *PCR_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*PCR_FieldTerminalPathValue)
}

func (s PCRPathSelectorDigestAlg) WithArrayOfValues(values []DigestAlg) *PCR_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*PCR_FieldTerminalPathArrayOfValues)
}

type PCRPathSelectorComment struct{}

func (PCRPathSelectorComment) FieldPath() *PCR_FieldTerminalPath {
	return &PCR_FieldTerminalPath{selector: PCR_FieldPathSelectorComment}
}

func (s PCRPathSelectorComment) WithValue(value string) *PCR_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*PCR_FieldTerminalPathValue)
}

func (s PCRPathSelectorComment) WithArrayOfValues(values []string) *PCR_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*PCR_FieldTerminalPathArrayOfValues)
}
