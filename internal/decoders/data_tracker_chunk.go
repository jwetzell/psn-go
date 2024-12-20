package decoders

import (
	"encoding/binary"
	"log/slog"
)

type DataTrackerChunkData struct {
	Pos       *DataTrackerXYZChunk
	Speed     *DataTrackerXYZChunk
	Ori       *DataTrackerXYZChunk
	Status    *DataTrackerStatusChunk
	Accel     *DataTrackerXYZChunk
	TrgtPos   *DataTrackerXYZChunk
	Timestamp *DataTrackerTimestampChunk
}

type DataTrackerChunk struct {
	Chunk Chunk
	Data  DataTrackerChunkData
}

func DecodeDataTrackerChunk(bytes []byte) DataTrackerChunk {
	chunk := DecodeChunk(bytes)
	data := DataTrackerChunkData{}

	if chunk.Header.HasSubchunks && chunk.ChunkData != nil && chunk.Header.DataLen > 0 {
		offset := 0

		for offset < int(chunk.Header.DataLen) {
			switch id := binary.LittleEndian.Uint16(chunk.ChunkData[offset : offset+2]); id {
			case 0x0000:
				pos := DecodeDataTrackerXYZChunk(chunk.ChunkData[offset:])
				data.Pos = &pos
				offset += 4
				if data.Pos.Chunk.Header.DataLen > 0 {
					offset = offset + int(data.Pos.Chunk.Header.DataLen)
				}
			case 0x0001:
				speed := DecodeDataTrackerXYZChunk(chunk.ChunkData[offset:])
				data.Speed = &speed
				offset += 4
				if data.Speed.Chunk.Header.DataLen > 0 {
					offset = offset + int(data.Speed.Chunk.Header.DataLen)
				}
			case 0x0002:
				ori := DecodeDataTrackerXYZChunk(chunk.ChunkData[offset:])
				data.Ori = &ori
				offset += 4
				if data.Ori.Chunk.Header.DataLen > 0 {
					offset = offset + int(data.Ori.Chunk.Header.DataLen)
				}
			case 0x0003:
				status := DecodeDataTrackerStatusChunk(chunk.ChunkData[offset:])
				data.Status = &status
				offset += 4
				if data.Status.Chunk.Header.DataLen > 0 {
					offset = offset + int(data.Status.Chunk.Header.DataLen)
				}
			case 0x0004:
				accel := DecodeDataTrackerXYZChunk(chunk.ChunkData[offset:])
				data.Accel = &accel
				offset += 4
				if data.Accel.Chunk.Header.DataLen > 0 {
					offset = offset + int(data.Accel.Chunk.Header.DataLen)
				}
			case 0x0005:
				trgtpos := DecodeDataTrackerXYZChunk(chunk.ChunkData[offset:])
				data.TrgtPos = &trgtpos
				offset += 4
				if data.TrgtPos.Chunk.Header.DataLen > 0 {
					offset = offset + int(data.TrgtPos.Chunk.Header.DataLen)
				}
			case 0x0006:
				timestamp := DecodeDataTrackerTimestampChunk(chunk.ChunkData[offset:])
				data.Timestamp = &timestamp
				offset += 4
				if data.Timestamp.Chunk.Header.DataLen > 0 {
					offset = offset + int(data.Timestamp.Chunk.Header.DataLen)
				}
			default:
				offset = int(chunk.Header.DataLen)
				slog.Error("unhandled data tracker packet chunk id", "id", id)
			}
		}
	}

	return DataTrackerChunk{
		Chunk: chunk,
		Data:  data,
	}
}
