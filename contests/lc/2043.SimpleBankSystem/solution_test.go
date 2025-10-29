package _043_SimpleBankSystem

import "testing"

func TestBankOperations(t *testing.T) {
	b := Constructor([]int64{10, 100, 20, 50, 30})

	if ok := b.Withdraw(3, 10); !ok {
		t.Fatalf("expected withdraw to succeed")
	}

	if ok := b.Transfer(5, 1, 20); !ok {
		t.Fatalf("expected transfer to succeed")
	}

	if ok := b.Deposit(5, 20); !ok {
		t.Fatalf("expected deposit to succeed")
	}

	if ok := b.Transfer(3, 4, 15); ok {
		t.Fatalf("expected transfer to fail due to insufficient funds")
	}

	if ok := b.Withdraw(10, 50); ok {
		t.Fatalf("expected withdraw to fail due to invalid account")
	}

	expected := []int64{30, 100, 10, 50, 30}
	for i, want := range expected {
		if got := b.balance[i]; got != want {
			t.Fatalf("account %d balance = %d, want %d", i+1, got, want)
		}
	}
}

func TestConstructorCopiesInput(t *testing.T) {
	input := []int64{1, 2, 3}
	b := Constructor(input)
	input[0] = 999

	if b.balance[0] != 1 {
		t.Fatalf("expected constructor to copy input slice; got %d", b.balance[0])
	}
}
