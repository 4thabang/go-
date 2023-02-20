package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		AssertBallance(t, wallet, got)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		AssertBallance(t, wallet, got)
		AssertNotError(t, err)
	})

	// we want this to return an error, this constitutes a successful test.
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(100))

		got := wallet.Balance()
		AssertBallance(t, wallet, got)
		AssertError(t, err, ErrCannotWithdraw)
	})
}

func AssertBallance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("expected: %d, got: %d", want, got)
	}
}

func AssertNotError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func AssertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}
