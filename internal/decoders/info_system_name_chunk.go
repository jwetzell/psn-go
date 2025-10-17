package decoders

import "github.com/jwetzell/psn-go/internal/chunks"

func DecodeInfoSystemNameChunk(bytes []byte) (chunks.InfoSystemNameChunk, error) {
	chunk, err := DecodeChunk(bytes)
	if err != nil {
		return chunks.InfoSystemNameChunk{}, err
	}
	data := chunks.InfoSystemNameChunkData{}

	if chunk.Header.DataLen > 0 {
		data.SystemName = string(chunk.ChunkData[0:chunk.Header.DataLen])
	}

	return chunks.InfoSystemNameChunk{
		Chunk: chunk,
		Data:  data,
	}, nil
}
