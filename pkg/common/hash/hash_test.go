package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	secret  string
	hash    string
	want    bool
	message string
}

var correctResultMessage = "The secrets should match"
var incorrectResultMessage = "The secrets should not math"

func TestCompareSecrets(t *testing.T) {
	secret1 := "12345678"
	hash1, _ := New(secret1)

	secret2 := "1234567890"
	hash2, _ := New("othersecret")

	tCases := []testCase{
		{
			secret:  secret1,
			hash:    hash1,
			want:    true,
			message: correctResultMessage,
		},
		{
			secret:  secret2,
			hash:    hash2,
			want:    false,
			message: incorrectResultMessage,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.message, func(t *testing.T) {
			assert.Equal(t, tc.want, CompareSecrets(tc.secret, tc.hash))
		})
	}
}
