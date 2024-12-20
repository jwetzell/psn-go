package psn

import (
	"github.com/jwetzell/psn-go/internal/decoders"
	"github.com/jwetzell/psn-go/internal/encoders"
)

type XYZData struct {
	X float32
	Y float32
	Z float32
}

type Tracker struct {
	Id        uint16
	Name      string
	Pos       *XYZData
	Speed     *XYZData
	Ori       *XYZData
	Validity  *float32
	Accel     *XYZData
	TrgtPos   *XYZData
	Timestamp *uint64
}

func (t *Tracker) SetPos(x float32, y float32, z float32) {
	if t.Pos == nil {
		t.Pos = &XYZData{
			X: x,
			Y: y,
			Z: z,
		}
	} else {
		t.Pos.X = x
		t.Pos.Y = y
		t.Pos.Z = z
	}
}

func (t *Tracker) SetSpeed(x float32, y float32, z float32) {
	if t.Speed == nil {
		t.Speed = &XYZData{
			X: x,
			Y: y,
			Z: z,
		}
	} else {
		t.Speed.X = x
		t.Speed.Y = y
		t.Speed.Z = z
	}
}

func (t *Tracker) SetOri(x float32, y float32, z float32) {
	if t.Ori == nil {
		t.Ori = &XYZData{
			X: x,
			Y: y,
			Z: z,
		}
	} else {
		t.Ori.X = x
		t.Ori.Y = y
		t.Ori.Z = z
	}
}

func (t *Tracker) SetStatus(validity float32) {
	t.Validity = &validity
}

func (t *Tracker) SetAccel(x float32, y float32, z float32) {
	if t.Accel == nil {
		t.Accel = &XYZData{
			X: x,
			Y: y,
			Z: z,
		}
	} else {
		t.Accel.X = x
		t.Accel.Y = y
		t.Accel.Z = z
	}
}

func (t *Tracker) SetTrgtPos(x float32, y float32, z float32) {
	if t.TrgtPos == nil {
		t.TrgtPos = &XYZData{
			X: x,
			Y: y,
			Z: z,
		}
	} else {
		t.TrgtPos.X = x
		t.TrgtPos.Y = y
		t.TrgtPos.Z = z
	}
}

func (t *Tracker) SetTimestamp(timestamp uint64) {
	t.Timestamp = &timestamp
}

func (t *Tracker) GetDataChunk() []byte {
	fieldChunks := [][]byte{}
	if t.Pos != nil {
		fieldChunks = append(fieldChunks, encoders.EncodeDataTrackerPosChunk(t.Pos.X, t.Pos.Y, t.Pos.Z))
	}

	if t.Speed != nil {
		fieldChunks = append(fieldChunks, encoders.EncodeDataTrackerSpeedChunk(t.Speed.X, t.Speed.Y, t.Speed.Z))
	}

	if t.Ori != nil {
		fieldChunks = append(fieldChunks, encoders.EncodeDataTrackerOriChunk(t.Ori.X, t.Ori.Y, t.Ori.Z))
	}

	if t.Validity != nil {
		fieldChunks = append(fieldChunks, encoders.EncodeDataTrackerStatusChunk(*t.Validity))
	}

	if t.Accel != nil {
		fieldChunks = append(fieldChunks, encoders.EncodeDataTrackerAccelChunk(t.Accel.X, t.Accel.Y, t.Accel.Z))
	}

	if t.TrgtPos != nil {
		fieldChunks = append(fieldChunks, encoders.EncodeDataTrackerTrgtPosChunk(t.TrgtPos.X, t.TrgtPos.Y, t.TrgtPos.Z))
	}

	if t.Timestamp != nil {
		fieldChunks = append(fieldChunks, encoders.EncodeDataTrackerTimestampChunk(*t.Timestamp))
	}

	return encoders.EncodeDataTrackerChunk(t.Id, fieldChunks)
}

func (t *Tracker) GetInfoChunk() []byte {
	return encoders.EncodeInfoTrackerChunk(t.Id, encoders.EncodeInfoTrackerNameChunk(t.Name))
}

func (t *Tracker) UpdateInfo(infoTrackerChunk decoders.InfoTrackerChunk) {
	t.Id = infoTrackerChunk.Chunk.Header.Id
	t.Name = infoTrackerChunk.Data.TrackerName.Data.TrackerName
}

func (t *Tracker) UpdateData(dataTrackerChunk decoders.DataTrackerChunk) {
	if dataTrackerChunk.Data.Pos != nil {
		t.SetPos(dataTrackerChunk.Data.Pos.Data.X, dataTrackerChunk.Data.Pos.Data.Y, dataTrackerChunk.Data.Pos.Data.Z)
	}

	if dataTrackerChunk.Data.Speed != nil {
		t.SetSpeed(dataTrackerChunk.Data.Speed.Data.X, dataTrackerChunk.Data.Speed.Data.Y, dataTrackerChunk.Data.Speed.Data.Z)
	}

	if dataTrackerChunk.Data.Ori != nil {
		t.SetOri(dataTrackerChunk.Data.Ori.Data.X, dataTrackerChunk.Data.Ori.Data.Y, dataTrackerChunk.Data.Ori.Data.Z)
	}

	if dataTrackerChunk.Data.Status != nil {
		t.SetStatus(dataTrackerChunk.Data.Status.Data.Validity)
	}

	if dataTrackerChunk.Data.Accel != nil {
		t.SetAccel(dataTrackerChunk.Data.Accel.Data.X, dataTrackerChunk.Data.Accel.Data.Y, dataTrackerChunk.Data.Accel.Data.Z)
	}

	if dataTrackerChunk.Data.TrgtPos != nil {
		t.SetTrgtPos(dataTrackerChunk.Data.TrgtPos.Data.X, dataTrackerChunk.Data.TrgtPos.Data.Y, dataTrackerChunk.Data.TrgtPos.Data.Z)
	}

	if dataTrackerChunk.Data.Timestamp != nil {
		t.SetTimestamp(dataTrackerChunk.Data.Timestamp.Data.Timestamp)
	}
}

func TrackerFromInfo(infoTrackerChunk decoders.InfoTrackerChunk) *Tracker {
	var tracker Tracker
	tracker.UpdateInfo(infoTrackerChunk)
	return &tracker
}

func TrackerFromData(dataTrackerChunk decoders.DataTrackerChunk) *Tracker {
	var tracker Tracker
	tracker.UpdateData(dataTrackerChunk)
	return &tracker
}
