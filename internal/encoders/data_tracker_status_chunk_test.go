package encoders

import (
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
	"github.com/stretchr/testify/require"
)

func TestDataTrackerStatusChunkEncoding(t *testing.T) {
	testCases := []struct {
		description string
		expected    []byte
		data        chunks.DataTrackerStatusChunkData
	}{
		{
			description: "basic status",
			expected: []byte{
				3, 0, 4, 0, 0, 0, 128, 63,
			},
			data: chunks.DataTrackerStatusChunkData{
				Validity: 1,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := EncodeDataTrackerStatusChunk(testCase.data.Validity)

			require.Equal(t, testCase.expected, actual)
		})
	}
}
