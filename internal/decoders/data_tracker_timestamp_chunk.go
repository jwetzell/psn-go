package decoders

import (
	"encoding/binary"
	"errors"
)

type DataTrackerTimestampChunkData struct {
	Timestamp uint64
}

type DataTrackerTimestampChunk struct {
	Chunk Chunk
	Data  DataTrackerTimestampChunkData
}

func DecodeDataTrackerTimestampChunk(bytes []byte) (DataTrackerTimestampChunk, error) {
	chunk, err := DecodeChunk(bytes)

	if err != nil {
		return DataTrackerTimestampChunk{}, err
	}

	if len(chunk.ChunkData) < 8 {
		return DataTrackerTimestampChunk{}, errors.New("DATA_TRACKER_TIMESTAMP chunk must be at least 8 bytes")
	}

	timestamp := binary.LittleEndian.Uint64(chunk.ChunkData[0:8])

	data := DataTrackerTimestampChunkData{
		Timestamp: timestamp,
	}

	return DataTrackerTimestampChunk{
			Chunk: chunk,
			Data:  data,
		},
		nil
}
