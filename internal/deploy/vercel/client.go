package vercel

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/pepsighan/graftini_backend/internal/deploy/config"
)

// client is a shared client to make API requests.
var client = resty.New()

// request to get a request with prefilled configuration to hit Vercel APIs.
func request(ctx context.Context) *resty.Request {
	return client.R().
		SetContext(ctx).
		SetAuthScheme("Bearer").
		SetAuthToken(config.VercelToken).
		SetQueryParam("teamId", config.VercelTeamID)
}

// route generates a full API endpoint from the given relative path. Do not start
// the path with /.
func route(path string) string {
	return fmt.Sprintf("https://api.vercel.com/%v", path)
}

// VercelFailure is an error response returned by vercel.
type VercelFailure struct {
	VercelError struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// Error implements the error interface.
func (e *VercelFailure) Error() string {
	return e.VercelError.Message
}
