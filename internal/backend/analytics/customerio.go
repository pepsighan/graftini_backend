package analytics

// CustomerIOEvent is a event as defined on Customer.io.
type CustomerIOEvent string

const (
	CustomerIOEvent_UserSignedUp                = "user_signed_up"
	CustomerIOEvent_EarlyAccessAllowed          = "early_access_allowed"
	CustomerIOEvent_ProjectDeployedForFirstTime = "project_deployed_for_the_first_time"
)

type ProjectDeployedForFirstTimeMeta struct {
	ProjectName string `json:"project_name"`
	AppURL      string `json:"app_url"`
	// ProjectID is a slug id present in the /dashboard/project/:slug-id URL.
	ProjectID string `json:"project_id"`
}
