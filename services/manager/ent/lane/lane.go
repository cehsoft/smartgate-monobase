// Code generated by entc, DO NOT EDIT.

package lane

import (
	"time"
)

const (
	// Label holds the string label denoting the lane type in the database.
	Label = "lane"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldGateID holds the string denoting the gate_id field in the database.
	FieldGateID = "gate_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeCams holds the string denoting the cams edge name in mutations.
	EdgeCams = "cams"
	// EdgeGate holds the string denoting the gate edge name in mutations.
	EdgeGate = "gate"
	// Table holds the table name of the lane in the database.
	Table = "lanes"
	// CamsTable is the table that holds the cams relation/edge.
	CamsTable = "cam_settings"
	// CamsInverseTable is the table name for the CamSetting entity.
	// It exists in this package in order to avoid circular dependency with the "camsetting" package.
	CamsInverseTable = "cam_settings"
	// CamsColumn is the table column denoting the cams relation/edge.
	CamsColumn = "lane_id"
	// GateTable is the table that holds the gate relation/edge.
	GateTable = "lanes"
	// GateInverseTable is the table name for the Gate entity.
	// It exists in this package in order to avoid circular dependency with the "gate" package.
	GateInverseTable = "gates"
	// GateColumn is the table column denoting the gate relation/edge.
	GateColumn = "gate_id"
)

// Columns holds all SQL columns for lane fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldGateID,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)