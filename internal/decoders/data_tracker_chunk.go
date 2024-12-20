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

func decodeDataTrackerChunkData(dataTrackerChunk Chunk) DataTrackerChunkData {

	data := DataTrackerChunkData{}

	if dataTrackerChunk.Header.HasSubchunks && dataTrackerChunk.ChunkData != nil && dataTrackerChunk.Header.DataLen > 0 {
		offset := 0

		for offset < int(dataTrackerChunk.Header.DataLen) {

			switch id := binary.LittleEndian.Uint16(dataTrackerChunk.ChunkData[offset : offset+2]); id {
			case 0:
				pos := DecodeDataTrackerXYZChunk(dataTrackerChunk.ChunkData[offset:])
				data.Pos = &pos
				offset += 4
				if data.Pos.Chunk.Header.DataLen > 0 {
					offset = offset + int(data.Pos.Chunk.Header.DataLen)
				}
			case 1:
				speed := DecodeDataTrackerXYZChunk(dataTrackerChunk.ChunkData[offset:])
				data.Speed = &speed
				offset += 4
				if data.Speed.Chunk.Header.DataLen > 0 {
					offset = offset + int(data.Speed.Chunk.Header.DataLen)
				}
			case 2:
				ori := DecodeDataTrackerXYZChunk(dataTrackerChunk.ChunkData[offset:])
				data.Ori = &ori
				offset += 4
				if data.Ori.Chunk.Header.DataLen > 0 {
					offset = offset + int(data.Ori.Chunk.Header.DataLen)
				}
			case 3:
				status := DecodeDataTrackerStatusChunk(dataTrackerChunk.ChunkData[offset:])
				data.Status = &status
				offset += 4
				if data.Status.Chunk.Header.DataLen > 0 {
					offset = offset + int(data.Status.Chunk.Header.DataLen)
				}
			case 4:
				accel := DecodeDataTrackerXYZChunk(dataTrackerChunk.ChunkData[offset:])
				data.Accel = &accel
				offset += 4
				if data.Accel.Chunk.Header.DataLen > 0 {
					offset = offset + int(data.Accel.Chunk.Header.DataLen)
				}
			case 5:
				trgtpos := DecodeDataTrackerXYZChunk(dataTrackerChunk.ChunkData[offset:])
				data.TrgtPos = &trgtpos
				offset += 4
				if data.TrgtPos.Chunk.Header.DataLen > 0 {
					offset = offset + int(data.TrgtPos.Chunk.Header.DataLen)
				}
			case 6:
				timestamp := DecodeDataTrackerTimestampChunk(dataTrackerChunk.ChunkData[offset:])
				data.Timestamp = &timestamp
				offset += 4
				if data.Timestamp.Chunk.Header.DataLen > 0 {
					offset = offset + int(data.Timestamp.Chunk.Header.DataLen)
				}
			default:
				offset = int(dataTrackerChunk.Header.DataLen)
				slog.Error("unhandled data tracker packet chunk id", "id", id)
			}
		}
	}
	return data
}

type DataTrackerChunk struct {
	Chunk Chunk
	Data  DataTrackerChunkData
}

func DecodeDataTrackerChunk(bytes []byte) DataTrackerChunk {
	chunk := DecodeChunk(bytes)
	data := decodeDataTrackerChunkData(chunk)

	return DataTrackerChunk{
		Chunk: chunk,
		Data:  data,
	}
}
