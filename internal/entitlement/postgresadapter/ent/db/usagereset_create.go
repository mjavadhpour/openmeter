// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/internal/entitlement/postgresadapter/ent/db/entitlement"
	"github.com/openmeterio/openmeter/internal/entitlement/postgresadapter/ent/db/usagereset"
)

// UsageResetCreate is the builder for creating a UsageReset entity.
type UsageResetCreate struct {
	config
	mutation *UsageResetMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetNamespace sets the "namespace" field.
func (urc *UsageResetCreate) SetNamespace(s string) *UsageResetCreate {
	urc.mutation.SetNamespace(s)
	return urc
}

// SetCreatedAt sets the "created_at" field.
func (urc *UsageResetCreate) SetCreatedAt(t time.Time) *UsageResetCreate {
	urc.mutation.SetCreatedAt(t)
	return urc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (urc *UsageResetCreate) SetNillableCreatedAt(t *time.Time) *UsageResetCreate {
	if t != nil {
		urc.SetCreatedAt(*t)
	}
	return urc
}

// SetUpdatedAt sets the "updated_at" field.
func (urc *UsageResetCreate) SetUpdatedAt(t time.Time) *UsageResetCreate {
	urc.mutation.SetUpdatedAt(t)
	return urc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (urc *UsageResetCreate) SetNillableUpdatedAt(t *time.Time) *UsageResetCreate {
	if t != nil {
		urc.SetUpdatedAt(*t)
	}
	return urc
}

// SetDeletedAt sets the "deleted_at" field.
func (urc *UsageResetCreate) SetDeletedAt(t time.Time) *UsageResetCreate {
	urc.mutation.SetDeletedAt(t)
	return urc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (urc *UsageResetCreate) SetNillableDeletedAt(t *time.Time) *UsageResetCreate {
	if t != nil {
		urc.SetDeletedAt(*t)
	}
	return urc
}

// SetEntitlementID sets the "entitlement_id" field.
func (urc *UsageResetCreate) SetEntitlementID(s string) *UsageResetCreate {
	urc.mutation.SetEntitlementID(s)
	return urc
}

// SetResetTime sets the "reset_time" field.
func (urc *UsageResetCreate) SetResetTime(t time.Time) *UsageResetCreate {
	urc.mutation.SetResetTime(t)
	return urc
}

// SetID sets the "id" field.
func (urc *UsageResetCreate) SetID(s string) *UsageResetCreate {
	urc.mutation.SetID(s)
	return urc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (urc *UsageResetCreate) SetNillableID(s *string) *UsageResetCreate {
	if s != nil {
		urc.SetID(*s)
	}
	return urc
}

// SetEntitlement sets the "entitlement" edge to the Entitlement entity.
func (urc *UsageResetCreate) SetEntitlement(e *Entitlement) *UsageResetCreate {
	return urc.SetEntitlementID(e.ID)
}

// Mutation returns the UsageResetMutation object of the builder.
func (urc *UsageResetCreate) Mutation() *UsageResetMutation {
	return urc.mutation
}

// Save creates the UsageReset in the database.
func (urc *UsageResetCreate) Save(ctx context.Context) (*UsageReset, error) {
	urc.defaults()
	return withHooks(ctx, urc.sqlSave, urc.mutation, urc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (urc *UsageResetCreate) SaveX(ctx context.Context) *UsageReset {
	v, err := urc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urc *UsageResetCreate) Exec(ctx context.Context) error {
	_, err := urc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urc *UsageResetCreate) ExecX(ctx context.Context) {
	if err := urc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (urc *UsageResetCreate) defaults() {
	if _, ok := urc.mutation.CreatedAt(); !ok {
		v := usagereset.DefaultCreatedAt()
		urc.mutation.SetCreatedAt(v)
	}
	if _, ok := urc.mutation.UpdatedAt(); !ok {
		v := usagereset.DefaultUpdatedAt()
		urc.mutation.SetUpdatedAt(v)
	}
	if _, ok := urc.mutation.ID(); !ok {
		v := usagereset.DefaultID()
		urc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (urc *UsageResetCreate) check() error {
	if _, ok := urc.mutation.Namespace(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`db: missing required field "UsageReset.namespace"`)}
	}
	if v, ok := urc.mutation.Namespace(); ok {
		if err := usagereset.NamespaceValidator(v); err != nil {
			return &ValidationError{Name: "namespace", err: fmt.Errorf(`db: validator failed for field "UsageReset.namespace": %w`, err)}
		}
	}
	if _, ok := urc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "UsageReset.created_at"`)}
	}
	if _, ok := urc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`db: missing required field "UsageReset.updated_at"`)}
	}
	if _, ok := urc.mutation.EntitlementID(); !ok {
		return &ValidationError{Name: "entitlement_id", err: errors.New(`db: missing required field "UsageReset.entitlement_id"`)}
	}
	if _, ok := urc.mutation.ResetTime(); !ok {
		return &ValidationError{Name: "reset_time", err: errors.New(`db: missing required field "UsageReset.reset_time"`)}
	}
	if _, ok := urc.mutation.EntitlementID(); !ok {
		return &ValidationError{Name: "entitlement", err: errors.New(`db: missing required edge "UsageReset.entitlement"`)}
	}
	return nil
}

func (urc *UsageResetCreate) sqlSave(ctx context.Context) (*UsageReset, error) {
	if err := urc.check(); err != nil {
		return nil, err
	}
	_node, _spec := urc.createSpec()
	if err := sqlgraph.CreateNode(ctx, urc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected UsageReset.ID type: %T", _spec.ID.Value)
		}
	}
	urc.mutation.id = &_node.ID
	urc.mutation.done = true
	return _node, nil
}

func (urc *UsageResetCreate) createSpec() (*UsageReset, *sqlgraph.CreateSpec) {
	var (
		_node = &UsageReset{config: urc.config}
		_spec = sqlgraph.NewCreateSpec(usagereset.Table, sqlgraph.NewFieldSpec(usagereset.FieldID, field.TypeString))
	)
	_spec.OnConflict = urc.conflict
	if id, ok := urc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := urc.mutation.Namespace(); ok {
		_spec.SetField(usagereset.FieldNamespace, field.TypeString, value)
		_node.Namespace = value
	}
	if value, ok := urc.mutation.CreatedAt(); ok {
		_spec.SetField(usagereset.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := urc.mutation.UpdatedAt(); ok {
		_spec.SetField(usagereset.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := urc.mutation.DeletedAt(); ok {
		_spec.SetField(usagereset.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := urc.mutation.ResetTime(); ok {
		_spec.SetField(usagereset.FieldResetTime, field.TypeTime, value)
		_node.ResetTime = value
	}
	if nodes := urc.mutation.EntitlementIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usagereset.EntitlementTable,
			Columns: []string{usagereset.EntitlementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entitlement.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.EntitlementID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UsageReset.Create().
//		SetNamespace(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UsageResetUpsert) {
//			SetNamespace(v+v).
//		}).
//		Exec(ctx)
func (urc *UsageResetCreate) OnConflict(opts ...sql.ConflictOption) *UsageResetUpsertOne {
	urc.conflict = opts
	return &UsageResetUpsertOne{
		create: urc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UsageReset.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (urc *UsageResetCreate) OnConflictColumns(columns ...string) *UsageResetUpsertOne {
	urc.conflict = append(urc.conflict, sql.ConflictColumns(columns...))
	return &UsageResetUpsertOne{
		create: urc,
	}
}

type (
	// UsageResetUpsertOne is the builder for "upsert"-ing
	//  one UsageReset node.
	UsageResetUpsertOne struct {
		create *UsageResetCreate
	}

	// UsageResetUpsert is the "OnConflict" setter.
	UsageResetUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *UsageResetUpsert) SetUpdatedAt(v time.Time) *UsageResetUpsert {
	u.Set(usagereset.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UsageResetUpsert) UpdateUpdatedAt() *UsageResetUpsert {
	u.SetExcluded(usagereset.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UsageResetUpsert) SetDeletedAt(v time.Time) *UsageResetUpsert {
	u.Set(usagereset.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UsageResetUpsert) UpdateDeletedAt() *UsageResetUpsert {
	u.SetExcluded(usagereset.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *UsageResetUpsert) ClearDeletedAt() *UsageResetUpsert {
	u.SetNull(usagereset.FieldDeletedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.UsageReset.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(usagereset.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UsageResetUpsertOne) UpdateNewValues() *UsageResetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(usagereset.FieldID)
		}
		if _, exists := u.create.mutation.Namespace(); exists {
			s.SetIgnore(usagereset.FieldNamespace)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(usagereset.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.EntitlementID(); exists {
			s.SetIgnore(usagereset.FieldEntitlementID)
		}
		if _, exists := u.create.mutation.ResetTime(); exists {
			s.SetIgnore(usagereset.FieldResetTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UsageReset.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UsageResetUpsertOne) Ignore() *UsageResetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UsageResetUpsertOne) DoNothing() *UsageResetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UsageResetCreate.OnConflict
// documentation for more info.
func (u *UsageResetUpsertOne) Update(set func(*UsageResetUpsert)) *UsageResetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UsageResetUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UsageResetUpsertOne) SetUpdatedAt(v time.Time) *UsageResetUpsertOne {
	return u.Update(func(s *UsageResetUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UsageResetUpsertOne) UpdateUpdatedAt() *UsageResetUpsertOne {
	return u.Update(func(s *UsageResetUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UsageResetUpsertOne) SetDeletedAt(v time.Time) *UsageResetUpsertOne {
	return u.Update(func(s *UsageResetUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UsageResetUpsertOne) UpdateDeletedAt() *UsageResetUpsertOne {
	return u.Update(func(s *UsageResetUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *UsageResetUpsertOne) ClearDeletedAt() *UsageResetUpsertOne {
	return u.Update(func(s *UsageResetUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *UsageResetUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for UsageResetCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UsageResetUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UsageResetUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("db: UsageResetUpsertOne.ID is not supported by MySQL driver. Use UsageResetUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UsageResetUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UsageResetCreateBulk is the builder for creating many UsageReset entities in bulk.
type UsageResetCreateBulk struct {
	config
	err      error
	builders []*UsageResetCreate
	conflict []sql.ConflictOption
}

// Save creates the UsageReset entities in the database.
func (urcb *UsageResetCreateBulk) Save(ctx context.Context) ([]*UsageReset, error) {
	if urcb.err != nil {
		return nil, urcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(urcb.builders))
	nodes := make([]*UsageReset, len(urcb.builders))
	mutators := make([]Mutator, len(urcb.builders))
	for i := range urcb.builders {
		func(i int, root context.Context) {
			builder := urcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UsageResetMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, urcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = urcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, urcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, urcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (urcb *UsageResetCreateBulk) SaveX(ctx context.Context) []*UsageReset {
	v, err := urcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urcb *UsageResetCreateBulk) Exec(ctx context.Context) error {
	_, err := urcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urcb *UsageResetCreateBulk) ExecX(ctx context.Context) {
	if err := urcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UsageReset.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UsageResetUpsert) {
//			SetNamespace(v+v).
//		}).
//		Exec(ctx)
func (urcb *UsageResetCreateBulk) OnConflict(opts ...sql.ConflictOption) *UsageResetUpsertBulk {
	urcb.conflict = opts
	return &UsageResetUpsertBulk{
		create: urcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UsageReset.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (urcb *UsageResetCreateBulk) OnConflictColumns(columns ...string) *UsageResetUpsertBulk {
	urcb.conflict = append(urcb.conflict, sql.ConflictColumns(columns...))
	return &UsageResetUpsertBulk{
		create: urcb,
	}
}

// UsageResetUpsertBulk is the builder for "upsert"-ing
// a bulk of UsageReset nodes.
type UsageResetUpsertBulk struct {
	create *UsageResetCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UsageReset.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(usagereset.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UsageResetUpsertBulk) UpdateNewValues() *UsageResetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(usagereset.FieldID)
			}
			if _, exists := b.mutation.Namespace(); exists {
				s.SetIgnore(usagereset.FieldNamespace)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(usagereset.FieldCreatedAt)
			}
			if _, exists := b.mutation.EntitlementID(); exists {
				s.SetIgnore(usagereset.FieldEntitlementID)
			}
			if _, exists := b.mutation.ResetTime(); exists {
				s.SetIgnore(usagereset.FieldResetTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UsageReset.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UsageResetUpsertBulk) Ignore() *UsageResetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UsageResetUpsertBulk) DoNothing() *UsageResetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UsageResetCreateBulk.OnConflict
// documentation for more info.
func (u *UsageResetUpsertBulk) Update(set func(*UsageResetUpsert)) *UsageResetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UsageResetUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UsageResetUpsertBulk) SetUpdatedAt(v time.Time) *UsageResetUpsertBulk {
	return u.Update(func(s *UsageResetUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UsageResetUpsertBulk) UpdateUpdatedAt() *UsageResetUpsertBulk {
	return u.Update(func(s *UsageResetUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UsageResetUpsertBulk) SetDeletedAt(v time.Time) *UsageResetUpsertBulk {
	return u.Update(func(s *UsageResetUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UsageResetUpsertBulk) UpdateDeletedAt() *UsageResetUpsertBulk {
	return u.Update(func(s *UsageResetUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *UsageResetUpsertBulk) ClearDeletedAt() *UsageResetUpsertBulk {
	return u.Update(func(s *UsageResetUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *UsageResetUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("db: OnConflict was set for builder %d. Set it on the UsageResetCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for UsageResetCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UsageResetUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
