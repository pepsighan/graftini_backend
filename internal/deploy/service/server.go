package service

import (
	context "context"
)

// Server is used to implement the GRPC deploy service.
type Server struct {
	UnimplementedDeployServer
}

func (s *Server) DeployProject(ctx context.Context, in *DeployRequest) (*DeployReply, error) {
	return nil, nil
}
