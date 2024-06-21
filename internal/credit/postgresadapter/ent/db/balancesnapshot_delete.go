// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/internal/credit/postgresadapter/ent/db/balancesnapshot"
	"github.com/openmeterio/openmeter/internal/credit/postgresadapter/ent/db/predicate"
)

// BalanceSnapshotDelete is the builder for deleting a BalanceSnapshot entity.
type BalanceSnapshotDelete struct {
	config
	hooks    []Hook
	mutation *BalanceSnapshotMutation
}

// Where appends a list predicates to the BalanceSnapshotDelete builder.
func (bsd *BalanceSnapshotDelete) Where(ps ...predicate.BalanceSnapshot) *BalanceSnapshotDelete {
	bsd.mutation.Where(ps...)
	return bsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bsd *BalanceSnapshotDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bsd.sqlExec, bsd.mutation, bsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bsd *BalanceSnapshotDelete) ExecX(ctx context.Context) int {
	n, err := bsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bsd *BalanceSnapshotDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(balancesnapshot.Table, sqlgraph.NewFieldSpec(balancesnapshot.FieldID, field.TypeInt))
	if ps := bsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bsd.mutation.done = true
	return affected, err
}

// BalanceSnapshotDeleteOne is the builder for deleting a single BalanceSnapshot entity.
type BalanceSnapshotDeleteOne struct {
	bsd *BalanceSnapshotDelete
}

// Where appends a list predicates to the BalanceSnapshotDelete builder.
func (bsdo *BalanceSnapshotDeleteOne) Where(ps ...predicate.BalanceSnapshot) *BalanceSnapshotDeleteOne {
	bsdo.bsd.mutation.Where(ps...)
	return bsdo
}

// Exec executes the deletion query.
func (bsdo *BalanceSnapshotDeleteOne) Exec(ctx context.Context) error {
	n, err := bsdo.bsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{balancesnapshot.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bsdo *BalanceSnapshotDeleteOne) ExecX(ctx context.Context) {
	if err := bsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
