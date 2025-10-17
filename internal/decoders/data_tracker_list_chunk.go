package decoders

import "github.com/jwetzell/psn-go/internal/chunks"

func DecodeDataTrackerListChunk(bytes []byte) (chunks.DataTrackerListChunk, error) {
	chunk, err := DecodeChunk(bytes)

	if err != nil {
		return chunks.DataTrackerListChunk{}, err
	}

	trackers := []chunks.DataTrackerChunk{}
	if chunk.Header.HasSubchunks && chunk.Header.DataLen > 0 {
		offset := 0
		for offset < int(chunk.Header.DataLen) {
			trackerChunk, err := DecodeDataTrackerChunk(chunk.ChunkData[offset:])
			if err != nil {
				return chunks.DataTrackerListChunk{}, err
			}
			offset += 4
			if trackerChunk.Chunk.Header.DataLen > 0 {
				offset += int(trackerChunk.Chunk.Header.DataLen)
			}
			trackers = append(trackers, trackerChunk)
		}
	}

	data := chunks.DataTrackerListChunkData{
		Trackers: trackers,
	}

	return chunks.DataTrackerListChunk{
			Chunk: chunk,
			Data:  data,
		},
		nil
}
