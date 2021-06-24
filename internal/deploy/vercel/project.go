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
