// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/invalidfieldmessage"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/contrib/entproto/internal/entprototest/ent/schema"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InvalidFieldMessageUpdate is the builder for updating InvalidFieldMessage entities.
type InvalidFieldMessageUpdate struct {
	config
	hooks    []Hook
	mutation *InvalidFieldMessageMutation
}

// Where appends a list predicates to the InvalidFieldMessageUpdate builder.
func (ifmu *InvalidFieldMessageUpdate) Where(ps ...predicate.InvalidFieldMessage) *InvalidFieldMessageUpdate {
	ifmu.mutation.Where(ps...)
	return ifmu
}

// SetJSON sets the "json" field.
func (ifmu *InvalidFieldMessageUpdate) SetJSON(sj *schema.SomeJSON) *InvalidFieldMessageUpdate {
	ifmu.mutation.SetJSON(sj)
	return ifmu
}

// Mutation returns the InvalidFieldMessageMutation object of the builder.
func (ifmu *InvalidFieldMessageUpdate) Mutation() *InvalidFieldMessageMutation {
	return ifmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ifmu *InvalidFieldMessageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ifmu.hooks) == 0 {
		affected, err = ifmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InvalidFieldMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ifmu.mutation = mutation
			affected, err = ifmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ifmu.hooks) - 1; i >= 0; i-- {
			if ifmu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ifmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ifmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ifmu *InvalidFieldMessageUpdate) SaveX(ctx context.Context) int {
	affected, err := ifmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ifmu *InvalidFieldMessageUpdate) Exec(ctx context.Context) error {
	_, err := ifmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ifmu *InvalidFieldMessageUpdate) ExecX(ctx context.Context) {
	if err := ifmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ifmu *InvalidFieldMessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   invalidfieldmessage.Table,
			Columns: invalidfieldmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: invalidfieldmessage.FieldID,
			},
		},
	}
	if ps := ifmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ifmu.mutation.JSON(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: invalidfieldmessage.FieldJSON,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ifmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invalidfieldmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// InvalidFieldMessageUpdateOne is the builder for updating a single InvalidFieldMessage entity.
type InvalidFieldMessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InvalidFieldMessageMutation
}

// SetJSON sets the "json" field.
func (ifmuo *InvalidFieldMessageUpdateOne) SetJSON(sj *schema.SomeJSON) *InvalidFieldMessageUpdateOne {
	ifmuo.mutation.SetJSON(sj)
	return ifmuo
}

// Mutation returns the InvalidFieldMessageMutation object of the builder.
func (ifmuo *InvalidFieldMessageUpdateOne) Mutation() *InvalidFieldMessageMutation {
	return ifmuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ifmuo *InvalidFieldMessageUpdateOne) Select(field string, fields ...string) *InvalidFieldMessageUpdateOne {
	ifmuo.fields = append([]string{field}, fields...)
	return ifmuo
}

// Save executes the query and returns the updated InvalidFieldMessage entity.
func (ifmuo *InvalidFieldMessageUpdateOne) Save(ctx context.Context) (*InvalidFieldMessage, error) {
	var (
		err  error
		node *InvalidFieldMessage
	)
	if len(ifmuo.hooks) == 0 {
		node, err = ifmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InvalidFieldMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ifmuo.mutation = mutation
			node, err = ifmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ifmuo.hooks) - 1; i >= 0; i-- {
			if ifmuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ifmuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ifmuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*InvalidFieldMessage)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from InvalidFieldMessageMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ifmuo *InvalidFieldMessageUpdateOne) SaveX(ctx context.Context) *InvalidFieldMessage {
	node, err := ifmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ifmuo *InvalidFieldMessageUpdateOne) Exec(ctx context.Context) error {
	_, err := ifmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ifmuo *InvalidFieldMessageUpdateOne) ExecX(ctx context.Context) {
	if err := ifmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ifmuo *InvalidFieldMessageUpdateOne) sqlSave(ctx context.Context) (_node *InvalidFieldMessage, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   invalidfieldmessage.Table,
			Columns: invalidfieldmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: invalidfieldmessage.FieldID,
			},
		},
	}
	id, ok := ifmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "InvalidFieldMessage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ifmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invalidfieldmessage.FieldID)
		for _, f := range fields {
			if !invalidfieldmessage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != invalidfieldmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ifmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ifmuo.mutation.JSON(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: invalidfieldmessage.FieldJSON,
		})
	}
	_node = &InvalidFieldMessage{config: ifmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ifmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invalidfieldmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
