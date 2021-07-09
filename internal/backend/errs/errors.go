// This file lists the custom errors that can be sent to the frontend.
package errs

import (
	"context"
	"errors"

	"github.com/labstack/gommon/log"
	"github.com/pepsighan/graftini_backend/internal/pkg/storage"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

var ErrUnauthorizedAccess = newGQLError("unauthorized_access", "you are not authorized to access")
var ErrServerError = newGQLError("server_error", "server error")
var ErrUnsupportedMimeType = newGQLError("unsupported_mime_type", "file of unsupported mime type uploaded")

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

	log.Errorf("server errored due to: %v", err)
	return ErrServerError
}

// PanicPresenter returns server error while reporting the error.
func PanicPresenter(ctx context.Context, err interface{}) error {
	// TODO: Report the error.
	log.Errorf("server panicked due to: %v", err)
	return ErrServerError
}
