package decoders

import (
	"encoding/binary"
	"log/slog"
)

type DataPacketChunkData struct {
	PacketHeader *PacketHeaderChunk
	TrackerList  *DataTrackerListChunk
}

type DataPacketChunk struct {
	Chunk Chunk
	Data  DataPacketChunkData
}

func DecodeDataPacketChunk(bytes []byte) DataPacketChunk {
	chunk := DecodeChunk(bytes)
	data := DataPacketChunkData{}

	if chunk.Header.HasSubchunks && chunk.ChunkData != nil && chunk.Header.DataLen > 0 {
		offset := 0

		for offset < int(chunk.Header.DataLen) {

			switch id := binary.LittleEndian.Uint16(chunk.ChunkData[offset : offset+2]); id {
			case 0:
				packet_header := DecodePacketHeaderChunk(chunk.ChunkData[offset:])
				data.PacketHeader = &packet_header
				offset += 4
				if packet_header.Chunk.Header.DataLen > 0 {
					offset = offset + int(packet_header.Chunk.Header.DataLen)
				}
			case 1:
				tracker_list := DecodeDataTrackerListChunk(chunk.ChunkData[offset:])
				data.TrackerList = &tracker_list
				offset += 4
				if tracker_list.Chunk.Header.DataLen > 0 {
					offset = offset + int(tracker_list.Chunk.Header.DataLen)
				}

			default:
				slog.Error("unhandled info packet id", "id", id)
			}
		}
	}

	return DataPacketChunk{
		Chunk: chunk,
		Data:  data,
	}
}
