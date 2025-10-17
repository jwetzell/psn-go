package decoders

import (
	"encoding/binary"
	"log/slog"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func DecodeInfoTrackerChunk(bytes []byte) (chunks.InfoTrackerChunk, error) {
	chunk, err := DecodeChunk(bytes)
	if err != nil {
		return chunks.InfoTrackerChunk{}, err
	}
	data := chunks.InfoTrackerChunkData{}

	if chunk.Header.HasSubchunks && chunk.ChunkData != nil && chunk.Header.DataLen > 0 {
		offset := 0

		for offset < int(chunk.Header.DataLen) {
			switch id := binary.LittleEndian.Uint16(chunk.ChunkData[offset : offset+2]); id {
			case 0x0000:
				tracker_name, err := DecodeInfoTrackerNameChunk(chunk.ChunkData[offset:])
				if err != nil {
					return chunks.InfoTrackerChunk{}, err
				}
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

	return chunks.InfoTrackerChunk{
			Chunk: chunk,
			Data:  data,
		},
		nil
}
