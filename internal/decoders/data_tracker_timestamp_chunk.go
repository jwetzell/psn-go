package decoders

import (
	"encoding/binary"
)

type DataTrackerTimestampChunkData struct {
	Timestamp uint64
}

type DataTrackerTimestampChunk struct {
	Chunk Chunk
	Data  DataTrackerTimestampChunkData
}

func DecodeDataTrackerTimestampChunk(bytes []byte) DataTrackerTimestampChunk {
	chunk := DecodeChunk(bytes)

	timestamp := binary.LittleEndian.Uint64(chunk.ChunkData[0:8])

	data := DataTrackerTimestampChunkData{
		Timestamp: timestamp,
	}

	return DataTrackerTimestampChunk{
		Chunk: chunk,
		Data:  data,
	}

}
