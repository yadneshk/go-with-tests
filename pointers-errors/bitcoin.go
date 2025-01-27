package pointerserrors

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("insufficient funds to withdraw")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(coins Bitcoin) {
	w.balance += coins
}

func (w *Wallet) Withdraw(coins Bitcoin) error {
	if w.balance >= coins {
		w.balance -= coins
	} else {
		return ErrInsufficientFunds
	}
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
