module github.com/cloudwan/edgelq-sdk

go 1.21

require (
	github.com/cloudwan/goten-sdk v0.9.22
	github.com/golang/protobuf v1.5.2
	github.com/google/cel-go v0.5.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/iancoleman/strcase v0.0.0-20180726023541-3605ed457bf7
	github.com/spf13/pflag v1.0.5
	go.opencensus.io v0.23.0
	golang.org/x/oauth2 v0.0.0-20210402161424-2e8d93401602
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.26.0
)

require (
	cloud.google.com/go v0.81.0 // indirect
	github.com/alecthomas/participle v0.5.0 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.1 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/appengine v1.6.7 // indirect
)

replace google.golang.org/protobuf => github.com/cloudwan/goten-protobuf v1.26.0
