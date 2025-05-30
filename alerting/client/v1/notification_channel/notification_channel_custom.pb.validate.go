// Code generated by protoc-gen-goten-validate
// File: edgelq/alerting/proto/v1/notification_channel_custom.proto
// DO NOT EDIT!!!

package notification_channel_client

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
	notification_channel "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/notification_channel"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
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
	_ = &notification_channel.NotificationChannel{}
	_ = &iam_project.Project{}
)

func (obj *TestNotificationChannelRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
