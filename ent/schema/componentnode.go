//+build ignore
package schema

// This file is manually written so that we can strictly type the schema of a markup before
// storing it in the DB.

// ComponentNode is the representation of a single component that is designed on the editor.
type ComponentNode struct {
	// ID that is unique within a tree.
	ID string `json:"id"`
	// Type of component that is drawn.
	Type string `json:"type"`
	// The kind of display the component is. Either block or inline.
	Display string `json:"display"`
	// Props that are used to define properties that his node has (properties of the associated
	// component).
	Props map[string]interface{} `json:"props"`
	// IsCanvas marks whether this node is a canvas i.e. other nodes can be drawn with it.
	IsCanvas bool `json:"isCanvas"`
	// ParentID is the parent of this component.
	ParentID *string `json:"parentId"`
	// ChildrenNodes are the nodes that are nested within this node. This can have values
	// when this node is a canvas.
	ChildrenNodes []string `json:"childrenNodes"`
}

// ComponentMap is a map of the component nodes. The key here is the ID of the component node.
// The root node of the map has `ROOT` id.
// This is a standardized format designed by Graft and a flattenned out version of the tree.
type ComponentMap map[string]ComponentNode
