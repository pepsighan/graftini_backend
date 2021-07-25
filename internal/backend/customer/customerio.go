package customer

import (
	"github.com/customerio/go-customerio"
	"github.com/pepsighan/graftini_backend/internal/backend/config"
)

// CustomerIOEvent is a event as defined on Customer.io.
type CustomerIOEvent string

const (
	CustomerIOEvent_UserSignedUp                CustomerIOEvent = "user_signed_up"
	CustomerIOEvent_EarlyAccessAllowed          CustomerIOEvent = "early_access_allowed"
	CustomerIOEvent_ProjectDeployedForFirstTime CustomerIOEvent = "project_deployed_for_the_first_time"
	CustomerIOEvent_SentContactUsQuery          CustomerIOEvent = "sent_contact_us_query"
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
