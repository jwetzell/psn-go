package decoders

type InfoSystemNameChunkData struct {
	SystemName string
}

func decodeInfoSystemNameChunkData(infoSystemNameChunk Chunk) InfoSystemNameChunkData {
	system_name := string(infoSystemNameChunk.ChunkData[0:infoSystemNameChunk.Header.DataLen])

	return InfoSystemNameChunkData{
		SystemName: system_name,
	}
}

type InfoSystemNameChunk struct {
	Chunk Chunk
	Data  InfoSystemNameChunkData
}

func DecodeInfoSystemNameChunk(bytes []byte) InfoSystemNameChunk {
	chunk := DecodeChunk(bytes)
	data := decodeInfoSystemNameChunkData(chunk)

	return InfoSystemNameChunk{
		Chunk: chunk,
		Data:  data,
	}

}
