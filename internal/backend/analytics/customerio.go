package analytics

import (
	"github.com/customerio/go-customerio"
	"github.com/pepsighan/graftini_backend/internal/backend/config"
	"github.com/pepsighan/graftini_backend/internal/pkg/logger"
)

// CustomerIOEvent is a event as defined on Customer.io.
type CustomerIOEvent string

const (
	CustomerIOEvent_UserSignedUp                CustomerIOEvent = "user_signed_up"
	CustomerIOEvent_EarlyAccessAllowed          CustomerIOEvent = "early_access_allowed"
	CustomerIOEvent_ProjectDeployedForFirstTime CustomerIOEvent = "project_deployed_for_the_first_time"
)

type ProjectDeployedForFirstTimeMeta struct {
	ProjectName string `json:"project_name"`
	AppURL      string `json:"app_url"`
	// ProjectID is a slug id present in the /dashboard/project/:slug-id URL.
	ProjectID string `json:"project_id"`
}

func newClient() *customerio.CustomerIO {
	return customerio.NewTrackClient(config.CustomerIOSiteID, config.CustomerIOAPIKey)
}

// LogUserSignedUp is an event that is logged when user signs up.
func LogUserSignedUp(email string) error {
	track := newClient()

	err := track.Track(email, string(CustomerIOEvent_UserSignedUp), map[string]interface{}{})
	if err != nil {
		return logger.Errorf("could not send %v event to customer.io: %w", CustomerIOEvent_UserSignedUp, err)
	}

	return nil
}

// LogProjectDeployedForFirstTime is an event that is logged when a project is deployed for the
// first time by a user.
func LogProjectDeployedForFirstTime(email string, meta *ProjectDeployedForFirstTimeMeta) error {
	track := newClient()

	err := track.Track(email, string(CustomerIOEvent_ProjectDeployedForFirstTime), map[string]interface{}{
		"project_name": meta.ProjectName,
		"app_url":      meta.AppURL,
		"project_id":   meta.ProjectID,
	})
	if err != nil {
		return logger.Errorf("could not send %v event to customer.io: %w", CustomerIOEvent_ProjectDeployedForFirstTime, err)
	}

	return nil
}
