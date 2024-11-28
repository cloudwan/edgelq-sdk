module github.com/cloudwan/edgelq-sdk

go 1.22.9

require (
	github.com/cloudwan/goten-sdk v1.7.1-0.20241128161314-a382ab6fa42b
	github.com/google/cel-go v0.20.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/alecthomas/participle v0.5.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.1 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/stoewer/go-strcase v1.2.0 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.9.0 // indirect
)

replace google.golang.org/protobuf => github.com/cloudwan/protobuf-go v1.34.2
