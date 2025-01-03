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
	Data  DataPacketChunkData
	Chunk Chunk
}

func DecodeDataPacketChunk(bytes []byte) (DataPacketChunk, error) {
	chunk, err := DecodeChunk(bytes)

	if err != nil {
		return DataPacketChunk{}, err
	}

	data := DataPacketChunkData{}

	if chunk.Header.HasSubchunks && chunk.ChunkData != nil && chunk.Header.DataLen > 0 {
		offset := 0

		for offset < int(chunk.Header.DataLen) {
			switch id := binary.LittleEndian.Uint16(chunk.ChunkData[offset : offset+2]); id {
			case 0x0000:
				packet_header, err := DecodePacketHeaderChunk(chunk.ChunkData[offset:])
				if err != nil {
					return DataPacketChunk{}, err
				}
				data.PacketHeader = &packet_header
				offset += 4
				if packet_header.Chunk.Header.DataLen > 0 {
					offset = offset + int(packet_header.Chunk.Header.DataLen)
				}
			case 0x0001:
				tracker_list, err := DecodeDataTrackerListChunk(chunk.ChunkData[offset:])
				if err != nil {
					return DataPacketChunk{}, err
				}
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
		},
		nil
}
