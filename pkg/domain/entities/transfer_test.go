package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransfer(t *testing.T) {
	origID := "7d24e8bf-d17e-4fc6-8b94-a3eb93b2b66b"
	destID := "34988bed-af65-4add-a308-0001656cd156"

	t.Run("Should create a tranfer successfully", func(t *testing.T) {
		amount := 200
		trans, err := NewTransfer(origID, destID, amount)

		assert.Nil(t, err)
		if err != nil {
			assert.Equal(t, trans.AccountOriginID, origID)
			assert.Equal(t, trans.AccountDestinationID, destID)
			assert.Equal(t, trans.Amount, amount)
			assert.NotNil(t, trans.ID)
		}
	})

	t.Run("Should not create a tranfer due to a 0 value amount", func(t *testing.T) {
		amount := 0
		trans, err := NewTransfer(origID, destID, amount)

		assert.Equal(t, err, ErrInvalidAmount)
		assert.Nil(t, trans)
	})

	t.Run("Should not create a tranfer because the origin account is equal de destination account", func(t *testing.T) {
		amount := 0
		trans, err := NewTransfer(origID, destID, amount)

		assert.Equal(t, err, ErrInvalidAmount)
		assert.Nil(t, trans)
	})

	t.Run("Should not create a tranfer due to a negative value amount", func(t *testing.T) {
		amount := 100
		destID = origID
		trans, err := NewTransfer(origID, destID, amount)

		assert.Equal(t, err, ErrOrigAccEqualDestAcc)
		assert.Nil(t, trans)
	})
}
