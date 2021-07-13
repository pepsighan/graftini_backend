package deployclient

import (
	"context"
	"fmt"
	"time"

	"github.com/pepsighan/graftini_backend/internal/backend/config"
	"github.com/pepsighan/graftini_backend/internal/deploy/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

func grpcConn() (*grpc.ClientConn, error) {
	// Wait for 100 seconds before the deploy server connects. It won't take as much
	// time but its just a buffer in case the cloud run spins up a new deploy server
	// slowly.
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var opts []grpc.DialOption

	// if config.Env.IsLocal() {
	opts = append(
		opts,
		grpc.WithInsecure(),
		// Try to reconnect as soon as possible.
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: backoff.Config{
				BaseDelay:  1.0 * time.Second,
				Multiplier: 1.6,
				Jitter:     0.2,
				MaxDelay:   10 * time.Second,
			},
		}),
	)
	// } else {
	// 	// Enable SSL connections.
	// 	systemRoots, err := x509.SystemCertPool()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	cred := credentials.NewTLS(&tls.Config{
	// 		RootCAs: systemRoots,
	// 	})
	// 	opts = append(opts, grpc.WithTransportCredentials(cred))
	// }

	conn, err := grpc.DialContext(
		ctx,
		config.DeployEndpoint,
		opts...,
	)

	if err != nil {
		return nil, fmt.Errorf("could not connect with the deploy server: %w", err)
	}

	return conn, nil
}

// GrpcClient creates a GRPC client for the deploy server.
// Only create a client when it is required.
func GrpcClient() (service.DeployClient, *grpc.ClientConn, error) {
	conn, err := grpcConn()

	if err != nil {
		return nil, nil, fmt.Errorf("could not connect with the deploy server: %w", err)
	}

	return service.NewDeployClient(conn), conn, nil
}
