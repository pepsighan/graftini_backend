//+build ignore
package schema

import (
	"bytes"
	"encoding/json"
	"log"
)

// This file is manually written so that we can strictly type the schema of a markup before
// storing it in the DB.

// MarkupNode is the representation of a single component that is designed on the editor.
type MarkupNode struct {
	// A unique ID within a markup.
	ID string `json:"id"`
	// The kind of Component that is drawn here.
	Component string `json:"component"`
	// Props that are used to define properties that his node has (properties of the associated
	// component).
	Props map[string]interface{} `json:"props"`
	// IsCanvas marks whether this node is a canvas i.e. other nodes can be drawn with it.
	IsCanvas bool `json:"isCanvas"`
	// ChildrenNodes are the nodes that are nested within this node. This can have values
	// when this node is a canvas.
	ChildrenNodes []string `json:"childrenNodes"`
}

// Markup is a map of the markup nodes. The key here is the ID of the markup node.
// The root node of the markup has `ROOT` id.
type Markup map[string]MarkupNode

const rootComponentID = "ROOT"
const rootComponent = "Root"

// defaultMarkup is markup where an empty canvas is drawn with a root node.
func defaultMarkup() string {
	markup, err := json.Marshal(Markup(map[string]MarkupNode{
		rootComponentID: {
			ID:        rootComponentID,
			Component: rootComponent,
			IsCanvas:  true,
			Props: map[string]interface{}{
				"backgroundColor": map[string]int{
					"r": 255,
					"g": 255,
					"b": 255,
					"a": 1,
				},
			},
		},
	}))
	if err != nil {
		log.Fatalf("wrong default markup is defined: %v", err)
	}

	return bytes.NewBuffer(markup).String()
}
