package server

import "github.com/init-tech-solution/service-spitc-stream/services/manager/mygrpc"

func paging(reqPage *mygrpc.RequestPaging) (uint, uint) {
	offset := uint(0)
	limit := uint(50)

	if reqPage == nil {
		return offset, limit
	}

	if reqPage.Limit != nil {
		limit = uint(*reqPage.Limit)
	}

	if reqPage.Offset != nil {
		offset = uint(*reqPage.Offset) * limit
	}

	return offset, limit
}
