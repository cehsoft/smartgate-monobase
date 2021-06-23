package uuid

import (
	"strings"

	"github.com/google/uuid"
)

func UUID(length int) string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	if length != 0 && length <= len(uuid) {
		return uuid[0 : length-1]
	}

	return uuid
}
