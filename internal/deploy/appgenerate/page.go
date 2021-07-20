package appgenerate

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/google/uuid"
	"github.com/pepsighan/graftini_backend/internal/deploy/config"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/schema"
	"github.com/pepsighan/graftini_backend/internal/pkg/imagekit"
)

type GenerateContext struct {
	Ent *ent.Client
}

// buildPage generates the NextJS page component.
func buildPage(ctx context.Context, pg *schema.PageSnapshot, generateCtx *GenerateContext) (string, error) {
	var sb strings.Builder
	sb.WriteString("import { Box, Text } from '@graftini/bricks';\n")
	sb.WriteString("import { defaultTextProps } from 'utils/text';\n\n")

	sb.WriteString("export default function Page")
	// Name the page component with Page followed by UUID. This is not user-readable anyways.
	// Make it simple and unique to implement.
	sb.WriteString(strings.ReplaceAll(pg.ID, "-", ""))
	sb.WriteString("() {\n")

	if err := buildPageMarkup(ctx, &sb, pg.ComponentMap, generateCtx); err != nil {
		return "", err
	}

	sb.WriteString("\n}")
	return sb.String(), nil
}

// buildPageMarkup generates the rendering markup for the page.
func buildPageMarkup(ctx context.Context, sb *strings.Builder, componentMap schema.ComponentMap, generateCtx *GenerateContext) error {
	sb.WriteString("return (<>")

	// Build the markup from the root.
	root := componentMap["ROOT"]
	for _, childID := range root.ChildrenNodes {
		err := buildSubTreeMarkup(ctx, sb, childID, componentMap, generateCtx)
		if err != nil {
			return err
		}
	}

	sb.WriteString("</>);")
	return nil
}

// buildSubTreeMarkup generates the markup for the component and its children.
func buildSubTreeMarkup(ctx context.Context, sb *strings.Builder, componentID string, componentMap schema.ComponentMap, generateCtx *GenerateContext) error {
	comp := componentMap[componentID]

	// Start tag of the component.
	sb.WriteString("<")
	sb.WriteString(comp.Type)

	if comp.Type == "Text" {
		// Add default text props to the text component. This is how
		// a default base style for the text component is done.
		sb.WriteString(" {...defaultTextProps}")
	}

	if err := buildProps(ctx, sb, &comp, generateCtx); err != nil {
		return err
	}

	sb.WriteString(">\n")

	// Render the children components.
	if comp.IsCanvas {
		for _, childID := range comp.ChildrenNodes {
			buildSubTreeMarkup(ctx, sb, childID, componentMap, generateCtx)
		}
	}

	// End tag of the component.
	sb.WriteString("\n</")
	sb.WriteString(comp.Type)
	sb.WriteString(">")

	return nil
}

// buildProps generates a series of prop assignments.
func buildProps(ctx context.Context, sb *strings.Builder, comp *schema.ComponentNode, generateCtx *GenerateContext) error {
	for k, v := range comp.Props {
		sb.WriteString(" ")

		switch k {
		case "imageId":
			sb.WriteString("imageUrl")
		default:
			sb.WriteString(k)
		}

		sb.WriteString("={")

		// The prop may be an object
		value, err := json.Marshal(v)
		if err != nil {
			return err
		}

		switch k {
		case "imageId":
			url, err := getImageURL(ctx, value, generateCtx)
			if err != nil {
				return err
			}
			sb.WriteString(url)
		default:
			sb.Write(value)
		}

		sb.WriteString("}")
	}

	return nil
}

// getImageURL gets the image url for the image ID.
func getImageURL(ctx context.Context, imageID []byte, generateCtx *GenerateContext) (string, error) {
	id, err := uuid.FromBytes(imageID)
	if err != nil {
		return "", err
	}

	file, err := generateCtx.Ent.File.Get(ctx, id)
	if err != nil {
		return "", err
	}

	return imagekit.GetImageKitURLForFile(config.ImageKitURLEndpoint, file.ID, file.Kind), nil
}
