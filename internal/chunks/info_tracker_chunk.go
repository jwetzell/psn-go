package chunks

type InfoTrackerChunkData struct {
	TrackerName *InfoTrackerNameChunk
}

type InfoTrackerChunk struct {
	Data  InfoTrackerChunkData
	Chunk Chunk
}
