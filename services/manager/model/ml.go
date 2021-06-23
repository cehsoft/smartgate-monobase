package model

import "time"

type ContainerTracking struct {
	ContainerID string
	ImageURL    string
	Score       float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
