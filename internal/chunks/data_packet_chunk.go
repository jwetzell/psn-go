package chunks

type DataPacketChunkData struct {
	PacketHeader *PacketHeaderChunk
	TrackerList  *DataTrackerListChunk
}

type DataPacketChunk struct {
	Data  DataPacketChunkData
	Chunk Chunk
}
