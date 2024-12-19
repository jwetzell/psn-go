package decoders

import "encoding/binary"

type PacketHeaderChunkData struct {
	PacketTimestamp  uint64
	VersionHigh      uint8
	VersionLow       uint8
	FrameId          uint8
	FramePacketCount uint8
}

func decodePacketHeaderChunkData(packetHeaderChunk Chunk) PacketHeaderChunkData {
	packet_timestamp := binary.LittleEndian.Uint64(packetHeaderChunk.ChunkData[0:8])
	version_high := packetHeaderChunk.ChunkData[8]
	version_low := packetHeaderChunk.ChunkData[9]
	frame_id := packetHeaderChunk.ChunkData[10]
	frame_packet_count := packetHeaderChunk.ChunkData[11]

	return PacketHeaderChunkData{
		PacketTimestamp:  packet_timestamp,
		VersionHigh:      version_high,
		VersionLow:       version_low,
		FrameId:          frame_id,
		FramePacketCount: frame_packet_count,
	}
}

type PacketHeaderChunk struct {
	Chunk Chunk
	Data  PacketHeaderChunkData
}

func DecodePacketHeaderChunk(bytes []byte) PacketHeaderChunk {
	chunk := DecodeChunk(bytes)
	data := decodePacketHeaderChunkData(chunk)

	return PacketHeaderChunk{
		Chunk: chunk,
		Data:  data,
	}

}
