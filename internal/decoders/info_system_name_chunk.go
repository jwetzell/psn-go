package decoders

type InfoSystemNameChunkData struct {
	SystemName string
}

type InfoSystemNameChunk struct {
	Chunk Chunk
	Data  InfoSystemNameChunkData
}

func DecodeInfoSystemNameChunk(bytes []byte) InfoSystemNameChunk {
	chunk := DecodeChunk(bytes)
	system_name := string(chunk.ChunkData[0:chunk.Header.DataLen])

	data := InfoSystemNameChunkData{
		SystemName: system_name,
	}

	return InfoSystemNameChunk{
		Chunk: chunk,
		Data:  data,
	}

}
