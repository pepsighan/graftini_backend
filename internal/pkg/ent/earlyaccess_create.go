// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/earlyaccess"
)

// EarlyAccessCreate is the builder for creating a EarlyAccess entity.
type EarlyAccessCreate struct {
	config
	mutation *EarlyAccessMutation
	hooks    []Hook
}

// SetEmail sets the "email" field.
func (eac *EarlyAccessCreate) SetEmail(s string) *EarlyAccessCreate {
	eac.mutation.SetEmail(s)
	return eac
}

// SetCreatedAt sets the "created_at" field.
func (eac *EarlyAccessCreate) SetCreatedAt(t time.Time) *EarlyAccessCreate {
	eac.mutation.SetCreatedAt(t)
	return eac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eac *EarlyAccessCreate) SetNillableCreatedAt(t *time.Time) *EarlyAccessCreate {
	if t != nil {
		eac.SetCreatedAt(*t)
	}
	return eac
}

// SetUpdatedAt sets the "updated_at" field.
func (eac *EarlyAccessCreate) SetUpdatedAt(t time.Time) *EarlyAccessCreate {
	eac.mutation.SetUpdatedAt(t)
	return eac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (eac *EarlyAccessCreate) SetNillableUpdatedAt(t *time.Time) *EarlyAccessCreate {
	if t != nil {
		eac.SetUpdatedAt(*t)
	}
	return eac
}

// Mutation returns the EarlyAccessMutation object of the builder.
func (eac *EarlyAccessCreate) Mutation() *EarlyAccessMutation {
	return eac.mutation
}

// Save creates the EarlyAccess in the database.
func (eac *EarlyAccessCreate) Save(ctx context.Context) (*EarlyAccess, error) {
	var (
		err  error
		node *EarlyAccess
	)
	eac.defaults()
	if len(eac.hooks) == 0 {
		if err = eac.check(); err != nil {
			return nil, err
		}
		node, err = eac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EarlyAccessMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eac.check(); err != nil {
				return nil, err
			}
			eac.mutation = mutation
			node, err = eac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(eac.hooks) - 1; i >= 0; i-- {
			mut = eac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (eac *EarlyAccessCreate) SaveX(ctx context.Context) *EarlyAccess {
	v, err := eac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (eac *EarlyAccessCreate) defaults() {
	if _, ok := eac.mutation.CreatedAt(); !ok {
		v := earlyaccess.DefaultCreatedAt()
		eac.mutation.SetCreatedAt(v)
	}
	if _, ok := eac.mutation.UpdatedAt(); !ok {
		v := earlyaccess.DefaultUpdatedAt()
		eac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eac *EarlyAccessCreate) check() error {
	if _, ok := eac.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New("ent: missing required field \"email\"")}
	}
	if _, ok := eac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := eac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	return nil
}

func (eac *EarlyAccessCreate) sqlSave(ctx context.Context) (*EarlyAccess, error) {
	_node, _spec := eac.createSpec()
	if err := sqlgraph.CreateNode(ctx, eac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (eac *EarlyAccessCreate) createSpec() (*EarlyAccess, *sqlgraph.CreateSpec) {
	var (
		_node = &EarlyAccess{config: eac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: earlyaccess.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: earlyaccess.FieldID,
			},
		}
	)
	if value, ok := eac.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: earlyaccess.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := eac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: earlyaccess.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := eac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: earlyaccess.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// EarlyAccessCreateBulk is the builder for creating many EarlyAccess entities in bulk.
type EarlyAccessCreateBulk struct {
	config
	builders []*EarlyAccessCreate
}

// Save creates the EarlyAccess entities in the database.
func (eacb *EarlyAccessCreateBulk) Save(ctx context.Context) ([]*EarlyAccess, error) {
	specs := make([]*sqlgraph.CreateSpec, len(eacb.builders))
	nodes := make([]*EarlyAccess, len(eacb.builders))
	mutators := make([]Mutator, len(eacb.builders))
	for i := range eacb.builders {
		func(i int, root context.Context) {
			builder := eacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EarlyAccessMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, eacb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eacb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, eacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eacb *EarlyAccessCreateBulk) SaveX(ctx context.Context) []*EarlyAccess {
	v, err := eacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
