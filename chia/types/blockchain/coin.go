// https://raw.githubusercontent.com/Chia-Network/chia-blockchain/main/chia/types/blockchain_format/coin.py

package blockchain

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

// 1000000000000 mojo = 1.               XCH =  Chia      = Teramojo = One Trillion Mojos
// 1000000000    mojo = 0.001            XCH =  MilliChia = Gigamojo
// 1000000       mojo = 0.000001         XCH =  MicroChia = Megamojo
// 1000          mojo = 0.000000001      XCH =  NanoChia  = Kilomojo
// 1             mojo = 0.000000000001   XCH =  PicoChia  = Mojo
const (
	Teramojo = 1_000_000_000_000
	Gigamojo = 1_000_000_000
	Megamojo = 1_000_000
	Kilomojo = 1_000
	Mojo     = 1

	Chia      = 1
	MilliChia = 0.001
	MicroChia = 0.000001
	NanoChia  = 0.000000001
	PicoChia  = 0.000000000001
)

// @TODO
// Coin
//
// This structure is used in the body for the reward and fees genesis coins.
//
// coin.py
//
// 		parent_coin_info: bytes32  # down with this sort of thing.
// 		puzzle_hash: bytes32
// 		amount: uint64
//
// 		def get_hash(self) -> bytes32:
// 			# This does not use streamable format for hashing, the amount is
// 			# serialized using CLVM integer format.

// 			# Note that int_to_bytes() will prepend a 0 to integers where the most
// 			# significant bit is set, to encode it as a positive number. This
// 			# despite "amount" being unsigned. This way, a CLVM program can generate
// 			# these hashes easily.
// 			return std_hash(self.parent_coin_info + self.puzzle_hash + int_to_bytes(self.amount))
//
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/coin.py
//
// hash.py
//
// 		def std_hash(b) -> bytes32:
//     		"""
//     		The standard hash used in many places.
//     		"""
//     		return bytes32(blspy.Util.hash256(bytes(b)))
//
// https://github.com/Chia-Network/chia-blockchain/blob/17b061d1d7b76ac01ca8c673331122536354f536/chia/util/hash.py
//
// 		std_hash(self.parent_coin_info + self.puzzle_hash + int_to_bytes(self.amount))
//
type Coin struct {
	ParentCoinInfo []byte `json:"parent_coin_info"`
	PuzzleHash     []byte `json:"puzzle_hash"`
	Amount         uint64 `json:"amount"`
}

func (c Coin) ID() []byte {
	// Convert c.Amount to a []byte
	// size := (bits.Len64(c.Amount) + 8) >> 3
	amount := make([]byte, 8)
	binary.BigEndian.PutUint64(amount, c.Amount)
	fmt.Println(amount)

	hasher := sha256.New()
	hasher.Write(c.ParentCoinInfo)
	hasher.Write(c.PuzzleHash)
	hasher.Write(amount)

	return hasher.Sum(nil)
}

func (c Coin) StringID() string {
	return hex.EncodeToString(c.ID())
}

// ConvertToMojo converts XCH to Mojo
func ConvertToMojo(amount int) float64 {
	return float64(amount) / Teramojo
}

// ConvertToXCH converts Mojo to XCH
func ConvertToXCH(amount float64) float64 {
	return amount * float64(Teramojo)
}
