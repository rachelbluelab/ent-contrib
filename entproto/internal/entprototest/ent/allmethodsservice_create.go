// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/allmethodsservice"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AllMethodsServiceCreate is the builder for creating a AllMethodsService entity.
type AllMethodsServiceCreate struct {
	config
	mutation *AllMethodsServiceMutation
	hooks    []Hook
}

// Mutation returns the AllMethodsServiceMutation object of the builder.
func (amsc *AllMethodsServiceCreate) Mutation() *AllMethodsServiceMutation {
	return amsc.mutation
}

// Save creates the AllMethodsService in the database.
func (amsc *AllMethodsServiceCreate) Save(ctx context.Context) (*AllMethodsService, error) {
	var (
		err  error
		node *AllMethodsService
	)
	if len(amsc.hooks) == 0 {
		if err = amsc.check(); err != nil {
			return nil, err
		}
		node, err = amsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AllMethodsServiceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = amsc.check(); err != nil {
				return nil, err
			}
			amsc.mutation = mutation
			if node, err = amsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(amsc.hooks) - 1; i >= 0; i-- {
			if amsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = amsc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, amsc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AllMethodsService)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AllMethodsServiceMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (amsc *AllMethodsServiceCreate) SaveX(ctx context.Context) *AllMethodsService {
	v, err := amsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (amsc *AllMethodsServiceCreate) Exec(ctx context.Context) error {
	_, err := amsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (amsc *AllMethodsServiceCreate) ExecX(ctx context.Context) {
	if err := amsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (amsc *AllMethodsServiceCreate) check() error {
	return nil
}

func (amsc *AllMethodsServiceCreate) sqlSave(ctx context.Context) (*AllMethodsService, error) {
	_node, _spec := amsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, amsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (amsc *AllMethodsServiceCreate) createSpec() (*AllMethodsService, *sqlgraph.CreateSpec) {
	var (
		_node = &AllMethodsService{config: amsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: allmethodsservice.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: allmethodsservice.FieldID,
			},
		}
	)
	return _node, _spec
}

// AllMethodsServiceCreateBulk is the builder for creating many AllMethodsService entities in bulk.
type AllMethodsServiceCreateBulk struct {
	config
	builders []*AllMethodsServiceCreate
}

// Save creates the AllMethodsService entities in the database.
func (amscb *AllMethodsServiceCreateBulk) Save(ctx context.Context) ([]*AllMethodsService, error) {
	specs := make([]*sqlgraph.CreateSpec, len(amscb.builders))
	nodes := make([]*AllMethodsService, len(amscb.builders))
	mutators := make([]Mutator, len(amscb.builders))
	for i := range amscb.builders {
		func(i int, root context.Context) {
			builder := amscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AllMethodsServiceMutation)
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
					_, err = mutators[i+1].Mutate(root, amscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, amscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, amscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (amscb *AllMethodsServiceCreateBulk) SaveX(ctx context.Context) []*AllMethodsService {
	v, err := amscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (amscb *AllMethodsServiceCreateBulk) Exec(ctx context.Context) error {
	_, err := amscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (amscb *AllMethodsServiceCreateBulk) ExecX(ctx context.Context) {
	if err := amscb.Exec(ctx); err != nil {
		panic(err)
	}
}
