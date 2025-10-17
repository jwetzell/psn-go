package decoders

import (
	"encoding/binary"
	"errors"
	"math"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func DecodeDataTrackerXYZChunk(bytes []byte) (chunks.DataTrackerXYZChunk, error) {
	chunk, err := DecodeChunk(bytes)

	if err != nil {
		return chunks.DataTrackerXYZChunk{}, err
	}

	if len(chunk.ChunkData) < 12 {
		return chunks.DataTrackerXYZChunk{}, errors.New("DATA_TRACKER_XYZ chunk must be at least 12 bytes")
	}

	xBits := binary.LittleEndian.Uint32(chunk.ChunkData[0:4])
	yBits := binary.LittleEndian.Uint32(chunk.ChunkData[4:8])
	zBits := binary.LittleEndian.Uint32(chunk.ChunkData[8:12])

	data := chunks.DataTrackerXYZChunkData{
		X: math.Float32frombits(xBits),
		Y: math.Float32frombits(yBits),
		Z: math.Float32frombits(zBits),
	}

	return chunks.DataTrackerXYZChunk{
		Chunk: chunk,
		Data:  data,
	}, nil
}
