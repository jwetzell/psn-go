package psn_test

import (
	"testing"

	"github.com/jwetzell/psn-go"
)

func TestTrackerSetPosNilPos(t *testing.T) {
	tracker := &psn.Tracker{}
	tracker.SetPos(1.0, 2.0, 3.0)

	if tracker.Pos == nil {
		t.Errorf("Expected Pos to be initialized, but it is nil")
	} else {
		if tracker.Pos.X != 1.0 || tracker.Pos.Y != 2.0 || tracker.Pos.Z != 3.0 {
			t.Errorf("Expected Pos to be (1.0, 2.0, 3.0), but got (%f, %f, %f)", tracker.Pos.X, tracker.Pos.Y, tracker.Pos.Z)
		}
	}
}

func TestTrackerSetPosNonNilPos(t *testing.T) {
	tracker := &psn.Tracker{
		Pos: &psn.XYZData{X: 0.0, Y: 0.0, Z: 0.0},
	}
	tracker.SetPos(4.0, 5.0, 6.0)

	if tracker.Pos == nil {
		t.Errorf("Expected Pos to be initialized, but it is nil")
	} else {
		if tracker.Pos.X != 4.0 || tracker.Pos.Y != 5.0 || tracker.Pos.Z != 6.0 {
			t.Errorf("Expected Pos to be (4.0, 5.0, 6.0), but got (%f, %f, %f)", tracker.Pos.X, tracker.Pos.Y, tracker.Pos.Z)
		}
	}
}

func TestTrackerSetSpeedNilSpeed(t *testing.T) {
	tracker := &psn.Tracker{}
	tracker.SetSpeed(1.0, 2.0, 3.0)

	if tracker.Speed == nil {
		t.Errorf("Expected Speed to be initialized, but it is nil")
	} else {
		if tracker.Speed.X != 1.0 || tracker.Speed.Y != 2.0 || tracker.Speed.Z != 3.0 {
			t.Errorf("Expected Speed to be (1.0, 2.0, 3.0), but got (%f, %f, %f)", tracker.Speed.X, tracker.Speed.Y, tracker.Speed.Z)
		}
	}
}

func TestTrackerSetSpeedNonNilSpeed(t *testing.T) {
	tracker := &psn.Tracker{
		Speed: &psn.XYZData{X: 0.0, Y: 0.0, Z: 0.0},
	}
	tracker.SetSpeed(4.0, 5.0, 6.0)

	if tracker.Speed == nil {
		t.Errorf("Expected Speed to be initialized, but it is nil")
	} else {
		if tracker.Speed.X != 4.0 || tracker.Speed.Y != 5.0 || tracker.Speed.Z != 6.0 {
			t.Errorf("Expected Speed to be (4.0, 5.0, 6.0), but got (%f, %f, %f)", tracker.Speed.X, tracker.Speed.Y, tracker.Speed.Z)
		}
	}
}

func TestTrackerSetOriNilOri(t *testing.T) {
	tracker := &psn.Tracker{}
	tracker.SetOri(1.0, 2.0, 3.0)

	if tracker.Ori == nil {
		t.Errorf("Expected Ori to be initialized, but it is nil")
	} else {
		if tracker.Ori.X != 1.0 || tracker.Ori.Y != 2.0 || tracker.Ori.Z != 3.0 {
			t.Errorf("Expected Ori to be (1.0, 2.0, 3.0), but got (%f, %f, %f)", tracker.Ori.X, tracker.Ori.Y, tracker.Ori.Z)
		}
	}
}

func TestTrackerSetOriNonNilOri(t *testing.T) {
	tracker := &psn.Tracker{
		Ori: &psn.XYZData{X: 0.0, Y: 0.0, Z: 0.0},
	}
	tracker.SetOri(4.0, 5.0, 6.0)

	if tracker.Ori == nil {
		t.Errorf("Expected Ori to be initialized, but it is nil")
	} else {
		if tracker.Ori.X != 4.0 || tracker.Ori.Y != 5.0 || tracker.Ori.Z != 6.0 {
			t.Errorf("Expected Ori to be (4.0, 5.0, 6.0), but got (%f, %f, %f)", tracker.Ori.X, tracker.Ori.Y, tracker.Ori.Z)
		}
	}
}

func TestTrackerSetStatusNilStatus(t *testing.T) {
	tracker := &psn.Tracker{}
	tracker.SetStatus(0.75)

	if tracker.Validity == nil {
		t.Errorf("Expected Validity to be initialized, but it is nil")
	} else {
		if *tracker.Validity != 0.75 {
			t.Errorf("Expected Validity to be 0.75, but got %f", *tracker.Validity)
		}
	}
}

func TestTrackerSetStatusNonNilStatus(t *testing.T) {
	tracker := &psn.Tracker{
		Validity: new(float32),
	}
	*tracker.Validity = 0.5
	tracker.SetStatus(0.85)

	if tracker.Validity == nil {
		t.Errorf("Expected Validity to be initialized, but it is nil")
	} else {
		if *tracker.Validity != 0.85 {
			t.Errorf("Expected Validity to be 0.85, but got %f", *tracker.Validity)
		}
	}
}

func TestTrackerSetAccelNilAccel(t *testing.T) {
	tracker := &psn.Tracker{}
	tracker.SetAccel(1.0, 2.0, 3.0)

	if tracker.Accel == nil {
		t.Errorf("Expected Accel to be initialized, but it is nil")
	} else {
		if tracker.Accel.X != 1.0 || tracker.Accel.Y != 2.0 || tracker.Accel.Z != 3.0 {
			t.Errorf("Expected Accel to be (1.0, 2.0, 3.0), but got (%f, %f, %f)", tracker.Accel.X, tracker.Accel.Y, tracker.Accel.Z)
		}
	}
}

func TestTrackerSetAccelNonNilAccel(t *testing.T) {
	tracker := &psn.Tracker{
		Accel: &psn.XYZData{X: 0.0, Y: 0.0, Z: 0.0},
	}
	tracker.SetAccel(4.0, 5.0, 6.0)

	if tracker.Accel == nil {
		t.Errorf("Expected Accel to be initialized, but it is nil")
	} else {
		if tracker.Accel.X != 4.0 || tracker.Accel.Y != 5.0 || tracker.Accel.Z != 6.0 {
			t.Errorf("Expected Accel to be (4.0, 5.0, 6.0), but got (%f, %f, %f)", tracker.Accel.X, tracker.Accel.Y, tracker.Accel.Z)
		}
	}
}

func TestTrackerSetTrgtPosNilTrgtPos(t *testing.T) {
	tracker := &psn.Tracker{}
	tracker.SetTrgtPos(1.0, 2.0, 3.0)

	if tracker.TrgtPos == nil {
		t.Errorf("Expected TrgtPos to be initialized, but it is nil")
	} else {
		if tracker.TrgtPos.X != 1.0 || tracker.TrgtPos.Y != 2.0 || tracker.TrgtPos.Z != 3.0 {
			t.Errorf("Expected TrgtPos to be (1.0, 2.0, 3.0), but got (%f, %f, %f)", tracker.TrgtPos.X, tracker.TrgtPos.Y, tracker.TrgtPos.Z)
		}
	}
}

func TestTrackerSetTrgtPosNonNilTrgtPos(t *testing.T) {
	tracker := &psn.Tracker{
		TrgtPos: &psn.XYZData{X: 0.0, Y: 0.0, Z: 0.0},
	}
	tracker.SetTrgtPos(4.0, 5.0, 6.0)

	if tracker.TrgtPos == nil {
		t.Errorf("Expected TrgtPos to be initialized, but it is nil")
	} else {
		if tracker.TrgtPos.X != 4.0 || tracker.TrgtPos.Y != 5.0 || tracker.TrgtPos.Z != 6.0 {
			t.Errorf("Expected TrgtPos to be (4.0, 5.0, 6.0), but got (%f, %f, %f)", tracker.TrgtPos.X, tracker.TrgtPos.Y, tracker.TrgtPos.Z)
		}
	}
}

func TestTrackerSetTimestampNilTimestamp(t *testing.T) {
	tracker := &psn.Tracker{}
	tracker.SetTimestamp(1)

	if tracker.Timestamp == nil {
		t.Errorf("Expected Timestamp to be initialized, but it is nil")
	} else {
		if *tracker.Timestamp != 1 {
			t.Errorf("Expected Timestamp to be 1, but got %d", *tracker.Timestamp)
		}
	}
}

func TestTrackerSetTimestampNonNilTimestamp(t *testing.T) {
	tracker := &psn.Tracker{
		Timestamp: new(uint64),
	}
	*tracker.Timestamp = 5
	tracker.SetTimestamp(8)

	if tracker.Timestamp == nil {
		t.Errorf("Expected Timestamp to be initialized, but it is nil")
	} else {
		if *tracker.Timestamp != 8 {
			t.Errorf("Expected Timestamp to be 8, but got %d", *tracker.Timestamp)
		}
	}
}
