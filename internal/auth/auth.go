package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo"
	"github.com/pepsighan/graftini_backend/ent"
	"github.com/pepsighan/graftini_backend/ent/user"
)

var ErrUnauthorizedAccess = errors.New("unauthorized access")

type AuthContext struct {
	echo.Context
	token string
}

// user gets the user if a valid authentication header is present in the request. If there is no token in the
// request then it deems it to be an unauthenticated request. If there is it treats the request likewise.
func (a *AuthContext) user(ctx context.Context, entClient *ent.Client, firebaseAuth *auth.Client) (*ent.User, error) {
	if a.token == "" {
		// There is no user in the request. If there is no user for the user, do nothing.
		return nil, nil
	}

	token, err := firebaseAuth.VerifyIDToken(ctx, a.token)
	if err != nil {
		// The request has an authorized token and if the token verification fails, then the user is unauthorized.
		return nil, fmt.Errorf("could not get user from auth: %w due to: %v", ErrUnauthorizedAccess, err)
	}

	// The following errors in the subsequent code are not actually unauthorized access errors because the following
	// errors are caused due to some other failures.
	user, err := entClient.User.Query().Where(user.FirebaseUIDEQ(token.UID)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("could not get user from database: %w", err)
	}

	// Try to save the user as this is the first login.
	if user == nil {
		userRecord, err := firebaseAuth.GetUser(ctx, token.UID)
		if err != nil {
			return nil, fmt.Errorf("could not get user from firebase: %w", err)
		}

		if len(userRecord.ProviderUserInfo) == 0 {
			return nil, fmt.Errorf("no provider user info found: %w", err)
		}

		// Store the user in the database for later. This is probably the first login.
		user, err = entClient.User.Create().
			SetEmail(userRecord.ProviderUserInfo[0].Email).
			SetFirebaseUID(token.UID).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not save user for the first time: %w", err)
		}
	}

	return user, nil
}

// GetUserFromBearerAuthInContext gets user from the context by using bearer auth. The user may be nil if no authorized header is sent.
// This can be used in any query be it @isAuthenicated or not.
func GetUserFromBearerAuthInContext(ctx context.Context, entClient *ent.Client, firebaseAuth *auth.Client) (*ent.User, error) {
	authContext, _ := ctx.Value(authContextKey).(*AuthContext)
	if authContext == nil {
		// This normally won't happen because the auth context is set in the context.
		return nil, nil
	}

	return authContext.user(ctx, entClient, firebaseAuth)
}

// RequiredAuthenticatedUser gets user from the context of an @isAuthenticated query or mutation. It will panic if it found
// no user. So keep in mind to use it only within @isAuthenticated query or mutation. For other situations, use
// [GetUserFromBearerAuthInContext].
func RequiredAuthenticatedUser(ctx context.Context) *ent.User {
	return ctx.Value(authUserContextKey).(*ent.User)
}

// WithBearerAuth extracts the authorization bearer header into the context that can be used later
// on as needed.
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

// WithAuthUser adds the user to the context and returns a new one.
func WithAuthUser(ctx context.Context, user *ent.User) context.Context {
	return context.WithValue(ctx, authUserContextKey, user)
}

type contextKey string

const authContextKey contextKey = "authContext"
const authUserContextKey contextKey = "authUserContext"
