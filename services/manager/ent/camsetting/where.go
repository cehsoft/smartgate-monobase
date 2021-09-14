// Code generated by entc, DO NOT EDIT.

package camsetting

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/init-tech-solution/service-spitc-stream/services/manager/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Position applies equality check predicate on the "position" field. It's identical to PositionEQ.
func Position(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPosition), v))
	})
}

// LaneID applies equality check predicate on the "lane_id" field. It's identical to LaneIDEQ.
func LaneID(v int) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLaneID), v))
	})
}

// RtspURL applies equality check predicate on the "rtsp_url" field. It's identical to RtspURLEQ.
func RtspURL(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRtspURL), v))
	})
}

// WebrtcURL applies equality check predicate on the "webrtc_url" field. It's identical to WebrtcURLEQ.
func WebrtcURL(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWebrtcURL), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.CamSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CamSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.CamSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CamSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// PositionEQ applies the EQ predicate on the "position" field.
func PositionEQ(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPosition), v))
	})
}

// PositionNEQ applies the NEQ predicate on the "position" field.
func PositionNEQ(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPosition), v))
	})
}

// PositionIn applies the In predicate on the "position" field.
func PositionIn(vs ...string) predicate.CamSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CamSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPosition), v...))
	})
}

// PositionNotIn applies the NotIn predicate on the "position" field.
func PositionNotIn(vs ...string) predicate.CamSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CamSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPosition), v...))
	})
}

// PositionGT applies the GT predicate on the "position" field.
func PositionGT(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPosition), v))
	})
}

// PositionGTE applies the GTE predicate on the "position" field.
func PositionGTE(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPosition), v))
	})
}

// PositionLT applies the LT predicate on the "position" field.
func PositionLT(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPosition), v))
	})
}

// PositionLTE applies the LTE predicate on the "position" field.
func PositionLTE(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPosition), v))
	})
}

// PositionContains applies the Contains predicate on the "position" field.
func PositionContains(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPosition), v))
	})
}

// PositionHasPrefix applies the HasPrefix predicate on the "position" field.
func PositionHasPrefix(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPosition), v))
	})
}

// PositionHasSuffix applies the HasSuffix predicate on the "position" field.
func PositionHasSuffix(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPosition), v))
	})
}

// PositionIsNil applies the IsNil predicate on the "position" field.
func PositionIsNil() predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPosition)))
	})
}

// PositionNotNil applies the NotNil predicate on the "position" field.
func PositionNotNil() predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPosition)))
	})
}

// PositionEqualFold applies the EqualFold predicate on the "position" field.
func PositionEqualFold(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPosition), v))
	})
}

// PositionContainsFold applies the ContainsFold predicate on the "position" field.
func PositionContainsFold(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPosition), v))
	})
}

// LaneIDEQ applies the EQ predicate on the "lane_id" field.
func LaneIDEQ(v int) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLaneID), v))
	})
}

// LaneIDNEQ applies the NEQ predicate on the "lane_id" field.
func LaneIDNEQ(v int) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLaneID), v))
	})
}

// LaneIDIn applies the In predicate on the "lane_id" field.
func LaneIDIn(vs ...int) predicate.CamSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CamSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLaneID), v...))
	})
}

// LaneIDNotIn applies the NotIn predicate on the "lane_id" field.
func LaneIDNotIn(vs ...int) predicate.CamSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CamSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLaneID), v...))
	})
}

// LaneIDIsNil applies the IsNil predicate on the "lane_id" field.
func LaneIDIsNil() predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLaneID)))
	})
}

// LaneIDNotNil applies the NotNil predicate on the "lane_id" field.
func LaneIDNotNil() predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLaneID)))
	})
}

// RtspURLEQ applies the EQ predicate on the "rtsp_url" field.
func RtspURLEQ(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRtspURL), v))
	})
}

// RtspURLNEQ applies the NEQ predicate on the "rtsp_url" field.
func RtspURLNEQ(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRtspURL), v))
	})
}

// RtspURLIn applies the In predicate on the "rtsp_url" field.
func RtspURLIn(vs ...string) predicate.CamSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CamSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRtspURL), v...))
	})
}

// RtspURLNotIn applies the NotIn predicate on the "rtsp_url" field.
func RtspURLNotIn(vs ...string) predicate.CamSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CamSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRtspURL), v...))
	})
}

// RtspURLGT applies the GT predicate on the "rtsp_url" field.
func RtspURLGT(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRtspURL), v))
	})
}

// RtspURLGTE applies the GTE predicate on the "rtsp_url" field.
func RtspURLGTE(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRtspURL), v))
	})
}

// RtspURLLT applies the LT predicate on the "rtsp_url" field.
func RtspURLLT(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRtspURL), v))
	})
}

// RtspURLLTE applies the LTE predicate on the "rtsp_url" field.
func RtspURLLTE(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRtspURL), v))
	})
}

// RtspURLContains applies the Contains predicate on the "rtsp_url" field.
func RtspURLContains(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRtspURL), v))
	})
}

// RtspURLHasPrefix applies the HasPrefix predicate on the "rtsp_url" field.
func RtspURLHasPrefix(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRtspURL), v))
	})
}

// RtspURLHasSuffix applies the HasSuffix predicate on the "rtsp_url" field.
func RtspURLHasSuffix(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRtspURL), v))
	})
}

// RtspURLEqualFold applies the EqualFold predicate on the "rtsp_url" field.
func RtspURLEqualFold(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRtspURL), v))
	})
}

// RtspURLContainsFold applies the ContainsFold predicate on the "rtsp_url" field.
func RtspURLContainsFold(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRtspURL), v))
	})
}

// WebrtcURLEQ applies the EQ predicate on the "webrtc_url" field.
func WebrtcURLEQ(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWebrtcURL), v))
	})
}

// WebrtcURLNEQ applies the NEQ predicate on the "webrtc_url" field.
func WebrtcURLNEQ(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWebrtcURL), v))
	})
}

// WebrtcURLIn applies the In predicate on the "webrtc_url" field.
func WebrtcURLIn(vs ...string) predicate.CamSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CamSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWebrtcURL), v...))
	})
}

// WebrtcURLNotIn applies the NotIn predicate on the "webrtc_url" field.
func WebrtcURLNotIn(vs ...string) predicate.CamSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CamSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWebrtcURL), v...))
	})
}

// WebrtcURLGT applies the GT predicate on the "webrtc_url" field.
func WebrtcURLGT(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWebrtcURL), v))
	})
}

// WebrtcURLGTE applies the GTE predicate on the "webrtc_url" field.
func WebrtcURLGTE(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWebrtcURL), v))
	})
}

// WebrtcURLLT applies the LT predicate on the "webrtc_url" field.
func WebrtcURLLT(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWebrtcURL), v))
	})
}

// WebrtcURLLTE applies the LTE predicate on the "webrtc_url" field.
func WebrtcURLLTE(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWebrtcURL), v))
	})
}

// WebrtcURLContains applies the Contains predicate on the "webrtc_url" field.
func WebrtcURLContains(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldWebrtcURL), v))
	})
}

// WebrtcURLHasPrefix applies the HasPrefix predicate on the "webrtc_url" field.
func WebrtcURLHasPrefix(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldWebrtcURL), v))
	})
}

// WebrtcURLHasSuffix applies the HasSuffix predicate on the "webrtc_url" field.
func WebrtcURLHasSuffix(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldWebrtcURL), v))
	})
}

// WebrtcURLEqualFold applies the EqualFold predicate on the "webrtc_url" field.
func WebrtcURLEqualFold(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldWebrtcURL), v))
	})
}

// WebrtcURLContainsFold applies the ContainsFold predicate on the "webrtc_url" field.
func WebrtcURLContainsFold(v string) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldWebrtcURL), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CamSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CamSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CamSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CamSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasLane applies the HasEdge predicate on the "lane" edge.
func HasLane() predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LaneTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LaneTable, LaneColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLaneWith applies the HasEdge predicate on the "lane" edge with a given conditions (other predicates).
func HasLaneWith(preds ...predicate.Lane) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LaneInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LaneTable, LaneColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSuggestions applies the HasEdge predicate on the "suggestions" edge.
func HasSuggestions() predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SuggestionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SuggestionsTable, SuggestionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSuggestionsWith applies the HasEdge predicate on the "suggestions" edge with a given conditions (other predicates).
func HasSuggestionsWith(preds ...predicate.ContainerTrackingSuggestion) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SuggestionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SuggestionsTable, SuggestionsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CamSetting) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CamSetting) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CamSetting) predicate.CamSetting {
	return predicate.CamSetting(func(s *sql.Selector) {
		p(s.Not())
	})
}