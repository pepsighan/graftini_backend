package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	authFirebase "firebase.google.com/go/v4/auth"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/pepsighan/graftini_backend/ent"
	"github.com/pepsighan/graftini_backend/graph"
	"github.com/pepsighan/graftini_backend/graph/generated"
	"github.com/pepsighan/graftini_backend/internal/auth"
	"github.com/pepsighan/graftini_backend/internal/config"
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

func graphqlHandler(client *ent.Client) echo.HandlerFunc {
	firebaseClient := firebaseAuth()

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers:  graph.NewResolver(client, firebaseClient),
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
	e.Logger.Fatal(e.Start(":" + config.Port))
}
