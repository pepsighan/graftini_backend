package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	firebase "firebase.google.com/go/v4"
	authFirebase "firebase.google.com/go/v4/auth"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/pepsighan/graftini_backend/internal/backend/auth"
	"github.com/pepsighan/graftini_backend/internal/backend/config"
	"github.com/pepsighan/graftini_backend/internal/backend/graph"
	"github.com/pepsighan/graftini_backend/internal/backend/graph/generated"
	"github.com/pepsighan/graftini_backend/internal/deploy/service"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"google.golang.org/grpc"
)

func firebaseAuth() *authFirebase.Client {
	// The configuration for firebase is taken from the environment.
	// Make sure to add `GOOGLE_APPLICATION_CREDENTIALS` as a JSON file path.
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Panicf("could not initialize firebase app: %v", err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Panicf("could not initialize firebase auth: %v", err)
	}

	return client
}

func grpcClient() (service.DeployClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(config.DeployEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return service.NewDeployClient(conn), conn
}

func graphqlHandler(client *ent.Client) echo.HandlerFunc {
	firebaseClient := firebaseAuth()

	deployClient, grpcConn := grpcClient()
	defer grpcConn.Close()

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers:  graph.NewResolver(client, firebaseClient, deployClient),
		Directives: graph.NewDirective(client, firebaseClient),
	}))

	return func(c echo.Context) error {
		ctx := auth.WithBearerAuth(c)
		h.ServeHTTP(c.Response().Writer, c.Request().WithContext(ctx))
		return nil
	}
}

func playgroundHandler() echo.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}

func main() {
	client, err := ent.Open("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Run the migrations if any on start. The migrations won't drop any fields by
	// default.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	e := echo.New()

	// Recover from panics within route handlers. This saves the app from crashes.
	e.Use(middleware.Recover())

	// Secure middleware provides protection against cross-site scripting (XSS) attack,
	// content type sniffing, clickjacking, insecure connection and other code injection attacks.
	e.Use(middleware.Secure())

	e.Use(middleware.Logger())

	// Do not allow CORs requests by default. If allowed origins are provided, then
	// use them.
	corsConfig := middleware.DefaultCORSConfig
	corsConfig.AllowOrigins = config.AllowedOrigins
	e.Use(middleware.CORSWithConfig(corsConfig))

	// Do not allow any request with body more than 2MB by default. This will
	// limit DoS attacks by file uploads.
	e.Use(middleware.BodyLimit(config.MaxBodySize))

	e.POST("/query", graphqlHandler(client))
	e.GET("/", playgroundHandler())

	// Start server
	go func() {
		if err := e.Start(":" + config.Port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("shutting down the server")
		}
	}()

	// Wait for interrupt or terminate signal to gracefully shutdown the server with a timeout
	// of 10 seconds.
	quit := make(chan os.Signal, 1)
	// SIGINT handles Ctrl+C locally.
	// SIGTERM handles Cloud Run termination signal.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Close database connections before shutting down.
	if err := client.Close(); err != nil {
		log.Printf("failed to close database connection: %v", err)
	}

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
