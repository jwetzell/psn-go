package decoders

import (
	"encoding/binary"
	"log/slog"
)

type DataPacketChunkData struct {
	PacketHeader *PacketHeaderChunk
	TrackerList  *DataTrackerListChunk
}

func decodeDataPacketChunkData(dataPacketChunk Chunk) DataPacketChunkData {
	data := DataPacketChunkData{}

	if dataPacketChunk.Header.HasSubchunks && dataPacketChunk.ChunkData != nil && dataPacketChunk.Header.DataLen > 0 {
		offset := 0

		for offset < int(dataPacketChunk.Header.DataLen) {

			switch id := binary.LittleEndian.Uint16(dataPacketChunk.ChunkData[offset : offset+2]); id {
			case 0:
				packet_header := DecodePacketHeaderChunk(dataPacketChunk.ChunkData[offset:])
				data.PacketHeader = &packet_header
				offset = offset + CHUNK_HEADER_SIZE
				if packet_header.Chunk.Header.DataLen > 0 {
					offset = offset + int(packet_header.Chunk.Header.DataLen)
				}
			case 1:
				tracker_list := DecodeDataTrackerListChunk(dataPacketChunk.ChunkData[offset:])
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

type DataPacketChunk struct {
	Chunk Chunk
	Data  DataPacketChunkData `json:"data"`
}

func DecodeDataPacketChunk(bytes []byte) DataPacketChunk {
	chunk := DecodeChunk(bytes)
	data := decodeDataPacketChunkData(chunk)

	return DataPacketChunk{
		Chunk: chunk,
		Data:  data,
	}
}
