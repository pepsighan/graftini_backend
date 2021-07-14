package appgenerate

import (
	"encoding/json"
	"strings"

	"github.com/pepsighan/graftini_backend/internal/pkg/ent/schema"
)

// buildPage generates the NextJS page component.
func buildPage(pg *schema.PageSnapshot) (string, error) {
	var sb strings.Builder
	sb.WriteString("import { Box, Text } from '@graftini/bricks';\n\n")

	sb.WriteString("export default function Page")
	// Name the page component with Page followed by UUID. This is not user-readable anyways.
	// Make it simple and unique to implement.
	sb.WriteString(strings.ReplaceAll(pg.ID, "-", ""))
	sb.WriteString("() {\n")

	if err := buildPageMarkup(&sb, pg.ComponentMap); err != nil {
		return "", err
	}

	sb.WriteString("\n}")
	return sb.String(), nil
}

// buildPageMarkup generates the rendering markup for the page.
func buildPageMarkup(sb *strings.Builder, componentMap schema.ComponentMap) error {
	sb.WriteString("return (<>")

	// Build the markup from the root.
	root := componentMap["ROOT"]
	for _, childID := range root.ChildrenNodes {
		err := buildSubTreeMarkup(sb, childID, componentMap)
		if err != nil {
			return err
		}
	}

	sb.WriteString("</>);")
	return nil
}

// buildSubTreeMarkup generates the markup for the component and its children.
func buildSubTreeMarkup(sb *strings.Builder, componentID string, componentMap schema.ComponentMap) error {
	comp := componentMap[componentID]

	// Start tag of the component.
	sb.WriteString("<")
	sb.WriteString(comp.Type)

	if err := buildProps(sb, &comp); err != nil {
		return err
	}

	sb.WriteString(">\n")

	// Render the children components.
	if comp.IsCanvas {
		for _, childID := range comp.ChildrenNodes {
			buildSubTreeMarkup(sb, childID, componentMap)
		}
	}

	// End tag of the component.
	sb.WriteString("\n</")
	sb.WriteString(comp.Type)
	sb.WriteString(">")

	return nil
}

// buildProps generates a series of prop assignments.
func buildProps(sb *strings.Builder, comp *schema.ComponentNode) error {
	for k, v := range comp.Props {
		sb.WriteString(" ")
		sb.WriteString(k)
		sb.WriteString("={")

		// The prop may be an object
		value, err := json.Marshal(v)
		if err != nil {
			return err
		}

		sb.Write(value)
		sb.WriteString("}")
	}

	return nil
}
