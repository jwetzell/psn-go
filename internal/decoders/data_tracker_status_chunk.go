package decoders

import (
	"encoding/binary"
	"math"
)

type DataTrackerStatusChunkData struct {
	Validity float32
}

type DataTrackerStatusChunk struct {
	Chunk Chunk
	Data  DataTrackerStatusChunkData
}

func DecodeDataTrackerStatusChunk(bytes []byte) DataTrackerStatusChunk {
	chunk := DecodeChunk(bytes)

	statusBits := binary.LittleEndian.Uint32(chunk.ChunkData[0:4])

	data := DataTrackerStatusChunkData{
		Validity: math.Float32frombits(statusBits),
	}

	return DataTrackerStatusChunk{
		Chunk: chunk,
		Data:  data,
	}
}
