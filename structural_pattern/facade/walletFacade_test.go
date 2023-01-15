package facade

import (
	"testing"
)

func TestWalletFacade(t *testing.T) {
	t.Log()
	walletFacade := newWalletFacade("abc", 1234)
	t.Log()

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}

	t.Log()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}
}
