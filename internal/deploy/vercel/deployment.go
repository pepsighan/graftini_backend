package vercel

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"github.com/pepsighan/graftini_backend/internal/deploy/config"
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

// ProjectFile is the metadata of the files that is used to create deployments.
type ProjectFile struct {
	File string `json:"file"`
	SHA  string `json:"sha"`
	Size int    `json:"size"`
}

// CreateNewDeployment creates a new deployment with the given file SHA1 hashes.
func CreateNewDeployment(ctx context.Context, projectName string, files []*ProjectFile) (*Deployment, error) {
	response, err := request(ctx).
		SetBody(map[string]interface{}{
			"name":  projectName,
			"files": files,
			"projectSettings": map[string]string{
				"framework": "nextjs",
			},
			"target": "production", // Only allow production vercel deployments.
			"build": map[string]interface{}{
				"env": map[string]string{
					"NPM_RC": config.GitHubNPMRepoToken,
				},
			},
		}).
		SetResult(Deployment{}).
		SetError(VercelFailure{}).
		Post(route("v12/now/deployments"))

	if err != nil {
		return nil, fmt.Errorf("could not create new deployment: %w", err)
	}

	fail, _ := response.Error().(*VercelFailure)
	if fail != nil {
		return nil, fmt.Errorf("could not create new deployment: %w", fail)
	}

	return response.Result().(*Deployment), nil
}

// UploadDeploymentFile uploads the given file path and returns the SHA1 hash and the content
// length for it.
func UploadDeploymentFile(ctx context.Context, filepath string) (string, int, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", 0, fmt.Errorf("could not upload deployment file: %w", err)
	}

	hash, err := calcSHA1Hash(bytes)
	if err != nil {
		return "", 0, fmt.Errorf("could not upload deployment file: %w", err)
	}

	response, err := request(ctx).
		SetHeader("x-now-digest", hash).
		SetContentLength(true).
		SetBody(bytes).
		SetError(VercelFailure{}).
		Post(route("v2/now/files"))

	if err != nil {
		return "", 0, fmt.Errorf("could not upload deployment file: %w", err)
	}

	fail, _ := response.Error().(*VercelFailure)
	if fail != nil {
		return "", 0, fmt.Errorf("could not upload deployment file: %w", fail)
	}

	return hash, len(bytes), nil
}

// GetDeployment gets the deployment.
func GetDeployment(ctx context.Context, deploymentID string) (*Deployment, error) {
	response, err := request(ctx).
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
func CancelDeployment(ctx context.Context, deploymentID string) (*Deployment, error) {
	response, err := request(ctx).
		SetResult(Deployment{}).
		SetError(VercelFailure{}).
		Patch(route(fmt.Sprintf("v12/now/deployments/%v/cancel", deploymentID)))

	if err != nil {
		return nil, fmt.Errorf("could not cancel deployment: %w", err)
	}

	fail, _ := response.Error().(*VercelFailure)
	if fail != nil {
		return nil, fmt.Errorf("could not cancel deployment: %w", fail)
	}

	return response.Result().(*Deployment), nil
}

// calcSHA1Hash calculates the SHA1 hash for the byte array.
func calcSHA1Hash(bytes []byte) (string, error) {
	hasher := sha1.New()

	_, err := hasher.Write(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
