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

func decodeDataTrackerXYZChunk(dataTrackerXYZChunk Chunk) DataTrackerXYZChunkData {
	xBits := binary.LittleEndian.Uint32(dataTrackerXYZChunk.ChunkData[0:4])
	yBits := binary.LittleEndian.Uint32(dataTrackerXYZChunk.ChunkData[4:8])
	zBits := binary.LittleEndian.Uint32(dataTrackerXYZChunk.ChunkData[8:12])

	return DataTrackerXYZChunkData{
		X: math.Float32frombits(xBits),
		Y: math.Float32frombits(yBits),
		Z: math.Float32frombits(zBits),
	}
}

type DataTrackerXYZChunk struct {
	Chunk Chunk
	Data  DataTrackerXYZChunkData
}

func DecodeDataTrackerXYZChunk(bytes []byte) DataTrackerXYZChunk {
	chunk := DecodeChunk(bytes)
	data := decodeDataTrackerXYZChunk(chunk)

	return DataTrackerXYZChunk{
		Chunk: chunk,
		Data:  data,
	}

}
