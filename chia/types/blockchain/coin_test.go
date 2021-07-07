package blockchain

import (
	"encoding/hex"
	"testing"
)

// https://www.chiaexplorer.com/blockchain/coin/0xf7adb0d7ccbe9477cbc1038ab989bbb3db8668a5226640351e8b8fd1b093105d
func TestCoinStringID(t *testing.T) {
	// This is the CoinID
	want := "f7adb0d7ccbe9477cbc1038ab989bbb3db8668a5226640351e8b8fd1b093105d"

	pci, err := hex.DecodeString("013e041f0a30d1c3906b3f45f05d58b6458ba7e8ef0d3ae74a5a4b84de7ab7b7")
	if err != nil {
		t.Errorf("Error decoding hex string: %s", err)
	}

	ph, err := hex.DecodeString("3481663ac2c20684d05525027e9207e4a9f7fb66a9539257c4f70e57045d3c9b")
	if err != nil {
		t.Errorf("Error decoding hex string: %s", err)
	}

	coin := &Coin{
		ParentCoinInfo: pci,
		PuzzleHash:     ph,
		Amount:         Mojo,
	}

	if coin.StringID() != want {
		t.Errorf("Want %s, got: %s", want, coin.StringID())
	}
}

func TestConvertToMojo(t *testing.T) {
	if ConvertToMojo(1) != 0.000000000001 {
		t.Errorf("Want 0.000000000001, got: %f", ConvertToMojo(1))
	}
}

func TestConvertToXCH(t *testing.T) {
	if ConvertToXCH(0.000000000001) != 1 {
		t.Errorf("Want 1, got: %f", ConvertToXCH(1))
	}
}
