package chunks

type InfoPacketChunkData struct {
	PacketHeader *PacketHeaderChunk
	SystemName   *InfoSystemNameChunk
	TrackerList  *InfoTrackerListChunk
}

type InfoPacketChunk struct {
	Data  InfoPacketChunkData
	Chunk Chunk
}
