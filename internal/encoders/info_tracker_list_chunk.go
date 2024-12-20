package encoders

func EncodeInfoTrackerListChunk(trackerChunks [][]byte) []byte {
	trackerChunkBytes := []byte{}

	for _, trackerChunk := range trackerChunks {
		trackerChunkBytes = append(trackerChunkBytes, trackerChunk...)
	}

	return EncodeChunk(0x0002, trackerChunkBytes, true)
}
