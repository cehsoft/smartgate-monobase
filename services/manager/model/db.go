package model

import "github.com/init-tech-solution/service-spitc-stream/services/manager/ent"

type ContainerTrackingSuggestion struct {
	ent.ContainerTrackingSuggestion
	FullCount int `db:"full_count"`
}
