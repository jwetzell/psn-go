package decoders

type DataTrackerListChunkData struct {
	Trackers []DataTrackerChunk
}
type DataTrackerListChunk struct {
	Chunk Chunk
	Data  DataTrackerListChunkData
}

func DecodeDataTrackerListChunk(bytes []byte) DataTrackerListChunk {
	chunk := DecodeChunk(bytes)

	trackers := []DataTrackerChunk{}
	if chunk.Header.HasSubchunks && chunk.Header.DataLen > 0 {
		offset := 0
		for offset < int(chunk.Header.DataLen) {
			trackerChunk := DecodeDataTrackerChunk(chunk.ChunkData[offset:])
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
	}
}
