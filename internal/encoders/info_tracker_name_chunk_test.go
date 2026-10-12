package encoders

import (
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
	"github.com/stretchr/testify/require"
)

func TestInfoTrackerNameChunkEncoding(t *testing.T) {
	testCases := []struct {
		description string
		expected    []byte
		data        chunks.InfoTrackerNameChunkData
	}{
		{
			description: "InfoTrackerNameChunk",
			expected: []byte{
				0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49,
			},
			data: chunks.InfoTrackerNameChunkData{
				TrackerName: "Tracker 1",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := EncodeInfoTrackerNameChunk(testCase.data.TrackerName)

			require.Equal(t, testCase.expected, actual)
		})
	}
}
