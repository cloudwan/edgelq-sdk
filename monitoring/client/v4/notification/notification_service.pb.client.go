// Code generated by protoc-gen-goten-client
// API: NotificationService
// DO NOT EDIT!!!

package notification_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	notification "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/notification"
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
	_ = &notification.Notification{}
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

// NotificationServiceClient is the client API for NotificationService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationServiceClient interface {
	GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*notification.Notification, error)
	BatchGetNotifications(ctx context.Context, in *BatchGetNotificationsRequest, opts ...grpc.CallOption) (*BatchGetNotificationsResponse, error)
	ListNotifications(ctx context.Context, in *ListNotificationsRequest, opts ...grpc.CallOption) (*ListNotificationsResponse, error)
	WatchNotification(ctx context.Context, in *WatchNotificationRequest, opts ...grpc.CallOption) (WatchNotificationClientStream, error)
	WatchNotifications(ctx context.Context, in *WatchNotificationsRequest, opts ...grpc.CallOption) (WatchNotificationsClientStream, error)
	CreateNotification(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*notification.Notification, error)
	UpdateNotification(ctx context.Context, in *UpdateNotificationRequest, opts ...grpc.CallOption) (*notification.Notification, error)
	DeleteNotification(ctx context.Context, in *DeleteNotificationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &client{cc}
}

func (c *client) GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*notification.Notification, error) {
	out := new(notification.Notification)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.NotificationService/GetNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetNotifications(ctx context.Context, in *BatchGetNotificationsRequest, opts ...grpc.CallOption) (*BatchGetNotificationsResponse, error) {
	out := new(BatchGetNotificationsResponse)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.NotificationService/BatchGetNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListNotifications(ctx context.Context, in *ListNotificationsRequest, opts ...grpc.CallOption) (*ListNotificationsResponse, error) {
	out := new(ListNotificationsResponse)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.NotificationService/ListNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchNotification(ctx context.Context, in *WatchNotificationRequest, opts ...grpc.CallOption) (WatchNotificationClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchNotification",
			ServerStreams: true,
		},
		"/ntt.monitoring.v4.NotificationService/WatchNotification", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchNotificationWatchNotificationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchNotificationClientStream interface {
	Recv() (*WatchNotificationResponse, error)
	grpc.ClientStream
}

type watchNotificationWatchNotificationClient struct {
	grpc.ClientStream
}

func (x *watchNotificationWatchNotificationClient) Recv() (*WatchNotificationResponse, error) {
	m := new(WatchNotificationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchNotifications(ctx context.Context, in *WatchNotificationsRequest, opts ...grpc.CallOption) (WatchNotificationsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchNotifications",
			ServerStreams: true,
		},
		"/ntt.monitoring.v4.NotificationService/WatchNotifications", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchNotificationsWatchNotificationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchNotificationsClientStream interface {
	Recv() (*WatchNotificationsResponse, error)
	grpc.ClientStream
}

type watchNotificationsWatchNotificationsClient struct {
	grpc.ClientStream
}

func (x *watchNotificationsWatchNotificationsClient) Recv() (*WatchNotificationsResponse, error) {
	m := new(WatchNotificationsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateNotification(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*notification.Notification, error) {
	out := new(notification.Notification)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.NotificationService/CreateNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateNotification(ctx context.Context, in *UpdateNotificationRequest, opts ...grpc.CallOption) (*notification.Notification, error) {
	out := new(notification.Notification)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.NotificationService/UpdateNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteNotification(ctx context.Context, in *DeleteNotificationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.NotificationService/DeleteNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
