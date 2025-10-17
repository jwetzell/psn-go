package decoders

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGoodDataTrackerListChunkDecoding(t *testing.T) {
	testCases := []struct {
		description string
		bytes       []byte
		expected    DataTrackerListChunk
	}{
		{
			description: "DataTrackerListChunk",
			bytes: []byte{
				1, 0, 104, 128, 1, 0, 100, 128, 0, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 1, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0,
				64, 0, 0, 64, 64, 2, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 3, 0, 4, 0, 0, 0, 128, 63, 4, 0, 12, 0, 0, 0, 128,
				63, 0, 0, 0, 64, 0, 0, 64, 64, 5, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 6, 0, 8, 0, 210, 2, 150, 73, 0, 0, 0, 0,
			},
			expected: DataTrackerListChunk{
				Chunk: Chunk{
					ChunkData: []byte{
						1, 0, 100, 128, 0, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 1, 0, 12, 0, 0, 0, 128, 63, 0, 0,
						0, 64, 0, 0, 64, 64, 2, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 3, 0, 4, 0, 0, 0, 128, 63, 4,
						0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 5, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64,
						6, 0, 8, 0, 210, 2, 150, 73, 0, 0, 0, 0,
					},
					Header: ChunkHeader{DataLen: 104, Id: 1, HasSubchunks: true},
				},
				Data: DataTrackerListChunkData{
					Trackers: []DataTrackerChunk{
						{
							Chunk: Chunk{
								ChunkData: []byte{
									0, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 1, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0,
									0, 64, 64, 2, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 3, 0, 4, 0, 0, 0, 128, 63, 4, 0,
									12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 5, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64,
									64, 6, 0, 8, 0, 210, 2, 150, 73, 0, 0, 0, 0,
								},
								Header: ChunkHeader{DataLen: 100, Id: 1, HasSubchunks: true},
							},
							Data: DataTrackerChunkData{
								Pos: &DataTrackerXYZChunk{
									Chunk: Chunk{
										ChunkData: []byte{0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
										Header:    ChunkHeader{DataLen: 12, Id: 0, HasSubchunks: false},
									},
									Data: DataTrackerXYZChunkData{
										X: 1.0,
										Y: 2.0,
										Z: 3.0,
									},
								},
								Speed: &DataTrackerXYZChunk{
									Chunk: Chunk{
										ChunkData: []byte{0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
										Header:    ChunkHeader{DataLen: 12, Id: 1, HasSubchunks: false},
									},
									Data: DataTrackerXYZChunkData{
										X: 1.0,
										Y: 2.0,
										Z: 3.0,
									},
								},
								Ori: &DataTrackerXYZChunk{
									Chunk: Chunk{
										ChunkData: []byte{0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
										Header:    ChunkHeader{DataLen: 12, Id: 2, HasSubchunks: false},
									},
									Data: DataTrackerXYZChunkData{
										X: 1.0,
										Y: 2.0,
										Z: 3.0,
									},
								},
								Status: &DataTrackerStatusChunk{
									Chunk: Chunk{
										ChunkData: []byte{0, 0, 128, 63},
										Header:    ChunkHeader{DataLen: 4, Id: 3, HasSubchunks: false},
									},
									Data: DataTrackerStatusChunkData{
										Validity: 1.0,
									},
								},
								Accel: &DataTrackerXYZChunk{
									Chunk: Chunk{
										ChunkData: []byte{0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
										Header:    ChunkHeader{DataLen: 12, Id: 4, HasSubchunks: false},
									},
									Data: DataTrackerXYZChunkData{
										X: 1.0,
										Y: 2.0,
										Z: 3.0,
									},
								},
								TrgtPos: &DataTrackerXYZChunk{
									Chunk: Chunk{
										ChunkData: []byte{0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
										Header:    ChunkHeader{DataLen: 12, Id: 5, HasSubchunks: false},
									},
									Data: DataTrackerXYZChunkData{
										X: 1.0,
										Y: 2.0,
										Z: 3.0,
									},
								},
								Timestamp: &DataTrackerTimestampChunk{
									Chunk: Chunk{
										ChunkData: []byte{210, 2, 150, 73, 0, 0, 0, 0},
										Header:    ChunkHeader{DataLen: 8, Id: 6, HasSubchunks: false},
									},
									Data: DataTrackerTimestampChunkData{
										Timestamp: 1234567890,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {

		actual, err := DecodeDataTrackerListChunk(testCase.bytes)

		if err != nil {
			t.Errorf("Test '%s' failed to decode chunk properly", testCase.description)
			fmt.Println(err)
		}

		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf("Test '%s' failed to decode chunk properly", testCase.description)
			fmt.Printf("expected: %+v\n", testCase.expected)
			fmt.Printf("actual: %+v\n", actual)
		}
	}
}
