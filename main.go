package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/pepsighan/nocodepress_backend/ent"
)

func main() {
	client, err := ent.Open("postgres", "host=<host> port=<port> user=<user> dbname=<database> password=<pass>")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
