package chunks

type DataTrackerListChunkData struct {
	Trackers []DataTrackerChunk
}
type DataTrackerListChunk struct {
	Chunk Chunk
	Data  DataTrackerListChunkData
}
