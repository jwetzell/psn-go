package decoders

import (
	"encoding/binary"
)

type DataTrackerTimestampChunkData struct {
	Timestamp uint64
}

func decodeDataTrackerTimestampChunk(dataTrackerTimestampChunk Chunk) DataTrackerTimestampChunkData {
	timestamp := binary.LittleEndian.Uint64(dataTrackerTimestampChunk.ChunkData[0:8])

	return DataTrackerTimestampChunkData{
		Timestamp: timestamp,
	}
}

type DataTrackerTimestampChunk struct {
	Chunk Chunk
	Data  DataTrackerTimestampChunkData
}

func DecodeDataTrackerTimestampChunk(bytes []byte) DataTrackerTimestampChunk {
	chunk := DecodeChunk(bytes)
	data := decodeDataTrackerTimestampChunk(chunk)

	return DataTrackerTimestampChunk{
		Chunk: chunk,
		Data:  data,
	}

}
