// Code generated by protoc-gen-goten-validate
// File: edgelq/devices/proto/v1alpha2/broker_custom.proto
// DO NOT EDIT!!!

package broker_client

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
	device "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/device"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/project"
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
	_ = &device.Device{}
	_ = &project.Project{}
)

func (obj *ListenForConnectionsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Msg.(type) {
	case *ListenForConnectionsRequest_RegisterListener_:
		if subobj, ok := interface{}(opt.RegisterListener).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ListenForConnectionsRequest", "registerListener", opt.RegisterListener, "nested object validation failed", err)
			}
		}
	case *ListenForConnectionsRequest_ChannelOpenError_:
		if subobj, ok := interface{}(opt.ChannelOpenError).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ListenForConnectionsRequest", "channelOpenError", opt.ChannelOpenError, "nested object validation failed", err)
			}
		}
	case *ListenForConnectionsRequest_KeepAlive_:
		if subobj, ok := interface{}(opt.KeepAlive).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ListenForConnectionsRequest", "keepAlive", opt.KeepAlive, "nested object validation failed", err)
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
func (obj *ListenForConnectionsRequest_RegisterListener) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListenForConnectionsRequest_ChannelOpenError) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListenForConnectionsRequest_KeepAlive) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListenForConnectionsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Msg.(type) {
	case *ListenForConnectionsResponse_ChannelRequested_:
		if subobj, ok := interface{}(opt.ChannelRequested).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ListenForConnectionsResponse", "channelRequested", opt.ChannelRequested, "nested object validation failed", err)
			}
		}
	case *ListenForConnectionsResponse_KeepAlive_:
		if subobj, ok := interface{}(opt.KeepAlive).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ListenForConnectionsResponse", "keepAlive", opt.KeepAlive, "nested object validation failed", err)
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
func (obj *ListenForConnectionsResponse_ChannelRequested) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListenForConnectionsResponse_KeepAlive) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *OpenConnectionChannelSocketRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Msg.(type) {
	case *OpenConnectionChannelSocketRequest_RegisterSocket_:
		if subobj, ok := interface{}(opt.RegisterSocket).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("OpenConnectionChannelSocketRequest", "registerSocket", opt.RegisterSocket, "nested object validation failed", err)
			}
		}
	case *OpenConnectionChannelSocketRequest_Data:
	case *OpenConnectionChannelSocketRequest_Error:
	case *OpenConnectionChannelSocketRequest_Ack:
		if subobj, ok := interface{}(opt.Ack).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("OpenConnectionChannelSocketRequest", "ack", opt.Ack, "nested object validation failed", err)
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
func (obj *OpenConnectionChannelSocketRequest_RegisterSocket) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *OpenConnectionChannelSocketResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Msg.(type) {
	case *OpenConnectionChannelSocketResponse_Data:
	case *OpenConnectionChannelSocketResponse_Ack:
		if subobj, ok := interface{}(opt.Ack).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("OpenConnectionChannelSocketResponse", "ack", opt.Ack, "nested object validation failed", err)
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
func (obj *ConnectToDeviceRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Msg.(type) {
	case *ConnectToDeviceRequest_OpenRequest_:
		if subobj, ok := interface{}(opt.OpenRequest).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ConnectToDeviceRequest", "openRequest", opt.OpenRequest, "nested object validation failed", err)
			}
		}
	case *ConnectToDeviceRequest_Data:
	case *ConnectToDeviceRequest_Ack:
		if subobj, ok := interface{}(opt.Ack).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ConnectToDeviceRequest", "ack", opt.Ack, "nested object validation failed", err)
			}
		}
	case *ConnectToDeviceRequest_KeepAlive_:
		if subobj, ok := interface{}(opt.KeepAlive).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ConnectToDeviceRequest", "keepAlive", opt.KeepAlive, "nested object validation failed", err)
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
func (obj *ConnectToDeviceRequest_OpenRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ConnectToDeviceRequest_KeepAlive) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ConnectToDeviceResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Msg.(type) {
	case *ConnectToDeviceResponse_OpenResponse_:
		if subobj, ok := interface{}(opt.OpenResponse).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ConnectToDeviceResponse", "openResponse", opt.OpenResponse, "nested object validation failed", err)
			}
		}
	case *ConnectToDeviceResponse_Data:
	case *ConnectToDeviceResponse_Error:
	case *ConnectToDeviceResponse_Ack:
		if subobj, ok := interface{}(opt.Ack).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ConnectToDeviceResponse", "ack", opt.Ack, "nested object validation failed", err)
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
func (obj *ConnectToDeviceResponse_OpenResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Ack) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *SSHService) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *SSHService_Hello) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *SSHService_TerminalSize) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *SSHService_ClientOut) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Msg.(type) {
	case *SSHService_ClientOut_Data:
	case *SSHService_ClientOut_SshHello:
		if subobj, ok := interface{}(opt.SshHello).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ClientOut", "sshHello", opt.SshHello, "nested object validation failed", err)
			}
		}
	case *SSHService_ClientOut_SshResizeTerminal:
		if subobj, ok := interface{}(opt.SshResizeTerminal).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ClientOut", "sshResizeTerminal", opt.SshResizeTerminal, "nested object validation failed", err)
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
func (obj *SSHService_ClientIn) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Msg.(type) {
	case *SSHService_ClientIn_Data:
	default:
		_ = opt
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *SCPService) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Msg.(type) {
	case *SCPService_Dir:
		if subobj, ok := interface{}(opt.Dir).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("SCPService", "dir", opt.Dir, "nested object validation failed", err)
			}
		}
	case *SCPService_File:
		if subobj, ok := interface{}(opt.File).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("SCPService", "file", opt.File, "nested object validation failed", err)
			}
		}
	case *SCPService_Eot:
	case *SCPService_Config:
		if subobj, ok := interface{}(opt.Config).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("SCPService", "config", opt.Config, "nested object validation failed", err)
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
func (obj *SCPService_Configure) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *SCPService_CreateDirectory) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *SCPService_CreateFile) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Msg.(type) {
	case *SCPService_CreateFile_Init:
		if subobj, ok := interface{}(opt.Init).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("CreateFile", "init", opt.Init, "nested object validation failed", err)
			}
		}
	case *SCPService_CreateFile_Data:
	case *SCPService_CreateFile_Eof:
	default:
		_ = opt
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *SCPService_CreateFile_Initialize) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *LogsService) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *LogsService_ToDevice) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *LogsService_ToClient) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *PodManagementService) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *SystemStateService) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
