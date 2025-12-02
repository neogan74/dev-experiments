package _043_SimpleBankSystem

// Bank simulates a set of bank accounts with basic operations.
type Bank struct {
	balance []int64
}

// Constructor creates a Bank with a defensive copy of the initial balances.
func Constructor(balance []int64) Bank {
	balances := make([]int64, len(balance))
	copy(balances, balance)
	return Bank{balance: balances}
}

// Transfer moves funds from account1 to account2 when both accounts exist and have sufficient funds.
func (b *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !b.validAccount(account1) || !b.validAccount(account2) {
		return false
	}
	idx1 := account1 - 1
	idx2 := account2 - 1
	if b.balance[idx1] < money {
		return false
	}
	b.balance[idx1] -= money
	b.balance[idx2] += money
	return true
}

// Deposit adds funds to the given account when it exists.
func (b *Bank) Deposit(account int, money int64) bool {
	if !b.validAccount(account) {
		return false
	}
	b.balance[account-1] += money
	return true
}

// Withdraw removes funds from the given account if it exists and has enough balance.
func (b *Bank) Withdraw(account int, money int64) bool {
	if !b.validAccount(account) {
		return false
	}
	idx := account - 1
	if b.balance[idx] < money {
		return false
	}
	b.balance[idx] -= money
	return true
}

func (b *Bank) validAccount(account int) bool {
	return account >= 1 && account <= len(b.balance)
}
