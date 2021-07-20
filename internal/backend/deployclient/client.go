package deployclient

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/pepsighan/graftini_backend/internal/backend/config"
	"github.com/pepsighan/graftini_backend/internal/deploy/service"
	"github.com/pepsighan/graftini_backend/internal/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func grpcConn() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption

	if config.Env.IsLocal() {
		opts = append(
			opts,
			grpc.WithInsecure(),
		)
	} else {
		// Enable SSL connections.
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		cred := credentials.NewTLS(&tls.Config{
			RootCAs: systemRoots,
		})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}

	conn, err := grpc.Dial(
		config.DeployEndpoint,
		opts...,
	)

	if err != nil {
		return nil, logger.Errorf("could not connect with the deploy server: %w", err)
	}

	return conn, nil
}

// GrpcClient creates a GRPC client for the deploy server.
// Only create a client when it is required.
func GrpcClient() (service.DeployClient, *grpc.ClientConn, error) {
	conn, err := grpcConn()

	if err != nil {
		return nil, nil, logger.Errorf("could not connect with the deploy server: %w", err)
	}

	return service.NewDeployClient(conn), conn, nil
}
