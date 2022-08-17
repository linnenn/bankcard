package backcard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBankNameMapList(t *testing.T) {
	resp := BankNameMapList()
	t.Log(resp)
}

func TestBankNameList(t *testing.T) {
	resp := BankNameList()
	t.Log(resp)
}

func TestBankInfo(t *testing.T) {
	var cards = []string{
		"6214850106608721", "6214850106608722",
	}
	for _, card := range cards {
		err, g := BankCardInfo(card)
		assert.Nil(t, err)
		t.Log(g)
	}
}
