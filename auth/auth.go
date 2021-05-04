package auth

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
}

// User gets the user if a valid authentication header is present in the request.
func (a *AuthContext) User(ctx context.Context, entClient *ent.Client) (*ent.User, error) {
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

	user, err := entClient.User.Query().Where(user.FirebaseUIDEQ(token.UID)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("could not get user from database: %w", err)
	}

	// Try to save the user as this is the first login.
	if user == nil {
		userRecord, err := client.GetUser(ctx, token.UID)
		if err != nil {
			return nil, fmt.Errorf("could not get user from firebase: %w", err)
		}

		// Store the user in the database for later. This is probably the first login.
		user, err = entClient.User.Create().SetEmail(userRecord.Email).SetFirebaseUID(token.UID).Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not save user for the first time: %w", err)
		}
	}

	return user, nil
}

// UserFromContext gets user from the context.
func UserFromContext(ctx context.Context, entClient *ent.Client) (*ent.User, error) {
	authContext, _ := ctx.Value(authContextKey).(*AuthContext)
	if authContext == nil {
		return nil, nil
	}

	return authContext.User(ctx, entClient)
}

// WithBearerAuth middleware extract the authorization bearer header into a struct that can be used
// later on as needed.
func WithBearerAuth(c echo.Context) context.Context {
	ctx := c.Request().Context()

	authorizationHeader := c.Request().Header.Get("Authorization")
	splits := strings.Split(authorizationHeader, " ")
	if len(splits) != 2 {
		return ctx
	}

	if strings.ToLower(splits[0]) != "bearer" {
		return ctx
	}

	return context.WithValue(ctx, authContextKey, &AuthContext{c, splits[1]})
}

type contextKey string

const authContextKey contextKey = "authContext"
