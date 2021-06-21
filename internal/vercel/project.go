package vercel

import "fmt"

// Project is a project that is on vercel.
type Project struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AccountID string `json:"accountId"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}

// CreateProject creates a new vercel project.
func CreateProject(name string) (*Project, error) {
	response, err := request().
		SetBody(map[string]string{
			"name": name,
		}).
		SetResult(Project{}).
		SetError(VercelFailure{}).
		Post(route("v8/projects"))

	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	err1, _ := response.Error().(*VercelFailure)
	if err != nil {
		return nil, err1
	}

	return response.Result().(*Project), nil
}

// DeleteProject deletes a vercel project.
func DeleteProject(projectID string) error {
	response, err := request().
		SetError(VercelFailure{}).
		Delete(route(fmt.Sprintf("v8/projects/%v", projectID)))

	if err != nil {
		return err
	}

	err1, _ := response.Error().(*VercelFailure)
	return err1
}
