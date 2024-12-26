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

type InfoPacketChunk struct {
	Data  InfoPacketChunkData
	Chunk Chunk
}

func DecodeInfoPacketChunk(bytes []byte) (InfoPacketChunk, error) {
	chunk, err := DecodeChunk(bytes)
	if err != nil {
		return InfoPacketChunk{}, err
	}
	data := InfoPacketChunkData{}

	if chunk.Header.HasSubchunks && chunk.ChunkData != nil && chunk.Header.DataLen > 0 {
		offset := 0

		for offset < int(chunk.Header.DataLen) {
			switch id := binary.LittleEndian.Uint16(chunk.ChunkData[offset : offset+2]); id {
			case 0x0000:
				packet_header, err := DecodePacketHeaderChunk(chunk.ChunkData[offset:])
				if err != nil {
					return InfoPacketChunk{}, err
				}
				data.PacketHeader = &packet_header
				offset += 4
				if packet_header.Chunk.Header.DataLen > 0 {
					offset = offset + int(packet_header.Chunk.Header.DataLen)
				}
			case 0x0001:
				system_name, err := DecodeInfoSystemNameChunk(chunk.ChunkData[offset:])
				if err != nil {
					return InfoPacketChunk{}, err
				}
				data.SystemName = &system_name
				offset += 4
				if system_name.Chunk.Header.DataLen > 0 {
					offset = offset + int(system_name.Chunk.Header.DataLen)
				}
			case 0x0002:
				tracker_list, err := DecodeInfoTrackerListChunk(chunk.ChunkData[offset:])
				if err != nil {
					return InfoPacketChunk{}, err
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

	return InfoPacketChunk{
			Chunk: chunk,
			Data:  data,
		},
		nil
}
