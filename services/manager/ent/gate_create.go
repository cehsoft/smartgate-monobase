// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/gate"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/lane"
)

// GateCreate is the builder for creating a Gate entity.
type GateCreate struct {
	config
	mutation *GateMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (gc *GateCreate) SetName(s string) *GateCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetCreatedAt sets the "created_at" field.
func (gc *GateCreate) SetCreatedAt(t time.Time) *GateCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GateCreate) SetNillableCreatedAt(t *time.Time) *GateCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// AddLaneIDs adds the "lanes" edge to the Lane entity by IDs.
func (gc *GateCreate) AddLaneIDs(ids ...int) *GateCreate {
	gc.mutation.AddLaneIDs(ids...)
	return gc
}

// AddLanes adds the "lanes" edges to the Lane entity.
func (gc *GateCreate) AddLanes(l ...*Lane) *GateCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return gc.AddLaneIDs(ids...)
}

// Mutation returns the GateMutation object of the builder.
func (gc *GateCreate) Mutation() *GateMutation {
	return gc.mutation
}

// Save creates the Gate in the database.
func (gc *GateCreate) Save(ctx context.Context) (*Gate, error) {
	var (
		err  error
		node *Gate
	)
	gc.defaults()
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			node, err = gc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GateCreate) SaveX(ctx context.Context) *Gate {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (gc *GateCreate) defaults() {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		v := gate.DefaultCreatedAt()
		gc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GateCreate) check() error {
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	return nil
}

func (gc *GateCreate) sqlSave(ctx context.Context) (*Gate, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gc *GateCreate) createSpec() (*Gate, *sqlgraph.CreateSpec) {
	var (
		_node = &Gate{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: gate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: gate.FieldID,
			},
		}
	)
	if value, ok := gc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gate.FieldName,
		})
		_node.Name = value
	}
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: gate.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := gc.mutation.LanesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gate.LanesTable,
			Columns: []string{gate.LanesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lane.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GateCreateBulk is the builder for creating many Gate entities in bulk.
type GateCreateBulk struct {
	config
	builders []*GateCreate
}

// Save creates the Gate entities in the database.
func (gcb *GateCreateBulk) Save(ctx context.Context) ([]*Gate, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Gate, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GateMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GateCreateBulk) SaveX(ctx context.Context) []*Gate {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
