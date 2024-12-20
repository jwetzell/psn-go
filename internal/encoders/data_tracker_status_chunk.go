package encoders

import (
	"encoding/binary"
	"math"
)

func EncodeDataTrackerStatusChunk(validity float32) []byte {
	bytes := make([]byte, 4)

	binary.LittleEndian.PutUint32(bytes, math.Float32bits(validity))

	return EncodeChunk(0x0003, bytes, false)
}
