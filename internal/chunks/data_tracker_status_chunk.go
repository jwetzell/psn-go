package chunks

type DataTrackerStatusChunkData struct {
	Validity float32
}

type DataTrackerStatusChunk struct {
	Chunk Chunk
	Data  DataTrackerStatusChunkData
}
