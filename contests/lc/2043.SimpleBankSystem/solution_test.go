package _043_SimpleBankSystem

import (
	"reflect"
	"testing"
)

func TestBankTransferSuccess(t *testing.T) {
	bank := NewBank([]int64{10, 100})

	if ok := bank.Transfer(1, 2, 5); !ok {
		t.Fatalf("Transfer returned false, want true")
	}

	want := []int64{5, 105}
	if !reflect.DeepEqual(bank.balance, want) {
		t.Fatalf("balance = %v, want %v", bank.balance, want)
	}
}

func TestBankTransferFailure(t *testing.T) {
	tests := []struct {
		name     string
		balance  []int64
		account1 int
		account2 int
		money    int64
	}{
		{
			name:     "insufficient funds",
			balance:  []int64{5, 20},
			account1: 1,
			account2: 2,
			money:    10,
		},
		{
			name:     "target account out of range",
			balance:  []int64{5, 20},
			account1: 1,
			account2: 3,
			money:    1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			bank := NewBank(append([]int64(nil), tt.balance...))

			if ok := bank.Transfer(tt.account1, tt.account2, tt.money); ok {
				t.Fatalf("Transfer returned true, want false")
			}

			if !reflect.DeepEqual(bank.balance, tt.balance) {
				t.Fatalf("balance mutated to %v, want %v", bank.balance, tt.balance)
			}
		})
	}
}

func TestBankDeposit(t *testing.T) {
	bank := NewBank([]int64{10})

	if ok := bank.Deposit(1, 5); !ok {
		t.Fatalf("Deposit returned false, want true")
	}

	want := []int64{15}
	if !reflect.DeepEqual(bank.balance, want) {
		t.Fatalf("balance = %v, want %v", bank.balance, want)
	}

	if ok := bank.Deposit(2, 5); ok {
		t.Fatalf("Deposit with invalid account returned true, want false")
	}

	if !reflect.DeepEqual(bank.balance, want) {
		t.Fatalf("balance mutated to %v, want %v", bank.balance, want)
	}
}

func TestBankWithdraw(t *testing.T) {
	bank := NewBank([]int64{10})

	if ok := bank.Withdraw(1, 7); !ok {
		t.Fatalf("Withdraw returned false, want true")
	}

	want := []int64{3}
	if !reflect.DeepEqual(bank.balance, want) {
		t.Fatalf("balance = %v, want %v", bank.balance, want)
	}

	if ok := bank.Withdraw(1, 4); ok {
		t.Fatalf("Withdraw with insufficient funds returned true, want false")
	}

	if ok := bank.Withdraw(2, 1); ok {
		t.Fatalf("Withdraw with invalid account returned true, want false")
	}

	if !reflect.DeepEqual(bank.balance, want) {
		t.Fatalf("balance mutated to %v, want %v", bank.balance, want)
	}
}
