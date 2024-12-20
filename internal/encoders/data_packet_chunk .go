package encoders

func EncodeDataPacketChunk(packetHeaderChunk []byte, trackerListChunk []byte) []byte {
	chunkData := append(packetHeaderChunk, trackerListChunk...)

	return EncodeChunk(0x6755, chunkData, true)
}
