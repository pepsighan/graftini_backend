//go:generate go run -mod=mod github.com/valyala/quicktemplate/qtc -dir=./templates

package appgenerate

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pepsighan/graftini_backend/internal/deploy/appgenerate/templates"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/schema"
)

// GenerateProject generates a code base for the project and returns the file path in which
// it was generated in.
func GenerateCodeBaseForProject(ctx context.Context, project *ent.Project) (CodeBasePath, error) {
	pages, err := project.QueryPages().All(ctx)
	if err != nil {
		return "", err
	}

	// Create a temporary directory
	projectPath, err := newCodeBasePath()
	if err != nil {
		return "", err
	}

	for _, page := range pages {
		if err := writePageToFile(page, "/some/path/for/page"); err != nil {
			return projectPath, err
		}
	}

	return projectPath, nil
}

// writePageToFile writes the page component based on the given page information.
func writePageToFile(p *ent.Page, path string) error {
	if p.ComponentMap == nil {
		return nil
	}

	body, err := generateComponentBody(*p.ComponentMap)
	if err != nil {
		return err
	}

	page := templates.Page(p.Name, body)
	fmt.Println(page)

	return nil
}

func generateComponentBody(c string) (string, error) {
	componentMap := &schema.ComponentMap{}
	err := json.Unmarshal([]byte(c), componentMap)
	if err != nil {
		return "", err
	}

	return templates.PageContent(componentMap), nil
}

// CodeBasePath is the path within which a project is generated.
// This path is a temporary directory which needs to be manually cleaned up
// after use. On Cloud Run, the filesystem actually resides on the RAM, so
// need to be careful to clean things up otherwise it can clog up memory
// between requests (as the container may still be kept alive).
// https://cloud.google.com/run/docs/reference/container-contract#filesystem
type CodeBasePath string

// newCodeBasePath creates a new code base path for a project to be generated in.
func newCodeBasePath() (CodeBasePath, error) {
	path, err := ioutil.TempDir("deployApps", "app")
	if err != nil {
		return "", err
	}

	return CodeBasePath(path), nil
}

// Cleanup removes all the files within the code base path.
func (c CodeBasePath) Cleanup() error {
	return os.RemoveAll(string(c))
}
