// Code generated by entc, DO NOT EDIT.

package deployment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/predicate"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/schema"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v schema.DeploymentStatus) predicate.Deployment {
	vc := string(v)
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), vc))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v schema.DeploymentStatus) predicate.Deployment {
	vc := string(v)
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), vc))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v schema.DeploymentStatus) predicate.Deployment {
	vc := string(v)
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), vc))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...schema.DeploymentStatus) predicate.Deployment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Deployment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...schema.DeploymentStatus) predicate.Deployment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Deployment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v schema.DeploymentStatus) predicate.Deployment {
	vc := string(v)
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), vc))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v schema.DeploymentStatus) predicate.Deployment {
	vc := string(v)
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), vc))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v schema.DeploymentStatus) predicate.Deployment {
	vc := string(v)
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), vc))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v schema.DeploymentStatus) predicate.Deployment {
	vc := string(v)
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), vc))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v schema.DeploymentStatus) predicate.Deployment {
	vc := string(v)
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), vc))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v schema.DeploymentStatus) predicate.Deployment {
	vc := string(v)
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), vc))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v schema.DeploymentStatus) predicate.Deployment {
	vc := string(v)
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), vc))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v schema.DeploymentStatus) predicate.Deployment {
	vc := string(v)
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), vc))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v schema.DeploymentStatus) predicate.Deployment {
	vc := string(v)
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), vc))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Deployment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deployment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Deployment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deployment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Deployment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deployment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Deployment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Deployment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasDeploymentsOf applies the HasEdge predicate on the "deployments_of" edge.
func HasDeploymentsOf() predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DeploymentsOfTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeploymentsOfTable, DeploymentsOfColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeploymentsOfWith applies the HasEdge predicate on the "deployments_of" edge with a given conditions (other predicates).
func HasDeploymentsOfWith(preds ...predicate.Project) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DeploymentsOfInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeploymentsOfTable, DeploymentsOfColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Deployment) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Deployment) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Deployment) predicate.Deployment {
	return predicate.Deployment(func(s *sql.Selector) {
		p(s.Not())
	})
}
