package decoders

import (
	"encoding/binary"
	"log/slog"
)

type InfoTrackerChunkData struct {
	TrackerName *InfoTrackerNameChunk
}

type InfoTrackerChunk struct {
	Chunk Chunk
	Data  InfoTrackerChunkData
}

func DecodeInfoTrackerChunk(bytes []byte) InfoTrackerChunk {
	chunk := DecodeChunk(bytes)
	data := InfoTrackerChunkData{}

	if chunk.Header.HasSubchunks && chunk.ChunkData != nil && chunk.Header.DataLen > 0 {
		offset := 0

		for offset < int(chunk.Header.DataLen) {
			switch id := binary.LittleEndian.Uint16(chunk.ChunkData[offset : offset+2]); id {
			case 0x0000:
				tracker_name := DecodeInfoTrackerNameChunk(chunk.ChunkData[offset:])
				data.TrackerName = &tracker_name
				offset += 4
				if tracker_name.Chunk.Header.DataLen > 0 {
					offset = offset + int(tracker_name.Chunk.Header.DataLen)
				}
			default:
				slog.Error("unhandled info tracker chunk id", "id", id)
			}
		}
	}

	return InfoTrackerChunk{
		Chunk: chunk,
		Data:  data,
	}
}
