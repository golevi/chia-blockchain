// https://github.com/Chia-Network/chia-blockchain/blob/1.2.1/chia/consensus/pot_iterations.py

package consensus

import (
	"math/big"

	"github.com/golevi/chia-blockchain/chia/util"
)

// CalculateIterationsQuality calculates the number of iterations from the
// quality. This is derived as the difficulty times the constant factor times a
// random number between 0 and 1 (based on quality string), divided by plot size.
//
// https://github.com/Chia-Network/chia-blockchain/blob/1.2.1/chia/consensus/pot_iterations.py#L46
//
// 		difficulty_constant_factor: uint128,
// 		quality_string: bytes32,
// 		size: int,
// 		difficulty: uint64,
// 		cc_sp_output_hash: bytes32,
//
func CalculateIterationsQuality(
	DifficultyConstantFactor big.Int,
	QualityString []byte,
	Size int,
	Difficulty uint64,
	CCSPOutputHash []byte,
) uint64 {
	// sp_quality_string: bytes32 = std_hash(quality_string + cc_sp_output_hash)

	var input []byte
	input = append(input, QualityString...)
	input = append(input, CCSPOutputHash...)
	SPQualityBytes := util.StandardHash(input)

	// iters = uint64(
	//     int(difficulty)
	//     * int(difficulty_constant_factor)
	//     * int.from_bytes(sp_quality_string, "big", signed=False)
	//     // (int(pow(2, 256)) * int(_expected_plot_size(size)))
	// )
	BigIterations := big.NewInt(1)

	BigDifficulty := big.NewInt(1)
	BigDifficulty.SetUint64(Difficulty)

	BigSPQuality := big.NewInt(1)
	BigSPQuality.SetBytes(SPQualityBytes)

	BigIterations.Mul(BigDifficulty, &DifficultyConstantFactor)
	BigIterations.Mul(BigIterations, BigSPQuality)

	BigPow := big.NewInt(2)
	BigPow.Exp(BigPow, big.NewInt(256), big.NewInt(0))
	BigPow.Mul(BigPow, big.NewInt(int64(ExpectedPlotSize(Size))))

	BigIterations.Div(BigIterations, BigPow)

	return BigIterations.Uint64()
}
