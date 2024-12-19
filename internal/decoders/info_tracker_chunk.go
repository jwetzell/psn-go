package decoders

import (
	"encoding/binary"
	"log/slog"
)

type InfoTrackerChunkData struct {
	TrackerName *InfoTrackerNameChunk
}

func decodeInfoTrackerChunkData(infoTrackerChunk Chunk) InfoTrackerChunkData {
	data := InfoTrackerChunkData{}

	if infoTrackerChunk.Header.HasSubchunks && infoTrackerChunk.ChunkData != nil && infoTrackerChunk.Header.DataLen > 0 {
		offset := 0

		for offset < int(infoTrackerChunk.Header.DataLen) {

			switch id := binary.LittleEndian.Uint16(infoTrackerChunk.ChunkData[offset : offset+2]); id {
			case 0:
				tracker_name := DecodeInfoTrackerNameChunk(infoTrackerChunk.ChunkData[offset:])
				data.TrackerName = &tracker_name
				offset = offset + CHUNK_HEADER_SIZE
				if tracker_name.Chunk.Header.DataLen > 0 {
					offset = offset + int(tracker_name.Chunk.Header.DataLen)
				}
			default:
				slog.Error("unhandled info tracker chunk id", "id", id)
			}
		}
	}
	return data
}

type InfoTrackerChunk struct {
	Chunk Chunk
	Data  InfoTrackerChunkData
}

func DecodeInfoTrackerChunk(bytes []byte) InfoTrackerChunk {
	chunk := DecodeChunk(bytes)
	data := decodeInfoTrackerChunkData(chunk)

	return InfoTrackerChunk{
		Chunk: chunk,
		Data:  data,
	}
}
