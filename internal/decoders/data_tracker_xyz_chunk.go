package decoders

import (
	"encoding/binary"
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

func DecodeDataTrackerXYZChunk(bytes []byte) DataTrackerXYZChunk {
	chunk := DecodeChunk(bytes)

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
	}

}
