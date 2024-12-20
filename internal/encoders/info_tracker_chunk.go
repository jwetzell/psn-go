package encoders

func EncodeInfoTrackerChunk(trackerId uint16, trackerNameChunk []byte) []byte {
	return EncodeChunk(trackerId, trackerNameChunk, true)
}
