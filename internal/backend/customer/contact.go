package customer

import (
	"context"

	"github.com/customerio/go-customerio"
	"github.com/pepsighan/graftini_backend/internal/backend/config"
	"github.com/pepsighan/graftini_backend/internal/pkg/logger"
)

// SendContactUsEmail sends a contact us email to the team that we
// received from the website.
func SendContactUsEmail(ctx context.Context, name, email, content string) error {
	client := customerio.NewAPIClient(config.CustomerIOAPPAPIKey)

	request := customerio.SendEmailRequest{
		To:                     "team@graftini.com",
		TransactionalMessageID: config.CustomerIOContactUsTransactionalMailID,
		MessageData: map[string]interface{}{
			"name":    name,
			"email":   email,
			"content": content,
		},
	}

	_, err := client.SendEmail(ctx, &request)
	if err != nil {
		return logger.Errorf("could not send contact us email: %w", err)
	}

	return nil
}
