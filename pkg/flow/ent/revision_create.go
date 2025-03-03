// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/ref"
	"github.com/direktiv/direktiv/pkg/flow/ent/revision"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// RevisionCreate is the builder for creating a Revision entity.
type RevisionCreate struct {
	config
	mutation *RevisionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (rc *RevisionCreate) SetCreatedAt(t time.Time) *RevisionCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RevisionCreate) SetNillableCreatedAt(t *time.Time) *RevisionCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetHash sets the "hash" field.
func (rc *RevisionCreate) SetHash(s string) *RevisionCreate {
	rc.mutation.SetHash(s)
	return rc
}

// SetSource sets the "source" field.
func (rc *RevisionCreate) SetSource(b []byte) *RevisionCreate {
	rc.mutation.SetSource(b)
	return rc
}

// SetMetadata sets the "metadata" field.
func (rc *RevisionCreate) SetMetadata(m map[string]interface{}) *RevisionCreate {
	rc.mutation.SetMetadata(m)
	return rc
}

// SetID sets the "id" field.
func (rc *RevisionCreate) SetID(u uuid.UUID) *RevisionCreate {
	rc.mutation.SetID(u)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *RevisionCreate) SetNillableID(u *uuid.UUID) *RevisionCreate {
	if u != nil {
		rc.SetID(*u)
	}
	return rc
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (rc *RevisionCreate) SetWorkflowID(id uuid.UUID) *RevisionCreate {
	rc.mutation.SetWorkflowID(id)
	return rc
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (rc *RevisionCreate) SetWorkflow(w *Workflow) *RevisionCreate {
	return rc.SetWorkflowID(w.ID)
}

// AddRefIDs adds the "refs" edge to the Ref entity by IDs.
func (rc *RevisionCreate) AddRefIDs(ids ...uuid.UUID) *RevisionCreate {
	rc.mutation.AddRefIDs(ids...)
	return rc
}

// AddRefs adds the "refs" edges to the Ref entity.
func (rc *RevisionCreate) AddRefs(r ...*Ref) *RevisionCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rc.AddRefIDs(ids...)
}

// AddInstanceIDs adds the "instances" edge to the Instance entity by IDs.
func (rc *RevisionCreate) AddInstanceIDs(ids ...uuid.UUID) *RevisionCreate {
	rc.mutation.AddInstanceIDs(ids...)
	return rc
}

// AddInstances adds the "instances" edges to the Instance entity.
func (rc *RevisionCreate) AddInstances(i ...*Instance) *RevisionCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return rc.AddInstanceIDs(ids...)
}

// Mutation returns the RevisionMutation object of the builder.
func (rc *RevisionCreate) Mutation() *RevisionMutation {
	return rc.mutation
}

// Save creates the Revision in the database.
func (rc *RevisionCreate) Save(ctx context.Context) (*Revision, error) {
	var (
		err  error
		node *Revision
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RevisionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RevisionCreate) SaveX(ctx context.Context) *Revision {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RevisionCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RevisionCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RevisionCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := revision.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.ID(); !ok {
		v := revision.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RevisionCreate) check() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Revision.created_at"`)}
	}
	if _, ok := rc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "Revision.hash"`)}
	}
	if _, ok := rc.mutation.Source(); !ok {
		return &ValidationError{Name: "source", err: errors.New(`ent: missing required field "Revision.source"`)}
	}
	if _, ok := rc.mutation.Metadata(); !ok {
		return &ValidationError{Name: "metadata", err: errors.New(`ent: missing required field "Revision.metadata"`)}
	}
	if _, ok := rc.mutation.WorkflowID(); !ok {
		return &ValidationError{Name: "workflow", err: errors.New(`ent: missing required edge "Revision.workflow"`)}
	}
	return nil
}

func (rc *RevisionCreate) sqlSave(ctx context.Context) (*Revision, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (rc *RevisionCreate) createSpec() (*Revision, *sqlgraph.CreateSpec) {
	var (
		_node = &Revision{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: revision.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: revision.FieldID,
			},
		}
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: revision.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.Hash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: revision.FieldHash,
		})
		_node.Hash = value
	}
	if value, ok := rc.mutation.Source(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: revision.FieldSource,
		})
		_node.Source = value
	}
	if value, ok := rc.mutation.Metadata(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: revision.FieldMetadata,
		})
		_node.Metadata = value
	}
	if nodes := rc.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   revision.WorkflowTable,
			Columns: []string{revision.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.workflow_revisions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.RefsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.RefsTable,
			Columns: []string{revision.RefsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ref.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.InstancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   revision.InstancesTable,
			Columns: []string{revision.InstancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
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

// RevisionCreateBulk is the builder for creating many Revision entities in bulk.
type RevisionCreateBulk struct {
	config
	builders []*RevisionCreate
}

// Save creates the Revision entities in the database.
func (rcb *RevisionCreateBulk) Save(ctx context.Context) ([]*Revision, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Revision, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RevisionMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RevisionCreateBulk) SaveX(ctx context.Context) []*Revision {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RevisionCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RevisionCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
