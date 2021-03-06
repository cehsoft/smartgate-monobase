// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/camsetting"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/gate"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/lane"
)

// LaneCreate is the builder for creating a Lane entity.
type LaneCreate struct {
	config
	mutation *LaneMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (lc *LaneCreate) SetName(s string) *LaneCreate {
	lc.mutation.SetName(s)
	return lc
}

// SetGateID sets the "gate_id" field.
func (lc *LaneCreate) SetGateID(i int) *LaneCreate {
	lc.mutation.SetGateID(i)
	return lc
}

// SetNillableGateID sets the "gate_id" field if the given value is not nil.
func (lc *LaneCreate) SetNillableGateID(i *int) *LaneCreate {
	if i != nil {
		lc.SetGateID(*i)
	}
	return lc
}

// SetCreatedAt sets the "created_at" field.
func (lc *LaneCreate) SetCreatedAt(t time.Time) *LaneCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LaneCreate) SetNillableCreatedAt(t *time.Time) *LaneCreate {
	if t != nil {
		lc.SetCreatedAt(*t)
	}
	return lc
}

// AddCamIDs adds the "cams" edge to the CamSetting entity by IDs.
func (lc *LaneCreate) AddCamIDs(ids ...int) *LaneCreate {
	lc.mutation.AddCamIDs(ids...)
	return lc
}

// AddCams adds the "cams" edges to the CamSetting entity.
func (lc *LaneCreate) AddCams(c ...*CamSetting) *LaneCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return lc.AddCamIDs(ids...)
}

// SetGate sets the "gate" edge to the Gate entity.
func (lc *LaneCreate) SetGate(g *Gate) *LaneCreate {
	return lc.SetGateID(g.ID)
}

// Mutation returns the LaneMutation object of the builder.
func (lc *LaneCreate) Mutation() *LaneMutation {
	return lc.mutation
}

// Save creates the Lane in the database.
func (lc *LaneCreate) Save(ctx context.Context) (*Lane, error) {
	var (
		err  error
		node *Lane
	)
	lc.defaults()
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LaneMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lc.check(); err != nil {
				return nil, err
			}
			lc.mutation = mutation
			if node, err = lc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			if lc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LaneCreate) SaveX(ctx context.Context) *Lane {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LaneCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LaneCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LaneCreate) defaults() {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		v := lane.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LaneCreate) check() error {
	if _, ok := lc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	return nil
}

func (lc *LaneCreate) sqlSave(ctx context.Context) (*Lane, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (lc *LaneCreate) createSpec() (*Lane, *sqlgraph.CreateSpec) {
	var (
		_node = &Lane{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: lane.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lane.FieldID,
			},
		}
	)
	_spec.OnConflict = lc.conflict
	if value, ok := lc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lane.FieldName,
		})
		_node.Name = value
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lane.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := lc.mutation.CamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   lane.CamsTable,
			Columns: []string{lane.CamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: camsetting.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.GateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lane.GateTable,
			Columns: []string{lane.GateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.GateID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Lane.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LaneUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (lc *LaneCreate) OnConflict(opts ...sql.ConflictOption) *LaneUpsertOne {
	lc.conflict = opts
	return &LaneUpsertOne{
		create: lc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Lane.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (lc *LaneCreate) OnConflictColumns(columns ...string) *LaneUpsertOne {
	lc.conflict = append(lc.conflict, sql.ConflictColumns(columns...))
	return &LaneUpsertOne{
		create: lc,
	}
}

type (
	// LaneUpsertOne is the builder for "upsert"-ing
	//  one Lane node.
	LaneUpsertOne struct {
		create *LaneCreate
	}

	// LaneUpsert is the "OnConflict" setter.
	LaneUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *LaneUpsert) SetName(v string) *LaneUpsert {
	u.Set(lane.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *LaneUpsert) UpdateName() *LaneUpsert {
	u.SetExcluded(lane.FieldName)
	return u
}

// SetGateID sets the "gate_id" field.
func (u *LaneUpsert) SetGateID(v int) *LaneUpsert {
	u.Set(lane.FieldGateID, v)
	return u
}

// UpdateGateID sets the "gate_id" field to the value that was provided on create.
func (u *LaneUpsert) UpdateGateID() *LaneUpsert {
	u.SetExcluded(lane.FieldGateID)
	return u
}

// ClearGateID clears the value of the "gate_id" field.
func (u *LaneUpsert) ClearGateID() *LaneUpsert {
	u.SetNull(lane.FieldGateID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *LaneUpsert) SetCreatedAt(v time.Time) *LaneUpsert {
	u.Set(lane.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *LaneUpsert) UpdateCreatedAt() *LaneUpsert {
	u.SetExcluded(lane.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Lane.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *LaneUpsertOne) UpdateNewValues() *LaneUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Lane.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *LaneUpsertOne) Ignore() *LaneUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LaneUpsertOne) DoNothing() *LaneUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LaneCreate.OnConflict
// documentation for more info.
func (u *LaneUpsertOne) Update(set func(*LaneUpsert)) *LaneUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LaneUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *LaneUpsertOne) SetName(v string) *LaneUpsertOne {
	return u.Update(func(s *LaneUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *LaneUpsertOne) UpdateName() *LaneUpsertOne {
	return u.Update(func(s *LaneUpsert) {
		s.UpdateName()
	})
}

// SetGateID sets the "gate_id" field.
func (u *LaneUpsertOne) SetGateID(v int) *LaneUpsertOne {
	return u.Update(func(s *LaneUpsert) {
		s.SetGateID(v)
	})
}

// UpdateGateID sets the "gate_id" field to the value that was provided on create.
func (u *LaneUpsertOne) UpdateGateID() *LaneUpsertOne {
	return u.Update(func(s *LaneUpsert) {
		s.UpdateGateID()
	})
}

// ClearGateID clears the value of the "gate_id" field.
func (u *LaneUpsertOne) ClearGateID() *LaneUpsertOne {
	return u.Update(func(s *LaneUpsert) {
		s.ClearGateID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *LaneUpsertOne) SetCreatedAt(v time.Time) *LaneUpsertOne {
	return u.Update(func(s *LaneUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *LaneUpsertOne) UpdateCreatedAt() *LaneUpsertOne {
	return u.Update(func(s *LaneUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *LaneUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LaneCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LaneUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LaneUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LaneUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LaneCreateBulk is the builder for creating many Lane entities in bulk.
type LaneCreateBulk struct {
	config
	builders []*LaneCreate
	conflict []sql.ConflictOption
}

// Save creates the Lane entities in the database.
func (lcb *LaneCreateBulk) Save(ctx context.Context) ([]*Lane, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Lane, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LaneMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = lcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LaneCreateBulk) SaveX(ctx context.Context) []*Lane {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LaneCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LaneCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Lane.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LaneUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (lcb *LaneCreateBulk) OnConflict(opts ...sql.ConflictOption) *LaneUpsertBulk {
	lcb.conflict = opts
	return &LaneUpsertBulk{
		create: lcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Lane.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (lcb *LaneCreateBulk) OnConflictColumns(columns ...string) *LaneUpsertBulk {
	lcb.conflict = append(lcb.conflict, sql.ConflictColumns(columns...))
	return &LaneUpsertBulk{
		create: lcb,
	}
}

// LaneUpsertBulk is the builder for "upsert"-ing
// a bulk of Lane nodes.
type LaneUpsertBulk struct {
	create *LaneCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Lane.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *LaneUpsertBulk) UpdateNewValues() *LaneUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Lane.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *LaneUpsertBulk) Ignore() *LaneUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LaneUpsertBulk) DoNothing() *LaneUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LaneCreateBulk.OnConflict
// documentation for more info.
func (u *LaneUpsertBulk) Update(set func(*LaneUpsert)) *LaneUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LaneUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *LaneUpsertBulk) SetName(v string) *LaneUpsertBulk {
	return u.Update(func(s *LaneUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *LaneUpsertBulk) UpdateName() *LaneUpsertBulk {
	return u.Update(func(s *LaneUpsert) {
		s.UpdateName()
	})
}

// SetGateID sets the "gate_id" field.
func (u *LaneUpsertBulk) SetGateID(v int) *LaneUpsertBulk {
	return u.Update(func(s *LaneUpsert) {
		s.SetGateID(v)
	})
}

// UpdateGateID sets the "gate_id" field to the value that was provided on create.
func (u *LaneUpsertBulk) UpdateGateID() *LaneUpsertBulk {
	return u.Update(func(s *LaneUpsert) {
		s.UpdateGateID()
	})
}

// ClearGateID clears the value of the "gate_id" field.
func (u *LaneUpsertBulk) ClearGateID() *LaneUpsertBulk {
	return u.Update(func(s *LaneUpsert) {
		s.ClearGateID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *LaneUpsertBulk) SetCreatedAt(v time.Time) *LaneUpsertBulk {
	return u.Update(func(s *LaneUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *LaneUpsertBulk) UpdateCreatedAt() *LaneUpsertBulk {
	return u.Update(func(s *LaneUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *LaneUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the LaneCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LaneCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LaneUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
