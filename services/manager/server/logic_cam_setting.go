package server

import (
	"context"

	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/camsetting"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/mygrpc"
)

func (svc *Server) ListCamSettings(ctx context.Context, req *mygrpc.ReqListCamSettings) (*mygrpc.ResListCamSettings, error) {
	settings, err := svc.db.CamSetting.Query().
		Where(camsetting.LaneIDEQ(int(req.LaneID))).
		All(ctx)
	if err != nil {
		return nil, err
	}

	resSettings := []*mygrpc.CamSetting{}
	for _, s := range settings {
		resSettings = append(resSettings, &mygrpc.CamSetting{
			ID:        int32(s.ID),
			LaneID:    int32(s.LaneID),
			Name:      s.Name,
			WebRTCURL: s.WebrtcURL,
			Position:  s.Position,
		})
	}

	return &mygrpc.ResListCamSettings{
		Settings: resSettings,
	}, nil
}

func (svc *Server) ListLanes(ctx context.Context, req *mygrpc.ReqListLanes) (*mygrpc.ResListLanes, error) {
	lanes, err := svc.db.Lane.Query().WithGate().All(ctx)
	if err != nil {
		return nil, err
	}

	resLanes := []*mygrpc.Lane{}
	for _, l := range lanes {
		resLanes = append(resLanes, &mygrpc.Lane{
			ID:       int32(l.ID),
			GateID:   int32(l.Edges.Gate.ID),
			Name:     l.Name,
			GateName: l.Edges.Gate.Name,
		})
	}

	return &mygrpc.ResListLanes{
		Lanes: resLanes,
	}, nil
}
