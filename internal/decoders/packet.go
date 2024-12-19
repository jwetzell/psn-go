package decoders

import (
	"encoding/binary"
	"errors"
	"log/slog"
)

func DecodePacketChunk(bytes []byte) (interface{}, error) {
	switch id := binary.LittleEndian.Uint16(bytes[0:2]); id {
	case 0x6756:
		return DecodeInfoPacketChunk(bytes), nil
	case 0x6755:
		return DecodeDataPacketChunk(bytes), nil
	default:
		slog.Error("unhandled packet id", "id", id)
		return Chunk{}, errors.New("unhandled packet id")
	}

}
