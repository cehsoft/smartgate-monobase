// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/containertracking"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/containertrackingsuggestion"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/predicate"
)

// ContainerTrackingUpdate is the builder for updating ContainerTracking entities.
type ContainerTrackingUpdate struct {
	config
	hooks    []Hook
	mutation *ContainerTrackingMutation
}

// Where appends a list predicates to the ContainerTrackingUpdate builder.
func (ctu *ContainerTrackingUpdate) Where(ps ...predicate.ContainerTracking) *ContainerTrackingUpdate {
	ctu.mutation.Where(ps...)
	return ctu
}

// SetContainerID sets the "container_id" field.
func (ctu *ContainerTrackingUpdate) SetContainerID(s string) *ContainerTrackingUpdate {
	ctu.mutation.SetContainerID(s)
	return ctu
}

// SetNillableContainerID sets the "container_id" field if the given value is not nil.
func (ctu *ContainerTrackingUpdate) SetNillableContainerID(s *string) *ContainerTrackingUpdate {
	if s != nil {
		ctu.SetContainerID(*s)
	}
	return ctu
}

// ClearContainerID clears the value of the "container_id" field.
func (ctu *ContainerTrackingUpdate) ClearContainerID() *ContainerTrackingUpdate {
	ctu.mutation.ClearContainerID()
	return ctu
}

// SetSessionID sets the "session_id" field.
func (ctu *ContainerTrackingUpdate) SetSessionID(s string) *ContainerTrackingUpdate {
	ctu.mutation.SetSessionID(s)
	return ctu
}

// SetNillableSessionID sets the "session_id" field if the given value is not nil.
func (ctu *ContainerTrackingUpdate) SetNillableSessionID(s *string) *ContainerTrackingUpdate {
	if s != nil {
		ctu.SetSessionID(*s)
	}
	return ctu
}

// ClearSessionID clears the value of the "session_id" field.
func (ctu *ContainerTrackingUpdate) ClearSessionID() *ContainerTrackingUpdate {
	ctu.mutation.ClearSessionID()
	return ctu
}

// SetCreatedAt sets the "created_at" field.
func (ctu *ContainerTrackingUpdate) SetCreatedAt(t time.Time) *ContainerTrackingUpdate {
	ctu.mutation.SetCreatedAt(t)
	return ctu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ctu *ContainerTrackingUpdate) SetNillableCreatedAt(t *time.Time) *ContainerTrackingUpdate {
	if t != nil {
		ctu.SetCreatedAt(*t)
	}
	return ctu
}

// AddSuggestionIDs adds the "suggestions" edge to the ContainerTrackingSuggestion entity by IDs.
func (ctu *ContainerTrackingUpdate) AddSuggestionIDs(ids ...int) *ContainerTrackingUpdate {
	ctu.mutation.AddSuggestionIDs(ids...)
	return ctu
}

// AddSuggestions adds the "suggestions" edges to the ContainerTrackingSuggestion entity.
func (ctu *ContainerTrackingUpdate) AddSuggestions(c ...*ContainerTrackingSuggestion) *ContainerTrackingUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctu.AddSuggestionIDs(ids...)
}

// Mutation returns the ContainerTrackingMutation object of the builder.
func (ctu *ContainerTrackingUpdate) Mutation() *ContainerTrackingMutation {
	return ctu.mutation
}

// ClearSuggestions clears all "suggestions" edges to the ContainerTrackingSuggestion entity.
func (ctu *ContainerTrackingUpdate) ClearSuggestions() *ContainerTrackingUpdate {
	ctu.mutation.ClearSuggestions()
	return ctu
}

// RemoveSuggestionIDs removes the "suggestions" edge to ContainerTrackingSuggestion entities by IDs.
func (ctu *ContainerTrackingUpdate) RemoveSuggestionIDs(ids ...int) *ContainerTrackingUpdate {
	ctu.mutation.RemoveSuggestionIDs(ids...)
	return ctu
}

// RemoveSuggestions removes "suggestions" edges to ContainerTrackingSuggestion entities.
func (ctu *ContainerTrackingUpdate) RemoveSuggestions(c ...*ContainerTrackingSuggestion) *ContainerTrackingUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctu.RemoveSuggestionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctu *ContainerTrackingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ctu.hooks) == 0 {
		affected, err = ctu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContainerTrackingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctu.mutation = mutation
			affected, err = ctu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ctu.hooks) - 1; i >= 0; i-- {
			if ctu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ctu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctu *ContainerTrackingUpdate) SaveX(ctx context.Context) int {
	affected, err := ctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctu *ContainerTrackingUpdate) Exec(ctx context.Context) error {
	_, err := ctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctu *ContainerTrackingUpdate) ExecX(ctx context.Context) {
	if err := ctu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ctu *ContainerTrackingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   containertracking.Table,
			Columns: containertracking.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: containertracking.FieldID,
			},
		},
	}
	if ps := ctu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctu.mutation.ContainerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertracking.FieldContainerID,
		})
	}
	if ctu.mutation.ContainerIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: containertracking.FieldContainerID,
		})
	}
	if value, ok := ctu.mutation.SessionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertracking.FieldSessionID,
		})
	}
	if ctu.mutation.SessionIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: containertracking.FieldSessionID,
		})
	}
	if value, ok := ctu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: containertracking.FieldCreatedAt,
		})
	}
	if ctu.mutation.SuggestionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   containertracking.SuggestionsTable,
			Columns: []string{containertracking.SuggestionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: containertrackingsuggestion.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.RemovedSuggestionsIDs(); len(nodes) > 0 && !ctu.mutation.SuggestionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   containertracking.SuggestionsTable,
			Columns: []string{containertracking.SuggestionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: containertrackingsuggestion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.SuggestionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   containertracking.SuggestionsTable,
			Columns: []string{containertracking.SuggestionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: containertrackingsuggestion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{containertracking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ContainerTrackingUpdateOne is the builder for updating a single ContainerTracking entity.
type ContainerTrackingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContainerTrackingMutation
}

// SetContainerID sets the "container_id" field.
func (ctuo *ContainerTrackingUpdateOne) SetContainerID(s string) *ContainerTrackingUpdateOne {
	ctuo.mutation.SetContainerID(s)
	return ctuo
}

// SetNillableContainerID sets the "container_id" field if the given value is not nil.
func (ctuo *ContainerTrackingUpdateOne) SetNillableContainerID(s *string) *ContainerTrackingUpdateOne {
	if s != nil {
		ctuo.SetContainerID(*s)
	}
	return ctuo
}

// ClearContainerID clears the value of the "container_id" field.
func (ctuo *ContainerTrackingUpdateOne) ClearContainerID() *ContainerTrackingUpdateOne {
	ctuo.mutation.ClearContainerID()
	return ctuo
}

// SetSessionID sets the "session_id" field.
func (ctuo *ContainerTrackingUpdateOne) SetSessionID(s string) *ContainerTrackingUpdateOne {
	ctuo.mutation.SetSessionID(s)
	return ctuo
}

// SetNillableSessionID sets the "session_id" field if the given value is not nil.
func (ctuo *ContainerTrackingUpdateOne) SetNillableSessionID(s *string) *ContainerTrackingUpdateOne {
	if s != nil {
		ctuo.SetSessionID(*s)
	}
	return ctuo
}

// ClearSessionID clears the value of the "session_id" field.
func (ctuo *ContainerTrackingUpdateOne) ClearSessionID() *ContainerTrackingUpdateOne {
	ctuo.mutation.ClearSessionID()
	return ctuo
}

// SetCreatedAt sets the "created_at" field.
func (ctuo *ContainerTrackingUpdateOne) SetCreatedAt(t time.Time) *ContainerTrackingUpdateOne {
	ctuo.mutation.SetCreatedAt(t)
	return ctuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ctuo *ContainerTrackingUpdateOne) SetNillableCreatedAt(t *time.Time) *ContainerTrackingUpdateOne {
	if t != nil {
		ctuo.SetCreatedAt(*t)
	}
	return ctuo
}

// AddSuggestionIDs adds the "suggestions" edge to the ContainerTrackingSuggestion entity by IDs.
func (ctuo *ContainerTrackingUpdateOne) AddSuggestionIDs(ids ...int) *ContainerTrackingUpdateOne {
	ctuo.mutation.AddSuggestionIDs(ids...)
	return ctuo
}

// AddSuggestions adds the "suggestions" edges to the ContainerTrackingSuggestion entity.
func (ctuo *ContainerTrackingUpdateOne) AddSuggestions(c ...*ContainerTrackingSuggestion) *ContainerTrackingUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctuo.AddSuggestionIDs(ids...)
}

// Mutation returns the ContainerTrackingMutation object of the builder.
func (ctuo *ContainerTrackingUpdateOne) Mutation() *ContainerTrackingMutation {
	return ctuo.mutation
}

// ClearSuggestions clears all "suggestions" edges to the ContainerTrackingSuggestion entity.
func (ctuo *ContainerTrackingUpdateOne) ClearSuggestions() *ContainerTrackingUpdateOne {
	ctuo.mutation.ClearSuggestions()
	return ctuo
}

// RemoveSuggestionIDs removes the "suggestions" edge to ContainerTrackingSuggestion entities by IDs.
func (ctuo *ContainerTrackingUpdateOne) RemoveSuggestionIDs(ids ...int) *ContainerTrackingUpdateOne {
	ctuo.mutation.RemoveSuggestionIDs(ids...)
	return ctuo
}

// RemoveSuggestions removes "suggestions" edges to ContainerTrackingSuggestion entities.
func (ctuo *ContainerTrackingUpdateOne) RemoveSuggestions(c ...*ContainerTrackingSuggestion) *ContainerTrackingUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctuo.RemoveSuggestionIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctuo *ContainerTrackingUpdateOne) Select(field string, fields ...string) *ContainerTrackingUpdateOne {
	ctuo.fields = append([]string{field}, fields...)
	return ctuo
}

// Save executes the query and returns the updated ContainerTracking entity.
func (ctuo *ContainerTrackingUpdateOne) Save(ctx context.Context) (*ContainerTracking, error) {
	var (
		err  error
		node *ContainerTracking
	)
	if len(ctuo.hooks) == 0 {
		node, err = ctuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContainerTrackingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctuo.mutation = mutation
			node, err = ctuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ctuo.hooks) - 1; i >= 0; i-- {
			if ctuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ctuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctuo *ContainerTrackingUpdateOne) SaveX(ctx context.Context) *ContainerTracking {
	node, err := ctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctuo *ContainerTrackingUpdateOne) Exec(ctx context.Context) error {
	_, err := ctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctuo *ContainerTrackingUpdateOne) ExecX(ctx context.Context) {
	if err := ctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ctuo *ContainerTrackingUpdateOne) sqlSave(ctx context.Context) (_node *ContainerTracking, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   containertracking.Table,
			Columns: containertracking.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: containertracking.FieldID,
			},
		},
	}
	id, ok := ctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ContainerTracking.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, containertracking.FieldID)
		for _, f := range fields {
			if !containertracking.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != containertracking.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctuo.mutation.ContainerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertracking.FieldContainerID,
		})
	}
	if ctuo.mutation.ContainerIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: containertracking.FieldContainerID,
		})
	}
	if value, ok := ctuo.mutation.SessionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: containertracking.FieldSessionID,
		})
	}
	if ctuo.mutation.SessionIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: containertracking.FieldSessionID,
		})
	}
	if value, ok := ctuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: containertracking.FieldCreatedAt,
		})
	}
	if ctuo.mutation.SuggestionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   containertracking.SuggestionsTable,
			Columns: []string{containertracking.SuggestionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: containertrackingsuggestion.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.RemovedSuggestionsIDs(); len(nodes) > 0 && !ctuo.mutation.SuggestionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   containertracking.SuggestionsTable,
			Columns: []string{containertracking.SuggestionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: containertrackingsuggestion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.SuggestionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   containertracking.SuggestionsTable,
			Columns: []string{containertracking.SuggestionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: containertrackingsuggestion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ContainerTracking{config: ctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{containertracking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
