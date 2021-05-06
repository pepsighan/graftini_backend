// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/google/uuid"
)

type NewPage struct {
	ProjectID uuid.UUID `json:"projectId"`
	Name      string    `json:"name"`
	Route     string    `json:"route"`
}

type NewProject struct {
	Name string `json:"name"`
}
