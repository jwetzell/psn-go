package decoders

import (
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
	"github.com/stretchr/testify/require"
)

func TestGoodInfoTrackerNameChunkDecoding(t *testing.T) {
	testCases := []struct {
		description string
		bytes       []byte
		expected    chunks.InfoTrackerNameChunk
	}{
		{
			description: "InfoTrackerNameChunk",
			bytes: []byte{
				0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49,
			},
			expected: chunks.InfoTrackerNameChunk{
				Chunk: chunks.Chunk{
					ChunkData: []byte{84, 114, 97, 99, 107, 101, 114, 32, 49},
					Header:    chunks.ChunkHeader{DataLen: 9, Id: 0, HasSubchunks: false},
				},
				Data: chunks.InfoTrackerNameChunkData{
					TrackerName: "Tracker 1",
				},
			},
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.description, func(t *testing.T) {
			actual, err := DecodeInfoTrackerNameChunk(testCase.bytes)

			require.NoError(t, err)
			require.Equal(t, testCase.expected, actual)
		})
	}
}
