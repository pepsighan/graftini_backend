package vercel

import "fmt"

// Deployment is a deployment on Vercel.
// Types documented at https://vercel.com/docs/api#endpoints/deployments/create-a-new-deployment/response-parameters.
type Deployment struct {
	ID         string               `json:"id"`
	URL        string               `json:"url"`
	Name       string               `json:"name"`
	ReadyState DeploymentReadyState `json:"readyState"`
	CreatedAt  int                  `json:"createdAt"`
}

// DeploymentReadyState is the current state the deployment is in.
type DeploymentReadyState string

const (
	DeploymentInitializing DeploymentReadyState = "INITIALIZING"
	DeploymentAnalyzing    DeploymentReadyState = "ANALYZING"
	DeploymentBuilding     DeploymentReadyState = "BUILDING"
	DeploymentDeploying    DeploymentReadyState = "DEPLOYING"
	DeploymentReady        DeploymentReadyState = "READY"
	DeploymentError        DeploymentReadyState = "ERROR"
	DeploymentCancelled    DeploymentReadyState = "CANCELED"
)

// CreateNewDeployment creates a new deployment with the given file SHA1 hashes.
func CreateNewDeployment(projectName string, fileSHAs []string) (*Deployment, error) {
	response, err := request().
		SetBody(map[string]interface{}{
			"name":  projectName,
			"files": fileSHAs,
			"projectSettings": map[string]string{
				"framework": "nextjs",
			},
		}).
		SetResult(Deployment{}).
		SetError(VercelFailure{}).
		Post(route("v12/now/deployments"))

	if err != nil {
		return nil, err
	}

	fail, _ := response.Error().(*VercelFailure)
	if fail != nil {
		return nil, fail
	}

	return response.Result().(*Deployment), nil
}

// CancelDeployment cancels the currently running deployment.
func CancelDeployment(deploymentID string) (*Deployment, error) {
	response, err := request().
		SetResult(Deployment{}).
		SetError(VercelFailure{}).
		Patch(route(fmt.Sprintf("v12/now/deployments/%v/cancel", deploymentID)))

	if err != nil {
		return nil, err
	}

	fail, _ := response.Error().(*VercelFailure)
	if fail != nil {
		return nil, fail
	}

	return response.Result().(*Deployment), nil
}
