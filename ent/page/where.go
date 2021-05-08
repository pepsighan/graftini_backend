// Code generated by entc, DO NOT EDIT.

package page

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/pepsighan/nocodepress_backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Route applies equality check predicate on the "route" field. It's identical to RouteEQ.
func Route(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoute), v))
	})
}

// Markup applies equality check predicate on the "markup" field. It's identical to MarkupEQ.
func Markup(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMarkup), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Page {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Page {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// RouteEQ applies the EQ predicate on the "route" field.
func RouteEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoute), v))
	})
}

// RouteNEQ applies the NEQ predicate on the "route" field.
func RouteNEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRoute), v))
	})
}

// RouteIn applies the In predicate on the "route" field.
func RouteIn(vs ...string) predicate.Page {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRoute), v...))
	})
}

// RouteNotIn applies the NotIn predicate on the "route" field.
func RouteNotIn(vs ...string) predicate.Page {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRoute), v...))
	})
}

// RouteGT applies the GT predicate on the "route" field.
func RouteGT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRoute), v))
	})
}

// RouteGTE applies the GTE predicate on the "route" field.
func RouteGTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRoute), v))
	})
}

// RouteLT applies the LT predicate on the "route" field.
func RouteLT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRoute), v))
	})
}

// RouteLTE applies the LTE predicate on the "route" field.
func RouteLTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRoute), v))
	})
}

// RouteContains applies the Contains predicate on the "route" field.
func RouteContains(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRoute), v))
	})
}

// RouteHasPrefix applies the HasPrefix predicate on the "route" field.
func RouteHasPrefix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRoute), v))
	})
}

// RouteHasSuffix applies the HasSuffix predicate on the "route" field.
func RouteHasSuffix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRoute), v))
	})
}

// RouteEqualFold applies the EqualFold predicate on the "route" field.
func RouteEqualFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRoute), v))
	})
}

// RouteContainsFold applies the ContainsFold predicate on the "route" field.
func RouteContainsFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRoute), v))
	})
}

// MarkupEQ applies the EQ predicate on the "markup" field.
func MarkupEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMarkup), v))
	})
}

// MarkupNEQ applies the NEQ predicate on the "markup" field.
func MarkupNEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMarkup), v))
	})
}

// MarkupIn applies the In predicate on the "markup" field.
func MarkupIn(vs ...string) predicate.Page {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMarkup), v...))
	})
}

// MarkupNotIn applies the NotIn predicate on the "markup" field.
func MarkupNotIn(vs ...string) predicate.Page {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMarkup), v...))
	})
}

// MarkupGT applies the GT predicate on the "markup" field.
func MarkupGT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMarkup), v))
	})
}

// MarkupGTE applies the GTE predicate on the "markup" field.
func MarkupGTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMarkup), v))
	})
}

// MarkupLT applies the LT predicate on the "markup" field.
func MarkupLT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMarkup), v))
	})
}

// MarkupLTE applies the LTE predicate on the "markup" field.
func MarkupLTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMarkup), v))
	})
}

// MarkupContains applies the Contains predicate on the "markup" field.
func MarkupContains(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMarkup), v))
	})
}

// MarkupHasPrefix applies the HasPrefix predicate on the "markup" field.
func MarkupHasPrefix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMarkup), v))
	})
}

// MarkupHasSuffix applies the HasSuffix predicate on the "markup" field.
func MarkupHasSuffix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMarkup), v))
	})
}

// MarkupEqualFold applies the EqualFold predicate on the "markup" field.
func MarkupEqualFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMarkup), v))
	})
}

// MarkupContainsFold applies the ContainsFold predicate on the "markup" field.
func MarkupContainsFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMarkup), v))
	})
}

// HasPageOf applies the HasEdge predicate on the "pageOf" edge.
func HasPageOf() predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PageOfTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PageOfTable, PageOfColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPageOfWith applies the HasEdge predicate on the "pageOf" edge with a given conditions (other predicates).
func HasPageOfWith(preds ...predicate.Project) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PageOfInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PageOfTable, PageOfColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Page) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Page) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
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
func Not(p predicate.Page) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		p(s.Not())
	})
}
