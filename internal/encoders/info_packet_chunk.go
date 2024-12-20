package encoders

func EncodeInfoPacketChunk(packetHeaderChunk []byte, systemNameChunk []byte, trackerListChunk []byte) []byte {
	chunkData := append(packetHeaderChunk, systemNameChunk...)
	chunkData = append(chunkData, trackerListChunk...)

	return EncodeChunk(0x6756, chunkData, true)
}
