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

// GrpcClient creates a GRPC client for the deploy server.
// Only create a client when it is required.
func GrpcClient() (service.DeployClient, *grpc.ClientConn, error) {
	// Wait for 100 seconds before the deploy server connects. It won't take as much
	// time but its just a buffer in case the cloud run spins up a new deploy server
	// slowly.
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		config.DeployEndpoint,
		grpc.WithInsecure(),
		grpc.WithBlock(),
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
	if err != nil {
		return nil, nil, fmt.Errorf("could not connect with the deploy server: %w", err)
	}

	return service.NewDeployClient(conn), conn, nil
}
