module github.com/cloudwan/edgelq-sdk

go 1.16

require (
	github.com/cloudwan/goten-sdk v0.4.35
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

replace google.golang.org/protobuf => github.com/cloudwan/goten-protobuf v1.26.0
