package decoders

import (
	"encoding/binary"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func DecodePacketHeaderChunk(bytes []byte) (chunks.PacketHeaderChunk, error) {
	chunk, err := DecodeChunk(bytes)
	if err != nil {
		return chunks.PacketHeaderChunk{}, err
	}

	packet_timestamp := binary.LittleEndian.Uint64(chunk.ChunkData[0:8])
	version_high := chunk.ChunkData[8]
	version_low := chunk.ChunkData[9]
	frame_id := chunk.ChunkData[10]
	frame_packet_count := chunk.ChunkData[11]

	data := chunks.PacketHeaderChunkData{
		PacketTimestamp:  packet_timestamp,
		VersionHigh:      version_high,
		VersionLow:       version_low,
		FrameId:          frame_id,
		FramePacketCount: frame_packet_count,
	}

	return chunks.PacketHeaderChunk{
			Chunk: chunk,
			Data:  data,
		},
		nil
}
