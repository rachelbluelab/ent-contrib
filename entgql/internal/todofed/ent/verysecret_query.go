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
	"fmt"
	"math"

	"entgo.io/contrib/entgql/internal/todofed/ent/predicate"
	"entgo.io/contrib/entgql/internal/todofed/ent/verysecret"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VerySecretQuery is the builder for querying VerySecret entities.
type VerySecretQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.VerySecret
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*VerySecret) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VerySecretQuery builder.
func (vsq *VerySecretQuery) Where(ps ...predicate.VerySecret) *VerySecretQuery {
	vsq.predicates = append(vsq.predicates, ps...)
	return vsq
}

// Limit adds a limit step to the query.
func (vsq *VerySecretQuery) Limit(limit int) *VerySecretQuery {
	vsq.limit = &limit
	return vsq
}

// Offset adds an offset step to the query.
func (vsq *VerySecretQuery) Offset(offset int) *VerySecretQuery {
	vsq.offset = &offset
	return vsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vsq *VerySecretQuery) Unique(unique bool) *VerySecretQuery {
	vsq.unique = &unique
	return vsq
}

// Order adds an order step to the query.
func (vsq *VerySecretQuery) Order(o ...OrderFunc) *VerySecretQuery {
	vsq.order = append(vsq.order, o...)
	return vsq
}

// First returns the first VerySecret entity from the query.
// Returns a *NotFoundError when no VerySecret was found.
func (vsq *VerySecretQuery) First(ctx context.Context) (*VerySecret, error) {
	nodes, err := vsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{verysecret.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vsq *VerySecretQuery) FirstX(ctx context.Context) *VerySecret {
	node, err := vsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VerySecret ID from the query.
// Returns a *NotFoundError when no VerySecret ID was found.
func (vsq *VerySecretQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{verysecret.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vsq *VerySecretQuery) FirstIDX(ctx context.Context) int {
	id, err := vsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VerySecret entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VerySecret entity is found.
// Returns a *NotFoundError when no VerySecret entities are found.
func (vsq *VerySecretQuery) Only(ctx context.Context) (*VerySecret, error) {
	nodes, err := vsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{verysecret.Label}
	default:
		return nil, &NotSingularError{verysecret.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vsq *VerySecretQuery) OnlyX(ctx context.Context) *VerySecret {
	node, err := vsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VerySecret ID in the query.
// Returns a *NotSingularError when more than one VerySecret ID is found.
// Returns a *NotFoundError when no entities are found.
func (vsq *VerySecretQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{verysecret.Label}
	default:
		err = &NotSingularError{verysecret.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vsq *VerySecretQuery) OnlyIDX(ctx context.Context) int {
	id, err := vsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VerySecrets.
func (vsq *VerySecretQuery) All(ctx context.Context) ([]*VerySecret, error) {
	if err := vsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return vsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (vsq *VerySecretQuery) AllX(ctx context.Context) []*VerySecret {
	nodes, err := vsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VerySecret IDs.
func (vsq *VerySecretQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := vsq.Select(verysecret.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vsq *VerySecretQuery) IDsX(ctx context.Context) []int {
	ids, err := vsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vsq *VerySecretQuery) Count(ctx context.Context) (int, error) {
	if err := vsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return vsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (vsq *VerySecretQuery) CountX(ctx context.Context) int {
	count, err := vsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vsq *VerySecretQuery) Exist(ctx context.Context) (bool, error) {
	if err := vsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return vsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (vsq *VerySecretQuery) ExistX(ctx context.Context) bool {
	exist, err := vsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VerySecretQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vsq *VerySecretQuery) Clone() *VerySecretQuery {
	if vsq == nil {
		return nil
	}
	return &VerySecretQuery{
		config:     vsq.config,
		limit:      vsq.limit,
		offset:     vsq.offset,
		order:      append([]OrderFunc{}, vsq.order...),
		predicates: append([]predicate.VerySecret{}, vsq.predicates...),
		// clone intermediate query.
		sql:    vsq.sql.Clone(),
		path:   vsq.path,
		unique: vsq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Password string `json:"password,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VerySecret.Query().
//		GroupBy(verysecret.FieldPassword).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (vsq *VerySecretQuery) GroupBy(field string, fields ...string) *VerySecretGroupBy {
	grbuild := &VerySecretGroupBy{config: vsq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := vsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return vsq.sqlQuery(ctx), nil
	}
	grbuild.label = verysecret.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Password string `json:"password,omitempty"`
//	}
//
//	client.VerySecret.Query().
//		Select(verysecret.FieldPassword).
//		Scan(ctx, &v)
//
func (vsq *VerySecretQuery) Select(fields ...string) *VerySecretSelect {
	vsq.fields = append(vsq.fields, fields...)
	selbuild := &VerySecretSelect{VerySecretQuery: vsq}
	selbuild.label = verysecret.Label
	selbuild.flds, selbuild.scan = &vsq.fields, selbuild.Scan
	return selbuild
}

func (vsq *VerySecretQuery) prepareQuery(ctx context.Context) error {
	for _, f := range vsq.fields {
		if !verysecret.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vsq.path != nil {
		prev, err := vsq.path(ctx)
		if err != nil {
			return err
		}
		vsq.sql = prev
	}
	return nil
}

func (vsq *VerySecretQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VerySecret, error) {
	var (
		nodes = []*VerySecret{}
		_spec = vsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*VerySecret).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &VerySecret{config: vsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(vsq.modifiers) > 0 {
		_spec.Modifiers = vsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range vsq.loadTotal {
		if err := vsq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vsq *VerySecretQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vsq.querySpec()
	if len(vsq.modifiers) > 0 {
		_spec.Modifiers = vsq.modifiers
	}
	_spec.Node.Columns = vsq.fields
	if len(vsq.fields) > 0 {
		_spec.Unique = vsq.unique != nil && *vsq.unique
	}
	return sqlgraph.CountNodes(ctx, vsq.driver, _spec)
}

func (vsq *VerySecretQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := vsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (vsq *VerySecretQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   verysecret.Table,
			Columns: verysecret.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: verysecret.FieldID,
			},
		},
		From:   vsq.sql,
		Unique: true,
	}
	if unique := vsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := vsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, verysecret.FieldID)
		for i := range fields {
			if fields[i] != verysecret.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vsq *VerySecretQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vsq.driver.Dialect())
	t1 := builder.Table(verysecret.Table)
	columns := vsq.fields
	if len(columns) == 0 {
		columns = verysecret.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vsq.sql != nil {
		selector = vsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vsq.unique != nil && *vsq.unique {
		selector.Distinct()
	}
	for _, p := range vsq.predicates {
		p(selector)
	}
	for _, p := range vsq.order {
		p(selector)
	}
	if offset := vsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VerySecretGroupBy is the group-by builder for VerySecret entities.
type VerySecretGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vsgb *VerySecretGroupBy) Aggregate(fns ...AggregateFunc) *VerySecretGroupBy {
	vsgb.fns = append(vsgb.fns, fns...)
	return vsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (vsgb *VerySecretGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := vsgb.path(ctx)
	if err != nil {
		return err
	}
	vsgb.sql = query
	return vsgb.sqlScan(ctx, v)
}

func (vsgb *VerySecretGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range vsgb.fields {
		if !verysecret.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := vsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vsgb *VerySecretGroupBy) sqlQuery() *sql.Selector {
	selector := vsgb.sql.Select()
	aggregation := make([]string, 0, len(vsgb.fns))
	for _, fn := range vsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(vsgb.fields)+len(vsgb.fns))
		for _, f := range vsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(vsgb.fields...)...)
}

// VerySecretSelect is the builder for selecting fields of VerySecret entities.
type VerySecretSelect struct {
	*VerySecretQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (vss *VerySecretSelect) Scan(ctx context.Context, v interface{}) error {
	if err := vss.prepareQuery(ctx); err != nil {
		return err
	}
	vss.sql = vss.VerySecretQuery.sqlQuery(ctx)
	return vss.sqlScan(ctx, v)
}

func (vss *VerySecretSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := vss.sql.Query()
	if err := vss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
