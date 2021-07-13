package types

import (
	"testing"

	"github.com/golevi/chia-blockchain/chia/types/blockchain"
)

func TestCoinRecordFields(t *testing.T) {
	hash := "aa36cee437c06c9d3985893fb9e82a1ccd4e805802468df5d1246c84b720fd79"
	ts := uint64(1626117708)

	cr := CoinRecord{
		Coin: blockchain.Coin{
			ParentCoinInfo: []byte("1"),
			PuzzleHash:     []byte("2"),
			Amount:         100,
		},
		ConfirmedBlockIndex: 1000,
		SpentBlockIndex:     10001,
		Spent:               false,
		Coinbase:            false,
		Timestamp:           ts,
	}

	if cr.Name() != hash {
		t.Errorf("Expected: %s, got: %s", string(hash), cr.Name())
	}

	if cr.Timestamp != ts {
		t.Errorf("Expected timestamp: %d, got: %d", ts, cr.Timestamp)
	}
}
