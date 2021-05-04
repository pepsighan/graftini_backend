package middleware

import (
	"context"
	"fmt"
	"strings"

	firebase "firebase.google.com/go/v4"
	"github.com/labstack/echo"
	"github.com/pepsighan/nocodepress_backend/ent"
	"github.com/pepsighan/nocodepress_backend/ent/user"
)

type AuthContext struct {
	echo.Context
	token string
	ent   *ent.Client
}

// User gets the user if a valid authentication header is present in the request.
func (a *AuthContext) User(ctx context.Context) (*ent.User, error) {
	if a.token == "" {
		// There is no user in the request.
		return nil, nil
	}

	// The configuration for firebase is taken from the environment.
	// Make sure to use `FIREBASE_CONFIG` as a JSON value of the credentials.
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("could not get user from auth: %w", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get user from auth: %w", err)
	}

	token, err := client.VerifyIDToken(ctx, a.token)
	if err != nil {
		return nil, fmt.Errorf("could not get user from auth: %w", err)
	}

	user, err := a.ent.User.Query().Where(user.FirebaseUIDEQ(token.UID)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("could not get user from database: %w", err)
	}

	// Return either the user or none if not found.
	return user, nil
}

// BearerAuth middleware extract the authorization bearer header into a struct that can be used
// later on as needed.
func BearerAuth(ent *ent.Client) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorizationHeader := c.Request().Header.Get("Authorization")
			splits := strings.Split(authorizationHeader, " ")
			if len(splits) != 2 {
				return hf(c)
			}

			if strings.ToLower(splits[0]) != "bearer" {
				return hf(c)
			}

			return hf(&AuthContext{c, splits[1], ent})
		}
	}
}
