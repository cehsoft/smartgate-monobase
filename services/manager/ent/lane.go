// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/gate"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/lane"
)

// Lane is the model entity for the Lane schema.
type Lane struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// GateID holds the value of the "gate_id" field.
	GateID int `json:"gate_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LaneQuery when eager-loading is set.
	Edges LaneEdges `json:"edges"`
}

// LaneEdges holds the relations/edges for other nodes in the graph.
type LaneEdges struct {
	// Cams holds the value of the cams edge.
	Cams []*CamSetting `json:"cams,omitempty"`
	// Gate holds the value of the gate edge.
	Gate *Gate `json:"gate,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CamsOrErr returns the Cams value or an error if the edge
// was not loaded in eager-loading.
func (e LaneEdges) CamsOrErr() ([]*CamSetting, error) {
	if e.loadedTypes[0] {
		return e.Cams, nil
	}
	return nil, &NotLoadedError{edge: "cams"}
}

// GateOrErr returns the Gate value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LaneEdges) GateOrErr() (*Gate, error) {
	if e.loadedTypes[1] {
		if e.Gate == nil {
			// The edge gate was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: gate.Label}
		}
		return e.Gate, nil
	}
	return nil, &NotLoadedError{edge: "gate"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Lane) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case lane.FieldID, lane.FieldGateID:
			values[i] = new(sql.NullInt64)
		case lane.FieldName:
			values[i] = new(sql.NullString)
		case lane.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Lane", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Lane fields.
func (l *Lane) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case lane.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case lane.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				l.Name = value.String
			}
		case lane.FieldGateID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gate_id", values[i])
			} else if value.Valid {
				l.GateID = int(value.Int64)
			}
		case lane.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryCams queries the "cams" edge of the Lane entity.
func (l *Lane) QueryCams() *CamSettingQuery {
	return (&LaneClient{config: l.config}).QueryCams(l)
}

// QueryGate queries the "gate" edge of the Lane entity.
func (l *Lane) QueryGate() *GateQuery {
	return (&LaneClient{config: l.config}).QueryGate(l)
}

// Update returns a builder for updating this Lane.
// Note that you need to call Lane.Unwrap() before calling this method if this Lane
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Lane) Update() *LaneUpdateOne {
	return (&LaneClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the Lane entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Lane) Unwrap() *Lane {
	tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Lane is not a transactional entity")
	}
	l.config.driver = tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Lane) String() string {
	var builder strings.Builder
	builder.WriteString("Lane(")
	builder.WriteString(fmt.Sprintf("id=%v", l.ID))
	builder.WriteString(", name=")
	builder.WriteString(l.Name)
	builder.WriteString(", gate_id=")
	builder.WriteString(fmt.Sprintf("%v", l.GateID))
	builder.WriteString(", created_at=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Lanes is a parsable slice of Lane.
type Lanes []*Lane

func (l Lanes) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}