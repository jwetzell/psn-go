package chunks

type DataTrackerTimestampChunkData struct {
	Timestamp uint64
}

type DataTrackerTimestampChunk struct {
	Chunk Chunk
	Data  DataTrackerTimestampChunkData
}
