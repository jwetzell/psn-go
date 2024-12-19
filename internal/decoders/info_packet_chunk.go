package decoders

import (
	"encoding/binary"
	"log/slog"
)

type InfoPacketChunkData struct {
	PacketHeader *PacketHeaderChunk
	SystemName   *InfoSystemNameChunk
	TrackerList  *InfoTrackerListChunk
}

type EmptyData struct{}

func decodeInfoPacketChunkData(infoPacketChunk Chunk) InfoPacketChunkData {
	data := InfoPacketChunkData{}

	if infoPacketChunk.Header.HasSubchunks && infoPacketChunk.ChunkData != nil && infoPacketChunk.Header.DataLen > 0 {
		offset := 0

		for offset < int(infoPacketChunk.Header.DataLen) {

			switch id := binary.LittleEndian.Uint16(infoPacketChunk.ChunkData[offset : offset+2]); id {
			case 0:
				packet_header := DecodePacketHeaderChunk(infoPacketChunk.ChunkData[offset:])
				data.PacketHeader = &packet_header
				offset = offset + CHUNK_HEADER_SIZE
				if packet_header.Chunk.Header.DataLen > 0 {
					offset = offset + int(packet_header.Chunk.Header.DataLen)
				}
			case 1:
				system_name := DecodeInfoSystemNameChunk(infoPacketChunk.ChunkData[offset:])
				data.SystemName = &system_name
				offset = offset + CHUNK_HEADER_SIZE
				if system_name.Chunk.Header.DataLen > 0 {
					offset = offset + int(system_name.Chunk.Header.DataLen)
				}
			case 2:
				tracker_list := DecodeInfoTrackerListChunk(infoPacketChunk.ChunkData[offset:])
				data.TrackerList = &tracker_list
				offset = offset + CHUNK_HEADER_SIZE
				if tracker_list.Chunk.Header.DataLen > 0 {
					offset = offset + int(tracker_list.Chunk.Header.DataLen)
				}

			default:
				slog.Error("unhandled info packet id", "id", id)
			}
		}
	}
	return data
}

type InfoPacketChunk struct {
	Chunk Chunk
	Data  InfoPacketChunkData
}

func DecodeInfoPacketChunk(bytes []byte) InfoPacketChunk {
	chunk := DecodeChunk(bytes)
	data := decodeInfoPacketChunkData(chunk)

	return InfoPacketChunk{
		Chunk: chunk,
		Data:  data,
	}
}
