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

	"entgo.io/contrib/entgql/internal/todouuid/ent/friendship"
	"entgo.io/contrib/entgql/internal/todouuid/ent/predicate"
	"entgo.io/contrib/entgql/internal/todouuid/ent/user"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FriendshipQuery is the builder for querying Friendship entities.
type FriendshipQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Friendship
	withUser   *UserQuery
	withFriend *UserQuery
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*Friendship) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FriendshipQuery builder.
func (fq *FriendshipQuery) Where(ps ...predicate.Friendship) *FriendshipQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit adds a limit step to the query.
func (fq *FriendshipQuery) Limit(limit int) *FriendshipQuery {
	fq.limit = &limit
	return fq
}

// Offset adds an offset step to the query.
func (fq *FriendshipQuery) Offset(offset int) *FriendshipQuery {
	fq.offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FriendshipQuery) Unique(unique bool) *FriendshipQuery {
	fq.unique = &unique
	return fq
}

// Order adds an order step to the query.
func (fq *FriendshipQuery) Order(o ...OrderFunc) *FriendshipQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryUser chains the current query on the "user" edge.
func (fq *FriendshipQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(friendship.Table, friendship.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, friendship.UserTable, friendship.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFriend chains the current query on the "friend" edge.
func (fq *FriendshipQuery) QueryFriend() *UserQuery {
	query := &UserQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(friendship.Table, friendship.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, friendship.FriendTable, friendship.FriendColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Friendship entity from the query.
// Returns a *NotFoundError when no Friendship was found.
func (fq *FriendshipQuery) First(ctx context.Context) (*Friendship, error) {
	nodes, err := fq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{friendship.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FriendshipQuery) FirstX(ctx context.Context) *Friendship {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Friendship ID from the query.
// Returns a *NotFoundError when no Friendship ID was found.
func (fq *FriendshipQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = fq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{friendship.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FriendshipQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Friendship entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Friendship entity is found.
// Returns a *NotFoundError when no Friendship entities are found.
func (fq *FriendshipQuery) Only(ctx context.Context) (*Friendship, error) {
	nodes, err := fq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{friendship.Label}
	default:
		return nil, &NotSingularError{friendship.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FriendshipQuery) OnlyX(ctx context.Context) *Friendship {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Friendship ID in the query.
// Returns a *NotSingularError when more than one Friendship ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FriendshipQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = fq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{friendship.Label}
	default:
		err = &NotSingularError{friendship.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FriendshipQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Friendships.
func (fq *FriendshipQuery) All(ctx context.Context) ([]*Friendship, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fq *FriendshipQuery) AllX(ctx context.Context) []*Friendship {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Friendship IDs.
func (fq *FriendshipQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := fq.Select(friendship.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FriendshipQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FriendshipQuery) Count(ctx context.Context) (int, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FriendshipQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FriendshipQuery) Exist(ctx context.Context) (bool, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FriendshipQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FriendshipQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FriendshipQuery) Clone() *FriendshipQuery {
	if fq == nil {
		return nil
	}
	return &FriendshipQuery{
		config:     fq.config,
		limit:      fq.limit,
		offset:     fq.offset,
		order:      append([]OrderFunc{}, fq.order...),
		predicates: append([]predicate.Friendship{}, fq.predicates...),
		withUser:   fq.withUser.Clone(),
		withFriend: fq.withFriend.Clone(),
		// clone intermediate query.
		sql:    fq.sql.Clone(),
		path:   fq.path,
		unique: fq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FriendshipQuery) WithUser(opts ...func(*UserQuery)) *FriendshipQuery {
	query := &UserQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withUser = query
	return fq
}

// WithFriend tells the query-builder to eager-load the nodes that are connected to
// the "friend" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FriendshipQuery) WithFriend(opts ...func(*UserQuery)) *FriendshipQuery {
	query := &UserQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withFriend = query
	return fq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Friendship.Query().
//		GroupBy(friendship.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fq *FriendshipQuery) GroupBy(field string, fields ...string) *FriendshipGroupBy {
	grbuild := &FriendshipGroupBy{config: fq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fq.sqlQuery(ctx), nil
	}
	grbuild.label = friendship.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Friendship.Query().
//		Select(friendship.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (fq *FriendshipQuery) Select(fields ...string) *FriendshipSelect {
	fq.fields = append(fq.fields, fields...)
	selbuild := &FriendshipSelect{FriendshipQuery: fq}
	selbuild.label = friendship.Label
	selbuild.flds, selbuild.scan = &fq.fields, selbuild.Scan
	return selbuild
}

func (fq *FriendshipQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fq.fields {
		if !friendship.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FriendshipQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Friendship, error) {
	var (
		nodes       = []*Friendship{}
		_spec       = fq.querySpec()
		loadedTypes = [2]bool{
			fq.withUser != nil,
			fq.withFriend != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Friendship).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Friendship{config: fq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(fq.modifiers) > 0 {
		_spec.Modifiers = fq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fq.withUser; query != nil {
		if err := fq.loadUser(ctx, query, nodes, nil,
			func(n *Friendship, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := fq.withFriend; query != nil {
		if err := fq.loadFriend(ctx, query, nodes, nil,
			func(n *Friendship, e *User) { n.Edges.Friend = e }); err != nil {
			return nil, err
		}
	}
	for i := range fq.loadTotal {
		if err := fq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *FriendshipQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Friendship, init func(*Friendship), assign func(*Friendship, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Friendship)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (fq *FriendshipQuery) loadFriend(ctx context.Context, query *UserQuery, nodes []*Friendship, init func(*Friendship), assign func(*Friendship, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Friendship)
	for i := range nodes {
		fk := nodes[i].FriendID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "friend_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (fq *FriendshipQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	if len(fq.modifiers) > 0 {
		_spec.Modifiers = fq.modifiers
	}
	_spec.Node.Columns = fq.fields
	if len(fq.fields) > 0 {
		_spec.Unique = fq.unique != nil && *fq.unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FriendshipQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fq *FriendshipQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   friendship.Table,
			Columns: friendship.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: friendship.FieldID,
			},
		},
		From:   fq.sql,
		Unique: true,
	}
	if unique := fq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, friendship.FieldID)
		for i := range fields {
			if fields[i] != friendship.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FriendshipQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(friendship.Table)
	columns := fq.fields
	if len(columns) == 0 {
		columns = friendship.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.unique != nil && *fq.unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FriendshipGroupBy is the group-by builder for Friendship entities.
type FriendshipGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FriendshipGroupBy) Aggregate(fns ...AggregateFunc) *FriendshipGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fgb *FriendshipGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fgb.path(ctx)
	if err != nil {
		return err
	}
	fgb.sql = query
	return fgb.sqlScan(ctx, v)
}

func (fgb *FriendshipGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fgb.fields {
		if !friendship.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fgb *FriendshipGroupBy) sqlQuery() *sql.Selector {
	selector := fgb.sql.Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fgb.fields)+len(fgb.fns))
		for _, f := range fgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fgb.fields...)...)
}

// FriendshipSelect is the builder for selecting fields of Friendship entities.
type FriendshipSelect struct {
	*FriendshipQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FriendshipSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	fs.sql = fs.FriendshipQuery.sqlQuery(ctx)
	return fs.sqlScan(ctx, v)
}

func (fs *FriendshipSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fs.sql.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
