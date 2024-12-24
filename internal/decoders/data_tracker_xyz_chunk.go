package decoders

import (
	"encoding/binary"
	"errors"
	"math"
)

type DataTrackerXYZChunkData struct {
	X float32
	Y float32
	Z float32
}

type DataTrackerXYZChunk struct {
	Chunk Chunk
	Data  DataTrackerXYZChunkData
}

func DecodeDataTrackerXYZChunk(bytes []byte) (DataTrackerXYZChunk, error) {
	chunk, err := DecodeChunk(bytes)

	if err != nil {
		return DataTrackerXYZChunk{}, err
	}

	if len(chunk.ChunkData) < 12 {
		return DataTrackerXYZChunk{}, errors.New("DATA_TRACKER_XYZ chunk must be at least 12 bytes")
	}

	xBits := binary.LittleEndian.Uint32(chunk.ChunkData[0:4])
	yBits := binary.LittleEndian.Uint32(chunk.ChunkData[4:8])
	zBits := binary.LittleEndian.Uint32(chunk.ChunkData[8:12])

	data := DataTrackerXYZChunkData{
		X: math.Float32frombits(xBits),
		Y: math.Float32frombits(yBits),
		Z: math.Float32frombits(zBits),
	}

	return DataTrackerXYZChunk{
		Chunk: chunk,
		Data:  data,
	}, nil
}
