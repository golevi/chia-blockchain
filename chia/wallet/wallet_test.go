package wallet

import "testing"

func TestWalletFields(t *testing.T) {
	wallet := Wallet{
		ID: 1,
	}

	if wallet.ID != 1 {
		t.Errorf("Wallet ID = %d, want 1", wallet.ID)
	}
}
