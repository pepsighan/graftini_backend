//go:generate go run -mod=mod github.com/valyala/quicktemplate/qtc -dir=./templates

package appgenerate

import (
	"encoding/json"
	"fmt"

	"github.com/pepsighan/graftini_backend/internal/deploy/appgenerate/templates"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/schema"
)

// WritePageToFile writes the page component based on the given page information.
func WritePageToFile(p *ent.Page) error {
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
