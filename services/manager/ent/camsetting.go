// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/camsetting"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/lane"
)

// CamSetting is the model entity for the CamSetting schema.
type CamSetting struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// LaneID holds the value of the "lane_id" field.
	LaneID int `json:"lane_id,omitempty"`
	// RtspURL holds the value of the "rtsp_url" field.
	RtspURL string `json:"rtsp_url,omitempty"`
	// WebrtcURL holds the value of the "webrtc_url" field.
	WebrtcURL string `json:"webrtc_url,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CamSettingQuery when eager-loading is set.
	Edges CamSettingEdges `json:"edges"`
}

// CamSettingEdges holds the relations/edges for other nodes in the graph.
type CamSettingEdges struct {
	// Lane holds the value of the lane edge.
	Lane *Lane `json:"lane,omitempty"`
	// Suggestions holds the value of the suggestions edge.
	Suggestions []*ContainerTrackingSuggestion `json:"suggestions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// LaneOrErr returns the Lane value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CamSettingEdges) LaneOrErr() (*Lane, error) {
	if e.loadedTypes[0] {
		if e.Lane == nil {
			// The edge lane was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: lane.Label}
		}
		return e.Lane, nil
	}
	return nil, &NotLoadedError{edge: "lane"}
}

// SuggestionsOrErr returns the Suggestions value or an error if the edge
// was not loaded in eager-loading.
func (e CamSettingEdges) SuggestionsOrErr() ([]*ContainerTrackingSuggestion, error) {
	if e.loadedTypes[1] {
		return e.Suggestions, nil
	}
	return nil, &NotLoadedError{edge: "suggestions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CamSetting) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case camsetting.FieldID, camsetting.FieldLaneID:
			values[i] = new(sql.NullInt64)
		case camsetting.FieldName, camsetting.FieldRtspURL, camsetting.FieldWebrtcURL:
			values[i] = new(sql.NullString)
		case camsetting.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CamSetting", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CamSetting fields.
func (cs *CamSetting) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case camsetting.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cs.ID = int(value.Int64)
		case camsetting.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cs.Name = value.String
			}
		case camsetting.FieldLaneID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field lane_id", values[i])
			} else if value.Valid {
				cs.LaneID = int(value.Int64)
			}
		case camsetting.FieldRtspURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field rtsp_url", values[i])
			} else if value.Valid {
				cs.RtspURL = value.String
			}
		case camsetting.FieldWebrtcURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field webrtc_url", values[i])
			} else if value.Valid {
				cs.WebrtcURL = value.String
			}
		case camsetting.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cs.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryLane queries the "lane" edge of the CamSetting entity.
func (cs *CamSetting) QueryLane() *LaneQuery {
	return (&CamSettingClient{config: cs.config}).QueryLane(cs)
}

// QuerySuggestions queries the "suggestions" edge of the CamSetting entity.
func (cs *CamSetting) QuerySuggestions() *ContainerTrackingSuggestionQuery {
	return (&CamSettingClient{config: cs.config}).QuerySuggestions(cs)
}

// Update returns a builder for updating this CamSetting.
// Note that you need to call CamSetting.Unwrap() before calling this method if this CamSetting
// was returned from a transaction, and the transaction was committed or rolled back.
func (cs *CamSetting) Update() *CamSettingUpdateOne {
	return (&CamSettingClient{config: cs.config}).UpdateOne(cs)
}

// Unwrap unwraps the CamSetting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cs *CamSetting) Unwrap() *CamSetting {
	tx, ok := cs.config.driver.(*txDriver)
	if !ok {
		panic("ent: CamSetting is not a transactional entity")
	}
	cs.config.driver = tx.drv
	return cs
}

// String implements the fmt.Stringer.
func (cs *CamSetting) String() string {
	var builder strings.Builder
	builder.WriteString("CamSetting(")
	builder.WriteString(fmt.Sprintf("id=%v", cs.ID))
	builder.WriteString(", name=")
	builder.WriteString(cs.Name)
	builder.WriteString(", lane_id=")
	builder.WriteString(fmt.Sprintf("%v", cs.LaneID))
	builder.WriteString(", rtsp_url=")
	builder.WriteString(cs.RtspURL)
	builder.WriteString(", webrtc_url=")
	builder.WriteString(cs.WebrtcURL)
	builder.WriteString(", created_at=")
	builder.WriteString(cs.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// CamSettings is a parsable slice of CamSetting.
type CamSettings []*CamSetting

func (cs CamSettings) config(cfg config) {
	for _i := range cs {
		cs[_i].config = cfg
	}
}
