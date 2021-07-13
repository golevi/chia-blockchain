// https://github.com/Chia-Network/chia-blockchain/blob/1.2.1/chia/consensus/pos_quality.py

package consensus

import "math"

// ExpectedPlotSize computes the expected size of the plot in bytes.
//
// Given the plot size parameter k (which is between 32 and 59), computes the
// expected size of the plot in bytes (times a constant factor). This is based
// on efficient encoding of the plot, and aims to be scale agnostic, so larger
// plots don't necessarily get more rewards per byte. The +1 is added to give
// half a bit more space per entry, which is necessary to store the entries in
// the plot.
//
//		return ((2 * k) + 1) * (2 ** (k - 1))
//
func ExpectedPlotSize(k int) uint64 {
	var a uint64 = uint64((2 * k) + 1)

	var ka float64 = float64(k - 1)
	var b uint64 = uint64(math.Pow(2, ka))

	return a * b
}
