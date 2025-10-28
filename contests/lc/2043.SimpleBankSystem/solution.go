package _043_SimpleBankSystem

type Bank struct {
	balance []int64
	n       int
}

func NewBank(balance []int64) Bank {
	return Bank{balance, len(balance)}
}

func (b *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 > b.n || account2 > b.n || b.balance[account1-1] < money {
		return false
	}
	b.balance[account1-1] -= money
	b.balance[account2-1] += money
	return true
}

func (b *Bank) Deposit(account int, money int64) bool {
	if account > b.n {
		return false
	}
	b.balance[account-1] += money
	return true
}

func (b *Bank) Withdraw(account int, money int64) bool {
	if account > b.n || b.balance[account-1] < money {
		return false
	}
	b.balance[account-1] -= money
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
