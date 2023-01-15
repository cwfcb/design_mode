package facade

import (
	"errors"
	"fmt"
)

// 复杂子系统的组成部分
type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

func (w *Wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return errors.New("Balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.balance = w.balance - amount
	return nil
}
