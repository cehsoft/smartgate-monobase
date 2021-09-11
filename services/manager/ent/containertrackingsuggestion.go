// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/camsetting"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/containertracking"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/containertrackingsuggestion"
)

// ContainerTrackingSuggestion is the model entity for the ContainerTrackingSuggestion schema.
type ContainerTrackingSuggestion struct {
	config `db:"-" json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ContainerID holds the value of the "container_id" field.
	ContainerID string `json:"container_id,omitempty" db:"container_id"`
	// Result holds the value of the "result" field.
	Result string `json:"result,omitempty"`
	// CamID holds the value of the "cam_id" field.
	CamID int `json:"cam_id,omitempty" db:"cam_id"`
	// TrackingID holds the value of the "tracking_id" field.
	TrackingID int `json:"tracking_id,omitempty" db:"tracking_id"`
	// TrackingType holds the value of the "tracking_type" field.
	TrackingType string `json:"tracking_type,omitempty" db:"tracking_type"`
	// Bic holds the value of the "bic" field.
	Bic string `json:"bic,omitempty"`
	// Serial holds the value of the "serial" field.
	Serial string `json:"serial,omitempty"`
	// Checksum holds the value of the "checksum" field.
	Checksum string `json:"checksum,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty" db:"image_url"`
	// Score holds the value of the "score" field.
	Score float32 `json:"score,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ContainerTrackingSuggestionQuery when eager-loading is set.
	Edges ContainerTrackingSuggestionEdges `json:"edges"`
}

// ContainerTrackingSuggestionEdges holds the relations/edges for other nodes in the graph.
type ContainerTrackingSuggestionEdges struct {
	// Cam holds the value of the cam edge.
	Cam *CamSetting `json:"cam,omitempty"`
	// Tracking holds the value of the tracking edge.
	Tracking *ContainerTracking `json:"tracking,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CamOrErr returns the Cam value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContainerTrackingSuggestionEdges) CamOrErr() (*CamSetting, error) {
	if e.loadedTypes[0] {
		if e.Cam == nil {
			// The edge cam was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: camsetting.Label}
		}
		return e.Cam, nil
	}
	return nil, &NotLoadedError{edge: "cam"}
}

// TrackingOrErr returns the Tracking value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContainerTrackingSuggestionEdges) TrackingOrErr() (*ContainerTracking, error) {
	if e.loadedTypes[1] {
		if e.Tracking == nil {
			// The edge tracking was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: containertracking.Label}
		}
		return e.Tracking, nil
	}
	return nil, &NotLoadedError{edge: "tracking"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ContainerTrackingSuggestion) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case containertrackingsuggestion.FieldScore:
			values[i] = new(sql.NullFloat64)
		case containertrackingsuggestion.FieldID, containertrackingsuggestion.FieldCamID, containertrackingsuggestion.FieldTrackingID:
			values[i] = new(sql.NullInt64)
		case containertrackingsuggestion.FieldContainerID, containertrackingsuggestion.FieldResult, containertrackingsuggestion.FieldTrackingType, containertrackingsuggestion.FieldBic, containertrackingsuggestion.FieldSerial, containertrackingsuggestion.FieldChecksum, containertrackingsuggestion.FieldImageURL:
			values[i] = new(sql.NullString)
		case containertrackingsuggestion.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ContainerTrackingSuggestion", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ContainerTrackingSuggestion fields.
func (cts *ContainerTrackingSuggestion) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case containertrackingsuggestion.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cts.ID = int(value.Int64)
		case containertrackingsuggestion.FieldContainerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field container_id", values[i])
			} else if value.Valid {
				cts.ContainerID = value.String
			}
		case containertrackingsuggestion.FieldResult:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value.Valid {
				cts.Result = value.String
			}
		case containertrackingsuggestion.FieldCamID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cam_id", values[i])
			} else if value.Valid {
				cts.CamID = int(value.Int64)
			}
		case containertrackingsuggestion.FieldTrackingID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tracking_id", values[i])
			} else if value.Valid {
				cts.TrackingID = int(value.Int64)
			}
		case containertrackingsuggestion.FieldTrackingType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tracking_type", values[i])
			} else if value.Valid {
				cts.TrackingType = value.String
			}
		case containertrackingsuggestion.FieldBic:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bic", values[i])
			} else if value.Valid {
				cts.Bic = value.String
			}
		case containertrackingsuggestion.FieldSerial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial", values[i])
			} else if value.Valid {
				cts.Serial = value.String
			}
		case containertrackingsuggestion.FieldChecksum:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field checksum", values[i])
			} else if value.Valid {
				cts.Checksum = value.String
			}
		case containertrackingsuggestion.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				cts.ImageURL = value.String
			}
		case containertrackingsuggestion.FieldScore:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field score", values[i])
			} else if value.Valid {
				cts.Score = float32(value.Float64)
			}
		case containertrackingsuggestion.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cts.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryCam queries the "cam" edge of the ContainerTrackingSuggestion entity.
func (cts *ContainerTrackingSuggestion) QueryCam() *CamSettingQuery {
	return (&ContainerTrackingSuggestionClient{config: cts.config}).QueryCam(cts)
}

// QueryTracking queries the "tracking" edge of the ContainerTrackingSuggestion entity.
func (cts *ContainerTrackingSuggestion) QueryTracking() *ContainerTrackingQuery {
	return (&ContainerTrackingSuggestionClient{config: cts.config}).QueryTracking(cts)
}

// Update returns a builder for updating this ContainerTrackingSuggestion.
// Note that you need to call ContainerTrackingSuggestion.Unwrap() before calling this method if this ContainerTrackingSuggestion
// was returned from a transaction, and the transaction was committed or rolled back.
func (cts *ContainerTrackingSuggestion) Update() *ContainerTrackingSuggestionUpdateOne {
	return (&ContainerTrackingSuggestionClient{config: cts.config}).UpdateOne(cts)
}

// Unwrap unwraps the ContainerTrackingSuggestion entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cts *ContainerTrackingSuggestion) Unwrap() *ContainerTrackingSuggestion {
	tx, ok := cts.config.driver.(*txDriver)
	if !ok {
		panic("ent: ContainerTrackingSuggestion is not a transactional entity")
	}
	cts.config.driver = tx.drv
	return cts
}

// String implements the fmt.Stringer.
func (cts *ContainerTrackingSuggestion) String() string {
	var builder strings.Builder
	builder.WriteString("ContainerTrackingSuggestion(")
	builder.WriteString(fmt.Sprintf("id=%v", cts.ID))
	builder.WriteString(", container_id=")
	builder.WriteString(cts.ContainerID)
	builder.WriteString(", result=")
	builder.WriteString(cts.Result)
	builder.WriteString(", cam_id=")
	builder.WriteString(fmt.Sprintf("%v", cts.CamID))
	builder.WriteString(", tracking_id=")
	builder.WriteString(fmt.Sprintf("%v", cts.TrackingID))
	builder.WriteString(", tracking_type=")
	builder.WriteString(cts.TrackingType)
	builder.WriteString(", bic=")
	builder.WriteString(cts.Bic)
	builder.WriteString(", serial=")
	builder.WriteString(cts.Serial)
	builder.WriteString(", checksum=")
	builder.WriteString(cts.Checksum)
	builder.WriteString(", image_url=")
	builder.WriteString(cts.ImageURL)
	builder.WriteString(", score=")
	builder.WriteString(fmt.Sprintf("%v", cts.Score))
	builder.WriteString(", created_at=")
	builder.WriteString(cts.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ContainerTrackingSuggestions is a parsable slice of ContainerTrackingSuggestion.
type ContainerTrackingSuggestions []*ContainerTrackingSuggestion

func (cts ContainerTrackingSuggestions) config(cfg config) {
	for _i := range cts {
		cts[_i].config = cfg
	}
}
