package decoders

import (
	"encoding/binary"
	"errors"
	"math"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func DecodeDataTrackerStatusChunk(bytes []byte) (chunks.DataTrackerStatusChunk, error) {
	chunk, err := DecodeChunk(bytes)

	if err != nil {
		return chunks.DataTrackerStatusChunk{}, err
	}

	if len(chunk.ChunkData) < 4 {
		return chunks.DataTrackerStatusChunk{}, errors.New("DATA_TRACKER_STATUS chunk must be at least 4 bytes")
	}

	statusBits := binary.LittleEndian.Uint32(chunk.ChunkData[0:4])

	data := chunks.DataTrackerStatusChunkData{
		Validity: math.Float32frombits(statusBits),
	}

	return chunks.DataTrackerStatusChunk{
			Chunk: chunk,
			Data:  data,
		},
		nil
}
