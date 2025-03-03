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
	"github.com/direktiv/direktiv/pkg/flow/ent/cloudevents"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// CloudEventsQuery is the builder for querying CloudEvents entities.
type CloudEventsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CloudEvents
	// eager-loading edges.
	withNamespace *NamespaceQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CloudEventsQuery builder.
func (ceq *CloudEventsQuery) Where(ps ...predicate.CloudEvents) *CloudEventsQuery {
	ceq.predicates = append(ceq.predicates, ps...)
	return ceq
}

// Limit adds a limit step to the query.
func (ceq *CloudEventsQuery) Limit(limit int) *CloudEventsQuery {
	ceq.limit = &limit
	return ceq
}

// Offset adds an offset step to the query.
func (ceq *CloudEventsQuery) Offset(offset int) *CloudEventsQuery {
	ceq.offset = &offset
	return ceq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ceq *CloudEventsQuery) Unique(unique bool) *CloudEventsQuery {
	ceq.unique = &unique
	return ceq
}

// Order adds an order step to the query.
func (ceq *CloudEventsQuery) Order(o ...OrderFunc) *CloudEventsQuery {
	ceq.order = append(ceq.order, o...)
	return ceq
}

// QueryNamespace chains the current query on the "namespace" edge.
func (ceq *CloudEventsQuery) QueryNamespace() *NamespaceQuery {
	query := &NamespaceQuery{config: ceq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ceq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ceq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cloudevents.Table, cloudevents.FieldID, selector),
			sqlgraph.To(namespace.Table, namespace.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cloudevents.NamespaceTable, cloudevents.NamespaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(ceq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CloudEvents entity from the query.
// Returns a *NotFoundError when no CloudEvents was found.
func (ceq *CloudEventsQuery) First(ctx context.Context) (*CloudEvents, error) {
	nodes, err := ceq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cloudevents.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ceq *CloudEventsQuery) FirstX(ctx context.Context) *CloudEvents {
	node, err := ceq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CloudEvents ID from the query.
// Returns a *NotFoundError when no CloudEvents ID was found.
func (ceq *CloudEventsQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ceq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cloudevents.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ceq *CloudEventsQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ceq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CloudEvents entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CloudEvents entity is found.
// Returns a *NotFoundError when no CloudEvents entities are found.
func (ceq *CloudEventsQuery) Only(ctx context.Context) (*CloudEvents, error) {
	nodes, err := ceq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cloudevents.Label}
	default:
		return nil, &NotSingularError{cloudevents.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ceq *CloudEventsQuery) OnlyX(ctx context.Context) *CloudEvents {
	node, err := ceq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CloudEvents ID in the query.
// Returns a *NotSingularError when more than one CloudEvents ID is found.
// Returns a *NotFoundError when no entities are found.
func (ceq *CloudEventsQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ceq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cloudevents.Label}
	default:
		err = &NotSingularError{cloudevents.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ceq *CloudEventsQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ceq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CloudEventsSlice.
func (ceq *CloudEventsQuery) All(ctx context.Context) ([]*CloudEvents, error) {
	if err := ceq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ceq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ceq *CloudEventsQuery) AllX(ctx context.Context) []*CloudEvents {
	nodes, err := ceq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CloudEvents IDs.
func (ceq *CloudEventsQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := ceq.Select(cloudevents.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ceq *CloudEventsQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ceq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ceq *CloudEventsQuery) Count(ctx context.Context) (int, error) {
	if err := ceq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ceq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ceq *CloudEventsQuery) CountX(ctx context.Context) int {
	count, err := ceq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ceq *CloudEventsQuery) Exist(ctx context.Context) (bool, error) {
	if err := ceq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ceq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ceq *CloudEventsQuery) ExistX(ctx context.Context) bool {
	exist, err := ceq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CloudEventsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ceq *CloudEventsQuery) Clone() *CloudEventsQuery {
	if ceq == nil {
		return nil
	}
	return &CloudEventsQuery{
		config:        ceq.config,
		limit:         ceq.limit,
		offset:        ceq.offset,
		order:         append([]OrderFunc{}, ceq.order...),
		predicates:    append([]predicate.CloudEvents{}, ceq.predicates...),
		withNamespace: ceq.withNamespace.Clone(),
		// clone intermediate query.
		sql:    ceq.sql.Clone(),
		path:   ceq.path,
		unique: ceq.unique,
	}
}

// WithNamespace tells the query-builder to eager-load the nodes that are connected to
// the "namespace" edge. The optional arguments are used to configure the query builder of the edge.
func (ceq *CloudEventsQuery) WithNamespace(opts ...func(*NamespaceQuery)) *CloudEventsQuery {
	query := &NamespaceQuery{config: ceq.config}
	for _, opt := range opts {
		opt(query)
	}
	ceq.withNamespace = query
	return ceq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EventId string `json:"eventId,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CloudEvents.Query().
//		GroupBy(cloudevents.FieldEventId).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ceq *CloudEventsQuery) GroupBy(field string, fields ...string) *CloudEventsGroupBy {
	group := &CloudEventsGroupBy{config: ceq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ceq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ceq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EventId string `json:"eventId,omitempty"`
//	}
//
//	client.CloudEvents.Query().
//		Select(cloudevents.FieldEventId).
//		Scan(ctx, &v)
//
func (ceq *CloudEventsQuery) Select(fields ...string) *CloudEventsSelect {
	ceq.fields = append(ceq.fields, fields...)
	return &CloudEventsSelect{CloudEventsQuery: ceq}
}

func (ceq *CloudEventsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ceq.fields {
		if !cloudevents.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ceq.path != nil {
		prev, err := ceq.path(ctx)
		if err != nil {
			return err
		}
		ceq.sql = prev
	}
	return nil
}

func (ceq *CloudEventsQuery) sqlAll(ctx context.Context) ([]*CloudEvents, error) {
	var (
		nodes       = []*CloudEvents{}
		withFKs     = ceq.withFKs
		_spec       = ceq.querySpec()
		loadedTypes = [1]bool{
			ceq.withNamespace != nil,
		}
	)
	if ceq.withNamespace != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, cloudevents.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CloudEvents{config: ceq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ceq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ceq.withNamespace; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*CloudEvents)
		for i := range nodes {
			if nodes[i].namespace_cloudevents == nil {
				continue
			}
			fk := *nodes[i].namespace_cloudevents
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(namespace.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "namespace_cloudevents" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Namespace = n
			}
		}
	}

	return nodes, nil
}

func (ceq *CloudEventsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ceq.querySpec()
	_spec.Node.Columns = ceq.fields
	if len(ceq.fields) > 0 {
		_spec.Unique = ceq.unique != nil && *ceq.unique
	}
	return sqlgraph.CountNodes(ctx, ceq.driver, _spec)
}

func (ceq *CloudEventsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ceq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ceq *CloudEventsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cloudevents.Table,
			Columns: cloudevents.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: cloudevents.FieldID,
			},
		},
		From:   ceq.sql,
		Unique: true,
	}
	if unique := ceq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ceq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cloudevents.FieldID)
		for i := range fields {
			if fields[i] != cloudevents.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ceq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ceq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ceq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ceq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ceq *CloudEventsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ceq.driver.Dialect())
	t1 := builder.Table(cloudevents.Table)
	columns := ceq.fields
	if len(columns) == 0 {
		columns = cloudevents.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ceq.sql != nil {
		selector = ceq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ceq.unique != nil && *ceq.unique {
		selector.Distinct()
	}
	for _, p := range ceq.predicates {
		p(selector)
	}
	for _, p := range ceq.order {
		p(selector)
	}
	if offset := ceq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ceq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CloudEventsGroupBy is the group-by builder for CloudEvents entities.
type CloudEventsGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cegb *CloudEventsGroupBy) Aggregate(fns ...AggregateFunc) *CloudEventsGroupBy {
	cegb.fns = append(cegb.fns, fns...)
	return cegb
}

// Scan applies the group-by query and scans the result into the given value.
func (cegb *CloudEventsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cegb.path(ctx)
	if err != nil {
		return err
	}
	cegb.sql = query
	return cegb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cegb *CloudEventsGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cegb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (cegb *CloudEventsGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cegb.fields) > 1 {
		return nil, errors.New("ent: CloudEventsGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cegb *CloudEventsGroupBy) StringsX(ctx context.Context) []string {
	v, err := cegb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cegb *CloudEventsGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cegb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cloudevents.Label}
	default:
		err = fmt.Errorf("ent: CloudEventsGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cegb *CloudEventsGroupBy) StringX(ctx context.Context) string {
	v, err := cegb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (cegb *CloudEventsGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cegb.fields) > 1 {
		return nil, errors.New("ent: CloudEventsGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cegb *CloudEventsGroupBy) IntsX(ctx context.Context) []int {
	v, err := cegb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cegb *CloudEventsGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cegb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cloudevents.Label}
	default:
		err = fmt.Errorf("ent: CloudEventsGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cegb *CloudEventsGroupBy) IntX(ctx context.Context) int {
	v, err := cegb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (cegb *CloudEventsGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cegb.fields) > 1 {
		return nil, errors.New("ent: CloudEventsGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cegb *CloudEventsGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cegb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cegb *CloudEventsGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cegb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cloudevents.Label}
	default:
		err = fmt.Errorf("ent: CloudEventsGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cegb *CloudEventsGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cegb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (cegb *CloudEventsGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cegb.fields) > 1 {
		return nil, errors.New("ent: CloudEventsGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cegb *CloudEventsGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cegb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cegb *CloudEventsGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cegb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cloudevents.Label}
	default:
		err = fmt.Errorf("ent: CloudEventsGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cegb *CloudEventsGroupBy) BoolX(ctx context.Context) bool {
	v, err := cegb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cegb *CloudEventsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cegb.fields {
		if !cloudevents.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cegb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cegb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cegb *CloudEventsGroupBy) sqlQuery() *sql.Selector {
	selector := cegb.sql.Select()
	aggregation := make([]string, 0, len(cegb.fns))
	for _, fn := range cegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cegb.fields)+len(cegb.fns))
		for _, f := range cegb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cegb.fields...)...)
}

// CloudEventsSelect is the builder for selecting fields of CloudEvents entities.
type CloudEventsSelect struct {
	*CloudEventsQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ces *CloudEventsSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ces.prepareQuery(ctx); err != nil {
		return err
	}
	ces.sql = ces.CloudEventsQuery.sqlQuery(ctx)
	return ces.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ces *CloudEventsSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ces.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ces *CloudEventsSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ces.fields) > 1 {
		return nil, errors.New("ent: CloudEventsSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ces.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ces *CloudEventsSelect) StringsX(ctx context.Context) []string {
	v, err := ces.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ces *CloudEventsSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ces.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cloudevents.Label}
	default:
		err = fmt.Errorf("ent: CloudEventsSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ces *CloudEventsSelect) StringX(ctx context.Context) string {
	v, err := ces.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ces *CloudEventsSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ces.fields) > 1 {
		return nil, errors.New("ent: CloudEventsSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ces.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ces *CloudEventsSelect) IntsX(ctx context.Context) []int {
	v, err := ces.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ces *CloudEventsSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ces.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cloudevents.Label}
	default:
		err = fmt.Errorf("ent: CloudEventsSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ces *CloudEventsSelect) IntX(ctx context.Context) int {
	v, err := ces.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ces *CloudEventsSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ces.fields) > 1 {
		return nil, errors.New("ent: CloudEventsSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ces.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ces *CloudEventsSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ces.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ces *CloudEventsSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ces.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cloudevents.Label}
	default:
		err = fmt.Errorf("ent: CloudEventsSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ces *CloudEventsSelect) Float64X(ctx context.Context) float64 {
	v, err := ces.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ces *CloudEventsSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ces.fields) > 1 {
		return nil, errors.New("ent: CloudEventsSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ces.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ces *CloudEventsSelect) BoolsX(ctx context.Context) []bool {
	v, err := ces.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ces *CloudEventsSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ces.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{cloudevents.Label}
	default:
		err = fmt.Errorf("ent: CloudEventsSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ces *CloudEventsSelect) BoolX(ctx context.Context) bool {
	v, err := ces.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ces *CloudEventsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ces.sql.Query()
	if err := ces.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
