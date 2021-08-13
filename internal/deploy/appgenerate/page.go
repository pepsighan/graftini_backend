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
	"github.com/pepsighan/graftini_backend/internal/pkg/logger"
)

type GenerateContext struct {
	Ent *ent.Client
}

// buildPage generates the NextJS page component.
func buildPage(ctx context.Context, pg *schema.PageSnapshot, generateCtx *GenerateContext) (string, error) {
	var sb strings.Builder
	sb.WriteString("import { Box, Text, rgbaToCss } from '@graftini/bricks';\n")
	sb.WriteString("import Head from 'next/head';\n")
	sb.WriteString("import { Global } from '@emotion/react';\n")
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

	if err := buildSEOForPage(sb, root.Props); err != nil {
		return err
	}

	if err := buildRootStyling(sb, root.Props); err != nil {
		return err
	}

	for _, childID := range root.ChildrenNodes {
		err := buildSubTreeMarkup(ctx, sb, childID, componentMap, true, generateCtx)
		if err != nil {
			return err
		}
	}

	sb.WriteString("</>);")
	return nil
}

// buildSubTreeMarkup generates the markup for the component and its children.
func buildSubTreeMarkup(
	ctx context.Context,
	sb *strings.Builder,
	componentID string,
	componentMap schema.ComponentMap,
	isRootChild bool,
	generateCtx *GenerateContext) error {

	comp := componentMap[componentID]

	// Start tag of the component.
	sb.WriteString("<")
	sb.WriteString(comp.Type)

	if comp.Type == "Text" {
		// Add default text props to the text component. This is how
		// a default base style for the text component is done.
		sb.WriteString(" {...defaultTextProps}")
	}

	if err := buildProps(ctx, sb, &comp, isRootChild, generateCtx); err != nil {
		return err
	}

	sb.WriteString(">\n")

	// Render the children components.
	if comp.IsCanvas {
		for _, childID := range comp.ChildrenNodes {
			err := buildSubTreeMarkup(ctx, sb, childID, componentMap, false, generateCtx)
			if err != nil {
				return err
			}
		}
	}

	// End tag of the component.
	sb.WriteString("\n</")
	sb.WriteString(comp.Type)
	sb.WriteString(">")

	return nil
}

// buildProps generates a series of prop assignments.
func buildProps(
	ctx context.Context,
	sb *strings.Builder,
	comp *schema.ComponentNode,
	isRootChild bool,
	generateCtx *GenerateContext) error {

	if comp.Type == "Text" {
		sb.WriteString(" content={")

		content, err := parseContent(ctx, comp.Props, generateCtx)
		if err != nil {
			return err
		}
		sb.Write(content)
		sb.WriteString("} ")

		width := comp.Props["width"]
		if width != nil {
			err := writePropAndValue(sb, "width", width)
			if err != nil {
				return err
			}
		} else {
			// If no width is provided, use the full-width for backwards compatibility.
			err := writePropAndValue(sb, "width", map[string]interface{}{
				"size": 100,
				"unit": "%",
			})
			if err != nil {
				return err
			}
		}

		// Add space between props.
		sb.WriteString(" ")

		height := comp.Props["height"]
		if height != nil {
			err := writeHeightDimension(isRootChild, sb, "height", height)
			if err != nil {
				return err
			}
		} else {
			// If no height is provided, then use "auto" for backgrounds comptability.
			err := writeHeightDimension(isRootChild, sb, "height", "auto")
			if err != nil {
				return err
			}
		}

		return nil
	}

	// The following is only applicable for box component.
	for k, v := range comp.Props {
		sb.WriteString(" ")

		switch k {
		case "link":
			if v != nil {
				// Link contains either of the two props.
				to, href, err := getLinkURL(ctx, v, generateCtx)
				if err != nil {
					return err
				}

				if to != nil {
					sb.WriteString("to={'")
					sb.WriteString(*to)
					sb.WriteString("'}")
				} else if href != nil {
					sb.WriteString("href={'")
					sb.WriteString(*href)
					sb.WriteString("'}")
				}
			}
		case "imageId":
			sb.WriteString("imageUrl={")
			if v != nil {
				url, err := getImageURL(ctx, v, generateCtx)
				if err != nil {
					return err
				}
				sb.WriteString("'")
				sb.WriteString(url)
				sb.WriteString("'")
			} else {
				sb.WriteString("null")
			}
			sb.WriteString("}")
		case "height":
			err := writeHeightDimension(isRootChild, sb, k, v)
			if err != nil {
				return err
			}
		case "minHeight", "maxHeight":
			err := writeHeightDimension(isRootChild, sb, k, v)
			if err != nil {
				return err
			}
		default:
			err := writePropAndValue(sb, k, v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// writeHeightDimension writes the dimension while transforming % units of height if it is a root child. In the
// editor, the editor itself is of 100vh height. Since we cannot do that to the Root once deployed,
// we are transfering that property to the children. So, 50% child of 100vh Root becomes 50vh of the child.
func writeHeightDimension(isRootChild bool, sb *strings.Builder, k string, v interface{}) error {
	if !isRootChild {
		return writePropAndValue(sb, k, v)
	}

	// This is `auto` or `none` value, nothing extra to do here.
	_, ok := v.(string)
	if ok {
		return writePropAndValue(sb, k, v)
	}

	obj, ok := v.(map[string]interface{})
	if !ok {
		return writePropAndValue(sb, k, v)
	}

	unit := obj["unit"]
	if unit != "%" {
		return writePropAndValue(sb, k, v)
	}

	obj["unit"] = "vh"
	return writePropAndValue(sb, k, obj)
}

// writePropAndValue writes the key and value as props and value as-is.
func writePropAndValue(sb *strings.Builder, k string, v interface{}) error {
	sb.WriteString(k)
	sb.WriteString("={")

	// The prop may be an object
	value, err := json.Marshal(v)
	if err != nil {
		return err
	}

	sb.Write(value)
	sb.WriteString("}")

	return nil
}

// getImageURL gets the image url for the image ID.
func getImageURL(ctx context.Context, imageID interface{}, generateCtx *GenerateContext) (string, error) {
	// imageID can only be a string type.
	id, err := uuid.Parse(imageID.(string))
	if err != nil {
		return "", err
	}

	file, err := generateCtx.Ent.File.Get(ctx, id)
	if err != nil {
		return "", err
	}

	return imagekit.GetImageKitURLForFile(config.ImageKitURLEndpoint, file.ID, file.Kind), nil
}

// getLinkURL gets the URL referred to by the pageID or the href.
func getLinkURL(ctx context.Context, link interface{}, generateCtx *GenerateContext) (*string, *string, error) {
	switch v := link.(type) {
	case map[string]interface{}:
		pageID, _ := v["pageId"].(string)
		href, _ := v["href"].(string)

		// Get the page link.
		if pageID != "" {
			id, err := uuid.Parse(pageID)
			if err != nil {
				return nil, nil, err
			}

			page, err := generateCtx.Ent.Page.Get(ctx, id)
			if err != nil {
				return nil, nil, err
			}

			return &page.Route, nil, nil
		}

		// Return the href as-is.
		if href != "" {
			return nil, &href, nil
		}

	default:
		return nil, nil, logger.Errorf("invalid link type")
	}

	// The link object is not well formed.
	return nil, nil, logger.Errorf("invalid link type")
}

type ProseMirrorDocument struct {
	Type    string                  `json:"type"`
	Content []*ProseMirrorParagraph `json:"content"`
}

type ProseMirrorParagraph struct {
	Type    string                 `json:"type"`
	Attrs   map[string]interface{} `json:"attrs"`
	Content []*ProseMirrorText     `json:"content"`
}

type ProseMirrorText struct {
	Type  string             `json:"type"`
	Marks []*ProseMirrorMark `json:"marks"`
	Text  string             `json:"text"`
}

type ProseMirrorMark struct {
	Type  string                 `json:"type"`
	Attrs map[string]interface{} `json:"attrs"`
}

// parseContent parses the content and modifies the links for use in the app.
func parseContent(ctx context.Context, props map[string]interface{}, generateCtx *GenerateContext) ([]byte, error) {
	var content interface{} = nil

	for k, v := range props {
		if k == "content" {
			content = v
			break
		}
	}

	if content == nil {
		return nil, logger.Errorf("content cannot be nil")
	}

	bytes, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	parsed := ProseMirrorDocument{}
	if err := json.Unmarshal(bytes, &parsed); err != nil {
		return nil, err
	}

	if parsed.Type != "doc" {
		return nil, logger.Errorf("content not of doc type")
	}

	for _, p := range parsed.Content {
		if p.Type != "paragraph" {
			return nil, logger.Errorf("doc content not of paragraph type")
		}

		for _, t := range p.Content {
			if t.Type != "text" {
				return nil, logger.Errorf("paragraph content not of text type")
			}

			for _, m := range t.Marks {
				if m.Type == "link" {
					// Modify the links if pageId and add them back to the marks.
					to, href, err := getLinkURL(ctx, m.Attrs, generateCtx)
					if err != nil {
						return nil, err
					}

					m.Attrs = map[string]interface{}{
						"to":   to,
						"href": href,
					}
				}
			}
		}
	}

	// Convert the parsed that has been modified into a json object.
	return json.Marshal(parsed)
}

// buildSEOForPage builds the SEO for the page.
func buildSEOForPage(sb *strings.Builder, props map[string]interface{}) error {
	seo := props["seo"]
	if seo == nil {
		return nil
	}

	switch s := seo.(type) {
	case map[string]interface{}:
		sb.WriteString("<Head>")

		title, _ := s["title"].(string)
		sb.WriteString("<title>")
		sb.WriteString(title)
		sb.WriteString("</title>\n")

		description, _ := s["description"].(string)
		sb.WriteString("<meta name=\"description\" content={'")
		sb.WriteString(description)
		sb.WriteString("'} />\n")

		sb.WriteString("</Head>\n")

	default:
		return logger.Errorf("invalid type of SEO object")
	}

	return nil
}

// buildRootStyling builds the styling for the root component. The root styling
// are put on the body itself.
func buildRootStyling(sb *strings.Builder, props map[string]interface{}) error {
	color := props["color"]
	if color == nil {
		return nil
	}

	rgba, err := json.Marshal(color)
	if err != nil {
		return err
	}

	sb.WriteString("<Global styles={` body { background-color: ${rgbaToCss(")
	sb.Write(rgba)
	sb.WriteString(")}; } `} />\n")

	return nil
}
