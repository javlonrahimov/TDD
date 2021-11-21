package main

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(bitcoin Bitcoin) {
	w.balance += bitcoin
}

func (w *Wallet) Withdraw(bitcoin Bitcoin) error {

	if bitcoin > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= bitcoin
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
