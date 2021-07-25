package customer

import (
	"time"

	"github.com/google/uuid"
	"github.com/pepsighan/graftini_backend/internal/pkg/logger"
	"go.uber.org/zap"
)

// LogUser either adds the user if it does not with with customer.io or updates
// them.
func LogUser(userID uuid.UUID, email string, createdAt time.Time) error {
	track := newClient()

	err := track.Identify(userID.String(), map[string]interface{}{
		"email":      email,
		"created_at": createdAt.Unix(),
	})

	if err != nil {
		return logger.Errorf("could not log user to customer.io: %w", err)
	}

	return nil
}

// LogUserSignedUp is an event that is logged when user signs up.
func LogUserSignedUp(email string) {
	track := newClient()

	err := track.Track(email, string(CustomerIOEvent_UserSignedUp), map[string]interface{}{})
	if err != nil {
		zap.S().Errorf("could not send %v event to customer.io: %w", CustomerIOEvent_UserSignedUp, err)
	}
}

// LogProjectDeployedForFirstTime is an event that is logged when a project is deployed for the
// first time by a user.
func LogProjectDeployedForFirstTime(email string, meta *ProjectDeployedForFirstTimeMeta) {
	track := newClient()

	err := track.Track(email, string(CustomerIOEvent_ProjectDeployedForFirstTime), map[string]interface{}{
		"project_name": meta.ProjectName,
		"app_url":      meta.AppURL,
		"project_id":   meta.ProjectID,
	})
	if err != nil {
		zap.S().Errorf("could not send %v event to customer.io: %w", CustomerIOEvent_ProjectDeployedForFirstTime, err)
	}
}
