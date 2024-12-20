package encoders

import "encoding/binary"

func EncodeChunk(id uint16, chunkData []byte, hasSubchunks bool) []byte {

	header := []byte{}

	header = binary.LittleEndian.AppendUint16(header, id)

	hasSubchunksBit := 0

	if hasSubchunks {
		hasSubchunksBit = 1
	}

	hasSubchunksBit = hasSubchunksBit << 15

	header = binary.LittleEndian.AppendUint16(header, uint16(len(chunkData)+hasSubchunksBit))

	bytes := header
	bytes = append(bytes, chunkData...)

	return bytes
}
