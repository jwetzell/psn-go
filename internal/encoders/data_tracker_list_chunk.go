package encoders

func EncodeDataTrackerListChunk(trackerChunks [][]byte) []byte {
	trackerChunkBytes := []byte{}

	for _, trackerChunk := range trackerChunks {
		trackerChunkBytes = append(trackerChunkBytes, trackerChunk...)
	}

	return EncodeChunk(0x0001, trackerChunkBytes, true)
}
