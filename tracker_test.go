package psn_test

import (
	"testing"

	"github.com/jwetzell/psn-go"
	"github.com/stretchr/testify/require"
)

func TestTrackerSetPosNilPos(t *testing.T) {
	tracker := &psn.Tracker{}
	tracker.SetPos(1.0, 2.0, 3.0)

	require.NotNil(t, tracker.Pos)

	require.Equal(t, float32(1.0), tracker.Pos.X)
	require.Equal(t, float32(2.0), tracker.Pos.Y)
	require.Equal(t, float32(3.0), tracker.Pos.Z)
}

func TestTrackerSetPosNonNilPos(t *testing.T) {
	tracker := &psn.Tracker{
		Pos: &psn.XYZData{X: 0.0, Y: 0.0, Z: 0.0},
	}
	tracker.SetPos(4.0, 5.0, 6.0)

	require.NotNil(t, tracker.Pos)
	require.Equal(t, float32(4.0), tracker.Pos.X)
	require.Equal(t, float32(5.0), tracker.Pos.Y)
	require.Equal(t, float32(6.0), tracker.Pos.Z)
}

func TestTrackerSetSpeedNilSpeed(t *testing.T) {
	tracker := &psn.Tracker{}
	tracker.SetSpeed(1.0, 2.0, 3.0)

	require.NotNil(t, tracker.Speed)
	require.Equal(t, float32(1.0), tracker.Speed.X)
	require.Equal(t, float32(2.0), tracker.Speed.Y)
	require.Equal(t, float32(3.0), tracker.Speed.Z)
}

func TestTrackerSetSpeedNonNilSpeed(t *testing.T) {
	tracker := &psn.Tracker{
		Speed: &psn.XYZData{X: 0.0, Y: 0.0, Z: 0.0},
	}
	tracker.SetSpeed(4.0, 5.0, 6.0)

	require.NotNil(t, tracker.Speed)
	require.Equal(t, float32(4.0), tracker.Speed.X)
	require.Equal(t, float32(5.0), tracker.Speed.Y)
	require.Equal(t, float32(6.0), tracker.Speed.Z)
}

func TestTrackerSetOriNilOri(t *testing.T) {
	tracker := &psn.Tracker{}
	tracker.SetOri(1.0, 2.0, 3.0)

	require.NotNil(t, tracker.Ori)
	require.Equal(t, float32(1.0), tracker.Ori.X)
	require.Equal(t, float32(2.0), tracker.Ori.Y)
	require.Equal(t, float32(3.0), tracker.Ori.Z)
}

func TestTrackerSetOriNonNilOri(t *testing.T) {
	tracker := &psn.Tracker{
		Ori: &psn.XYZData{X: 0.0, Y: 0.0, Z: 0.0},
	}
	tracker.SetOri(4.0, 5.0, 6.0)

	require.NotNil(t, tracker.Ori)
	require.Equal(t, float32(4.0), tracker.Ori.X)
	require.Equal(t, float32(5.0), tracker.Ori.Y)
	require.Equal(t, float32(6.0), tracker.Ori.Z)
}

func TestTrackerSetStatusNilStatus(t *testing.T) {
	tracker := &psn.Tracker{}
	tracker.SetStatus(0.75)

	require.NotNil(t, tracker.Validity)
	require.Equal(t, float32(0.75), *tracker.Validity)
}

func TestTrackerSetStatusNonNilStatus(t *testing.T) {
	tracker := &psn.Tracker{
		Validity: new(float32),
	}
	*tracker.Validity = 0.5
	tracker.SetStatus(0.85)

	require.NotNil(t, tracker.Validity)
	require.Equal(t, float32(0.85), *tracker.Validity)
}

func TestTrackerSetAccelNilAccel(t *testing.T) {
	tracker := &psn.Tracker{}
	tracker.SetAccel(1.0, 2.0, 3.0)

	require.NotNil(t, tracker.Accel)
	require.Equal(t, float32(1.0), tracker.Accel.X)
	require.Equal(t, float32(2.0), tracker.Accel.Y)
	require.Equal(t, float32(3.0), tracker.Accel.Z)
}

func TestTrackerSetAccelNonNilAccel(t *testing.T) {
	tracker := &psn.Tracker{
		Accel: &psn.XYZData{X: 0.0, Y: 0.0, Z: 0.0},
	}
	tracker.SetAccel(4.0, 5.0, 6.0)

	require.NotNil(t, tracker.Accel)
	require.Equal(t, float32(4.0), tracker.Accel.X)
	require.Equal(t, float32(5.0), tracker.Accel.Y)
	require.Equal(t, float32(6.0), tracker.Accel.Z)
}

func TestTrackerSetTrgtPosNilTrgtPos(t *testing.T) {
	tracker := &psn.Tracker{}
	tracker.SetTrgtPos(1.0, 2.0, 3.0)

	require.NotNil(t, tracker.TrgtPos)
	require.Equal(t, float32(1.0), tracker.TrgtPos.X)
	require.Equal(t, float32(2.0), tracker.TrgtPos.Y)
	require.Equal(t, float32(3.0), tracker.TrgtPos.Z)
}

func TestTrackerSetTrgtPosNonNilTrgtPos(t *testing.T) {
	tracker := &psn.Tracker{
		TrgtPos: &psn.XYZData{X: 0.0, Y: 0.0, Z: 0.0},
	}
	tracker.SetTrgtPos(4.0, 5.0, 6.0)

	require.NotNil(t, tracker.TrgtPos)
	require.Equal(t, float32(4.0), tracker.TrgtPos.X)
	require.Equal(t, float32(5.0), tracker.TrgtPos.Y)
	require.Equal(t, float32(6.0), tracker.TrgtPos.Z)
}

func TestTrackerSetTimestampNilTimestamp(t *testing.T) {
	tracker := &psn.Tracker{}
	tracker.SetTimestamp(1)

	require.NotNil(t, tracker.Timestamp)
	require.Equal(t, uint64(1), *tracker.Timestamp)
}

func TestTrackerSetTimestampNonNilTimestamp(t *testing.T) {
	tracker := &psn.Tracker{
		Timestamp: new(uint64),
	}
	*tracker.Timestamp = 5
	tracker.SetTimestamp(8)

	require.NotNil(t, tracker.Timestamp)
	require.Equal(t, uint64(8), *tracker.Timestamp)
}
