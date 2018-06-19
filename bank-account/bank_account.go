// Package account simulates a bank account
package account

import "sync"

var mutex = &sync.Mutex{}

// Account is the type for a bank account
type Account struct {
	closed  bool
	balance int64
}

// Open opens a new bank account with an initial deposit
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	a := &Account{balance: initialDeposit}

	return a
}

// Close closes an account and rturns the final payout
func (a *Account) Close() (payout int64, ok bool) {
	mutex.Lock()
	defer mutex.Unlock()
	if a.closed {
		return 0, false
	}

	payout = a.balance
	ok = true

	a.balance = 0
	a.closed = true

	return
}

// Balance returns the balance of the account unless the account is closed
func (a *Account) Balance() (balance int64, ok bool) {
	if a.closed {
		return 0, false
	}

	return a.balance, true
}

// Deposit deposits or withdraws money into/from an account
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	if a.closed {
		return 0, false
	}

	mutex.Lock()
	defer mutex.Unlock()
	if amount < 0 && a.balance+amount < 0 {
		return 0, false
	}

	a.balance += amount

	return a.balance, true
}
