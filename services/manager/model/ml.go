package model

import "time"

type ContainerTracking struct {
	ContainerID string
	ImageURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
