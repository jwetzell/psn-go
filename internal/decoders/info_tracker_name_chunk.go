package decoders

type InfoTrackerNameChunkData struct {
	TrackerName string
}

func decodeInfoTrackerNameChunkData(infoTrackerNameChunk Chunk) InfoTrackerNameChunkData {
	tracker_name := string(infoTrackerNameChunk.ChunkData[0:infoTrackerNameChunk.Header.DataLen])

	return InfoTrackerNameChunkData{
		TrackerName: tracker_name,
	}
}

type InfoTrackerNameChunk struct {
	Chunk Chunk
	Data  InfoTrackerNameChunkData
}

func DecodeInfoTrackerNameChunk(bytes []byte) InfoTrackerNameChunk {
	chunk := DecodeChunk(bytes)
	data := decodeInfoTrackerNameChunkData(chunk)

	return InfoTrackerNameChunk{
		Chunk: chunk,
		Data:  data,
	}

}
