// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/adminium/dex/storage/ent/db/offlinesession"
	"github.com/adminium/dex/storage/ent/db/predicate"
)

// OfflineSessionDelete is the builder for deleting a OfflineSession entity.
type OfflineSessionDelete struct {
	config
	hooks    []Hook
	mutation *OfflineSessionMutation
}

// Where appends a list predicates to the OfflineSessionDelete builder.
func (osd *OfflineSessionDelete) Where(ps ...predicate.OfflineSession) *OfflineSessionDelete {
	osd.mutation.Where(ps...)
	return osd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (osd *OfflineSessionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, OfflineSessionMutation](ctx, osd.sqlExec, osd.mutation, osd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (osd *OfflineSessionDelete) ExecX(ctx context.Context) int {
	n, err := osd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (osd *OfflineSessionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(offlinesession.Table, sqlgraph.NewFieldSpec(offlinesession.FieldID, field.TypeString))
	if ps := osd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, osd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	osd.mutation.done = true
	return affected, err
}

// OfflineSessionDeleteOne is the builder for deleting a single OfflineSession entity.
type OfflineSessionDeleteOne struct {
	osd *OfflineSessionDelete
}

// Where appends a list predicates to the OfflineSessionDelete builder.
func (osdo *OfflineSessionDeleteOne) Where(ps ...predicate.OfflineSession) *OfflineSessionDeleteOne {
	osdo.osd.mutation.Where(ps...)
	return osdo
}

// Exec executes the deletion query.
func (osdo *OfflineSessionDeleteOne) Exec(ctx context.Context) error {
	n, err := osdo.osd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{offlinesession.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (osdo *OfflineSessionDeleteOne) ExecX(ctx context.Context) {
	if err := osdo.Exec(ctx); err != nil {
		panic(err)
	}
}
