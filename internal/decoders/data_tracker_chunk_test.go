package decoders

import (
	"reflect"
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func TestGoodDataTrackerChunkDecoding(t *testing.T) {
	testCases := []struct {
		description string
		bytes       []byte
		expected    chunks.DataTrackerChunk
	}{
		{
			description: "DataTracker",
			bytes: []byte{
				1, 0, 100, 128, 0, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 1, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 2, 0, 12,
				0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 3, 0, 4, 0, 0, 0, 128, 63, 4, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64,
				0, 0, 64, 64, 5, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 6, 0, 8, 0, 210, 2, 150, 73, 0, 0, 0, 0,
			},
			expected: chunks.DataTrackerChunk{
				Chunk: chunks.Chunk{
					ChunkData: []byte{
						0, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 1, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0,
						0, 64, 64, 2, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 3, 0, 4, 0, 0, 0, 128, 63, 4, 0,
						12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 5, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64,
						64, 6, 0, 8, 0, 210, 2, 150, 73, 0, 0, 0, 0,
					},
					Header: chunks.ChunkHeader{DataLen: 100, Id: 1, HasSubchunks: true},
				},
				Data: chunks.DataTrackerChunkData{
					Pos: &chunks.DataTrackerXYZChunk{
						Chunk: chunks.Chunk{
							ChunkData: []byte{0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
							Header:    chunks.ChunkHeader{DataLen: 12, Id: 0, HasSubchunks: false},
						},
						Data: chunks.DataTrackerXYZChunkData{
							X: 1.0,
							Y: 2.0,
							Z: 3.0,
						},
					},
					Speed: &chunks.DataTrackerXYZChunk{
						Chunk: chunks.Chunk{
							ChunkData: []byte{0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
							Header:    chunks.ChunkHeader{DataLen: 12, Id: 1, HasSubchunks: false},
						},
						Data: chunks.DataTrackerXYZChunkData{
							X: 1.0,
							Y: 2.0,
							Z: 3.0,
						},
					},
					Ori: &chunks.DataTrackerXYZChunk{
						Chunk: chunks.Chunk{
							ChunkData: []byte{0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
							Header:    chunks.ChunkHeader{DataLen: 12, Id: 2, HasSubchunks: false},
						},
						Data: chunks.DataTrackerXYZChunkData{
							X: 1.0,
							Y: 2.0,
							Z: 3.0,
						},
					},
					Status: &chunks.DataTrackerStatusChunk{
						Chunk: chunks.Chunk{
							ChunkData: []byte{0, 0, 128, 63},
							Header:    chunks.ChunkHeader{DataLen: 4, Id: 3, HasSubchunks: false},
						},
						Data: chunks.DataTrackerStatusChunkData{
							Validity: 1.0,
						},
					},
					Accel: &chunks.DataTrackerXYZChunk{
						Chunk: chunks.Chunk{
							ChunkData: []byte{0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
							Header:    chunks.ChunkHeader{DataLen: 12, Id: 4, HasSubchunks: false},
						},
						Data: chunks.DataTrackerXYZChunkData{
							X: 1.0,
							Y: 2.0,
							Z: 3.0,
						},
					},
					TrgtPos: &chunks.DataTrackerXYZChunk{
						Chunk: chunks.Chunk{
							ChunkData: []byte{0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
							Header:    chunks.ChunkHeader{DataLen: 12, Id: 5, HasSubchunks: false},
						},
						Data: chunks.DataTrackerXYZChunkData{
							X: 1.0,
							Y: 2.0,
							Z: 3.0,
						},
					},
					Timestamp: &chunks.DataTrackerTimestampChunk{
						Chunk: chunks.Chunk{
							ChunkData: []byte{210, 2, 150, 73, 0, 0, 0, 0},
							Header:    chunks.ChunkHeader{DataLen: 8, Id: 6, HasSubchunks: false},
						},
						Data: chunks.DataTrackerTimestampChunkData{
							Timestamp: 1234567890,
						},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual, err := DecodeDataTrackerChunk(testCase.bytes)

			if err != nil {
				t.Errorf("failed to decode chunk properly, error: %v", err)
			}

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("failed to decode chunk properly, expected: %+v, actual: %+v", testCase.expected, actual)
			}
		})
	}
}
