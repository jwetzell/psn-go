package encoders

func EncodeInfoSystemNameChunk(systemName string) []byte {
	nameBytes := []uint8(systemName)

	return EncodeChunk(0x0001, nameBytes, false)
}
