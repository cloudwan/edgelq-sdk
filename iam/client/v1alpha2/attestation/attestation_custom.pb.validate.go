// Code generated by protoc-gen-goten-validate
// File: edgelq/iam/proto/v1alpha2/attestation_custom.proto
// DO NOT EDIT!!!

package attestation_client

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

	gotenvalidate "github.com/cloudwan/goten-sdk/runtime/validate"
)

// proto imports
import (
	attestation_domain "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/attestation_domain"
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/common"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
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
	_ = gotenvalidate.NewValidationError
)

// make sure we're using proto imports
var (
	_ = &attestation_domain.AttestationDomain{}
	_ = &iam_common.PCR{}
	_ = &project.Project{}
)

func (obj *VerifyRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Msg.(type) {
	case *VerifyRequest_AskForChallenge_:
		if subobj, ok := interface{}(opt.AskForChallenge).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("VerifyRequest", "askForChallenge", opt.AskForChallenge, "nested object validation failed", err)
			}
		}
	case *VerifyRequest_ChallengeResponse_:
		if subobj, ok := interface{}(opt.ChallengeResponse).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("VerifyRequest", "challengeResponse", opt.ChallengeResponse, "nested object validation failed", err)
			}
		}
	default:
		_ = opt
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *VerifyRequest_AskForChallenge) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *VerifyRequest_ChallengeResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Quotes {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ChallengeResponse", "quotes", obj.Quotes[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.Pcrs {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ChallengeResponse", "pcrs", obj.Pcrs[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *VerifyRequest_ChallengeResponse_Quote) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *VerifyResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Msg.(type) {
	case *VerifyResponse_Challenge_:
		if subobj, ok := interface{}(opt.Challenge).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("VerifyResponse", "challenge", opt.Challenge, "nested object validation failed", err)
			}
		}
	case *VerifyResponse_AttestationSuccessful_:
		if subobj, ok := interface{}(opt.AttestationSuccessful).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("VerifyResponse", "attestationSuccessful", opt.AttestationSuccessful, "nested object validation failed", err)
			}
		}
	default:
		_ = opt
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *VerifyResponse_Challenge) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *VerifyResponse_AttestationSuccessful) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
