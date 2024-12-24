package decoders

type DataTrackerListChunkData struct {
	Trackers []DataTrackerChunk
}
type DataTrackerListChunk struct {
	Chunk Chunk
	Data  DataTrackerListChunkData
}

func DecodeDataTrackerListChunk(bytes []byte) (DataTrackerListChunk, error) {
	chunk, err := DecodeChunk(bytes)

	if err != nil {
		return DataTrackerListChunk{}, err
	}

	trackers := []DataTrackerChunk{}
	if chunk.Header.HasSubchunks && chunk.Header.DataLen > 0 {
		offset := 0
		for offset < int(chunk.Header.DataLen) {
			trackerChunk, err := DecodeDataTrackerChunk(chunk.ChunkData[offset:])
			if err != nil {
				return DataTrackerListChunk{}, err
			}
			offset += 4
			if trackerChunk.Chunk.Header.DataLen > 0 {
				offset += int(trackerChunk.Chunk.Header.DataLen)
			}
			trackers = append(trackers, trackerChunk)
		}
	}

	data := DataTrackerListChunkData{
		Trackers: trackers,
	}

	return DataTrackerListChunk{
			Chunk: chunk,
			Data:  data,
		},
		nil
}
