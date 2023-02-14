package utils

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"time"

	"go.opencensus.io/plugin/ocgrpc"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/keepalive"
)

func dialHelper(ctx context.Context, endpoint string, creds credentials.PerRPCCredentials) *grpc.ClientConn {
	addr, _, _ := net.SplitHostPort(endpoint)
	tlsCaPool, err := x509.SystemCertPool()
	if err != nil {
		panic(fmt.Errorf("error getting system cert pool: %s", err))
	}
	tlsConfig := &tls.Config{
		ServerName: addr,
		RootCAs:    tlsCaPool,
	}
	transportCreds := credentials.NewTLS(tlsConfig)

	options := []grpc.DialOption{
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}),
		grpc.WithTransportCredentials(transportCreds),
		grpc.WithPerRPCCredentials(creds),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    time.Minute * 1,
			Timeout: time.Minute / 2,
		}),
	}
	grpcConn, err := grpc.DialContext(ctx, endpoint, options...)
	if err != nil {
		panic(fmt.Sprintf("Error dialing devices: %s", err))
	}
	return grpcConn
}

func DialJWT(ctx context.Context, endpoint, jwtKey string) *grpc.ClientConn {
	jwtCredentials, err := oauth.NewJWTAccessFromKey([]byte(jwtKey))
	if err != nil {
		panic(fmt.Errorf("error parsing ServiceAccount json credentials: %s", err))
	}
	return dialHelper(ctx, endpoint, jwtCredentials)
}

func Dial(ctx context.Context, endpoint, accessToken, credsFile string) *grpc.ClientConn {
	var creds credentials.PerRPCCredentials
	if accessToken != "" {
		creds = oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: accessToken,
			TokenType:   "Bearer",
		})
	} else {
		jwtCredentials, err := oauth.NewJWTAccessFromFile(credsFile)
		if err != nil {
			panic(fmt.Errorf("error parsing ServiceAccount file credentials: %s", err))
		}
		creds = jwtCredentials
	}
	return dialHelper(ctx, endpoint, creds)
}
