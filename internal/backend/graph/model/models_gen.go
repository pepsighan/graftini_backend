// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/google/uuid"
)

type ContactUsMessage struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Content string `json:"content"`
}

type DuplicatePage struct {
	ProjectID  uuid.UUID `json:"projectId"`
	Name       string    `json:"name"`
	Route      string    `json:"route"`
	CopyPageID uuid.UUID `json:"copyPageId"`
}

type NewGraphQLQuery struct {
	ProjectID    uuid.UUID `json:"projectId"`
	VariableName string    `json:"variableName"`
	GqlAst       string    `json:"gqlAst"`
}

type NewPage struct {
	ProjectID    uuid.UUID `json:"projectId"`
	Name         string    `json:"name"`
	Route        string    `json:"route"`
	ComponentMap string    `json:"componentMap"`
}

type NewProject struct {
	Name                    string `json:"name"`
	DefaultPageComponentMap string `json:"defaultPageComponentMap"`
}

type UpdatePage struct {
	ProjectID uuid.UUID `json:"projectId"`
	PageID    uuid.UUID `json:"pageId"`
	Name      string    `json:"name"`
	Route     string    `json:"route"`
}

type UpdatePageDesign struct {
	PageID       uuid.UUID `json:"pageId"`
	ComponentMap string    `json:"componentMap"`
}

type UpdateProject struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	GraphqlEndpoint *string   `json:"graphqlEndpoint"`
}

type UpdateProjectDesign struct {
	ProjectID uuid.UUID           `json:"projectId"`
	Pages     []*UpdatePageDesign `json:"pages"`
}
