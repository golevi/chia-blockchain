package consensus

import (
	"encoding/hex"
	"math/big"
	"testing"
)

// TestCalculateIterationsQuality makes sure we're getting the right number from
// the CalculateIterationsQuality function.
//
// https://github.com/Chia-Network/chia-blockchain/blob/1.2.1/tests/core/consensus/test_pot_iterations.py#L79
func TestCalculateIterationsQuality(t *testing.T) {
	size := 32
	difficulty := 2400

	difficultyConstantFactor := big.NewInt(1)
	difficultyConstantFactor.SetString("147573952589676412928", 10)

	qualityString, err := hex.DecodeString("6bc7e96778d56f99640d7d606682543a1449fe84c53d9b3bd764decacc29a10c")
	if err != nil {
		t.Error(err)
	}

	CCSPOutputHash, err := hex.DecodeString("396fef5662016a16c8849b58bcbd6362368792f637b2a7a2abd91db2f35b9a80")
	if err != nil {
		t.Error(err)
	}

	result := CalculateIterationsQuality(
		*difficultyConstantFactor,
		qualityString,
		size,
		uint64(difficulty),
		CCSPOutputHash,
	)

	expected := uint64(1110877509098)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	// Test 2

	qualityString, err = hex.DecodeString("53822f1368aa717084a8cb99edf143217ade0c7ae5dc2ae569f296897a947c60")
	if err != nil {
		t.Error(err)
	}

	CCSPOutputHash, err = hex.DecodeString("d2a53fe65883fdc929ceeae12583fbd1d9cdabfaae9ecb990761cd840834a885")
	if err != nil {
		t.Error(err)
	}

	result = CalculateIterationsQuality(
		*difficultyConstantFactor,
		qualityString,
		size,
		uint64(difficulty),
		CCSPOutputHash,
	)

	expected = uint64(1950187686686)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
