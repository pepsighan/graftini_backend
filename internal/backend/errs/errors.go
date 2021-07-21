// This file lists the custom errors that can be sent to the frontend.
package errs

import (
	"context"
	"errors"

	"github.com/pepsighan/graftini_backend/internal/pkg/storage"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
)

var ErrUnauthorizedAccess = newGQLError("unauthorized_access", "you are not authorized to access")
var ErrServerError = newGQLError("server_error", "server error")
var ErrUnsupportedMimeType = newGQLError("unsupported_mime_type", "file of unsupported mime type uploaded")
var ErrProjectLimitExceeded = newGQLError("project_limit_exceeded", "you have exceeded your project limit")

func newGQLError(kind, message string) *gqlerror.Error {
	return &gqlerror.Error{
		Message: message,
		Extensions: map[string]interface{}{
			"type": kind,
		},
	}
}

// ErrorPresenter only returns the error which are listed above. Any other error will
// result in ErrServerError to be returned (as we deem non-listed errors to be useless or hidden
// from the user's sight).
func ErrorPresenter(ctx context.Context, err error) *gqlerror.Error {
	// Only the above listed errors (except the server error) can be
	// shown to the users. They are user safe.
	if errors.Is(err, ErrUnauthorizedAccess) {
		return err.(*gqlerror.Error)
	}

	// Some errors may not be exactly in the above format. For those transform them
	// to above format.
	if errors.Is(err, storage.ErrUnsupportedMimeType) {
		return ErrUnsupportedMimeType
	}

	// Even though this error might have been logged at the source, we are logging it
	// here again because there may be other errors coming out of the libraries used.
	zap.S().Errorf("server errored due to: %v", err)
	return ErrServerError
}

// PanicPresenter returns server error while reporting the error.
func PanicPresenter(ctx context.Context, err interface{}) error {
	zap.S().Errorf("server panicked due to: %v", err)
	return ErrServerError
}
