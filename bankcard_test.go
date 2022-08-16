package bankcard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBankNames(t *testing.T) {
	resp := BankNames()
	t.Log(resp)
}

func TestBankInfo(t *testing.T) {
	var cards = []string{
		"6214850106608721", "6214850106608721",
	}
	for _, card := range cards {
		err, g := BankInfo(card)
		assert.Nil(t, err)
		t.Log(g)
	}
}
