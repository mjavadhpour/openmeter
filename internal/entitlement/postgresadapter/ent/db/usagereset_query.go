// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/internal/entitlement/postgresadapter/ent/db/entitlement"
	"github.com/openmeterio/openmeter/internal/entitlement/postgresadapter/ent/db/predicate"
	"github.com/openmeterio/openmeter/internal/entitlement/postgresadapter/ent/db/usagereset"
)

// UsageResetQuery is the builder for querying UsageReset entities.
type UsageResetQuery struct {
	config
	ctx             *QueryContext
	order           []usagereset.OrderOption
	inters          []Interceptor
	predicates      []predicate.UsageReset
	withEntitlement *EntitlementQuery
	modifiers       []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UsageResetQuery builder.
func (urq *UsageResetQuery) Where(ps ...predicate.UsageReset) *UsageResetQuery {
	urq.predicates = append(urq.predicates, ps...)
	return urq
}

// Limit the number of records to be returned by this query.
func (urq *UsageResetQuery) Limit(limit int) *UsageResetQuery {
	urq.ctx.Limit = &limit
	return urq
}

// Offset to start from.
func (urq *UsageResetQuery) Offset(offset int) *UsageResetQuery {
	urq.ctx.Offset = &offset
	return urq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (urq *UsageResetQuery) Unique(unique bool) *UsageResetQuery {
	urq.ctx.Unique = &unique
	return urq
}

// Order specifies how the records should be ordered.
func (urq *UsageResetQuery) Order(o ...usagereset.OrderOption) *UsageResetQuery {
	urq.order = append(urq.order, o...)
	return urq
}

// QueryEntitlement chains the current query on the "entitlement" edge.
func (urq *UsageResetQuery) QueryEntitlement() *EntitlementQuery {
	query := (&EntitlementClient{config: urq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := urq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := urq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usagereset.Table, usagereset.FieldID, selector),
			sqlgraph.To(entitlement.Table, entitlement.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, usagereset.EntitlementTable, usagereset.EntitlementColumn),
		)
		fromU = sqlgraph.SetNeighbors(urq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UsageReset entity from the query.
// Returns a *NotFoundError when no UsageReset was found.
func (urq *UsageResetQuery) First(ctx context.Context) (*UsageReset, error) {
	nodes, err := urq.Limit(1).All(setContextOp(ctx, urq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usagereset.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (urq *UsageResetQuery) FirstX(ctx context.Context) *UsageReset {
	node, err := urq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UsageReset ID from the query.
// Returns a *NotFoundError when no UsageReset ID was found.
func (urq *UsageResetQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = urq.Limit(1).IDs(setContextOp(ctx, urq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usagereset.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (urq *UsageResetQuery) FirstIDX(ctx context.Context) string {
	id, err := urq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UsageReset entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UsageReset entity is found.
// Returns a *NotFoundError when no UsageReset entities are found.
func (urq *UsageResetQuery) Only(ctx context.Context) (*UsageReset, error) {
	nodes, err := urq.Limit(2).All(setContextOp(ctx, urq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usagereset.Label}
	default:
		return nil, &NotSingularError{usagereset.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (urq *UsageResetQuery) OnlyX(ctx context.Context) *UsageReset {
	node, err := urq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UsageReset ID in the query.
// Returns a *NotSingularError when more than one UsageReset ID is found.
// Returns a *NotFoundError when no entities are found.
func (urq *UsageResetQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = urq.Limit(2).IDs(setContextOp(ctx, urq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usagereset.Label}
	default:
		err = &NotSingularError{usagereset.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (urq *UsageResetQuery) OnlyIDX(ctx context.Context) string {
	id, err := urq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UsageResets.
func (urq *UsageResetQuery) All(ctx context.Context) ([]*UsageReset, error) {
	ctx = setContextOp(ctx, urq.ctx, "All")
	if err := urq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UsageReset, *UsageResetQuery]()
	return withInterceptors[[]*UsageReset](ctx, urq, qr, urq.inters)
}

// AllX is like All, but panics if an error occurs.
func (urq *UsageResetQuery) AllX(ctx context.Context) []*UsageReset {
	nodes, err := urq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UsageReset IDs.
func (urq *UsageResetQuery) IDs(ctx context.Context) (ids []string, err error) {
	if urq.ctx.Unique == nil && urq.path != nil {
		urq.Unique(true)
	}
	ctx = setContextOp(ctx, urq.ctx, "IDs")
	if err = urq.Select(usagereset.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (urq *UsageResetQuery) IDsX(ctx context.Context) []string {
	ids, err := urq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (urq *UsageResetQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, urq.ctx, "Count")
	if err := urq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, urq, querierCount[*UsageResetQuery](), urq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (urq *UsageResetQuery) CountX(ctx context.Context) int {
	count, err := urq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (urq *UsageResetQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, urq.ctx, "Exist")
	switch _, err := urq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (urq *UsageResetQuery) ExistX(ctx context.Context) bool {
	exist, err := urq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UsageResetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (urq *UsageResetQuery) Clone() *UsageResetQuery {
	if urq == nil {
		return nil
	}
	return &UsageResetQuery{
		config:          urq.config,
		ctx:             urq.ctx.Clone(),
		order:           append([]usagereset.OrderOption{}, urq.order...),
		inters:          append([]Interceptor{}, urq.inters...),
		predicates:      append([]predicate.UsageReset{}, urq.predicates...),
		withEntitlement: urq.withEntitlement.Clone(),
		// clone intermediate query.
		sql:  urq.sql.Clone(),
		path: urq.path,
	}
}

// WithEntitlement tells the query-builder to eager-load the nodes that are connected to
// the "entitlement" edge. The optional arguments are used to configure the query builder of the edge.
func (urq *UsageResetQuery) WithEntitlement(opts ...func(*EntitlementQuery)) *UsageResetQuery {
	query := (&EntitlementClient{config: urq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	urq.withEntitlement = query
	return urq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Namespace string `json:"namespace,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UsageReset.Query().
//		GroupBy(usagereset.FieldNamespace).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (urq *UsageResetQuery) GroupBy(field string, fields ...string) *UsageResetGroupBy {
	urq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UsageResetGroupBy{build: urq}
	grbuild.flds = &urq.ctx.Fields
	grbuild.label = usagereset.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Namespace string `json:"namespace,omitempty"`
//	}
//
//	client.UsageReset.Query().
//		Select(usagereset.FieldNamespace).
//		Scan(ctx, &v)
func (urq *UsageResetQuery) Select(fields ...string) *UsageResetSelect {
	urq.ctx.Fields = append(urq.ctx.Fields, fields...)
	sbuild := &UsageResetSelect{UsageResetQuery: urq}
	sbuild.label = usagereset.Label
	sbuild.flds, sbuild.scan = &urq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UsageResetSelect configured with the given aggregations.
func (urq *UsageResetQuery) Aggregate(fns ...AggregateFunc) *UsageResetSelect {
	return urq.Select().Aggregate(fns...)
}

func (urq *UsageResetQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range urq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, urq); err != nil {
				return err
			}
		}
	}
	for _, f := range urq.ctx.Fields {
		if !usagereset.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if urq.path != nil {
		prev, err := urq.path(ctx)
		if err != nil {
			return err
		}
		urq.sql = prev
	}
	return nil
}

func (urq *UsageResetQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UsageReset, error) {
	var (
		nodes       = []*UsageReset{}
		_spec       = urq.querySpec()
		loadedTypes = [1]bool{
			urq.withEntitlement != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UsageReset).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UsageReset{config: urq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(urq.modifiers) > 0 {
		_spec.Modifiers = urq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, urq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := urq.withEntitlement; query != nil {
		if err := urq.loadEntitlement(ctx, query, nodes, nil,
			func(n *UsageReset, e *Entitlement) { n.Edges.Entitlement = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (urq *UsageResetQuery) loadEntitlement(ctx context.Context, query *EntitlementQuery, nodes []*UsageReset, init func(*UsageReset), assign func(*UsageReset, *Entitlement)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*UsageReset)
	for i := range nodes {
		fk := nodes[i].EntitlementID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(entitlement.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "entitlement_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (urq *UsageResetQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := urq.querySpec()
	if len(urq.modifiers) > 0 {
		_spec.Modifiers = urq.modifiers
	}
	_spec.Node.Columns = urq.ctx.Fields
	if len(urq.ctx.Fields) > 0 {
		_spec.Unique = urq.ctx.Unique != nil && *urq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, urq.driver, _spec)
}

func (urq *UsageResetQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(usagereset.Table, usagereset.Columns, sqlgraph.NewFieldSpec(usagereset.FieldID, field.TypeString))
	_spec.From = urq.sql
	if unique := urq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if urq.path != nil {
		_spec.Unique = true
	}
	if fields := urq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usagereset.FieldID)
		for i := range fields {
			if fields[i] != usagereset.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if urq.withEntitlement != nil {
			_spec.Node.AddColumnOnce(usagereset.FieldEntitlementID)
		}
	}
	if ps := urq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := urq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := urq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := urq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (urq *UsageResetQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(urq.driver.Dialect())
	t1 := builder.Table(usagereset.Table)
	columns := urq.ctx.Fields
	if len(columns) == 0 {
		columns = usagereset.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if urq.sql != nil {
		selector = urq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if urq.ctx.Unique != nil && *urq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range urq.modifiers {
		m(selector)
	}
	for _, p := range urq.predicates {
		p(selector)
	}
	for _, p := range urq.order {
		p(selector)
	}
	if offset := urq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := urq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (urq *UsageResetQuery) ForUpdate(opts ...sql.LockOption) *UsageResetQuery {
	if urq.driver.Dialect() == dialect.Postgres {
		urq.Unique(false)
	}
	urq.modifiers = append(urq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return urq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (urq *UsageResetQuery) ForShare(opts ...sql.LockOption) *UsageResetQuery {
	if urq.driver.Dialect() == dialect.Postgres {
		urq.Unique(false)
	}
	urq.modifiers = append(urq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return urq
}

// UsageResetGroupBy is the group-by builder for UsageReset entities.
type UsageResetGroupBy struct {
	selector
	build *UsageResetQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (urgb *UsageResetGroupBy) Aggregate(fns ...AggregateFunc) *UsageResetGroupBy {
	urgb.fns = append(urgb.fns, fns...)
	return urgb
}

// Scan applies the selector query and scans the result into the given value.
func (urgb *UsageResetGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, urgb.build.ctx, "GroupBy")
	if err := urgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UsageResetQuery, *UsageResetGroupBy](ctx, urgb.build, urgb, urgb.build.inters, v)
}

func (urgb *UsageResetGroupBy) sqlScan(ctx context.Context, root *UsageResetQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(urgb.fns))
	for _, fn := range urgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*urgb.flds)+len(urgb.fns))
		for _, f := range *urgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*urgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := urgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UsageResetSelect is the builder for selecting fields of UsageReset entities.
type UsageResetSelect struct {
	*UsageResetQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (urs *UsageResetSelect) Aggregate(fns ...AggregateFunc) *UsageResetSelect {
	urs.fns = append(urs.fns, fns...)
	return urs
}

// Scan applies the selector query and scans the result into the given value.
func (urs *UsageResetSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, urs.ctx, "Select")
	if err := urs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UsageResetQuery, *UsageResetSelect](ctx, urs.UsageResetQuery, urs, urs.inters, v)
}

func (urs *UsageResetSelect) sqlScan(ctx context.Context, root *UsageResetQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(urs.fns))
	for _, fn := range urs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*urs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := urs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
