package decoders

import (
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
	"github.com/stretchr/testify/require"
)

func TestGoodDataTrackerTimestampChunk(t *testing.T) {

	testCases := []struct {
		description string
		bytes       []byte
		expected    chunks.DataTrackerTimestampChunk
	}{
		{
			description: "DataTrackerTimestampChunk",
			bytes:       []byte{6, 0, 8, 0, 210, 2, 150, 73, 0, 0, 0, 0},
			expected: chunks.DataTrackerTimestampChunk{
				Chunk: chunks.Chunk{
					ChunkData: []byte{210, 2, 150, 73, 0, 0, 0, 0},
					Header:    chunks.ChunkHeader{DataLen: 8, Id: 6, HasSubchunks: false},
				},
				Data: chunks.DataTrackerTimestampChunkData{
					Timestamp: 1234567890,
				},
			},
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.description, func(t *testing.T) {
			actual, err := DecodeDataTrackerTimestampChunk(testCase.bytes)

			require.NoError(t, err)

			require.Equal(t, testCase.expected, actual)
		})
	}

}
