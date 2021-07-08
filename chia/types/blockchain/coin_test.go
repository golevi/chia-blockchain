package blockchain

import (
	"encoding/hex"
	"testing"
)

// https://www.chiaexplorer.com/blockchain/coin/0xf7adb0d7ccbe9477cbc1038ab989bbb3db8668a5226640351e8b8fd1b093105d
func TestCoinStringID(t *testing.T) {

	coins := &[]struct {
		Want   string
		Parent string
		Puzzle string
		Amount uint64
	}{
		{
			Want:   "f7adb0d7ccbe9477cbc1038ab989bbb3db8668a5226640351e8b8fd1b093105d",
			Parent: "013e041f0a30d1c3906b3f45f05d58b6458ba7e8ef0d3ae74a5a4b84de7ab7b7",
			Puzzle: "3481663ac2c20684d05525027e9207e4a9f7fb66a9539257c4f70e57045d3c9b",
			Amount: Mojo,
		},
		{
			Want:   "21506f38a385cc6640e54270f3ad16bec49f2538d31da9bc369c1f299245b287",
			Parent: "ccd5bb71183532bff220ba46c268991a00000000000000000000000000083d97",
			Puzzle: "5edd2a1ab38b2ad31b326d147adbc1693183e82fad9335346cafb1a1417c420c",
			Amount: 1.75 * Teramojo,
		},
		{
			Want:   "4b878f285b49e3198dcf78f7627bf720f9dc19bc59dc88ff5e3a48718fce2868",
			Parent: "3ff07eb358e8255a65c30a2dce0e5fbb00000000000000000000000000083e9d",
			Puzzle: "907491ca39c35bc1f9a6eda33f7c0f97a9f583975088dad7216f1edd79f522ae",
			Amount: 0.25 * Teramojo,
		},
		{
			Want:   "8efcf70686577ac6dc0df7717b87833190f3d6cf17ad435e24adbb3ee3b5ae47",
			Parent: "55ca386a9a232d455e6dda65e499f07a775710696b5c07a486011c4a28813fa0",
			Puzzle: "98f7a41f496927e3ead26888a68991cd81338543a5e5d1a8ad87b1d7b8ce9337",
			Amount: 6 * Teramojo,
		},
	}

	for _, c := range *coins {
		want := c.Want

		pci, err := hex.DecodeString(c.Parent)
		if err != nil {
			t.Errorf("Error decoding hex string: %s", err)
		}

		ph, err := hex.DecodeString(c.Puzzle)
		if err != nil {
			t.Errorf("Error decoding hex string: %s", err)
		}

		coin := &Coin{
			ParentCoinInfo: pci,
			PuzzleHash:     ph,
			Amount:         c.Amount,
		}

		if coin.StringID() != want {
			t.Errorf("Want %s, got: %s", want, coin.StringID())
		}
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
