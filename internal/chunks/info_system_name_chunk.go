package chunks

type InfoSystemNameChunkData struct {
	SystemName string
}

type InfoSystemNameChunk struct {
	Data  InfoSystemNameChunkData
	Chunk Chunk
}
