package decoders

type InfoTrackerListChunkData struct {
	Trackers []InfoTrackerChunk
}

func decodeInfoTrackerListChunkData(infoTrackerListChunk Chunk) InfoTrackerListChunkData {
	var trackers []InfoTrackerChunk
	if infoTrackerListChunk.Header.HasSubchunks && infoTrackerListChunk.Header.DataLen > 0 {
		offset := 0
		for offset < int(infoTrackerListChunk.Header.DataLen) {
			trackerChunk := DecodeInfoTrackerChunk(infoTrackerListChunk.ChunkData[offset:])
			offset += CHUNK_HEADER_SIZE
			if trackerChunk.Chunk.Header.DataLen > 0 {
				offset += int(trackerChunk.Chunk.Header.DataLen)
			}
			trackers = append(trackers, trackerChunk)
		}
	}

	return InfoTrackerListChunkData{
		Trackers: trackers,
	}
}

type InfoTrackerListChunk struct {
	Chunk Chunk
	Data  InfoTrackerListChunkData
}

func DecodeInfoTrackerListChunk(bytes []byte) InfoTrackerListChunk {
	chunk := DecodeChunk(bytes)
	data := decodeInfoTrackerListChunkData(chunk)

	return InfoTrackerListChunk{
		Chunk: chunk,
		Data:  data,
	}

}
