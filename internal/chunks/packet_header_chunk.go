package chunks

type PacketHeaderChunkData struct {
	PacketTimestamp  uint64
	VersionHigh      uint8
	VersionLow       uint8
	FrameId          uint8
	FramePacketCount uint8
}

type PacketHeaderChunk struct {
	Chunk Chunk
	Data  PacketHeaderChunkData
}
