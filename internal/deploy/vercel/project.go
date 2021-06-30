package vercel

import (
	"context"
	"fmt"
)

// Project is a project that is on vercel.
type Project struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AccountID string `json:"accountId"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}

// ProjectDomain is a domain attached to a project on vercel.
type ProjectDomain struct {
	Name      string `json:"name"`
	ProjectID string `json:"projectId"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}

// GetProject gets a project if it exists otherwise returns nil.
func GetProject(ctx context.Context, name string) (*Project, error) {
	response, err := request(ctx).
		SetResult(Project{}).
		SetError(VercelFailure{}).
		Get(route(fmt.Sprintf("v8/projects/%v", name)))

	if err != nil {
		return nil, fmt.Errorf("could not get the project: %w", err)
	}

	fail, _ := response.Error().(*VercelFailure)
	if fail != nil {
		if fail.VercelError.Code == "not_found" {
			return nil, nil
		}

		return nil, fmt.Errorf("could not get the project: %w", fail)
	}

	return response.Result().(*Project), nil
}

// CreateProject creates a new vercel project.
func CreateProject(ctx context.Context, name string) (*Project, error) {
	response, err := request(ctx).
		SetBody(map[string]string{
			"name": name,
		}).
		SetResult(Project{}).
		SetError(VercelFailure{}).
		Post(route("v8/projects"))

	if err != nil {
		return nil, fmt.Errorf("could not create project: %w", err)
	}

	fail, _ := response.Error().(*VercelFailure)
	if fail != nil {
		return nil, fmt.Errorf("could not create project: %w", fail)
	}

	return response.Result().(*Project), nil
}

// DeleteProject deletes a vercel project.
func DeleteProject(ctx context.Context, projectID string) error {
	response, err := request(ctx).
		SetError(VercelFailure{}).
		Delete(route(fmt.Sprintf("v8/projects/%v", projectID)))

	if err != nil {
		return fmt.Errorf("could not delete project: %w", err)
	}

	fail, _ := response.Error().(*VercelFailure)
	return fail
}

// DoesDomainExistInProject checks if the domain name exists in the project.
func DoesDomainExistInProject(ctx context.Context, projectID string, domainName string) (bool, error) {
	response, err := request(ctx).
		SetResult(ProjectDomain{}).
		SetError(VercelFailure{}).
		Get(route(fmt.Sprintf("v8/projects/%v/domains/%v", projectID, domainName)))

	if err != nil {
		return false, fmt.Errorf("could not check if domain exists: %w", err)
	}

	fail, _ := response.Error().(*VercelFailure)
	if fail != nil {
		if fail.VercelError.Code == "not_found" {
			return false, nil
		}

		return false, fmt.Errorf("could not check if domain exists: %w", err)
	}

	return true, nil
}

// AddDomainToProject adds the given domain name to the project. The domain name must be
// without any protocol://.
func AddDomainToProject(ctx context.Context, projectID string, domainName string) error {
	response, err := request(ctx).
		SetBody(map[string]string{
			"name": domainName,
		}).
		SetError(VercelFailure{}).
		Post(route(fmt.Sprintf("v8/projects/%v/domains", projectID)))

	if err != nil {
		return fmt.Errorf("could not add domain to project: %v", err)
	}

	fail, _ := response.Error().(*VercelFailure)
	if fail != nil {
		return fmt.Errorf("could not add domain to project: %v", err)
	}

	return nil
}

// RemoveDomainFromProject removes the given domain name from the project. The domain name must
// be without any protocol://.
func RemoveDomainFromProject(ctx context.Context, projectID string, domainName string) error {
	response, err := request(ctx).
		SetError(VercelFailure{}).
		Delete(route(fmt.Sprintf("v8/projects/%v/domains/%v", projectID, domainName)))

	if err != nil {
		return fmt.Errorf("could not remove domain from project: %v", err)
	}

	fail, _ := response.Error().(*VercelFailure)
	if fail != nil {
		return fmt.Errorf("could not remove domain from project: %v", err)
	}

	return nil
}
