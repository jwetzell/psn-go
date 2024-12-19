package decoders

import (
	"encoding/binary"
	"math"
)

type DataTrackerStatusChunkData struct {
	Validity float32
}

func decodeDataTrackerStatusChunk(dataTrackerStatusChunk Chunk) DataTrackerStatusChunkData {
	statusBits := binary.LittleEndian.Uint32(dataTrackerStatusChunk.ChunkData[0:4])

	return DataTrackerStatusChunkData{
		Validity: math.Float32frombits(statusBits),
	}
}

type DataTrackerStatusChunk struct {
	Chunk Chunk
	Data  DataTrackerStatusChunkData
}

func DecodeDataTrackerStatusChunk(bytes []byte) DataTrackerStatusChunk {
	chunk := DecodeChunk(bytes)
	data := decodeDataTrackerStatusChunk(chunk)

	return DataTrackerStatusChunk{
		Chunk: chunk,
		Data:  data,
	}

}
