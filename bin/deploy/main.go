package main

import (
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/pepsighan/graftini_backend/internal/deploy/config"
	"github.com/pepsighan/graftini_backend/internal/deploy/service"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	client, err := ent.Open("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	s := grpc.NewServer()
	service.RegisterDeployServer(s, &service.Server{
		Ent: client,
	})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
