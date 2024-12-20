package encoders

import (
	"encoding/binary"
)

func EncodeDataTrackerTimestampChunk(timestamp uint64) []byte {
	bytes := make([]byte, 8)

	binary.LittleEndian.PutUint64(bytes, timestamp)

	return EncodeChunk(0x0006, bytes, false)
}
