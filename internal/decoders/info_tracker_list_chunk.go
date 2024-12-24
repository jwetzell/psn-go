package decoders

type InfoTrackerListChunkData struct {
	Trackers []InfoTrackerChunk
}

type InfoTrackerListChunk struct {
	Chunk Chunk
	Data  InfoTrackerListChunkData
}

func DecodeInfoTrackerListChunk(bytes []byte) (InfoTrackerListChunk, error) {
	chunk, err := DecodeChunk(bytes)
	if err != nil {
		return InfoTrackerListChunk{}, err
	}

	trackers := []InfoTrackerChunk{}
	if chunk.Header.HasSubchunks && chunk.Header.DataLen > 0 {
		offset := 0
		for offset < int(chunk.Header.DataLen) {
			trackerChunk, err := DecodeInfoTrackerChunk(chunk.ChunkData[offset:])
			if err != nil {
				return InfoTrackerListChunk{}, err
			}
			offset += 4
			if trackerChunk.Chunk.Header.DataLen > 0 {
				offset += int(trackerChunk.Chunk.Header.DataLen)
			}
			trackers = append(trackers, trackerChunk)
		}
	}

	data := InfoTrackerListChunkData{
		Trackers: trackers,
	}

	return InfoTrackerListChunk{
			Chunk: chunk,
			Data:  data,
		},
		nil
}
