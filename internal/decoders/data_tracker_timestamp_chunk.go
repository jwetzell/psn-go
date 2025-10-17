package decoders

import (
	"encoding/binary"
	"errors"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func DecodeDataTrackerTimestampChunk(bytes []byte) (chunks.DataTrackerTimestampChunk, error) {
	chunk, err := DecodeChunk(bytes)

	if err != nil {
		return chunks.DataTrackerTimestampChunk{}, err
	}

	if len(chunk.ChunkData) < 8 {
		return chunks.DataTrackerTimestampChunk{}, errors.New("DATA_TRACKER_TIMESTAMP chunk must be at least 8 bytes")
	}

	timestamp := binary.LittleEndian.Uint64(chunk.ChunkData[0:8])

	data := chunks.DataTrackerTimestampChunkData{
		Timestamp: timestamp,
	}

	return chunks.DataTrackerTimestampChunk{
			Chunk: chunk,
			Data:  data,
		},
		nil
}
