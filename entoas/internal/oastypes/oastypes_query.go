// Code generated by ent, DO NOT EDIT.

package oastypes

import (
	"context"
	"fmt"
	"math"

	"entgo.io/contrib/entoas/internal/oastypes/oastypes"
	"entgo.io/contrib/entoas/internal/oastypes/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OASTypesQuery is the builder for querying OASTypes entities.
type OASTypesQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OASTypes
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OASTypesQuery builder.
func (otq *OASTypesQuery) Where(ps ...predicate.OASTypes) *OASTypesQuery {
	otq.predicates = append(otq.predicates, ps...)
	return otq
}

// Limit adds a limit step to the query.
func (otq *OASTypesQuery) Limit(limit int) *OASTypesQuery {
	otq.limit = &limit
	return otq
}

// Offset adds an offset step to the query.
func (otq *OASTypesQuery) Offset(offset int) *OASTypesQuery {
	otq.offset = &offset
	return otq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (otq *OASTypesQuery) Unique(unique bool) *OASTypesQuery {
	otq.unique = &unique
	return otq
}

// Order adds an order step to the query.
func (otq *OASTypesQuery) Order(o ...OrderFunc) *OASTypesQuery {
	otq.order = append(otq.order, o...)
	return otq
}

// First returns the first OASTypes entity from the query.
// Returns a *NotFoundError when no OASTypes was found.
func (otq *OASTypesQuery) First(ctx context.Context) (*OASTypes, error) {
	nodes, err := otq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{oastypes.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (otq *OASTypesQuery) FirstX(ctx context.Context) *OASTypes {
	node, err := otq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OASTypes ID from the query.
// Returns a *NotFoundError when no OASTypes ID was found.
func (otq *OASTypesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = otq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{oastypes.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (otq *OASTypesQuery) FirstIDX(ctx context.Context) int {
	id, err := otq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OASTypes entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OASTypes entity is found.
// Returns a *NotFoundError when no OASTypes entities are found.
func (otq *OASTypesQuery) Only(ctx context.Context) (*OASTypes, error) {
	nodes, err := otq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{oastypes.Label}
	default:
		return nil, &NotSingularError{oastypes.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (otq *OASTypesQuery) OnlyX(ctx context.Context) *OASTypes {
	node, err := otq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OASTypes ID in the query.
// Returns a *NotSingularError when more than one OASTypes ID is found.
// Returns a *NotFoundError when no entities are found.
func (otq *OASTypesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = otq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{oastypes.Label}
	default:
		err = &NotSingularError{oastypes.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (otq *OASTypesQuery) OnlyIDX(ctx context.Context) int {
	id, err := otq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OASTypesSlice.
func (otq *OASTypesQuery) All(ctx context.Context) ([]*OASTypes, error) {
	if err := otq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return otq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (otq *OASTypesQuery) AllX(ctx context.Context) []*OASTypes {
	nodes, err := otq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OASTypes IDs.
func (otq *OASTypesQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := otq.Select(oastypes.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (otq *OASTypesQuery) IDsX(ctx context.Context) []int {
	ids, err := otq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (otq *OASTypesQuery) Count(ctx context.Context) (int, error) {
	if err := otq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return otq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (otq *OASTypesQuery) CountX(ctx context.Context) int {
	count, err := otq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (otq *OASTypesQuery) Exist(ctx context.Context) (bool, error) {
	if err := otq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return otq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (otq *OASTypesQuery) ExistX(ctx context.Context) bool {
	exist, err := otq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OASTypesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (otq *OASTypesQuery) Clone() *OASTypesQuery {
	if otq == nil {
		return nil
	}
	return &OASTypesQuery{
		config:     otq.config,
		limit:      otq.limit,
		offset:     otq.offset,
		order:      append([]OrderFunc{}, otq.order...),
		predicates: append([]predicate.OASTypes{}, otq.predicates...),
		// clone intermediate query.
		sql:    otq.sql.Clone(),
		path:   otq.path,
		unique: otq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Int int `json:"int,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OASTypes.Query().
//		GroupBy(oastypes.FieldInt).
//		Aggregate(oastypes.Count()).
//		Scan(ctx, &v)
//
func (otq *OASTypesQuery) GroupBy(field string, fields ...string) *OASTypesGroupBy {
	grbuild := &OASTypesGroupBy{config: otq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := otq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return otq.sqlQuery(ctx), nil
	}
	grbuild.label = oastypes.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Int int `json:"int,omitempty"`
//	}
//
//	client.OASTypes.Query().
//		Select(oastypes.FieldInt).
//		Scan(ctx, &v)
//
func (otq *OASTypesQuery) Select(fields ...string) *OASTypesSelect {
	otq.fields = append(otq.fields, fields...)
	selbuild := &OASTypesSelect{OASTypesQuery: otq}
	selbuild.label = oastypes.Label
	selbuild.flds, selbuild.scan = &otq.fields, selbuild.Scan
	return selbuild
}

func (otq *OASTypesQuery) prepareQuery(ctx context.Context) error {
	for _, f := range otq.fields {
		if !oastypes.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("oastypes: invalid field %q for query", f)}
		}
	}
	if otq.path != nil {
		prev, err := otq.path(ctx)
		if err != nil {
			return err
		}
		otq.sql = prev
	}
	return nil
}

func (otq *OASTypesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OASTypes, error) {
	var (
		nodes = []*OASTypes{}
		_spec = otq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OASTypes).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OASTypes{config: otq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, otq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (otq *OASTypesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := otq.querySpec()
	_spec.Node.Columns = otq.fields
	if len(otq.fields) > 0 {
		_spec.Unique = otq.unique != nil && *otq.unique
	}
	return sqlgraph.CountNodes(ctx, otq.driver, _spec)
}

func (otq *OASTypesQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := otq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("oastypes: check existence: %w", err)
	}
	return n > 0, nil
}

func (otq *OASTypesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   oastypes.Table,
			Columns: oastypes.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: oastypes.FieldID,
			},
		},
		From:   otq.sql,
		Unique: true,
	}
	if unique := otq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := otq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oastypes.FieldID)
		for i := range fields {
			if fields[i] != oastypes.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := otq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := otq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := otq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := otq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (otq *OASTypesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(otq.driver.Dialect())
	t1 := builder.Table(oastypes.Table)
	columns := otq.fields
	if len(columns) == 0 {
		columns = oastypes.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if otq.sql != nil {
		selector = otq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if otq.unique != nil && *otq.unique {
		selector.Distinct()
	}
	for _, p := range otq.predicates {
		p(selector)
	}
	for _, p := range otq.order {
		p(selector)
	}
	if offset := otq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := otq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OASTypesGroupBy is the group-by builder for OASTypes entities.
type OASTypesGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (otgb *OASTypesGroupBy) Aggregate(fns ...AggregateFunc) *OASTypesGroupBy {
	otgb.fns = append(otgb.fns, fns...)
	return otgb
}

// Scan applies the group-by query and scans the result into the given value.
func (otgb *OASTypesGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := otgb.path(ctx)
	if err != nil {
		return err
	}
	otgb.sql = query
	return otgb.sqlScan(ctx, v)
}

func (otgb *OASTypesGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range otgb.fields {
		if !oastypes.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := otgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := otgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (otgb *OASTypesGroupBy) sqlQuery() *sql.Selector {
	selector := otgb.sql.Select()
	aggregation := make([]string, 0, len(otgb.fns))
	for _, fn := range otgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(otgb.fields)+len(otgb.fns))
		for _, f := range otgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(otgb.fields...)...)
}

// OASTypesSelect is the builder for selecting fields of OASTypes entities.
type OASTypesSelect struct {
	*OASTypesQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ots *OASTypesSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ots.prepareQuery(ctx); err != nil {
		return err
	}
	ots.sql = ots.OASTypesQuery.sqlQuery(ctx)
	return ots.sqlScan(ctx, v)
}

func (ots *OASTypesSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ots.sql.Query()
	if err := ots.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
