package encoders

import (
	"encoding/binary"
	"math"
)

func EncodeDataTrackerPosChunk(x float32, y float32, z float32) []byte {
	bytes := make([]byte, 12)

	binary.LittleEndian.PutUint32(bytes[0:4], math.Float32bits(x))
	binary.LittleEndian.PutUint32(bytes[4:8], math.Float32bits(y))
	binary.LittleEndian.PutUint32(bytes[8:12], math.Float32bits(z))

	return EncodeChunk(0x0000, bytes, false)
}
