package encoders

import (
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
	"github.com/stretchr/testify/require"
)

func TestDataTrackerAccelChunkEncoding(t *testing.T) {
	testCases := []struct {
		description string
		expected    []byte
		data        chunks.DataTrackerXYZChunkData
	}{
		{
			description: "basic acceleration",
			expected: []byte{
				4, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64,
			},
			data: chunks.DataTrackerXYZChunkData{
				X: 1,
				Y: 2,
				Z: 3,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := EncodeDataTrackerAccelChunk(testCase.data.X, testCase.data.Y, testCase.data.Z)

			require.Equal(t, testCase.expected, actual)
		})
	}
}
