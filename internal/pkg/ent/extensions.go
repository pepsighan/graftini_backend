//+build ignore
package ent

// This package will house any hand-written extension to the query builder of the models.

import (
	"github.com/google/uuid"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/project"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/user"
)

// ByIDAndOwnedBy filters the projects by ID as well has who owns it.
func (pq *ProjectQuery) ByIDAndOwnedBy(id uuid.UUID, ownedBy uuid.UUID) *ProjectQuery {
	return pq.Where(project.And(
		project.IDEQ(id),
		project.HasOwnerWith(user.IDEQ(ownedBy)),
	))
}
