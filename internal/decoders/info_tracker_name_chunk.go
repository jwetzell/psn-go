package decoders

type InfoTrackerNameChunkData struct {
	TrackerName string
}

type InfoTrackerNameChunk struct {
	Data  InfoTrackerNameChunkData
	Chunk Chunk
}

func DecodeInfoTrackerNameChunk(bytes []byte) (InfoTrackerNameChunk, error) {
	chunk, err := DecodeChunk(bytes)

	if err != nil {
		return InfoTrackerNameChunk{}, err
	}

	data := InfoTrackerNameChunkData{}

	if chunk.Header.DataLen > 0 {
		data.TrackerName = string(chunk.ChunkData[0:chunk.Header.DataLen])
	}

	return InfoTrackerNameChunk{
			Chunk: chunk,
			Data:  data,
		},
		nil
}
