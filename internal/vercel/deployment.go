package vercel

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

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

// UploadDeploymentFile uploads the given file and returns the SHA1 hash for it.
func UploadDeploymentFile(file *os.File) (string, error) {
	hash, err := calcSHA1Hash(file)
	if err != nil {
		return "", fmt.Errorf("could not upload deployment file: %w", err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("could not upload deployment file: %w", err)
	}

	response, err := request().
		SetHeader("x-now-digest", hash).
		SetContentLength(true).
		SetBody(bytes).
		SetError(VercelFailure{}).
		Post(route("v2/now/files"))

	if err != nil {
		return "", fmt.Errorf("could not upload deployment file: %w", err)
	}

	fail, _ := response.Error().(*VercelFailure)
	if fail != nil {
		return "", fmt.Errorf("could not upload deployment file: %w", fail)
	}

	return hash, nil
}

// GetDeployment gets the deployment.
func GetDeployment(deploymentID string) (*Deployment, error) {
	response, err := request().
		SetResult(Deployment{}).
		SetError(VercelFailure{}).
		Get(fmt.Sprintf("v11/now/deployments/%v", deploymentID))

	if err != nil {
		return nil, fmt.Errorf("could not get deployment: %w", err)
	}

	fail, _ := response.Error().(*VercelFailure)
	if fail != nil {
		return nil, fmt.Errorf("could not get deployment: %w", fail)
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

// calcSHA1Hash calculates the SHA1 hash for the given file.
func calcSHA1Hash(file *os.File) (string, error) {
	hasher := sha1.New()

	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return string(hasher.Sum(nil)[:]), nil
}
