package server

import (
	"sort"
	"time"

	"github.com/agnivade/levenshtein"
	"github.com/init-tech-solution/service-spitc-stream/services/lib/copi"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent"
)

func splitContainerID(containerID string) []string {
	if len(containerID) < 5 {
		return nil
	}

	bic := containerID[0:4]
	serial := containerID[4 : len(containerID)-1]
	checksum := containerID[len(containerID)-1:]
	return []string{bic, serial, checksum}
}

func dedupOCRs(inputOCRs []*ent.ContainerTrackingSuggestion) ([]*ent.ContainerTrackingSuggestion, error) {
	ocrs := []*ent.ContainerTrackingSuggestion{}

	err := copi.Dup(inputOCRs, &ocrs)
	if err != nil {
		return ocrs, err
	}

	sort.SliceStable(ocrs, func(i, j int) bool {
		return ocrs[i].CreatedAt.After(ocrs[j].CreatedAt)
	})

	serialOCRs := map[string][]*ent.ContainerTrackingSuggestion{}

	prev := &ent.ContainerTrackingSuggestion{}
	for _, ocr := range ocrs {
		label := ocr.Serial
		distance := levenshtein.ComputeDistance(prev.Serial, label)
		if distance < 3 {
			label = prev.Serial
		}

		if serialOCRs[label] == nil {
			serialOCRs[label] = []*ent.ContainerTrackingSuggestion{}
		}

		lastIdx := len(serialOCRs[label]) - 1
		if lastIdx > -1 {
			lastOCR := serialOCRs[label][lastIdx]
			if lastOCR.CreatedAt.Sub(ocr.CreatedAt) < time.Duration(5*time.Minute) {
				if ocr.Score > lastOCR.Score {
					serialOCRs[label][lastIdx] = ocr
				}
			} else {
				serialOCRs[label] = append(serialOCRs[label], ocr)
			}

		} else {
			serialOCRs[label] = append(serialOCRs[label], ocr)
		}

		prev = ocr
	}

	result := []*ent.ContainerTrackingSuggestion{}

	for _, ocrsGroup := range serialOCRs {
		result = append(result, ocrsGroup...)
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].CreatedAt.After(result[j].CreatedAt)
	})

	return ocrs, nil
}
