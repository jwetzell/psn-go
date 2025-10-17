package chunks

type InfoTrackerNameChunkData struct {
	TrackerName string
}

type InfoTrackerNameChunk struct {
	Data  InfoTrackerNameChunkData
	Chunk Chunk
}
