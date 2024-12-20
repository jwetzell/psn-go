package decoders

import "encoding/binary"

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

func DecodePacketHeaderChunk(bytes []byte) PacketHeaderChunk {
	chunk := DecodeChunk(bytes)

	packet_timestamp := binary.LittleEndian.Uint64(chunk.ChunkData[0:8])
	version_high := chunk.ChunkData[8]
	version_low := chunk.ChunkData[9]
	frame_id := chunk.ChunkData[10]
	frame_packet_count := chunk.ChunkData[11]

	data := PacketHeaderChunkData{
		PacketTimestamp:  packet_timestamp,
		VersionHigh:      version_high,
		VersionLow:       version_low,
		FrameId:          frame_id,
		FramePacketCount: frame_packet_count,
	}

	return PacketHeaderChunk{
		Chunk: chunk,
		Data:  data,
	}

}
