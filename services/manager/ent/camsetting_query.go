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
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/camsetting"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/lane"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/predicate"
)

// CamSettingQuery is the builder for querying CamSetting entities.
type CamSettingQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CamSetting
	// eager-loading edges.
	withLane *LaneQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CamSettingQuery builder.
func (csq *CamSettingQuery) Where(ps ...predicate.CamSetting) *CamSettingQuery {
	csq.predicates = append(csq.predicates, ps...)
	return csq
}

// Limit adds a limit step to the query.
func (csq *CamSettingQuery) Limit(limit int) *CamSettingQuery {
	csq.limit = &limit
	return csq
}

// Offset adds an offset step to the query.
func (csq *CamSettingQuery) Offset(offset int) *CamSettingQuery {
	csq.offset = &offset
	return csq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (csq *CamSettingQuery) Unique(unique bool) *CamSettingQuery {
	csq.unique = &unique
	return csq
}

// Order adds an order step to the query.
func (csq *CamSettingQuery) Order(o ...OrderFunc) *CamSettingQuery {
	csq.order = append(csq.order, o...)
	return csq
}

// QueryLane chains the current query on the "lane" edge.
func (csq *CamSettingQuery) QueryLane() *LaneQuery {
	query := &LaneQuery{config: csq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := csq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := csq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(camsetting.Table, camsetting.FieldID, selector),
			sqlgraph.To(lane.Table, lane.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, camsetting.LaneTable, camsetting.LaneColumn),
		)
		fromU = sqlgraph.SetNeighbors(csq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CamSetting entity from the query.
// Returns a *NotFoundError when no CamSetting was found.
func (csq *CamSettingQuery) First(ctx context.Context) (*CamSetting, error) {
	nodes, err := csq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{camsetting.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (csq *CamSettingQuery) FirstX(ctx context.Context) *CamSetting {
	node, err := csq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CamSetting ID from the query.
// Returns a *NotFoundError when no CamSetting ID was found.
func (csq *CamSettingQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = csq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{camsetting.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (csq *CamSettingQuery) FirstIDX(ctx context.Context) int {
	id, err := csq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CamSetting entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one CamSetting entity is not found.
// Returns a *NotFoundError when no CamSetting entities are found.
func (csq *CamSettingQuery) Only(ctx context.Context) (*CamSetting, error) {
	nodes, err := csq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{camsetting.Label}
	default:
		return nil, &NotSingularError{camsetting.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (csq *CamSettingQuery) OnlyX(ctx context.Context) *CamSetting {
	node, err := csq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CamSetting ID in the query.
// Returns a *NotSingularError when exactly one CamSetting ID is not found.
// Returns a *NotFoundError when no entities are found.
func (csq *CamSettingQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = csq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{camsetting.Label}
	default:
		err = &NotSingularError{camsetting.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (csq *CamSettingQuery) OnlyIDX(ctx context.Context) int {
	id, err := csq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CamSettings.
func (csq *CamSettingQuery) All(ctx context.Context) ([]*CamSetting, error) {
	if err := csq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return csq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (csq *CamSettingQuery) AllX(ctx context.Context) []*CamSetting {
	nodes, err := csq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CamSetting IDs.
func (csq *CamSettingQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := csq.Select(camsetting.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (csq *CamSettingQuery) IDsX(ctx context.Context) []int {
	ids, err := csq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (csq *CamSettingQuery) Count(ctx context.Context) (int, error) {
	if err := csq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return csq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (csq *CamSettingQuery) CountX(ctx context.Context) int {
	count, err := csq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (csq *CamSettingQuery) Exist(ctx context.Context) (bool, error) {
	if err := csq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return csq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (csq *CamSettingQuery) ExistX(ctx context.Context) bool {
	exist, err := csq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CamSettingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (csq *CamSettingQuery) Clone() *CamSettingQuery {
	if csq == nil {
		return nil
	}
	return &CamSettingQuery{
		config:     csq.config,
		limit:      csq.limit,
		offset:     csq.offset,
		order:      append([]OrderFunc{}, csq.order...),
		predicates: append([]predicate.CamSetting{}, csq.predicates...),
		withLane:   csq.withLane.Clone(),
		// clone intermediate query.
		sql:  csq.sql.Clone(),
		path: csq.path,
	}
}

// WithLane tells the query-builder to eager-load the nodes that are connected to
// the "lane" edge. The optional arguments are used to configure the query builder of the edge.
func (csq *CamSettingQuery) WithLane(opts ...func(*LaneQuery)) *CamSettingQuery {
	query := &LaneQuery{config: csq.config}
	for _, opt := range opts {
		opt(query)
	}
	csq.withLane = query
	return csq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CamSetting.Query().
//		GroupBy(camsetting.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (csq *CamSettingQuery) GroupBy(field string, fields ...string) *CamSettingGroupBy {
	group := &CamSettingGroupBy{config: csq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := csq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return csq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.CamSetting.Query().
//		Select(camsetting.FieldName).
//		Scan(ctx, &v)
//
func (csq *CamSettingQuery) Select(field string, fields ...string) *CamSettingSelect {
	csq.fields = append([]string{field}, fields...)
	return &CamSettingSelect{CamSettingQuery: csq}
}

func (csq *CamSettingQuery) prepareQuery(ctx context.Context) error {
	for _, f := range csq.fields {
		if !camsetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if csq.path != nil {
		prev, err := csq.path(ctx)
		if err != nil {
			return err
		}
		csq.sql = prev
	}
	return nil
}

func (csq *CamSettingQuery) sqlAll(ctx context.Context) ([]*CamSetting, error) {
	var (
		nodes       = []*CamSetting{}
		_spec       = csq.querySpec()
		loadedTypes = [1]bool{
			csq.withLane != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CamSetting{config: csq.config}
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
	if err := sqlgraph.QueryNodes(ctx, csq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := csq.withLane; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CamSetting)
		for i := range nodes {
			fk := nodes[i].LaneID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(lane.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "lane_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Lane = n
			}
		}
	}

	return nodes, nil
}

func (csq *CamSettingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := csq.querySpec()
	return sqlgraph.CountNodes(ctx, csq.driver, _spec)
}

func (csq *CamSettingQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := csq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (csq *CamSettingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   camsetting.Table,
			Columns: camsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: camsetting.FieldID,
			},
		},
		From:   csq.sql,
		Unique: true,
	}
	if unique := csq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := csq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, camsetting.FieldID)
		for i := range fields {
			if fields[i] != camsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := csq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := csq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := csq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := csq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (csq *CamSettingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(csq.driver.Dialect())
	t1 := builder.Table(camsetting.Table)
	selector := builder.Select(t1.Columns(camsetting.Columns...)...).From(t1)
	if csq.sql != nil {
		selector = csq.sql
		selector.Select(selector.Columns(camsetting.Columns...)...)
	}
	for _, p := range csq.predicates {
		p(selector)
	}
	for _, p := range csq.order {
		p(selector)
	}
	if offset := csq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := csq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CamSettingGroupBy is the group-by builder for CamSetting entities.
type CamSettingGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (csgb *CamSettingGroupBy) Aggregate(fns ...AggregateFunc) *CamSettingGroupBy {
	csgb.fns = append(csgb.fns, fns...)
	return csgb
}

// Scan applies the group-by query and scans the result into the given value.
func (csgb *CamSettingGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := csgb.path(ctx)
	if err != nil {
		return err
	}
	csgb.sql = query
	return csgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (csgb *CamSettingGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := csgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (csgb *CamSettingGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(csgb.fields) > 1 {
		return nil, errors.New("ent: CamSettingGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := csgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (csgb *CamSettingGroupBy) StringsX(ctx context.Context) []string {
	v, err := csgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (csgb *CamSettingGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = csgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{camsetting.Label}
	default:
		err = fmt.Errorf("ent: CamSettingGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (csgb *CamSettingGroupBy) StringX(ctx context.Context) string {
	v, err := csgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (csgb *CamSettingGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(csgb.fields) > 1 {
		return nil, errors.New("ent: CamSettingGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := csgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (csgb *CamSettingGroupBy) IntsX(ctx context.Context) []int {
	v, err := csgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (csgb *CamSettingGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = csgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{camsetting.Label}
	default:
		err = fmt.Errorf("ent: CamSettingGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (csgb *CamSettingGroupBy) IntX(ctx context.Context) int {
	v, err := csgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (csgb *CamSettingGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(csgb.fields) > 1 {
		return nil, errors.New("ent: CamSettingGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := csgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (csgb *CamSettingGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := csgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (csgb *CamSettingGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = csgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{camsetting.Label}
	default:
		err = fmt.Errorf("ent: CamSettingGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (csgb *CamSettingGroupBy) Float64X(ctx context.Context) float64 {
	v, err := csgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (csgb *CamSettingGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(csgb.fields) > 1 {
		return nil, errors.New("ent: CamSettingGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := csgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (csgb *CamSettingGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := csgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (csgb *CamSettingGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = csgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{camsetting.Label}
	default:
		err = fmt.Errorf("ent: CamSettingGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (csgb *CamSettingGroupBy) BoolX(ctx context.Context) bool {
	v, err := csgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (csgb *CamSettingGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range csgb.fields {
		if !camsetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := csgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := csgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (csgb *CamSettingGroupBy) sqlQuery() *sql.Selector {
	selector := csgb.sql
	columns := make([]string, 0, len(csgb.fields)+len(csgb.fns))
	columns = append(columns, csgb.fields...)
	for _, fn := range csgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(csgb.fields...)
}

// CamSettingSelect is the builder for selecting fields of CamSetting entities.
type CamSettingSelect struct {
	*CamSettingQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (css *CamSettingSelect) Scan(ctx context.Context, v interface{}) error {
	if err := css.prepareQuery(ctx); err != nil {
		return err
	}
	css.sql = css.CamSettingQuery.sqlQuery(ctx)
	return css.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (css *CamSettingSelect) ScanX(ctx context.Context, v interface{}) {
	if err := css.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (css *CamSettingSelect) Strings(ctx context.Context) ([]string, error) {
	if len(css.fields) > 1 {
		return nil, errors.New("ent: CamSettingSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := css.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (css *CamSettingSelect) StringsX(ctx context.Context) []string {
	v, err := css.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (css *CamSettingSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = css.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{camsetting.Label}
	default:
		err = fmt.Errorf("ent: CamSettingSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (css *CamSettingSelect) StringX(ctx context.Context) string {
	v, err := css.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (css *CamSettingSelect) Ints(ctx context.Context) ([]int, error) {
	if len(css.fields) > 1 {
		return nil, errors.New("ent: CamSettingSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := css.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (css *CamSettingSelect) IntsX(ctx context.Context) []int {
	v, err := css.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (css *CamSettingSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = css.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{camsetting.Label}
	default:
		err = fmt.Errorf("ent: CamSettingSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (css *CamSettingSelect) IntX(ctx context.Context) int {
	v, err := css.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (css *CamSettingSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(css.fields) > 1 {
		return nil, errors.New("ent: CamSettingSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := css.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (css *CamSettingSelect) Float64sX(ctx context.Context) []float64 {
	v, err := css.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (css *CamSettingSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = css.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{camsetting.Label}
	default:
		err = fmt.Errorf("ent: CamSettingSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (css *CamSettingSelect) Float64X(ctx context.Context) float64 {
	v, err := css.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (css *CamSettingSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(css.fields) > 1 {
		return nil, errors.New("ent: CamSettingSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := css.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (css *CamSettingSelect) BoolsX(ctx context.Context) []bool {
	v, err := css.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (css *CamSettingSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = css.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{camsetting.Label}
	default:
		err = fmt.Errorf("ent: CamSettingSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (css *CamSettingSelect) BoolX(ctx context.Context) bool {
	v, err := css.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (css *CamSettingSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := css.sqlQuery().Query()
	if err := css.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (css *CamSettingSelect) sqlQuery() sql.Querier {
	selector := css.sql
	selector.Select(selector.Columns(css.fields...)...)
	return selector
}
