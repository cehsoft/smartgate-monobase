package ocr

import (
	"sort"
	"time"

	"github.com/agnivade/levenshtein"
	"github.com/init-tech-solution/service-spitc-stream/services/lib/copi"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent"
)

func DedupOCRs(inputOCRs []*ent.ContainerTrackingSuggestion) ([]*ent.ContainerTrackingSuggestion, error) {
	ocrs := []*ent.ContainerTrackingSuggestion{}

	err := copi.Dup(inputOCRs, &ocrs)
	if err != nil {
		return ocrs, err
	}

	serialOCRs := map[string][]*ent.ContainerTrackingSuggestion{}

	sort.Slice(ocrs, func(i, j int) bool {
		return ocrs[i].CreatedAt.After(ocrs[j].CreatedAt)
	})

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

	ocrs = []*ent.ContainerTrackingSuggestion{}

	for _, ocrsGroup := range serialOCRs {
		ocrs = append(ocrs, ocrsGroup...)
	}

	sort.Slice(ocrs, func(i, j int) bool {
		return ocrs[i].CreatedAt.After(ocrs[j].CreatedAt)
	})

	return ocrs, nil
}
