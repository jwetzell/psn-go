package chunks

type DataTrackerXYZChunkData struct {
	X float32
	Y float32
	Z float32
}

type DataTrackerXYZChunk struct {
	Chunk Chunk
	Data  DataTrackerXYZChunkData
}
