package chunks

type DataTrackerChunkData struct {
	Pos       *DataTrackerXYZChunk
	Speed     *DataTrackerXYZChunk
	Ori       *DataTrackerXYZChunk
	Status    *DataTrackerStatusChunk
	Accel     *DataTrackerXYZChunk
	TrgtPos   *DataTrackerXYZChunk
	Timestamp *DataTrackerTimestampChunk
}

type DataTrackerChunk struct {
	Data  DataTrackerChunkData
	Chunk Chunk
}
