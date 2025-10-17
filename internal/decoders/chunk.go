package decoders

import (
	"encoding/binary"
	"errors"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func DecodeChunk(bytes []byte) (chunks.Chunk, error) {

	if len(bytes) < 4 {
		return chunks.Chunk{}, errors.New("chunk must be at least 4 bytes")
	}

	id := binary.LittleEndian.Uint16(bytes[0:2])
	lengthAndFlag := binary.LittleEndian.Uint16(bytes[2:4])

	data_len := lengthAndFlag

	has_subchunks := lengthAndFlag > 32768

	if has_subchunks {
		data_len = data_len - 32768
	}

	header := chunks.ChunkHeader{
		Id:           id,
		DataLen:      data_len,
		HasSubchunks: has_subchunks,
	}

	chunk_data := bytes[4 : 4+header.DataLen]

	return chunks.Chunk{
		Header:    header,
		ChunkData: chunk_data,
	}, nil
}
