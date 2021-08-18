package ocr

import (
	"sort"
	"time"

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

	for _, ocr := range ocrs {
		if serialOCRs[ocr.Serial] == nil {
			serialOCRs[ocr.Serial] = []*ent.ContainerTrackingSuggestion{}
		}

		lastIdx := len(serialOCRs[ocr.Serial]) - 1
		if lastIdx > -1 {
			lastOCR := serialOCRs[ocr.Serial][lastIdx]
			if lastOCR.CreatedAt.Sub(ocr.CreatedAt) < time.Duration(5*time.Minute) {
				if ocr.Score > lastOCR.Score {
					serialOCRs[ocr.Serial][lastIdx] = ocr
				}
			} else {
				serialOCRs[ocr.Serial] = append(serialOCRs[ocr.Serial], ocr)
			}

		} else {
			serialOCRs[ocr.Serial] = append(serialOCRs[ocr.Serial], ocr)
		}
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
