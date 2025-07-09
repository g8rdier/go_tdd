package pointers

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	got := wallet.Balance()
	want := 0

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
