// Code generated by protoc-gen-goten-client
// API: TpmAttestationCertService
// DO NOT EDIT!!!

package tpm_attestation_cert_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	tpm_attestation_cert "github.com/cloudwan/edgelq-sdk/devices/resources/v1/tpm_attestation_cert"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = protoimpl.DescBuilder{}
	_ = new(context.Context)
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &tpm_attestation_cert.TpmAttestationCert{}
	_ = &emptypb.Empty{}
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TpmAttestationCertServiceClient is the client API for TpmAttestationCertService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TpmAttestationCertServiceClient interface {
	GetTpmAttestationCert(ctx context.Context, in *GetTpmAttestationCertRequest, opts ...grpc.CallOption) (*tpm_attestation_cert.TpmAttestationCert, error)
	BatchGetTpmAttestationCerts(ctx context.Context, in *BatchGetTpmAttestationCertsRequest, opts ...grpc.CallOption) (*BatchGetTpmAttestationCertsResponse, error)
	ListTpmAttestationCerts(ctx context.Context, in *ListTpmAttestationCertsRequest, opts ...grpc.CallOption) (*ListTpmAttestationCertsResponse, error)
	WatchTpmAttestationCert(ctx context.Context, in *WatchTpmAttestationCertRequest, opts ...grpc.CallOption) (WatchTpmAttestationCertClientStream, error)
	WatchTpmAttestationCerts(ctx context.Context, in *WatchTpmAttestationCertsRequest, opts ...grpc.CallOption) (WatchTpmAttestationCertsClientStream, error)
	CreateTpmAttestationCert(ctx context.Context, in *CreateTpmAttestationCertRequest, opts ...grpc.CallOption) (*tpm_attestation_cert.TpmAttestationCert, error)
	UpdateTpmAttestationCert(ctx context.Context, in *UpdateTpmAttestationCertRequest, opts ...grpc.CallOption) (*tpm_attestation_cert.TpmAttestationCert, error)
	DeleteTpmAttestationCert(ctx context.Context, in *DeleteTpmAttestationCertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewTpmAttestationCertServiceClient(cc grpc.ClientConnInterface) TpmAttestationCertServiceClient {
	return &client{cc}
}

func (c *client) GetTpmAttestationCert(ctx context.Context, in *GetTpmAttestationCertRequest, opts ...grpc.CallOption) (*tpm_attestation_cert.TpmAttestationCert, error) {
	out := new(tpm_attestation_cert.TpmAttestationCert)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.TpmAttestationCertService/GetTpmAttestationCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetTpmAttestationCerts(ctx context.Context, in *BatchGetTpmAttestationCertsRequest, opts ...grpc.CallOption) (*BatchGetTpmAttestationCertsResponse, error) {
	out := new(BatchGetTpmAttestationCertsResponse)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.TpmAttestationCertService/BatchGetTpmAttestationCerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListTpmAttestationCerts(ctx context.Context, in *ListTpmAttestationCertsRequest, opts ...grpc.CallOption) (*ListTpmAttestationCertsResponse, error) {
	out := new(ListTpmAttestationCertsResponse)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.TpmAttestationCertService/ListTpmAttestationCerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchTpmAttestationCert(ctx context.Context, in *WatchTpmAttestationCertRequest, opts ...grpc.CallOption) (WatchTpmAttestationCertClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchTpmAttestationCert",
			ServerStreams: true,
		},
		"/ntt.devices.v1.TpmAttestationCertService/WatchTpmAttestationCert", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchTpmAttestationCertWatchTpmAttestationCertClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchTpmAttestationCertClientStream interface {
	Recv() (*WatchTpmAttestationCertResponse, error)
	grpc.ClientStream
}

type watchTpmAttestationCertWatchTpmAttestationCertClient struct {
	grpc.ClientStream
}

func (x *watchTpmAttestationCertWatchTpmAttestationCertClient) Recv() (*WatchTpmAttestationCertResponse, error) {
	m := new(WatchTpmAttestationCertResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchTpmAttestationCerts(ctx context.Context, in *WatchTpmAttestationCertsRequest, opts ...grpc.CallOption) (WatchTpmAttestationCertsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchTpmAttestationCerts",
			ServerStreams: true,
		},
		"/ntt.devices.v1.TpmAttestationCertService/WatchTpmAttestationCerts", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchTpmAttestationCertsWatchTpmAttestationCertsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchTpmAttestationCertsClientStream interface {
	Recv() (*WatchTpmAttestationCertsResponse, error)
	grpc.ClientStream
}

type watchTpmAttestationCertsWatchTpmAttestationCertsClient struct {
	grpc.ClientStream
}

func (x *watchTpmAttestationCertsWatchTpmAttestationCertsClient) Recv() (*WatchTpmAttestationCertsResponse, error) {
	m := new(WatchTpmAttestationCertsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateTpmAttestationCert(ctx context.Context, in *CreateTpmAttestationCertRequest, opts ...grpc.CallOption) (*tpm_attestation_cert.TpmAttestationCert, error) {
	out := new(tpm_attestation_cert.TpmAttestationCert)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.TpmAttestationCertService/CreateTpmAttestationCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateTpmAttestationCert(ctx context.Context, in *UpdateTpmAttestationCertRequest, opts ...grpc.CallOption) (*tpm_attestation_cert.TpmAttestationCert, error) {
	out := new(tpm_attestation_cert.TpmAttestationCert)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.TpmAttestationCertService/UpdateTpmAttestationCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteTpmAttestationCert(ctx context.Context, in *DeleteTpmAttestationCertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.TpmAttestationCertService/DeleteTpmAttestationCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
