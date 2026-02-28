package encoders

import (
	"reflect"
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func TestInfoTrackerChunkEncoding(t *testing.T) {
	testCases := []struct {
		description string
		expected    []byte
		chunk       chunks.InfoTrackerChunk
	}{
		{
			description: "InfoTrackerChunk",
			expected: []byte{
				1, 0, 13, 128, 0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49,
			},
			chunk: chunks.InfoTrackerChunk{
				Chunk: chunks.Chunk{
					ChunkData: []byte{0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49},
					Header:    chunks.ChunkHeader{DataLen: 13, Id: 1, HasSubchunks: true},
				},
				Data: chunks.InfoTrackerChunkData{
					TrackerName: &chunks.InfoTrackerNameChunk{
						Chunk: chunks.Chunk{
							ChunkData: []byte{84, 114, 97, 99, 107, 101, 114, 32, 49},
							Header:    chunks.ChunkHeader{DataLen: 9, Id: 0, HasSubchunks: false},
						},
						Data: chunks.InfoTrackerNameChunkData{
							TrackerName: "Tracker 1",
						},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := EncodeInfoTrackerChunk(testCase.chunk.Chunk.Header.Id, EncodeInfoTrackerNameChunk(testCase.chunk.Data.TrackerName.Data.TrackerName))

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("failed to encode chunk properly, expected: %v, actual: %v\n", testCase.expected, actual)
			}
		})
	}
}
