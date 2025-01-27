package pointerserrors

import (
	"testing"
)

func TestBitCoin(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		got := wallet.Balance()
		expected := Bitcoin(10)

		if got != expected {
			t.Errorf("got %s expected %s", got, expected)
		}
	})

	assertError := func(t testing.TB, err error, expectedErr string) {
		t.Helper()
		if err == nil {
			t.Fatalf("didn't get an error but wanted one")
		}
		if err.Error() != expectedErr {
			t.Errorf("got %q want %q", err.Error(), expectedErr)
		}
	}

	assertBalance := func(t testing.TB, coins Bitcoin, expected Bitcoin) {
		t.Helper()
		if coins != expected {
			t.Errorf("got %d want %d", coins, expected)
		}
	}

	t.Run("Withdraw > Balance", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		err := wallet.Withdraw(20)

		got := wallet.Balance()
		expected := Bitcoin(10)

		assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t, got, expected)
	})
	t.Run("Withdraw < Balance", func(t *testing.T) {
		wallet := Wallet{balance: 30}
		err := wallet.Withdraw(20)

		got := wallet.Balance()
		expected := Bitcoin(10)

		if err != nil {
			t.Errorf("not expecting error: insufficient balance")
		}

		if got != expected {
			t.Errorf("got %s expected %s", got, expected)
		}
		assertBalance(t, got, expected)
	})

}
