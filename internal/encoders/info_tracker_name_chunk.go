package encoders

func EncodeInfoTrackerNameChunk(trackerName string) []byte {
	nameBytes := []uint8(trackerName)

	return EncodeChunk(0x0000, nameBytes, false)
}
