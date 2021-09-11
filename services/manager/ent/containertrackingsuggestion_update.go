// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/camsetting"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/containertracking"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/containertrackingsuggestion"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/predicate"
)

// ContainerTrackingSuggestionUpdate is the builder for updating ContainerTrackingSuggestion entities.
type ContainerTrackingSuggestionUpdate struct {
	config
	hooks    []Hook
	mutation *ContainerTrackingSuggestionMutation
}

// Where appends a list predicates to the ContainerTrackingSuggestionUpdate builder.
func (ctsu *ContainerTrackingSuggestionUpdate) Where(ps ...predicate.ContainerTrackingSuggestion) *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.Where(ps...)
	return ctsu
}

// SetContainerID sets the "container_id" field.
func (ctsu *ContainerTrackingSuggestionUpdate) SetContainerID(s string) *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.SetContainerID(s)
	return ctsu
}

// SetNillableContainerID sets the "container_id" field if the given value is not nil.
func (ctsu *ContainerTrackingSuggestionUpdate) SetNillableContainerID(s *string) *ContainerTrackingSuggestionUpdate {
	if s != nil {
		ctsu.SetContainerID(*s)
	}
	return ctsu
}

// ClearContainerID clears the value of the "container_id" field.
func (ctsu *ContainerTrackingSuggestionUpdate) ClearContainerID() *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.ClearContainerID()
	return ctsu
}

// SetResult sets the "result" field.
func (ctsu *ContainerTrackingSuggestionUpdate) SetResult(s string) *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.SetResult(s)
	return ctsu
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (ctsu *ContainerTrackingSuggestionUpdate) SetNillableResult(s *string) *ContainerTrackingSuggestionUpdate {
	if s != nil {
		ctsu.SetResult(*s)
	}
	return ctsu
}

// ClearResult clears the value of the "result" field.
func (ctsu *ContainerTrackingSuggestionUpdate) ClearResult() *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.ClearResult()
	return ctsu
}

// SetCamID sets the "cam_id" field.
func (ctsu *ContainerTrackingSuggestionUpdate) SetCamID(i int) *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.SetCamID(i)
	return ctsu
}

// SetNillableCamID sets the "cam_id" field if the given value is not nil.
func (ctsu *ContainerTrackingSuggestionUpdate) SetNillableCamID(i *int) *ContainerTrackingSuggestionUpdate {
	if i != nil {
		ctsu.SetCamID(*i)
	}
	return ctsu
}

// ClearCamID clears the value of the "cam_id" field.
func (ctsu *ContainerTrackingSuggestionUpdate) ClearCamID() *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.ClearCamID()
	return ctsu
}

// SetTrackingID sets the "tracking_id" field.
func (ctsu *ContainerTrackingSuggestionUpdate) SetTrackingID(i int) *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.SetTrackingID(i)
	return ctsu
}

// SetNillableTrackingID sets the "tracking_id" field if the given value is not nil.
func (ctsu *ContainerTrackingSuggestionUpdate) SetNillableTrackingID(i *int) *ContainerTrackingSuggestionUpdate {
	if i != nil {
		ctsu.SetTrackingID(*i)
	}
	return ctsu
}

// ClearTrackingID clears the value of the "tracking_id" field.
func (ctsu *ContainerTrackingSuggestionUpdate) ClearTrackingID() *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.ClearTrackingID()
	return ctsu
}

// SetTrackingType sets the "tracking_type" field.
func (ctsu *ContainerTrackingSuggestionUpdate) SetTrackingType(s string) *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.SetTrackingType(s)
	return ctsu
}

// SetNillableTrackingType sets the "tracking_type" field if the given value is not nil.
func (ctsu *ContainerTrackingSuggestionUpdate) SetNillableTrackingType(s *string) *ContainerTrackingSuggestionUpdate {
	if s != nil {
		ctsu.SetTrackingType(*s)
	}
	return ctsu
}

// ClearTrackingType clears the value of the "tracking_type" field.
func (ctsu *ContainerTrackingSuggestionUpdate) ClearTrackingType() *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.ClearTrackingType()
	return ctsu
}

// SetBic sets the "bic" field.
func (ctsu *ContainerTrackingSuggestionUpdate) SetBic(s string) *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.SetBic(s)
	return ctsu
}

// SetNillableBic sets the "bic" field if the given value is not nil.
func (ctsu *ContainerTrackingSuggestionUpdate) SetNillableBic(s *string) *ContainerTrackingSuggestionUpdate {
	if s != nil {
		ctsu.SetBic(*s)
	}
	return ctsu
}

// ClearBic clears the value of the "bic" field.
func (ctsu *ContainerTrackingSuggestionUpdate) ClearBic() *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.ClearBic()
	return ctsu
}

// SetSerial sets the "serial" field.
func (ctsu *ContainerTrackingSuggestionUpdate) SetSerial(s string) *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.SetSerial(s)
	return ctsu
}

// SetNillableSerial sets the "serial" field if the given value is not nil.
func (ctsu *ContainerTrackingSuggestionUpdate) SetNillableSerial(s *string) *ContainerTrackingSuggestionUpdate {
	if s != nil {
		ctsu.SetSerial(*s)
	}
	return ctsu
}

// ClearSerial clears the value of the "serial" field.
func (ctsu *ContainerTrackingSuggestionUpdate) ClearSerial() *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.ClearSerial()
	return ctsu
}

// SetChecksum sets the "checksum" field.
func (ctsu *ContainerTrackingSuggestionUpdate) SetChecksum(s string) *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.SetChecksum(s)
	return ctsu
}

// SetNillableChecksum sets the "checksum" field if the given value is not nil.
func (ctsu *ContainerTrackingSuggestionUpdate) SetNillableChecksum(s *string) *ContainerTrackingSuggestionUpdate {
	if s != nil {
		ctsu.SetChecksum(*s)
	}
	return ctsu
}

// ClearChecksum clears the value of the "checksum" field.
func (ctsu *ContainerTrackingSuggestionUpdate) ClearChecksum() *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.ClearChecksum()
	return ctsu
}

// SetImageURL sets the "image_url" field.
func (ctsu *ContainerTrackingSuggestionUpdate) SetImageURL(s string) *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.SetImageURL(s)
	return ctsu
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (ctsu *ContainerTrackingSuggestionUpdate) SetNillableImageURL(s *string) *ContainerTrackingSuggestionUpdate {
	if s != nil {
		ctsu.SetImageURL(*s)
	}
	return ctsu
}

// ClearImageURL clears the value of the "image_url" field.
func (ctsu *ContainerTrackingSuggestionUpdate) ClearImageURL() *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.ClearImageURL()
	return ctsu
}

// SetScore sets the "score" field.
func (ctsu *ContainerTrackingSuggestionUpdate) SetScore(f float32) *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.ResetScore()
	ctsu.mutation.SetScore(f)
	return ctsu
}

// AddScore adds f to the "score" field.
func (ctsu *ContainerTrackingSuggestionUpdate) AddScore(f float32) *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.AddScore(f)
	return ctsu
}

// SetCreatedAt sets the "created_at" field.
func (ctsu *ContainerTrackingSuggestionUpdate) SetCreatedAt(t time.Time) *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.SetCreatedAt(t)
	return ctsu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ctsu *ContainerTrackingSuggestionUpdate) SetNillableCreatedAt(t *time.Time) *ContainerTrackingSuggestionUpdate {
	if t != nil {
		ctsu.SetCreatedAt(*t)
	}
	return ctsu
}

// SetCam sets the "cam" edge to the CamSetting entity.
func (ctsu *ContainerTrackingSuggestionUpdate) SetCam(c *CamSetting) *ContainerTrackingSuggestionUpdate {
	return ctsu.SetCamID(c.ID)
}

// SetTracking sets the "tracking" edge to the ContainerTracking entity.
func (ctsu *ContainerTrackingSuggestionUpdate) SetTracking(c *ContainerTracking) *ContainerTrackingSuggestionUpdate {
	return ctsu.SetTrackingID(c.ID)
}

// Mutation returns the ContainerTrackingSuggestionMutation object of the builder.
func (ctsu *ContainerTrackingSuggestionUpdate) Mutation() *ContainerTrackingSuggestionMutation {
	return ctsu.mutation
}

// ClearCam clears the "cam" edge to the CamSetting entity.
func (ctsu *ContainerTrackingSuggestionUpdate) ClearCam() *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.ClearCam()
	return ctsu
}

// ClearTracking clears the "tracking" edge to the ContainerTracking entity.
func (ctsu *ContainerTrackingSuggestionUpdate) ClearTracking() *ContainerTrackingSuggestionUpdate {
	ctsu.mutation.ClearTracking()
	return ctsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctsu *ContainerTrackingSuggestionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ctsu.hooks) == 0 {
		affected, err = ctsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContainerTrackingSuggestionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctsu.mutation = mutation
			affected, err = ctsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ctsu.hooks) - 1; i >= 0; i-- {
			if ctsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ctsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctsu *ContainerTrackingSuggestionUpdate) SaveX(ctx context.Context) int {
	affected, err := ctsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctsu *ContainerTrackingSuggestionUpdate) Exec(ctx context.Context) error {
	_, err := ctsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctsu *ContainerTrackingSuggestionUpdate) ExecX(ctx context.Context) {
	if err := ctsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ctsu *ContainerTrackingSuggestionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   containertrackingsuggestion.Table,
			Columns: containertrackingsuggestion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: containertrackingsuggestion.FieldID,
			},
		},
	}
	if ps := ctsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctsu.mutation.ContainerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertrackingsuggestion.FieldContainerID,
		})
	}
	if ctsu.mutation.ContainerIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: containertrackingsuggestion.FieldContainerID,
		})
	}
	if value, ok := ctsu.mutation.Result(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertrackingsuggestion.FieldResult,
		})
	}
	if ctsu.mutation.ResultCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: containertrackingsuggestion.FieldResult,
		})
	}
	if value, ok := ctsu.mutation.TrackingType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertrackingsuggestion.FieldTrackingType,
		})
	}
	if ctsu.mutation.TrackingTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: containertrackingsuggestion.FieldTrackingType,
		})
	}
	if value, ok := ctsu.mutation.Bic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertrackingsuggestion.FieldBic,
		})
	}
	if ctsu.mutation.BicCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: containertrackingsuggestion.FieldBic,
		})
	}
	if value, ok := ctsu.mutation.Serial(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertrackingsuggestion.FieldSerial,
		})
	}
	if ctsu.mutation.SerialCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: containertrackingsuggestion.FieldSerial,
		})
	}
	if value, ok := ctsu.mutation.Checksum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertrackingsuggestion.FieldChecksum,
		})
	}
	if ctsu.mutation.ChecksumCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: containertrackingsuggestion.FieldChecksum,
		})
	}
	if value, ok := ctsu.mutation.ImageURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertrackingsuggestion.FieldImageURL,
		})
	}
	if ctsu.mutation.ImageURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: containertrackingsuggestion.FieldImageURL,
		})
	}
	if value, ok := ctsu.mutation.Score(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: containertrackingsuggestion.FieldScore,
		})
	}
	if value, ok := ctsu.mutation.AddedScore(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: containertrackingsuggestion.FieldScore,
		})
	}
	if value, ok := ctsu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: containertrackingsuggestion.FieldCreatedAt,
		})
	}
	if ctsu.mutation.CamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   containertrackingsuggestion.CamTable,
			Columns: []string{containertrackingsuggestion.CamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: camsetting.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctsu.mutation.CamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   containertrackingsuggestion.CamTable,
			Columns: []string{containertrackingsuggestion.CamColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ctsu.mutation.TrackingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   containertrackingsuggestion.TrackingTable,
			Columns: []string{containertrackingsuggestion.TrackingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: containertracking.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctsu.mutation.TrackingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   containertrackingsuggestion.TrackingTable,
			Columns: []string{containertrackingsuggestion.TrackingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: containertracking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ctsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{containertrackingsuggestion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ContainerTrackingSuggestionUpdateOne is the builder for updating a single ContainerTrackingSuggestion entity.
type ContainerTrackingSuggestionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContainerTrackingSuggestionMutation
}

// SetContainerID sets the "container_id" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetContainerID(s string) *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.SetContainerID(s)
	return ctsuo
}

// SetNillableContainerID sets the "container_id" field if the given value is not nil.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetNillableContainerID(s *string) *ContainerTrackingSuggestionUpdateOne {
	if s != nil {
		ctsuo.SetContainerID(*s)
	}
	return ctsuo
}

// ClearContainerID clears the value of the "container_id" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) ClearContainerID() *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.ClearContainerID()
	return ctsuo
}

// SetResult sets the "result" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetResult(s string) *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.SetResult(s)
	return ctsuo
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetNillableResult(s *string) *ContainerTrackingSuggestionUpdateOne {
	if s != nil {
		ctsuo.SetResult(*s)
	}
	return ctsuo
}

// ClearResult clears the value of the "result" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) ClearResult() *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.ClearResult()
	return ctsuo
}

// SetCamID sets the "cam_id" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetCamID(i int) *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.SetCamID(i)
	return ctsuo
}

// SetNillableCamID sets the "cam_id" field if the given value is not nil.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetNillableCamID(i *int) *ContainerTrackingSuggestionUpdateOne {
	if i != nil {
		ctsuo.SetCamID(*i)
	}
	return ctsuo
}

// ClearCamID clears the value of the "cam_id" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) ClearCamID() *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.ClearCamID()
	return ctsuo
}

// SetTrackingID sets the "tracking_id" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetTrackingID(i int) *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.SetTrackingID(i)
	return ctsuo
}

// SetNillableTrackingID sets the "tracking_id" field if the given value is not nil.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetNillableTrackingID(i *int) *ContainerTrackingSuggestionUpdateOne {
	if i != nil {
		ctsuo.SetTrackingID(*i)
	}
	return ctsuo
}

// ClearTrackingID clears the value of the "tracking_id" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) ClearTrackingID() *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.ClearTrackingID()
	return ctsuo
}

// SetTrackingType sets the "tracking_type" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetTrackingType(s string) *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.SetTrackingType(s)
	return ctsuo
}

// SetNillableTrackingType sets the "tracking_type" field if the given value is not nil.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetNillableTrackingType(s *string) *ContainerTrackingSuggestionUpdateOne {
	if s != nil {
		ctsuo.SetTrackingType(*s)
	}
	return ctsuo
}

// ClearTrackingType clears the value of the "tracking_type" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) ClearTrackingType() *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.ClearTrackingType()
	return ctsuo
}

// SetBic sets the "bic" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetBic(s string) *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.SetBic(s)
	return ctsuo
}

// SetNillableBic sets the "bic" field if the given value is not nil.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetNillableBic(s *string) *ContainerTrackingSuggestionUpdateOne {
	if s != nil {
		ctsuo.SetBic(*s)
	}
	return ctsuo
}

// ClearBic clears the value of the "bic" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) ClearBic() *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.ClearBic()
	return ctsuo
}

// SetSerial sets the "serial" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetSerial(s string) *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.SetSerial(s)
	return ctsuo
}

// SetNillableSerial sets the "serial" field if the given value is not nil.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetNillableSerial(s *string) *ContainerTrackingSuggestionUpdateOne {
	if s != nil {
		ctsuo.SetSerial(*s)
	}
	return ctsuo
}

// ClearSerial clears the value of the "serial" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) ClearSerial() *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.ClearSerial()
	return ctsuo
}

// SetChecksum sets the "checksum" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetChecksum(s string) *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.SetChecksum(s)
	return ctsuo
}

// SetNillableChecksum sets the "checksum" field if the given value is not nil.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetNillableChecksum(s *string) *ContainerTrackingSuggestionUpdateOne {
	if s != nil {
		ctsuo.SetChecksum(*s)
	}
	return ctsuo
}

// ClearChecksum clears the value of the "checksum" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) ClearChecksum() *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.ClearChecksum()
	return ctsuo
}

// SetImageURL sets the "image_url" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetImageURL(s string) *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.SetImageURL(s)
	return ctsuo
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetNillableImageURL(s *string) *ContainerTrackingSuggestionUpdateOne {
	if s != nil {
		ctsuo.SetImageURL(*s)
	}
	return ctsuo
}

// ClearImageURL clears the value of the "image_url" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) ClearImageURL() *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.ClearImageURL()
	return ctsuo
}

// SetScore sets the "score" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetScore(f float32) *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.ResetScore()
	ctsuo.mutation.SetScore(f)
	return ctsuo
}

// AddScore adds f to the "score" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) AddScore(f float32) *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.AddScore(f)
	return ctsuo
}

// SetCreatedAt sets the "created_at" field.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetCreatedAt(t time.Time) *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.SetCreatedAt(t)
	return ctsuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetNillableCreatedAt(t *time.Time) *ContainerTrackingSuggestionUpdateOne {
	if t != nil {
		ctsuo.SetCreatedAt(*t)
	}
	return ctsuo
}

// SetCam sets the "cam" edge to the CamSetting entity.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetCam(c *CamSetting) *ContainerTrackingSuggestionUpdateOne {
	return ctsuo.SetCamID(c.ID)
}

// SetTracking sets the "tracking" edge to the ContainerTracking entity.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SetTracking(c *ContainerTracking) *ContainerTrackingSuggestionUpdateOne {
	return ctsuo.SetTrackingID(c.ID)
}

// Mutation returns the ContainerTrackingSuggestionMutation object of the builder.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) Mutation() *ContainerTrackingSuggestionMutation {
	return ctsuo.mutation
}

// ClearCam clears the "cam" edge to the CamSetting entity.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) ClearCam() *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.ClearCam()
	return ctsuo
}

// ClearTracking clears the "tracking" edge to the ContainerTracking entity.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) ClearTracking() *ContainerTrackingSuggestionUpdateOne {
	ctsuo.mutation.ClearTracking()
	return ctsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) Select(field string, fields ...string) *ContainerTrackingSuggestionUpdateOne {
	ctsuo.fields = append([]string{field}, fields...)
	return ctsuo
}

// Save executes the query and returns the updated ContainerTrackingSuggestion entity.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) Save(ctx context.Context) (*ContainerTrackingSuggestion, error) {
	var (
		err  error
		node *ContainerTrackingSuggestion
	)
	if len(ctsuo.hooks) == 0 {
		node, err = ctsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContainerTrackingSuggestionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctsuo.mutation = mutation
			node, err = ctsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ctsuo.hooks) - 1; i >= 0; i-- {
			if ctsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ctsuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctsuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) SaveX(ctx context.Context) *ContainerTrackingSuggestion {
	node, err := ctsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) Exec(ctx context.Context) error {
	_, err := ctsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctsuo *ContainerTrackingSuggestionUpdateOne) ExecX(ctx context.Context) {
	if err := ctsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ctsuo *ContainerTrackingSuggestionUpdateOne) sqlSave(ctx context.Context) (_node *ContainerTrackingSuggestion, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   containertrackingsuggestion.Table,
			Columns: containertrackingsuggestion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: containertrackingsuggestion.FieldID,
			},
		},
	}
	id, ok := ctsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ContainerTrackingSuggestion.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ctsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, containertrackingsuggestion.FieldID)
		for _, f := range fields {
			if !containertrackingsuggestion.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != containertrackingsuggestion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctsuo.mutation.ContainerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertrackingsuggestion.FieldContainerID,
		})
	}
	if ctsuo.mutation.ContainerIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: containertrackingsuggestion.FieldContainerID,
		})
	}
	if value, ok := ctsuo.mutation.Result(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertrackingsuggestion.FieldResult,
		})
	}
	if ctsuo.mutation.ResultCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: containertrackingsuggestion.FieldResult,
		})
	}
	if value, ok := ctsuo.mutation.TrackingType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertrackingsuggestion.FieldTrackingType,
		})
	}
	if ctsuo.mutation.TrackingTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: containertrackingsuggestion.FieldTrackingType,
		})
	}
	if value, ok := ctsuo.mutation.Bic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertrackingsuggestion.FieldBic,
		})
	}
	if ctsuo.mutation.BicCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: containertrackingsuggestion.FieldBic,
		})
	}
	if value, ok := ctsuo.mutation.Serial(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertrackingsuggestion.FieldSerial,
		})
	}
	if ctsuo.mutation.SerialCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: containertrackingsuggestion.FieldSerial,
		})
	}
	if value, ok := ctsuo.mutation.Checksum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertrackingsuggestion.FieldChecksum,
		})
	}
	if ctsuo.mutation.ChecksumCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: containertrackingsuggestion.FieldChecksum,
		})
	}
	if value, ok := ctsuo.mutation.ImageURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertrackingsuggestion.FieldImageURL,
		})
	}
	if ctsuo.mutation.ImageURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: containertrackingsuggestion.FieldImageURL,
		})
	}
	if value, ok := ctsuo.mutation.Score(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: containertrackingsuggestion.FieldScore,
		})
	}
	if value, ok := ctsuo.mutation.AddedScore(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: containertrackingsuggestion.FieldScore,
		})
	}
	if value, ok := ctsuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: containertrackingsuggestion.FieldCreatedAt,
		})
	}
	if ctsuo.mutation.CamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   containertrackingsuggestion.CamTable,
			Columns: []string{containertrackingsuggestion.CamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: camsetting.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctsuo.mutation.CamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   containertrackingsuggestion.CamTable,
			Columns: []string{containertrackingsuggestion.CamColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ctsuo.mutation.TrackingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   containertrackingsuggestion.TrackingTable,
			Columns: []string{containertrackingsuggestion.TrackingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: containertracking.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctsuo.mutation.TrackingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   containertrackingsuggestion.TrackingTable,
			Columns: []string{containertrackingsuggestion.TrackingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: containertracking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ContainerTrackingSuggestion{config: ctsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{containertrackingsuggestion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
