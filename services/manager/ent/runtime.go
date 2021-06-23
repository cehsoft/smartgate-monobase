// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/containertracking"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/containertrackingsuggestion"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	containertrackingFields := schema.ContainerTracking{}.Fields()
	_ = containertrackingFields
	// containertrackingDescCreatedAt is the schema descriptor for created_at field.
	containertrackingDescCreatedAt := containertrackingFields[1].Descriptor()
	// containertracking.DefaultCreatedAt holds the default value on creation for the created_at field.
	containertracking.DefaultCreatedAt = containertrackingDescCreatedAt.Default.(func() time.Time)
	containertrackingsuggestionFields := schema.ContainerTrackingSuggestion{}.Fields()
	_ = containertrackingsuggestionFields
	// containertrackingsuggestionDescCreatedAt is the schema descriptor for created_at field.
	containertrackingsuggestionDescCreatedAt := containertrackingsuggestionFields[4].Descriptor()
	// containertrackingsuggestion.DefaultCreatedAt holds the default value on creation for the created_at field.
	containertrackingsuggestion.DefaultCreatedAt = containertrackingsuggestionDescCreatedAt.Default.(func() time.Time)
}
