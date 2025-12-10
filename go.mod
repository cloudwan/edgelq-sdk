module github.com/cloudwan/edgelq-sdk

go 1.22.7

toolchain go1.22.10

require (
	github.com/cloudwan/goten-sdk v1.11.15
	github.com/spf13/pflag v1.0.5
	go.opencensus.io v0.23.0
	golang.org/x/oauth2 v0.23.0
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c
	google.golang.org/grpc v1.68.0
	google.golang.org/protobuf v1.34.2
)

require (
	cloud.google.com/go v0.81.0 // indirect
	github.com/alecthomas/participle/v2 v2.1.4 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.1 // indirect
	github.com/lithammer/shortuuid/v4 v4.2.0 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace google.golang.org/protobuf => github.com/cloudwan/protobuf-go v1.34.2
