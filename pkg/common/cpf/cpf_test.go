package cpf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	cpf     string
	want    bool
	message string
}

var correctResultMessage string = "Should validate as a correct cpf"
var incorrectResultMessage string = "Should validate as a incorrect cpf"

func TestValidate(t *testing.T) {
	tCases := []testCase{
		{
			cpf:     "390.910.470-38",
			want:    true,
			message: correctResultMessage,
		},
		//? wrong last digit
		{
			cpf:     "989.477.100-99",
			want:    false,
			message: incorrectResultMessage,
		},
		//? wrong first digit after '-'
		{
			cpf:     "989.477.100-71",
			want:    false,
			message: incorrectResultMessage,
		},
		{
			cpf:     "39091047038",
			want:    false,
			message: incorrectResultMessage,
		},
		{
			cpf:     "390910470-38",
			want:    false,
			message: incorrectResultMessage,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.message, func(t *testing.T) {
			assert.Equal(t, tc.want, Validate(tc.cpf))
		})
	}
}
