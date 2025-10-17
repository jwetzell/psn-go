package decoders

import "github.com/jwetzell/psn-go/internal/chunks"

func DecodeInfoTrackerNameChunk(bytes []byte) (chunks.InfoTrackerNameChunk, error) {
	chunk, err := DecodeChunk(bytes)

	if err != nil {
		return chunks.InfoTrackerNameChunk{}, err
	}

	data := chunks.InfoTrackerNameChunkData{}

	if chunk.Header.DataLen > 0 {
		data.TrackerName = string(chunk.ChunkData[0:chunk.Header.DataLen])
	}

	return chunks.InfoTrackerNameChunk{
			Chunk: chunk,
			Data:  data,
		},
		nil
}
