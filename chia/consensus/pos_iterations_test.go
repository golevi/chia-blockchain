package consensus

import (
	"math/big"
	"testing"

	"github.com/golevi/chia-blockchain/chia/util"
)

// TestCalculateIterationsQuality makes sure we're getting the right number from
// the CalculateIterationsQuality function.
//
// https://github.com/Chia-Network/chia-blockchain/blob/1.2.1/tests/core/consensus/test_pot_iterations.py#L79
func TestCalculateIterationsQuality(t *testing.T) {
	var qualityBytes []byte
	slotIndex := big.NewInt(0)
	k := big.NewInt(32)
	farmerIndex := big.NewInt(100)
	qualityBytes = append(qualityBytes, slotIndex.Bytes()...)
	qualityBytes = append(qualityBytes, k.Bytes()...)
	qualityBytes = append(qualityBytes, farmerIndex.Bytes()...)

	spIndex := big.NewInt(1)
	var spHashInput []byte
	spHashInput = append(spHashInput, slotIndex.Bytes()...)
	spHashInput = append(spHashInput, spIndex.Bytes()...)

	// sp_hash = std_hash(slot_index.to_bytes(4, "big") + sp_index.to_bytes(4, "big"))
	spHash := util.StandardHash(spHashInput)

	// quality = std_hash(slot_index.to_bytes(4, "big") + k.to_bytes(1, "big") + bytes(farmer_index))
	quality := util.StandardHash(qualityBytes)

	// sp_interval_iters = uint64(100000000 / 32)
	difficulty := uint64(500000000000)

	difficultyConstantFactor := big.NewInt(2)
	difficultyConstantFactor.Exp(difficultyConstantFactor, big.NewInt(25), nil)

	result := CalculateIterationsQuality(
		*difficultyConstantFactor,
		quality,
		int(k.Int64()),
		difficulty,
		spHash,
	)

	expected := uint64(116799048)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
