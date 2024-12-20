package encoders

import "encoding/binary"

func EncodePacketHeaderChunk(timestamp uint64, versionHigh uint8, versionLow uint8, frameId uint8, framePacketCount uint8) []byte {
	packetHeader := make([]byte, 12)

	binary.LittleEndian.PutUint64(packetHeader, timestamp)
	packetHeader[8] = versionHigh
	packetHeader[9] = versionLow
	packetHeader[10] = frameId
	packetHeader[11] = framePacketCount
	return EncodeChunk(0x0000, packetHeader, false)
}
