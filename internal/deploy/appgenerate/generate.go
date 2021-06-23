//go:generate go run -mod=mod github.com/valyala/quicktemplate/qtc -dir=./templates

package appgenerate

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pepsighan/graftini_backend/internal/deploy/appgenerate/templates"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/schema"
)

// GenerateProject generates a code base for the project and returns the file path in which
// it was generated in.
func GenerateCodeBaseForProject(ctx context.Context, project *ent.Project) (string, error) {
	pages, err := project.QueryPages().All(ctx)
	if err != nil {
		return "", err
	}

	generatedPath := "/some/project/path"

	for _, page := range pages {
		if err := writePageToFile(page, "/some/path/for/page"); err != nil {
			return "", err
		}
	}

	return generatedPath, nil
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
