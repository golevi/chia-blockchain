// https://raw.githubusercontent.com/Chia-Network/chia-blockchain/main/chia/types/blockchain_format/coin.py

package blockchain

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"math/bits"
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
	hasher := sha256.New()
	hasher.Write(c.ParentCoinInfo)
	hasher.Write(c.PuzzleHash)
	hasher.Write(c.getAmount())

	return hasher.Sum(nil)
}

func (c Coin) StringID() string {
	return hex.EncodeToString(c.ID())
}

// amount
//
//		 def int_to_bytes(v):
//		     byte_count = (v.bit_length() + 8) >> 3
//		     if v == 0:
//		         return b""
//		     r = v.to_bytes(byte_count, "big", signed=True)
//		     # make sure the string returned is minimal
//		     # ie. no leading 00 or ff bytes that are unnecessary
//		     while len(r) > 1 and r[0] == (0xFF if r[1] & 0x80 else 0):
//		         r = r[1:]
//		     return r
//
func (c Coin) getAmount() []byte {
	if c.Amount == 0 {
		return []byte{}
	}
	byteCount := bits.Len64(c.Amount) * 8
	r := make([]byte, byteCount)

	binary.BigEndian.PutUint64(r, c.Amount)

	for len(r) > 1 {
		r = r[1:]
	}

	return r
}

// ConvertToMojo converts XCH to Mojo
func ConvertToMojo(amount int) float64 {
	return float64(amount) / Teramojo
}

// ConvertToXCH converts Mojo to XCH
func ConvertToXCH(amount float64) float64 {
	return amount * float64(Teramojo)
}
