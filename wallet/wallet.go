package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (b *Wallet) Deposit(s Bitcoin) {
	b.balance = b.balance + s
}

func (b Wallet) Balance() Bitcoin {
	return b.balance
}
func (b *Wallet) Withdraw(s Bitcoin) error {
	if b.balance < s {
		return ErrNoFunds
	}
	b.balance = b.balance - s
	return nil
}

var ErrNoFunds = errors.New("no money")
