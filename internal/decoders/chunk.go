package decoders

import (
	"encoding/binary"
)

type ChunkHeader struct {
	Id           uint16
	DataLen      uint16
	HasSubchunks bool
}

type Chunk struct {
	Header    ChunkHeader
	ChunkData []byte
}

func DecodeChunk(bytes []byte) Chunk {
	id := binary.LittleEndian.Uint16(bytes[0:2])
	lengthAndFlag := binary.LittleEndian.Uint16(bytes[2:4])

	data_len := lengthAndFlag

	has_subchunks := lengthAndFlag > 32768

	if has_subchunks {
		data_len = data_len - 32768
	}

	header := ChunkHeader{
		Id:           id,
		DataLen:      data_len,
		HasSubchunks: has_subchunks,
	}

	chunk_data := bytes[4 : 4+header.DataLen]

	return Chunk{
		Header:    header,
		ChunkData: chunk_data,
	}
}
