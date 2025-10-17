package decoders

import "github.com/jwetzell/psn-go/internal/chunks"

func DecodeInfoTrackerListChunk(bytes []byte) (chunks.InfoTrackerListChunk, error) {
	chunk, err := DecodeChunk(bytes)
	if err != nil {
		return chunks.InfoTrackerListChunk{}, err
	}

	trackers := []chunks.InfoTrackerChunk{}
	if chunk.Header.HasSubchunks && chunk.Header.DataLen > 0 {
		offset := 0
		for offset < int(chunk.Header.DataLen) {
			trackerChunk, err := DecodeInfoTrackerChunk(chunk.ChunkData[offset:])
			if err != nil {
				return chunks.InfoTrackerListChunk{}, err
			}
			offset += 4
			if trackerChunk.Chunk.Header.DataLen > 0 {
				offset += int(trackerChunk.Chunk.Header.DataLen)
			}
			trackers = append(trackers, trackerChunk)
		}
	}

	data := chunks.InfoTrackerListChunkData{
		Trackers: trackers,
	}

	return chunks.InfoTrackerListChunk{
			Chunk: chunk,
			Data:  data,
		},
		nil
}
