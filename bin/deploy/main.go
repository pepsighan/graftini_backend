package main

import (
	"log"
	"net"

	"github.com/getsentry/sentry-go"
	_ "github.com/lib/pq"
	"github.com/pepsighan/graftini_backend/internal/deploy/config"
	"github.com/pepsighan/graftini_backend/internal/deploy/server"
	"github.com/pepsighan/graftini_backend/internal/deploy/service"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"google.golang.org/grpc"
)

func setupSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: config.SentryDSN,
	})
	if err != nil {
		log.Fatalf("could not initialize sentry: %v", err)
	}
}

func main() {
	setupSentry()

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
	service.RegisterDeployServer(s, &server.Server{
		Ent: client,
	})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
