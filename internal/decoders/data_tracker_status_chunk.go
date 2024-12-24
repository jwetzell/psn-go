package decoders

import (
	"encoding/binary"
	"errors"
	"math"
)

type DataTrackerStatusChunkData struct {
	Validity float32
}

type DataTrackerStatusChunk struct {
	Chunk Chunk
	Data  DataTrackerStatusChunkData
}

func DecodeDataTrackerStatusChunk(bytes []byte) (DataTrackerStatusChunk, error) {
	chunk, err := DecodeChunk(bytes)

	if err != nil {
		return DataTrackerStatusChunk{}, err
	}

	if len(chunk.ChunkData) < 4 {
		return DataTrackerStatusChunk{}, errors.New("DATA_TRACKER_STATUS chunk must be at least 4 bytes")
	}

	statusBits := binary.LittleEndian.Uint32(chunk.ChunkData[0:4])

	data := DataTrackerStatusChunkData{
		Validity: math.Float32frombits(statusBits),
	}

	return DataTrackerStatusChunk{
			Chunk: chunk,
			Data:  data,
		},
		nil
}
