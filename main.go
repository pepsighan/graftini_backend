package main

import (
	"context"
	"log"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	authFirebase "firebase.google.com/go/v4/auth"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/pepsighan/nocodepress_backend/auth"
	"github.com/pepsighan/nocodepress_backend/ent"
	"github.com/pepsighan/nocodepress_backend/graph"
	"github.com/pepsighan/nocodepress_backend/graph/generated"
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
	h := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{Resolvers: graph.NewResolver(client, firebaseAuth())},
	))

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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := ent.Open("postgres", os.Getenv("POSTGRES_URL"))
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
	allowedOrigins, ok := os.LookupEnv("ALLOWED_ORIGINS")
	if ok {
		corsConfig.AllowOrigins = strings.Split(allowedOrigins, ",")
	} else {
		corsConfig.AllowOrigins = []string{}
	}

	e.Use(middleware.CORSWithConfig(corsConfig))

	// Do not allow any request with body more than 2MB by default. This will
	// limit DoS attacks by file uploads.
	maxBodySize, ok := os.LookupEnv("MAX_BODY_SIZE")
	if ok {
		e.Use(middleware.BodyLimit(maxBodySize))
	} else {
		e.Use(middleware.BodyLimit("2M"))
	}

	e.POST("/query", graphqlHandler(client))
	e.GET("/", playgroundHandler())
	e.Logger.Fatal(e.Start(":1323"))
}
