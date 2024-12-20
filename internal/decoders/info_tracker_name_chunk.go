package decoders

type InfoTrackerNameChunkData struct {
	TrackerName string
}

type InfoTrackerNameChunk struct {
	Chunk Chunk
	Data  InfoTrackerNameChunkData
}

func DecodeInfoTrackerNameChunk(bytes []byte) InfoTrackerNameChunk {
	chunk := DecodeChunk(bytes)

	tracker_name := string(chunk.ChunkData[0:chunk.Header.DataLen])

	data := InfoTrackerNameChunkData{
		TrackerName: tracker_name,
	}

	return InfoTrackerNameChunk{
		Chunk: chunk,
		Data:  data,
	}
}
