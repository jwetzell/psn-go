package decoders

import (
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
	"github.com/stretchr/testify/require"
)

func TestGoodDataTrackerStatusChunk(t *testing.T) {

	testCases := []struct {
		description string
		bytes       []byte
		expected    chunks.DataTrackerStatusChunk
	}{
		{
			description: "DataTrackerStatusChunk",
			bytes:       []byte{3, 0, 4, 0, 0, 0, 128, 63},
			expected: chunks.DataTrackerStatusChunk{
				Chunk: chunks.Chunk{
					ChunkData: []byte{0, 0, 128, 63},
					Header:    chunks.ChunkHeader{DataLen: 4, Id: 3, HasSubchunks: false},
				},
				Data: chunks.DataTrackerStatusChunkData{
					Validity: 1.0,
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual, err := DecodeDataTrackerStatusChunk(testCase.bytes)

			require.NoError(t, err)

			require.Equal(t, testCase.expected, actual)
		})
	}

}
