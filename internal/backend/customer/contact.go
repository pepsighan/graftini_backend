package customer

import (
	"context"

	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/user"
	"github.com/pepsighan/graftini_backend/internal/pkg/logger"
)

// SendContactUsEmail sends a contact us email to the team that we received from the website.
func SendContactUsEmail(ctx context.Context, name, email, content string, entClient *ent.Client) error {
	user, err := entClient.User.Query().
		Where(user.EmailEQ(email)).
		First(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return err
	}

	client := newClient()

	if user != nil {
		err := client.Track(user.ID.String(), string(CustomerIOEvent_SentContactUsQuery), map[string]interface{}{
			"name":    name,
			"email":   email,
			"content": content,
		})

		if err != nil {
			return logger.Errorf("could not send contact us email: %w", err)
		}

		return nil
	}

	err = client.TrackAnonymous(string(CustomerIOEvent_SentContactUsQuery), map[string]interface{}{
		"name":    name,
		"email":   email,
		"content": content,
	})
	if err != nil {
		return logger.Errorf("could not send contact us email: %w", err)
	}

	return nil
}
