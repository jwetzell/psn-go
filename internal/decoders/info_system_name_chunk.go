package decoders

type InfoSystemNameChunkData struct {
	SystemName string
}

type InfoSystemNameChunk struct {
	Chunk Chunk
	Data  InfoSystemNameChunkData
}

func DecodeInfoSystemNameChunk(bytes []byte) (InfoSystemNameChunk, error) {
	chunk, err := DecodeChunk(bytes)
	if err != nil {
		return InfoSystemNameChunk{}, err
	}
	data := InfoSystemNameChunkData{}

	if chunk.Header.DataLen > 0 {
		data.SystemName = string(chunk.ChunkData[0:chunk.Header.DataLen])
	}

	return InfoSystemNameChunk{
		Chunk: chunk,
		Data:  data,
	}, nil
}
