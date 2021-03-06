// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/graphqlquery"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/predicate"
	"github.com/pepsighan/graftini_backend/internal/pkg/ent/project"
)

// GraphQLQueryQuery is the builder for querying GraphQLQuery entities.
type GraphQLQueryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GraphQLQuery
	// eager-loading edges.
	withQueryOf *ProjectQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GraphQLQueryQuery builder.
func (gqqq *GraphQLQueryQuery) Where(ps ...predicate.GraphQLQuery) *GraphQLQueryQuery {
	gqqq.predicates = append(gqqq.predicates, ps...)
	return gqqq
}

// Limit adds a limit step to the query.
func (gqqq *GraphQLQueryQuery) Limit(limit int) *GraphQLQueryQuery {
	gqqq.limit = &limit
	return gqqq
}

// Offset adds an offset step to the query.
func (gqqq *GraphQLQueryQuery) Offset(offset int) *GraphQLQueryQuery {
	gqqq.offset = &offset
	return gqqq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gqqq *GraphQLQueryQuery) Unique(unique bool) *GraphQLQueryQuery {
	gqqq.unique = &unique
	return gqqq
}

// Order adds an order step to the query.
func (gqqq *GraphQLQueryQuery) Order(o ...OrderFunc) *GraphQLQueryQuery {
	gqqq.order = append(gqqq.order, o...)
	return gqqq
}

// QueryQueryOf chains the current query on the "query_of" edge.
func (gqqq *GraphQLQueryQuery) QueryQueryOf() *ProjectQuery {
	query := &ProjectQuery{config: gqqq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gqqq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gqqq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(graphqlquery.Table, graphqlquery.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, graphqlquery.QueryOfTable, graphqlquery.QueryOfColumn),
		)
		fromU = sqlgraph.SetNeighbors(gqqq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GraphQLQuery entity from the query.
// Returns a *NotFoundError when no GraphQLQuery was found.
func (gqqq *GraphQLQueryQuery) First(ctx context.Context) (*GraphQLQuery, error) {
	nodes, err := gqqq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{graphqlquery.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gqqq *GraphQLQueryQuery) FirstX(ctx context.Context) *GraphQLQuery {
	node, err := gqqq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GraphQLQuery ID from the query.
// Returns a *NotFoundError when no GraphQLQuery ID was found.
func (gqqq *GraphQLQueryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gqqq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{graphqlquery.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gqqq *GraphQLQueryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := gqqq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GraphQLQuery entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one GraphQLQuery entity is not found.
// Returns a *NotFoundError when no GraphQLQuery entities are found.
func (gqqq *GraphQLQueryQuery) Only(ctx context.Context) (*GraphQLQuery, error) {
	nodes, err := gqqq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{graphqlquery.Label}
	default:
		return nil, &NotSingularError{graphqlquery.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gqqq *GraphQLQueryQuery) OnlyX(ctx context.Context) *GraphQLQuery {
	node, err := gqqq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GraphQLQuery ID in the query.
// Returns a *NotSingularError when exactly one GraphQLQuery ID is not found.
// Returns a *NotFoundError when no entities are found.
func (gqqq *GraphQLQueryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gqqq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{graphqlquery.Label}
	default:
		err = &NotSingularError{graphqlquery.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gqqq *GraphQLQueryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := gqqq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GraphQLQueries.
func (gqqq *GraphQLQueryQuery) All(ctx context.Context) ([]*GraphQLQuery, error) {
	if err := gqqq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gqqq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gqqq *GraphQLQueryQuery) AllX(ctx context.Context) []*GraphQLQuery {
	nodes, err := gqqq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GraphQLQuery IDs.
func (gqqq *GraphQLQueryQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := gqqq.Select(graphqlquery.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gqqq *GraphQLQueryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := gqqq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gqqq *GraphQLQueryQuery) Count(ctx context.Context) (int, error) {
	if err := gqqq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gqqq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gqqq *GraphQLQueryQuery) CountX(ctx context.Context) int {
	count, err := gqqq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gqqq *GraphQLQueryQuery) Exist(ctx context.Context) (bool, error) {
	if err := gqqq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gqqq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gqqq *GraphQLQueryQuery) ExistX(ctx context.Context) bool {
	exist, err := gqqq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GraphQLQueryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gqqq *GraphQLQueryQuery) Clone() *GraphQLQueryQuery {
	if gqqq == nil {
		return nil
	}
	return &GraphQLQueryQuery{
		config:      gqqq.config,
		limit:       gqqq.limit,
		offset:      gqqq.offset,
		order:       append([]OrderFunc{}, gqqq.order...),
		predicates:  append([]predicate.GraphQLQuery{}, gqqq.predicates...),
		withQueryOf: gqqq.withQueryOf.Clone(),
		// clone intermediate query.
		sql:  gqqq.sql.Clone(),
		path: gqqq.path,
	}
}

// WithQueryOf tells the query-builder to eager-load the nodes that are connected to
// the "query_of" edge. The optional arguments are used to configure the query builder of the edge.
func (gqqq *GraphQLQueryQuery) WithQueryOf(opts ...func(*ProjectQuery)) *GraphQLQueryQuery {
	query := &ProjectQuery{config: gqqq.config}
	for _, opt := range opts {
		opt(query)
	}
	gqqq.withQueryOf = query
	return gqqq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		VariableName string `json:"variable_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GraphQLQuery.Query().
//		GroupBy(graphqlquery.FieldVariableName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gqqq *GraphQLQueryQuery) GroupBy(field string, fields ...string) *GraphQLQueryGroupBy {
	group := &GraphQLQueryGroupBy{config: gqqq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gqqq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gqqq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		VariableName string `json:"variable_name,omitempty"`
//	}
//
//	client.GraphQLQuery.Query().
//		Select(graphqlquery.FieldVariableName).
//		Scan(ctx, &v)
//
func (gqqq *GraphQLQueryQuery) Select(field string, fields ...string) *GraphQLQuerySelect {
	gqqq.fields = append([]string{field}, fields...)
	return &GraphQLQuerySelect{GraphQLQueryQuery: gqqq}
}

func (gqqq *GraphQLQueryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gqqq.fields {
		if !graphqlquery.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gqqq.path != nil {
		prev, err := gqqq.path(ctx)
		if err != nil {
			return err
		}
		gqqq.sql = prev
	}
	return nil
}

func (gqqq *GraphQLQueryQuery) sqlAll(ctx context.Context) ([]*GraphQLQuery, error) {
	var (
		nodes       = []*GraphQLQuery{}
		withFKs     = gqqq.withFKs
		_spec       = gqqq.querySpec()
		loadedTypes = [1]bool{
			gqqq.withQueryOf != nil,
		}
	)
	if gqqq.withQueryOf != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, graphqlquery.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GraphQLQuery{config: gqqq.config}
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
	if err := sqlgraph.QueryNodes(ctx, gqqq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := gqqq.withQueryOf; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*GraphQLQuery)
		for i := range nodes {
			if nodes[i].project_queries == nil {
				continue
			}
			fk := *nodes[i].project_queries
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(project.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "project_queries" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.QueryOf = n
			}
		}
	}

	return nodes, nil
}

func (gqqq *GraphQLQueryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gqqq.querySpec()
	return sqlgraph.CountNodes(ctx, gqqq.driver, _spec)
}

func (gqqq *GraphQLQueryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gqqq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gqqq *GraphQLQueryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   graphqlquery.Table,
			Columns: graphqlquery.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: graphqlquery.FieldID,
			},
		},
		From:   gqqq.sql,
		Unique: true,
	}
	if unique := gqqq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gqqq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, graphqlquery.FieldID)
		for i := range fields {
			if fields[i] != graphqlquery.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gqqq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gqqq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gqqq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gqqq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gqqq *GraphQLQueryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gqqq.driver.Dialect())
	t1 := builder.Table(graphqlquery.Table)
	selector := builder.Select(t1.Columns(graphqlquery.Columns...)...).From(t1)
	if gqqq.sql != nil {
		selector = gqqq.sql
		selector.Select(selector.Columns(graphqlquery.Columns...)...)
	}
	for _, p := range gqqq.predicates {
		p(selector)
	}
	for _, p := range gqqq.order {
		p(selector)
	}
	if offset := gqqq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gqqq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GraphQLQueryGroupBy is the group-by builder for GraphQLQuery entities.
type GraphQLQueryGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gqqgb *GraphQLQueryGroupBy) Aggregate(fns ...AggregateFunc) *GraphQLQueryGroupBy {
	gqqgb.fns = append(gqqgb.fns, fns...)
	return gqqgb
}

// Scan applies the group-by query and scans the result into the given value.
func (gqqgb *GraphQLQueryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gqqgb.path(ctx)
	if err != nil {
		return err
	}
	gqqgb.sql = query
	return gqqgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gqqgb *GraphQLQueryGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gqqgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (gqqgb *GraphQLQueryGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gqqgb.fields) > 1 {
		return nil, errors.New("ent: GraphQLQueryGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gqqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gqqgb *GraphQLQueryGroupBy) StringsX(ctx context.Context) []string {
	v, err := gqqgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gqqgb *GraphQLQueryGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gqqgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{graphqlquery.Label}
	default:
		err = fmt.Errorf("ent: GraphQLQueryGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gqqgb *GraphQLQueryGroupBy) StringX(ctx context.Context) string {
	v, err := gqqgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (gqqgb *GraphQLQueryGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gqqgb.fields) > 1 {
		return nil, errors.New("ent: GraphQLQueryGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gqqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gqqgb *GraphQLQueryGroupBy) IntsX(ctx context.Context) []int {
	v, err := gqqgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gqqgb *GraphQLQueryGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gqqgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{graphqlquery.Label}
	default:
		err = fmt.Errorf("ent: GraphQLQueryGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gqqgb *GraphQLQueryGroupBy) IntX(ctx context.Context) int {
	v, err := gqqgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (gqqgb *GraphQLQueryGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gqqgb.fields) > 1 {
		return nil, errors.New("ent: GraphQLQueryGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gqqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gqqgb *GraphQLQueryGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gqqgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gqqgb *GraphQLQueryGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gqqgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{graphqlquery.Label}
	default:
		err = fmt.Errorf("ent: GraphQLQueryGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gqqgb *GraphQLQueryGroupBy) Float64X(ctx context.Context) float64 {
	v, err := gqqgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (gqqgb *GraphQLQueryGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gqqgb.fields) > 1 {
		return nil, errors.New("ent: GraphQLQueryGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gqqgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gqqgb *GraphQLQueryGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gqqgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gqqgb *GraphQLQueryGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gqqgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{graphqlquery.Label}
	default:
		err = fmt.Errorf("ent: GraphQLQueryGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gqqgb *GraphQLQueryGroupBy) BoolX(ctx context.Context) bool {
	v, err := gqqgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gqqgb *GraphQLQueryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gqqgb.fields {
		if !graphqlquery.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gqqgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gqqgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gqqgb *GraphQLQueryGroupBy) sqlQuery() *sql.Selector {
	selector := gqqgb.sql
	columns := make([]string, 0, len(gqqgb.fields)+len(gqqgb.fns))
	columns = append(columns, gqqgb.fields...)
	for _, fn := range gqqgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(gqqgb.fields...)
}

// GraphQLQuerySelect is the builder for selecting fields of GraphQLQuery entities.
type GraphQLQuerySelect struct {
	*GraphQLQueryQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gqqs *GraphQLQuerySelect) Scan(ctx context.Context, v interface{}) error {
	if err := gqqs.prepareQuery(ctx); err != nil {
		return err
	}
	gqqs.sql = gqqs.GraphQLQueryQuery.sqlQuery(ctx)
	return gqqs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gqqs *GraphQLQuerySelect) ScanX(ctx context.Context, v interface{}) {
	if err := gqqs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gqqs *GraphQLQuerySelect) Strings(ctx context.Context) ([]string, error) {
	if len(gqqs.fields) > 1 {
		return nil, errors.New("ent: GraphQLQuerySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gqqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gqqs *GraphQLQuerySelect) StringsX(ctx context.Context) []string {
	v, err := gqqs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gqqs *GraphQLQuerySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gqqs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{graphqlquery.Label}
	default:
		err = fmt.Errorf("ent: GraphQLQuerySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gqqs *GraphQLQuerySelect) StringX(ctx context.Context) string {
	v, err := gqqs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gqqs *GraphQLQuerySelect) Ints(ctx context.Context) ([]int, error) {
	if len(gqqs.fields) > 1 {
		return nil, errors.New("ent: GraphQLQuerySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gqqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gqqs *GraphQLQuerySelect) IntsX(ctx context.Context) []int {
	v, err := gqqs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gqqs *GraphQLQuerySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gqqs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{graphqlquery.Label}
	default:
		err = fmt.Errorf("ent: GraphQLQuerySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gqqs *GraphQLQuerySelect) IntX(ctx context.Context) int {
	v, err := gqqs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gqqs *GraphQLQuerySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gqqs.fields) > 1 {
		return nil, errors.New("ent: GraphQLQuerySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gqqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gqqs *GraphQLQuerySelect) Float64sX(ctx context.Context) []float64 {
	v, err := gqqs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gqqs *GraphQLQuerySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gqqs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{graphqlquery.Label}
	default:
		err = fmt.Errorf("ent: GraphQLQuerySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gqqs *GraphQLQuerySelect) Float64X(ctx context.Context) float64 {
	v, err := gqqs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gqqs *GraphQLQuerySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gqqs.fields) > 1 {
		return nil, errors.New("ent: GraphQLQuerySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gqqs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gqqs *GraphQLQuerySelect) BoolsX(ctx context.Context) []bool {
	v, err := gqqs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gqqs *GraphQLQuerySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gqqs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{graphqlquery.Label}
	default:
		err = fmt.Errorf("ent: GraphQLQuerySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gqqs *GraphQLQuerySelect) BoolX(ctx context.Context) bool {
	v, err := gqqs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gqqs *GraphQLQuerySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gqqs.sqlQuery().Query()
	if err := gqqs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gqqs *GraphQLQuerySelect) sqlQuery() sql.Querier {
	selector := gqqs.sql
	selector.Select(selector.Columns(gqqs.fields...)...)
	return selector
}
