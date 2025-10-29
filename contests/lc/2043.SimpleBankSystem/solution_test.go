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

func TestBankEdgeCases(t *testing.T) {
	b := Constructor([]int64{100})

	if ok := b.Deposit(1, 0); !ok {
		t.Fatalf("expected zero deposit to succeed")
	}

	if ok := b.Withdraw(1, 0); !ok {
		t.Fatalf("expected zero withdraw to succeed")
	}

	if ok := b.Transfer(1, 1, 50); !ok {
		t.Fatalf("expected self-transfer to succeed")
	}

	if b.balance[0] != 100 {
		t.Fatalf("expected balance to remain unchanged after self-transfer; got %d", b.balance[0])
	}

	if ok := b.Deposit(2, 10); ok {
		t.Fatalf("expected deposit to fail due to invalid account")
	}
}

func BenchmarkBankOperations(b *testing.B) {
	template := []int64{100, 200, 300, 400, 500}
	for i := 0; i < b.N; i++ {
		bank := Constructor(template)
		if !bank.Transfer(1, 2, 50) {
			b.Fatal("transfer failed")
		}
		if !bank.Deposit(3, 75) {
			b.Fatal("deposit failed")
		}
		if !bank.Withdraw(4, 40) {
			b.Fatal("withdraw failed")
		}
	}
}
