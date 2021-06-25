package appgenerate

import (
	"encoding/json"
	"strings"

	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/schema"
)

// buildPage generates the NextJS page component.
func buildPage(pg *ent.Page) (string, error) {
	componentMap, err := parseComponentMap(pg.ComponentMap)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString("import { Box, Text } from '@graftini/bricks';\n\n")

	sb.WriteString("export default function ")
	// Name the page component with the UUID. This is not user-readable anyways.
	// Make it simple and unique to implement.
	sb.WriteString(strings.ReplaceAll(pg.ID.String(), "-", ""))
	sb.WriteString("() {\n")
	buildPageMarkup(&sb, componentMap)
	sb.WriteString("}")
	return sb.String(), nil
}

// buildPageMarkup generates the rendering markup for the page.
func buildPageMarkup(sb *strings.Builder, componentMap schema.ComponentMap) {
	sb.WriteString("return (<>")

	// Build the markup from the root.
	root := componentMap["ROOT"]
	for _, childID := range root.ChildrenNodes {
		buildSubTreeMarkup(sb, childID, componentMap)
	}

	sb.WriteString("</>);")
}

// buildSubTreeMarkup generates the markup for the component and its children.
func buildSubTreeMarkup(sb *strings.Builder, componentID string, componentMap schema.ComponentMap) {
	comp := componentMap[componentID]

	// Start tag of the component.
	sb.WriteString("<")
	sb.WriteString(comp.Type)
	sb.WriteString(">")

	// Render the children components.
	if comp.IsCanvas {
		for _, childID := range comp.ChildrenNodes {
			buildSubTreeMarkup(sb, childID, componentMap)
		}
	}

	// End tag of the component.
	sb.WriteString("</")
	sb.WriteString(comp.Type)
	sb.WriteString(">")
}

// parseComponentMap parses the string to a component map.
func parseComponentMap(c string) (schema.ComponentMap, error) {
	componentMap := schema.ComponentMap{}

	err := json.Unmarshal([]byte(c), &componentMap)
	if err != nil {
		return nil, err
	}

	return componentMap, nil
}
