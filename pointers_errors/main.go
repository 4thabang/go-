package main

import (
	"errors"
	"fmt"
)

var (
	ErrCannotWithdraw = errors.New("cannot withdraw more than you have")
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(money Bitcoin) {
	w.balance += money
}

func (w *Wallet) Withdraw(money Bitcoin) error {
	if money > w.balance {
		return ErrCannotWithdraw
	}
	w.balance -= money
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
