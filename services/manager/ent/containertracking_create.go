// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/containertracking"
)

// ContainerTrackingCreate is the builder for creating a ContainerTracking entity.
type ContainerTrackingCreate struct {
	config
	mutation *ContainerTrackingMutation
	hooks    []Hook
}

// SetContainerID sets the "container_id" field.
func (ctc *ContainerTrackingCreate) SetContainerID(s string) *ContainerTrackingCreate {
	ctc.mutation.SetContainerID(s)
	return ctc
}

// SetImageURL sets the "image_url" field.
func (ctc *ContainerTrackingCreate) SetImageURL(s string) *ContainerTrackingCreate {
	ctc.mutation.SetImageURL(s)
	return ctc
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (ctc *ContainerTrackingCreate) SetNillableImageURL(s *string) *ContainerTrackingCreate {
	if s != nil {
		ctc.SetImageURL(*s)
	}
	return ctc
}

// SetManual sets the "manual" field.
func (ctc *ContainerTrackingCreate) SetManual(b bool) *ContainerTrackingCreate {
	ctc.mutation.SetManual(b)
	return ctc
}

// SetNillableManual sets the "manual" field if the given value is not nil.
func (ctc *ContainerTrackingCreate) SetNillableManual(b *bool) *ContainerTrackingCreate {
	if b != nil {
		ctc.SetManual(*b)
	}
	return ctc
}

// SetCreatedAt sets the "created_at" field.
func (ctc *ContainerTrackingCreate) SetCreatedAt(t time.Time) *ContainerTrackingCreate {
	ctc.mutation.SetCreatedAt(t)
	return ctc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ctc *ContainerTrackingCreate) SetNillableCreatedAt(t *time.Time) *ContainerTrackingCreate {
	if t != nil {
		ctc.SetCreatedAt(*t)
	}
	return ctc
}

// Mutation returns the ContainerTrackingMutation object of the builder.
func (ctc *ContainerTrackingCreate) Mutation() *ContainerTrackingMutation {
	return ctc.mutation
}

// Save creates the ContainerTracking in the database.
func (ctc *ContainerTrackingCreate) Save(ctx context.Context) (*ContainerTracking, error) {
	var (
		err  error
		node *ContainerTracking
	)
	ctc.defaults()
	if len(ctc.hooks) == 0 {
		if err = ctc.check(); err != nil {
			return nil, err
		}
		node, err = ctc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContainerTrackingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ctc.check(); err != nil {
				return nil, err
			}
			ctc.mutation = mutation
			node, err = ctc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ctc.hooks) - 1; i >= 0; i-- {
			mut = ctc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ctc *ContainerTrackingCreate) SaveX(ctx context.Context) *ContainerTracking {
	v, err := ctc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ctc *ContainerTrackingCreate) defaults() {
	if _, ok := ctc.mutation.Manual(); !ok {
		v := containertracking.DefaultManual
		ctc.mutation.SetManual(v)
	}
	if _, ok := ctc.mutation.CreatedAt(); !ok {
		v := containertracking.DefaultCreatedAt()
		ctc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctc *ContainerTrackingCreate) check() error {
	if _, ok := ctc.mutation.ContainerID(); !ok {
		return &ValidationError{Name: "container_id", err: errors.New("ent: missing required field \"container_id\"")}
	}
	if _, ok := ctc.mutation.Manual(); !ok {
		return &ValidationError{Name: "manual", err: errors.New("ent: missing required field \"manual\"")}
	}
	if _, ok := ctc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	return nil
}

func (ctc *ContainerTrackingCreate) sqlSave(ctx context.Context) (*ContainerTracking, error) {
	_node, _spec := ctc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ctc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ctc *ContainerTrackingCreate) createSpec() (*ContainerTracking, *sqlgraph.CreateSpec) {
	var (
		_node = &ContainerTracking{config: ctc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: containertracking.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: containertracking.FieldID,
			},
		}
	)
	if value, ok := ctc.mutation.ContainerID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertracking.FieldContainerID,
		})
		_node.ContainerID = value
	}
	if value, ok := ctc.mutation.ImageURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertracking.FieldImageURL,
		})
		_node.ImageURL = value
	}
	if value, ok := ctc.mutation.Manual(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: containertracking.FieldManual,
		})
		_node.Manual = value
	}
	if value, ok := ctc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: containertracking.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	return _node, _spec
}

// ContainerTrackingCreateBulk is the builder for creating many ContainerTracking entities in bulk.
type ContainerTrackingCreateBulk struct {
	config
	builders []*ContainerTrackingCreate
}

// Save creates the ContainerTracking entities in the database.
func (ctcb *ContainerTrackingCreateBulk) Save(ctx context.Context) ([]*ContainerTracking, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ctcb.builders))
	nodes := make([]*ContainerTracking, len(ctcb.builders))
	mutators := make([]Mutator, len(ctcb.builders))
	for i := range ctcb.builders {
		func(i int, root context.Context) {
			builder := ctcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContainerTrackingMutation)
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
					_, err = mutators[i+1].Mutate(root, ctcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ctcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ctcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ctcb *ContainerTrackingCreateBulk) SaveX(ctx context.Context) []*ContainerTracking {
	v, err := ctcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}