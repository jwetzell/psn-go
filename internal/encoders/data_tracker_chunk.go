package encoders

func EncodeDataTrackerChunk(trackerId uint16, fieldChunks [][]byte) []byte {
	fieldChunkBytes := []byte{}

	for _, fieldChunk := range fieldChunks {
		fieldChunkBytes = append(fieldChunkBytes, fieldChunk...)
	}

	return EncodeChunk(trackerId, fieldChunkBytes, len(fieldChunks) > 0)
}
