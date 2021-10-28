package server

import (
	"strings"
	"time"

	"github.com/agnivade/levenshtein"
	"github.com/init-tech-solution/service-spitc-stream/services/lib/uuid"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/mygrpc"
	"github.com/jmoiron/sqlx"
	"github.com/mitchellh/copystructure"
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

func newSessionWithOCR(sessToTypes map[string]map[string][]*ent.ContainerTrackingSuggestion, sessID string, ocr *ent.ContainerTrackingSuggestion) {
	if sessToTypes[sessID] == nil {
		sessToTypes[sessID] = map[string][]*ent.ContainerTrackingSuggestion{
			"contID":    {},
			"vehicleID": {},
			"romoocID":  {},
		}

		sessToTypes[sessID][ocr.TrackingType] = append(sessToTypes[sessID][ocr.TrackingType], ocr)
	}
}

func createVirtualSessions(inputOCRs []*ent.ContainerTrackingSuggestion) (map[string]map[string][]*ent.ContainerTrackingSuggestion, error) {
	copied, err := copystructure.Copy(inputOCRs)
	if err != nil {
		return nil, nil
	}

	ocrs := copied.([]*ent.ContainerTrackingSuggestion)

	sessToTypes := map[string]map[string][]*ent.ContainerTrackingSuggestion{} // temp sessions map with type and ocr results
	sessID := "virtual_" + uuid.UUID(20)                                      // current sessID
	prevTime := time.Time{}                                                   // tracking time

	for _, current := range ocrs {
		if current.Result == "" {
			prevTime = current.CreatedAt
			continue
		}

		// timebox is over -> create new session
		if current.CreatedAt.Sub(prevTime) > time.Duration(30*time.Second) {
			sessID = "virtual_" + uuid.UUID(20)             // create new session
			newSessionWithOCR(sessToTypes, sessID, current) // append new result to new created session

			prevTime = current.CreatedAt // update track time
			continue
		}

		if sessToTypes[sessID] == nil {
			// don't have session -> create new
			newSessionWithOCR(sessToTypes, sessID, current)

			prevTime = current.CreatedAt
			continue
		}

		typedResults := sessToTypes[sessID][current.TrackingType]
		lenTResults := len(typedResults)
		if lenTResults == 0 {
			// session not has this type result -> append new ocr result
			sessToTypes[sessID][current.TrackingType] = append(sessToTypes[sessID][current.TrackingType], current)

			prevTime = current.CreatedAt
			continue
		}

		lastTResultIdx := lenTResults - 1
		firstTResult := typedResults[0]
		lastTResult := typedResults[lastTResultIdx]
		// update track time
		diffScore := levenshtein.ComputeDistance(lastTResult.Result, current.Result) // matching current result with previous one
		if diffScore > 3 {                                                           // results diff to much -> create new session
			sessID = "virtual_" + uuid.UUID(20)             // create new session
			newSessionWithOCR(sessToTypes, sessID, current) // append new result to new created session

			prevTime = current.CreatedAt
			continue
		}

		sessToTypes[sessID][current.TrackingType] = append(sessToTypes[sessID][current.TrackingType], current)

		if current.Score > firstTResult.Score { // swap highest result to first index, this is not sort array
			temp := sessToTypes[sessID][current.TrackingType][0]          // store current first
			sessToTypes[sessID][current.TrackingType][0] = current        // assign new first
			sessToTypes[sessID][current.TrackingType][lenTResults] = temp // reassign last
		}

		prevTime = current.CreatedAt
	}

	return sessToTypes, nil
}
