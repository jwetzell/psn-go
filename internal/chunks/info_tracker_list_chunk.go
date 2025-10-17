package chunks

type InfoTrackerListChunkData struct {
	Trackers []InfoTrackerChunk
}

type InfoTrackerListChunk struct {
	Chunk Chunk
	Data  InfoTrackerListChunkData
}
