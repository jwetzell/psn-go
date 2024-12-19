package decoders

type DataTrackerListChunkData struct {
	Trackers []DataTrackerChunk
}

func decodeDataTrackerListChunkData(dataTrackerListChunk Chunk) DataTrackerListChunkData {
	var trackers []DataTrackerChunk
	if dataTrackerListChunk.Header.HasSubchunks && dataTrackerListChunk.Header.DataLen > 0 {
		offset := 0
		for offset < int(dataTrackerListChunk.Header.DataLen) {
			trackerChunk := DecodeDataTrackerChunk(dataTrackerListChunk.ChunkData[offset:])
			offset += CHUNK_HEADER_SIZE
			if trackerChunk.Chunk.Header.DataLen > 0 {
				offset += int(trackerChunk.Chunk.Header.DataLen)
			}
			trackers = append(trackers, trackerChunk)
		}
	}

	return DataTrackerListChunkData{
		Trackers: trackers,
	}
}

type DataTrackerListChunk struct {
	Chunk Chunk
	Data  DataTrackerListChunkData
}

func DecodeDataTrackerListChunk(bytes []byte) DataTrackerListChunk {
	chunk := DecodeChunk(bytes)
	data := decodeDataTrackerListChunkData(chunk)

	return DataTrackerListChunk{
		Chunk: chunk,
		Data:  data,
	}

}
