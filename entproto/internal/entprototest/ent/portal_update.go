// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/category"
	"entgo.io/contrib/entproto/internal/entprototest/ent/portal"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PortalUpdate is the builder for updating Portal entities.
type PortalUpdate struct {
	config
	hooks    []Hook
	mutation *PortalMutation
}

// Where appends a list predicates to the PortalUpdate builder.
func (pu *PortalUpdate) Where(ps ...predicate.Portal) *PortalUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PortalUpdate) SetName(s string) *PortalUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetDescription sets the "description" field.
func (pu *PortalUpdate) SetDescription(s string) *PortalUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (pu *PortalUpdate) SetCategoryID(id int) *PortalUpdate {
	pu.mutation.SetCategoryID(id)
	return pu
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (pu *PortalUpdate) SetNillableCategoryID(id *int) *PortalUpdate {
	if id != nil {
		pu = pu.SetCategoryID(*id)
	}
	return pu
}

// SetCategory sets the "category" edge to the Category entity.
func (pu *PortalUpdate) SetCategory(c *Category) *PortalUpdate {
	return pu.SetCategoryID(c.ID)
}

// Mutation returns the PortalMutation object of the builder.
func (pu *PortalUpdate) Mutation() *PortalMutation {
	return pu.mutation
}

// ClearCategory clears the "category" edge to the Category entity.
func (pu *PortalUpdate) ClearCategory() *PortalUpdate {
	pu.mutation.ClearCategory()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PortalUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PortalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PortalUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PortalUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PortalUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PortalUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   portal.Table,
			Columns: portal.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: portal.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: portal.FieldName,
		})
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: portal.FieldDescription,
		})
	}
	if pu.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   portal.CategoryTable,
			Columns: []string{portal.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   portal.CategoryTable,
			Columns: []string{portal.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{portal.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PortalUpdateOne is the builder for updating a single Portal entity.
type PortalUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PortalMutation
}

// SetName sets the "name" field.
func (puo *PortalUpdateOne) SetName(s string) *PortalUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetDescription sets the "description" field.
func (puo *PortalUpdateOne) SetDescription(s string) *PortalUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (puo *PortalUpdateOne) SetCategoryID(id int) *PortalUpdateOne {
	puo.mutation.SetCategoryID(id)
	return puo
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (puo *PortalUpdateOne) SetNillableCategoryID(id *int) *PortalUpdateOne {
	if id != nil {
		puo = puo.SetCategoryID(*id)
	}
	return puo
}

// SetCategory sets the "category" edge to the Category entity.
func (puo *PortalUpdateOne) SetCategory(c *Category) *PortalUpdateOne {
	return puo.SetCategoryID(c.ID)
}

// Mutation returns the PortalMutation object of the builder.
func (puo *PortalUpdateOne) Mutation() *PortalMutation {
	return puo.mutation
}

// ClearCategory clears the "category" edge to the Category entity.
func (puo *PortalUpdateOne) ClearCategory() *PortalUpdateOne {
	puo.mutation.ClearCategory()
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PortalUpdateOne) Select(field string, fields ...string) *PortalUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Portal entity.
func (puo *PortalUpdateOne) Save(ctx context.Context) (*Portal, error) {
	var (
		err  error
		node *Portal
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PortalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Portal)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PortalMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PortalUpdateOne) SaveX(ctx context.Context) *Portal {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PortalUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PortalUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PortalUpdateOne) sqlSave(ctx context.Context) (_node *Portal, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   portal.Table,
			Columns: portal.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: portal.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Portal.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, portal.FieldID)
		for _, f := range fields {
			if !portal.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != portal.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: portal.FieldName,
		})
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: portal.FieldDescription,
		})
	}
	if puo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   portal.CategoryTable,
			Columns: []string{portal.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   portal.CategoryTable,
			Columns: []string{portal.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Portal{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{portal.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
