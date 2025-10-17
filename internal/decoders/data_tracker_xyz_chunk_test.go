package decoders

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGoodDataTrackerXYZChunk(t *testing.T) {

	testCases := []struct {
		description string
		bytes       []byte
		expected    DataTrackerXYZChunk
	}{
		{
			description: "DataTrackerPosChunk",
			bytes:       []byte{0, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
			expected: DataTrackerXYZChunk{
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
		},
		{
			description: "DataTrackerSpeedChunk",
			bytes:       []byte{1, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
			expected: DataTrackerXYZChunk{
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
		},
		{
			description: "DataTrackerOriChunk",
			bytes:       []byte{2, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
			expected: DataTrackerXYZChunk{
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
		},
		{
			description: "DataTrackerAccelChunk",
			bytes:       []byte{4, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
			expected: DataTrackerXYZChunk{
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
		},
		{
			description: "DataTrackerTrgtPosChunk",
			bytes:       []byte{5, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
			expected: DataTrackerXYZChunk{
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
		},
	}

	for _, testCase := range testCases {

		actual, err := DecodeDataTrackerXYZChunk(testCase.bytes)

		if err != nil {
			t.Errorf("Test '%s' failed to decode chunk properly", testCase.description)
			fmt.Println(err)
		}

		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf("Test '%s' failed to decode chunk properly", testCase.description)
			fmt.Printf("expected: %v\n", testCase.expected)
			fmt.Printf("actual: %v\n", actual)
		}
	}

}
