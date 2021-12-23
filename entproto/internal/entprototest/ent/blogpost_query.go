// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/contrib/entproto/internal/entprototest/ent/blogpost"
	"entgo.io/contrib/entproto/internal/entprototest/ent/category"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/contrib/entproto/internal/entprototest/ent/user"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlogPostQuery is the builder for querying BlogPost entities.
type BlogPostQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BlogPost
	// eager-loading edges.
	withAuthor     *UserQuery
	withCategories *CategoryQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BlogPostQuery builder.
func (bpq *BlogPostQuery) Where(ps ...predicate.BlogPost) *BlogPostQuery {
	bpq.predicates = append(bpq.predicates, ps...)
	return bpq
}

// Limit adds a limit step to the query.
func (bpq *BlogPostQuery) Limit(limit int) *BlogPostQuery {
	bpq.limit = &limit
	return bpq
}

// Offset adds an offset step to the query.
func (bpq *BlogPostQuery) Offset(offset int) *BlogPostQuery {
	bpq.offset = &offset
	return bpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bpq *BlogPostQuery) Unique(unique bool) *BlogPostQuery {
	bpq.unique = &unique
	return bpq
}

// Order adds an order step to the query.
func (bpq *BlogPostQuery) Order(o ...OrderFunc) *BlogPostQuery {
	bpq.order = append(bpq.order, o...)
	return bpq
}

// QueryAuthor chains the current query on the "author" edge.
func (bpq *BlogPostQuery) QueryAuthor() *UserQuery {
	query := &UserQuery{config: bpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(blogpost.Table, blogpost.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, blogpost.AuthorTable, blogpost.AuthorColumn),
		)
		fromU = sqlgraph.SetNeighbors(bpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCategories chains the current query on the "categories" edge.
func (bpq *BlogPostQuery) QueryCategories() *CategoryQuery {
	query := &CategoryQuery{config: bpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(blogpost.Table, blogpost.FieldID, selector),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, blogpost.CategoriesTable, blogpost.CategoriesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(bpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BlogPost entity from the query.
// Returns a *NotFoundError when no BlogPost was found.
func (bpq *BlogPostQuery) First(ctx context.Context) (*BlogPost, error) {
	nodes, err := bpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{blogpost.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bpq *BlogPostQuery) FirstX(ctx context.Context) *BlogPost {
	node, err := bpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BlogPost ID from the query.
// Returns a *NotFoundError when no BlogPost ID was found.
func (bpq *BlogPostQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{blogpost.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bpq *BlogPostQuery) FirstIDX(ctx context.Context) int {
	id, err := bpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BlogPost entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one BlogPost entity is not found.
// Returns a *NotFoundError when no BlogPost entities are found.
func (bpq *BlogPostQuery) Only(ctx context.Context) (*BlogPost, error) {
	nodes, err := bpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{blogpost.Label}
	default:
		return nil, &NotSingularError{blogpost.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bpq *BlogPostQuery) OnlyX(ctx context.Context) *BlogPost {
	node, err := bpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BlogPost ID in the query.
// Returns a *NotSingularError when exactly one BlogPost ID is not found.
// Returns a *NotFoundError when no entities are found.
func (bpq *BlogPostQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = &NotSingularError{blogpost.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bpq *BlogPostQuery) OnlyIDX(ctx context.Context) int {
	id, err := bpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BlogPosts.
func (bpq *BlogPostQuery) All(ctx context.Context) ([]*BlogPost, error) {
	if err := bpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bpq *BlogPostQuery) AllX(ctx context.Context) []*BlogPost {
	nodes, err := bpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BlogPost IDs.
func (bpq *BlogPostQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bpq.Select(blogpost.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bpq *BlogPostQuery) IDsX(ctx context.Context) []int {
	ids, err := bpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bpq *BlogPostQuery) Count(ctx context.Context) (int, error) {
	if err := bpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bpq *BlogPostQuery) CountX(ctx context.Context) int {
	count, err := bpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bpq *BlogPostQuery) Exist(ctx context.Context) (bool, error) {
	if err := bpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bpq *BlogPostQuery) ExistX(ctx context.Context) bool {
	exist, err := bpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BlogPostQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bpq *BlogPostQuery) Clone() *BlogPostQuery {
	if bpq == nil {
		return nil
	}
	return &BlogPostQuery{
		config:         bpq.config,
		limit:          bpq.limit,
		offset:         bpq.offset,
		order:          append([]OrderFunc{}, bpq.order...),
		predicates:     append([]predicate.BlogPost{}, bpq.predicates...),
		withAuthor:     bpq.withAuthor.Clone(),
		withCategories: bpq.withCategories.Clone(),
		// clone intermediate query.
		sql:  bpq.sql.Clone(),
		path: bpq.path,
	}
}

// WithAuthor tells the query-builder to eager-load the nodes that are connected to
// the "author" edge. The optional arguments are used to configure the query builder of the edge.
func (bpq *BlogPostQuery) WithAuthor(opts ...func(*UserQuery)) *BlogPostQuery {
	query := &UserQuery{config: bpq.config}
	for _, opt := range opts {
		opt(query)
	}
	bpq.withAuthor = query
	return bpq
}

// WithCategories tells the query-builder to eager-load the nodes that are connected to
// the "categories" edge. The optional arguments are used to configure the query builder of the edge.
func (bpq *BlogPostQuery) WithCategories(opts ...func(*CategoryQuery)) *BlogPostQuery {
	query := &CategoryQuery{config: bpq.config}
	for _, opt := range opts {
		opt(query)
	}
	bpq.withCategories = query
	return bpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BlogPost.Query().
//		GroupBy(blogpost.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bpq *BlogPostQuery) GroupBy(field string, fields ...string) *BlogPostGroupBy {
	group := &BlogPostGroupBy{config: bpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bpq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.BlogPost.Query().
//		Select(blogpost.FieldTitle).
//		Scan(ctx, &v)
//
func (bpq *BlogPostQuery) Select(fields ...string) *BlogPostSelect {
	bpq.fields = append(bpq.fields, fields...)
	return &BlogPostSelect{BlogPostQuery: bpq}
}

func (bpq *BlogPostQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bpq.fields {
		if !blogpost.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bpq.path != nil {
		prev, err := bpq.path(ctx)
		if err != nil {
			return err
		}
		bpq.sql = prev
	}
	return nil
}

func (bpq *BlogPostQuery) sqlAll(ctx context.Context) ([]*BlogPost, error) {
	var (
		nodes       = []*BlogPost{}
		withFKs     = bpq.withFKs
		_spec       = bpq.querySpec()
		loadedTypes = [2]bool{
			bpq.withAuthor != nil,
			bpq.withCategories != nil,
		}
	)
	if bpq.withAuthor != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, blogpost.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &BlogPost{config: bpq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, bpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bpq.withAuthor; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*BlogPost)
		for i := range nodes {
			if nodes[i].blog_post_author == nil {
				continue
			}
			fk := *nodes[i].blog_post_author
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "blog_post_author" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Author = n
			}
		}
	}

	if query := bpq.withCategories; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*BlogPost, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Categories = []*Category{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*BlogPost)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   blogpost.CategoriesTable,
				Columns: blogpost.CategoriesPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(blogpost.CategoriesPrimaryKey[1], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, bpq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "categories": %w`, err)
		}
		query.Where(category.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "categories" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Categories = append(nodes[i].Edges.Categories, n)
			}
		}
	}

	return nodes, nil
}

func (bpq *BlogPostQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bpq.querySpec()
	_spec.Node.Columns = bpq.fields
	if len(bpq.fields) > 0 {
		_spec.Unique = bpq.unique != nil && *bpq.unique
	}
	return sqlgraph.CountNodes(ctx, bpq.driver, _spec)
}

func (bpq *BlogPostQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (bpq *BlogPostQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blogpost.Table,
			Columns: blogpost.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: blogpost.FieldID,
			},
		},
		From:   bpq.sql,
		Unique: true,
	}
	if unique := bpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blogpost.FieldID)
		for i := range fields {
			if fields[i] != blogpost.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bpq *BlogPostQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bpq.driver.Dialect())
	t1 := builder.Table(blogpost.Table)
	columns := bpq.fields
	if len(columns) == 0 {
		columns = blogpost.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bpq.sql != nil {
		selector = bpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bpq.unique != nil && *bpq.unique {
		selector.Distinct()
	}
	for _, p := range bpq.predicates {
		p(selector)
	}
	for _, p := range bpq.order {
		p(selector)
	}
	if offset := bpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BlogPostGroupBy is the group-by builder for BlogPost entities.
type BlogPostGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bpgb *BlogPostGroupBy) Aggregate(fns ...AggregateFunc) *BlogPostGroupBy {
	bpgb.fns = append(bpgb.fns, fns...)
	return bpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bpgb *BlogPostGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bpgb.path(ctx)
	if err != nil {
		return err
	}
	bpgb.sql = query
	return bpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bpgb *BlogPostGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := bpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (bpgb *BlogPostGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(bpgb.fields) > 1 {
		return nil, errors.New("ent: BlogPostGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := bpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bpgb *BlogPostGroupBy) StringsX(ctx context.Context) []string {
	v, err := bpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bpgb *BlogPostGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = fmt.Errorf("ent: BlogPostGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bpgb *BlogPostGroupBy) StringX(ctx context.Context) string {
	v, err := bpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (bpgb *BlogPostGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(bpgb.fields) > 1 {
		return nil, errors.New("ent: BlogPostGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := bpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bpgb *BlogPostGroupBy) IntsX(ctx context.Context) []int {
	v, err := bpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bpgb *BlogPostGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = fmt.Errorf("ent: BlogPostGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bpgb *BlogPostGroupBy) IntX(ctx context.Context) int {
	v, err := bpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (bpgb *BlogPostGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(bpgb.fields) > 1 {
		return nil, errors.New("ent: BlogPostGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := bpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bpgb *BlogPostGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := bpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bpgb *BlogPostGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = fmt.Errorf("ent: BlogPostGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bpgb *BlogPostGroupBy) Float64X(ctx context.Context) float64 {
	v, err := bpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (bpgb *BlogPostGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(bpgb.fields) > 1 {
		return nil, errors.New("ent: BlogPostGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := bpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bpgb *BlogPostGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := bpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bpgb *BlogPostGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = fmt.Errorf("ent: BlogPostGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bpgb *BlogPostGroupBy) BoolX(ctx context.Context) bool {
	v, err := bpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bpgb *BlogPostGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bpgb.fields {
		if !blogpost.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bpgb *BlogPostGroupBy) sqlQuery() *sql.Selector {
	selector := bpgb.sql.Select()
	aggregation := make([]string, 0, len(bpgb.fns))
	for _, fn := range bpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bpgb.fields)+len(bpgb.fns))
		for _, f := range bpgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bpgb.fields...)...)
}

// BlogPostSelect is the builder for selecting fields of BlogPost entities.
type BlogPostSelect struct {
	*BlogPostQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bps *BlogPostSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bps.prepareQuery(ctx); err != nil {
		return err
	}
	bps.sql = bps.BlogPostQuery.sqlQuery(ctx)
	return bps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bps *BlogPostSelect) ScanX(ctx context.Context, v interface{}) {
	if err := bps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (bps *BlogPostSelect) Strings(ctx context.Context) ([]string, error) {
	if len(bps.fields) > 1 {
		return nil, errors.New("ent: BlogPostSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := bps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bps *BlogPostSelect) StringsX(ctx context.Context) []string {
	v, err := bps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (bps *BlogPostSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = fmt.Errorf("ent: BlogPostSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bps *BlogPostSelect) StringX(ctx context.Context) string {
	v, err := bps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (bps *BlogPostSelect) Ints(ctx context.Context) ([]int, error) {
	if len(bps.fields) > 1 {
		return nil, errors.New("ent: BlogPostSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := bps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bps *BlogPostSelect) IntsX(ctx context.Context) []int {
	v, err := bps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (bps *BlogPostSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = fmt.Errorf("ent: BlogPostSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bps *BlogPostSelect) IntX(ctx context.Context) int {
	v, err := bps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (bps *BlogPostSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(bps.fields) > 1 {
		return nil, errors.New("ent: BlogPostSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := bps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bps *BlogPostSelect) Float64sX(ctx context.Context) []float64 {
	v, err := bps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (bps *BlogPostSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = fmt.Errorf("ent: BlogPostSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bps *BlogPostSelect) Float64X(ctx context.Context) float64 {
	v, err := bps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (bps *BlogPostSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(bps.fields) > 1 {
		return nil, errors.New("ent: BlogPostSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := bps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bps *BlogPostSelect) BoolsX(ctx context.Context) []bool {
	v, err := bps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (bps *BlogPostSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = fmt.Errorf("ent: BlogPostSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bps *BlogPostSelect) BoolX(ctx context.Context) bool {
	v, err := bps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bps *BlogPostSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bps.sql.Query()
	if err := bps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
