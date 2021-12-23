// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/entgql/internal/todopulid/ent/category"
	"entgo.io/contrib/entgql/internal/todopulid/ent/predicate"
	"entgo.io/contrib/entgql/internal/todopulid/ent/schema/pulid"
	"entgo.io/contrib/entgql/internal/todopulid/ent/todo"
	"entgo.io/contrib/entgql/internal/todopulid/ent/verysecret"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TodoUpdate is the builder for updating Todo entities.
type TodoUpdate struct {
	config
	hooks    []Hook
	mutation *TodoMutation
}

// Where appends a list predicates to the TodoUpdate builder.
func (tu *TodoUpdate) Where(ps ...predicate.Todo) *TodoUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetStatus sets the "status" field.
func (tu *TodoUpdate) SetStatus(t todo.Status) *TodoUpdate {
	tu.mutation.SetStatus(t)
	return tu
}

// SetPriority sets the "priority" field.
func (tu *TodoUpdate) SetPriority(i int) *TodoUpdate {
	tu.mutation.ResetPriority()
	tu.mutation.SetPriority(i)
	return tu
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tu *TodoUpdate) SetNillablePriority(i *int) *TodoUpdate {
	if i != nil {
		tu.SetPriority(*i)
	}
	return tu
}

// AddPriority adds i to the "priority" field.
func (tu *TodoUpdate) AddPriority(i int) *TodoUpdate {
	tu.mutation.AddPriority(i)
	return tu
}

// SetText sets the "text" field.
func (tu *TodoUpdate) SetText(s string) *TodoUpdate {
	tu.mutation.SetText(s)
	return tu
}

// SetBlob sets the "blob" field.
func (tu *TodoUpdate) SetBlob(b []byte) *TodoUpdate {
	tu.mutation.SetBlob(b)
	return tu
}

// ClearBlob clears the value of the "blob" field.
func (tu *TodoUpdate) ClearBlob() *TodoUpdate {
	tu.mutation.ClearBlob()
	return tu
}

// SetParentID sets the "parent" edge to the Todo entity by ID.
func (tu *TodoUpdate) SetParentID(id pulid.ID) *TodoUpdate {
	tu.mutation.SetParentID(id)
	return tu
}

// SetNillableParentID sets the "parent" edge to the Todo entity by ID if the given value is not nil.
func (tu *TodoUpdate) SetNillableParentID(id *pulid.ID) *TodoUpdate {
	if id != nil {
		tu = tu.SetParentID(*id)
	}
	return tu
}

// SetParent sets the "parent" edge to the Todo entity.
func (tu *TodoUpdate) SetParent(t *Todo) *TodoUpdate {
	return tu.SetParentID(t.ID)
}

// AddChildIDs adds the "children" edge to the Todo entity by IDs.
func (tu *TodoUpdate) AddChildIDs(ids ...pulid.ID) *TodoUpdate {
	tu.mutation.AddChildIDs(ids...)
	return tu
}

// AddChildren adds the "children" edges to the Todo entity.
func (tu *TodoUpdate) AddChildren(t ...*Todo) *TodoUpdate {
	ids := make([]pulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddChildIDs(ids...)
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (tu *TodoUpdate) SetCategoryID(id pulid.ID) *TodoUpdate {
	tu.mutation.SetCategoryID(id)
	return tu
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (tu *TodoUpdate) SetNillableCategoryID(id *pulid.ID) *TodoUpdate {
	if id != nil {
		tu = tu.SetCategoryID(*id)
	}
	return tu
}

// SetCategory sets the "category" edge to the Category entity.
func (tu *TodoUpdate) SetCategory(c *Category) *TodoUpdate {
	return tu.SetCategoryID(c.ID)
}

// SetSecretID sets the "secret" edge to the VerySecret entity by ID.
func (tu *TodoUpdate) SetSecretID(id pulid.ID) *TodoUpdate {
	tu.mutation.SetSecretID(id)
	return tu
}

// SetNillableSecretID sets the "secret" edge to the VerySecret entity by ID if the given value is not nil.
func (tu *TodoUpdate) SetNillableSecretID(id *pulid.ID) *TodoUpdate {
	if id != nil {
		tu = tu.SetSecretID(*id)
	}
	return tu
}

// SetSecret sets the "secret" edge to the VerySecret entity.
func (tu *TodoUpdate) SetSecret(v *VerySecret) *TodoUpdate {
	return tu.SetSecretID(v.ID)
}

// Mutation returns the TodoMutation object of the builder.
func (tu *TodoUpdate) Mutation() *TodoMutation {
	return tu.mutation
}

// ClearParent clears the "parent" edge to the Todo entity.
func (tu *TodoUpdate) ClearParent() *TodoUpdate {
	tu.mutation.ClearParent()
	return tu
}

// ClearChildren clears all "children" edges to the Todo entity.
func (tu *TodoUpdate) ClearChildren() *TodoUpdate {
	tu.mutation.ClearChildren()
	return tu
}

// RemoveChildIDs removes the "children" edge to Todo entities by IDs.
func (tu *TodoUpdate) RemoveChildIDs(ids ...pulid.ID) *TodoUpdate {
	tu.mutation.RemoveChildIDs(ids...)
	return tu
}

// RemoveChildren removes "children" edges to Todo entities.
func (tu *TodoUpdate) RemoveChildren(t ...*Todo) *TodoUpdate {
	ids := make([]pulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveChildIDs(ids...)
}

// ClearCategory clears the "category" edge to the Category entity.
func (tu *TodoUpdate) ClearCategory() *TodoUpdate {
	tu.mutation.ClearCategory()
	return tu
}

// ClearSecret clears the "secret" edge to the VerySecret entity.
func (tu *TodoUpdate) ClearSecret() *TodoUpdate {
	tu.mutation.ClearSecret()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TodoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TodoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TodoUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TodoUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TodoUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TodoUpdate) check() error {
	if v, ok := tu.mutation.Status(); ok {
		if err := todo.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Todo.status": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Text(); ok {
		if err := todo.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Todo.text": %w`, err)}
		}
	}
	return nil
}

func (tu *TodoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   todo.Table,
			Columns: todo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: todo.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: todo.FieldStatus,
		})
	}
	if value, ok := tu.mutation.Priority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: todo.FieldPriority,
		})
	}
	if value, ok := tu.mutation.AddedPriority(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: todo.FieldPriority,
		})
	}
	if value, ok := tu.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: todo.FieldText,
		})
	}
	if value, ok := tu.mutation.Blob(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: todo.FieldBlob,
		})
	}
	if tu.mutation.BlobCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: todo.FieldBlob,
		})
	}
	if tu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.ParentTable,
			Columns: []string{todo.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: todo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.ParentTable,
			Columns: []string{todo.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: todo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !tu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.CategoryTable,
			Columns: []string{todo.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.CategoryTable,
			Columns: []string{todo.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.SecretCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   todo.SecretTable,
			Columns: []string{todo.SecretColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: verysecret.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.SecretIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   todo.SecretTable,
			Columns: []string{todo.SecretColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: verysecret.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TodoUpdateOne is the builder for updating a single Todo entity.
type TodoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TodoMutation
}

// SetStatus sets the "status" field.
func (tuo *TodoUpdateOne) SetStatus(t todo.Status) *TodoUpdateOne {
	tuo.mutation.SetStatus(t)
	return tuo
}

// SetPriority sets the "priority" field.
func (tuo *TodoUpdateOne) SetPriority(i int) *TodoUpdateOne {
	tuo.mutation.ResetPriority()
	tuo.mutation.SetPriority(i)
	return tuo
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillablePriority(i *int) *TodoUpdateOne {
	if i != nil {
		tuo.SetPriority(*i)
	}
	return tuo
}

// AddPriority adds i to the "priority" field.
func (tuo *TodoUpdateOne) AddPriority(i int) *TodoUpdateOne {
	tuo.mutation.AddPriority(i)
	return tuo
}

// SetText sets the "text" field.
func (tuo *TodoUpdateOne) SetText(s string) *TodoUpdateOne {
	tuo.mutation.SetText(s)
	return tuo
}

// SetBlob sets the "blob" field.
func (tuo *TodoUpdateOne) SetBlob(b []byte) *TodoUpdateOne {
	tuo.mutation.SetBlob(b)
	return tuo
}

// ClearBlob clears the value of the "blob" field.
func (tuo *TodoUpdateOne) ClearBlob() *TodoUpdateOne {
	tuo.mutation.ClearBlob()
	return tuo
}

// SetParentID sets the "parent" edge to the Todo entity by ID.
func (tuo *TodoUpdateOne) SetParentID(id pulid.ID) *TodoUpdateOne {
	tuo.mutation.SetParentID(id)
	return tuo
}

// SetNillableParentID sets the "parent" edge to the Todo entity by ID if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableParentID(id *pulid.ID) *TodoUpdateOne {
	if id != nil {
		tuo = tuo.SetParentID(*id)
	}
	return tuo
}

// SetParent sets the "parent" edge to the Todo entity.
func (tuo *TodoUpdateOne) SetParent(t *Todo) *TodoUpdateOne {
	return tuo.SetParentID(t.ID)
}

// AddChildIDs adds the "children" edge to the Todo entity by IDs.
func (tuo *TodoUpdateOne) AddChildIDs(ids ...pulid.ID) *TodoUpdateOne {
	tuo.mutation.AddChildIDs(ids...)
	return tuo
}

// AddChildren adds the "children" edges to the Todo entity.
func (tuo *TodoUpdateOne) AddChildren(t ...*Todo) *TodoUpdateOne {
	ids := make([]pulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddChildIDs(ids...)
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (tuo *TodoUpdateOne) SetCategoryID(id pulid.ID) *TodoUpdateOne {
	tuo.mutation.SetCategoryID(id)
	return tuo
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableCategoryID(id *pulid.ID) *TodoUpdateOne {
	if id != nil {
		tuo = tuo.SetCategoryID(*id)
	}
	return tuo
}

// SetCategory sets the "category" edge to the Category entity.
func (tuo *TodoUpdateOne) SetCategory(c *Category) *TodoUpdateOne {
	return tuo.SetCategoryID(c.ID)
}

// SetSecretID sets the "secret" edge to the VerySecret entity by ID.
func (tuo *TodoUpdateOne) SetSecretID(id pulid.ID) *TodoUpdateOne {
	tuo.mutation.SetSecretID(id)
	return tuo
}

// SetNillableSecretID sets the "secret" edge to the VerySecret entity by ID if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableSecretID(id *pulid.ID) *TodoUpdateOne {
	if id != nil {
		tuo = tuo.SetSecretID(*id)
	}
	return tuo
}

// SetSecret sets the "secret" edge to the VerySecret entity.
func (tuo *TodoUpdateOne) SetSecret(v *VerySecret) *TodoUpdateOne {
	return tuo.SetSecretID(v.ID)
}

// Mutation returns the TodoMutation object of the builder.
func (tuo *TodoUpdateOne) Mutation() *TodoMutation {
	return tuo.mutation
}

// ClearParent clears the "parent" edge to the Todo entity.
func (tuo *TodoUpdateOne) ClearParent() *TodoUpdateOne {
	tuo.mutation.ClearParent()
	return tuo
}

// ClearChildren clears all "children" edges to the Todo entity.
func (tuo *TodoUpdateOne) ClearChildren() *TodoUpdateOne {
	tuo.mutation.ClearChildren()
	return tuo
}

// RemoveChildIDs removes the "children" edge to Todo entities by IDs.
func (tuo *TodoUpdateOne) RemoveChildIDs(ids ...pulid.ID) *TodoUpdateOne {
	tuo.mutation.RemoveChildIDs(ids...)
	return tuo
}

// RemoveChildren removes "children" edges to Todo entities.
func (tuo *TodoUpdateOne) RemoveChildren(t ...*Todo) *TodoUpdateOne {
	ids := make([]pulid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveChildIDs(ids...)
}

// ClearCategory clears the "category" edge to the Category entity.
func (tuo *TodoUpdateOne) ClearCategory() *TodoUpdateOne {
	tuo.mutation.ClearCategory()
	return tuo
}

// ClearSecret clears the "secret" edge to the VerySecret entity.
func (tuo *TodoUpdateOne) ClearSecret() *TodoUpdateOne {
	tuo.mutation.ClearSecret()
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TodoUpdateOne) Select(field string, fields ...string) *TodoUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Todo entity.
func (tuo *TodoUpdateOne) Save(ctx context.Context) (*Todo, error) {
	var (
		err  error
		node *Todo
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TodoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TodoUpdateOne) SaveX(ctx context.Context) *Todo {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TodoUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TodoUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TodoUpdateOne) check() error {
	if v, ok := tuo.mutation.Status(); ok {
		if err := todo.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Todo.status": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Text(); ok {
		if err := todo.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Todo.text": %w`, err)}
		}
	}
	return nil
}

func (tuo *TodoUpdateOne) sqlSave(ctx context.Context) (_node *Todo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   todo.Table,
			Columns: todo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: todo.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Todo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todo.FieldID)
		for _, f := range fields {
			if !todo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != todo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: todo.FieldStatus,
		})
	}
	if value, ok := tuo.mutation.Priority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: todo.FieldPriority,
		})
	}
	if value, ok := tuo.mutation.AddedPriority(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: todo.FieldPriority,
		})
	}
	if value, ok := tuo.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: todo.FieldText,
		})
	}
	if value, ok := tuo.mutation.Blob(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: todo.FieldBlob,
		})
	}
	if tuo.mutation.BlobCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: todo.FieldBlob,
		})
	}
	if tuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.ParentTable,
			Columns: []string{todo.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: todo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.ParentTable,
			Columns: []string{todo.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: todo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !tuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.CategoryTable,
			Columns: []string{todo.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.CategoryTable,
			Columns: []string{todo.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.SecretCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   todo.SecretTable,
			Columns: []string{todo.SecretColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: verysecret.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.SecretIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   todo.SecretTable,
			Columns: []string{todo.SecretColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: verysecret.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Todo{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
