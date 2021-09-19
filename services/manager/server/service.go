package server

import (
	"strings"

	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/mygrpc"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	db    *ent.Client
	rawdb *sqlx.DB

	clients map[string]chan *ent.ContainerTrackingSuggestion

	mygrpc.UnimplementedMyGRPCServer
}

var _ mygrpc.MyGRPCServer = &Server{}

func CreateServer(db *ent.Client, rawdb *sqlx.DB) *Server {
	return &Server{
		db:      db,
		rawdb:   rawdb,
		clients: map[string]chan *ent.ContainerTrackingSuggestion{},
	}
}

// Utility functions

func splitContainerID(containerID string) []string {
	if len(containerID) < 5 {
		return nil
	}

	bic := containerID[0:4]
	serial := containerID[4 : len(containerID)-1]
	checksum := containerID[len(containerID)-1:]
	return []string{bic, serial, checksum}
}

// extractTrackingMetaFromImgURL returns trackingType, trackingSession
func extractTrackingMetaFromImgURL(imgURL string) (string, string) {
	urlPaths := strings.Split(imgURL, "/")
	if len(urlPaths) < 2 {
		return "", ""
	}
	fileName := urlPaths[len(urlPaths)-1]
	nameParts := strings.Split(fileName, "_")
	if len(nameParts) < 2 {
		return urlPaths[len(urlPaths)-2], ""
	}
	return urlPaths[len(urlPaths)-2], nameParts[0]
}
