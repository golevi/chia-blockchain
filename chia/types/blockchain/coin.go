// https://raw.githubusercontent.com/Chia-Network/chia-blockchain/main/chia/types/blockchain_format/coin.py

package blockchain

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
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
//
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

// getAmount
//
// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/coin.py#L26
//
//		 # Note that int_to_bytes() will prepend a 0 to integers where the most
//       # significant bit is set, to encode it as a positive number. This
//     	 # despite "amount" being unsigned. This way, a CLVM program can generate
//       # these hashes easily.
//
// https://github.com/Chia-Network/clvm/blob/main/clvm/casts.py#L8
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

	// Could check if amount is one mojo and return 1 byte slice with 1 to speed
	// things up. If I do this, then I probably need more tests.
	if c.Amount == 1 {
		return []byte{1}
	}

	// Create a byte slice to contain the amount.
	//
	// Size returns how many bytes Write would generate to encode the value v,
	// which must be a fixed-size value or a slice of fixed-size values, or a
	// pointer to such data. If v is neither of these, Size returns -1.
	//
	// https://pkg.go.dev/encoding/binary#Size
	r := make([]byte, binary.Size(c.Amount))

	// Chia uses big endian for this.
	//
	// https://github.com/Chia-Network/clvm/blob/main/clvm/casts.py#L12
	// https://en.wikipedia.org/wiki/Endianness
	// https://pkg.go.dev/encoding/binary
	binary.BigEndian.PutUint64(r, c.Amount)

	// These are the comments from the python implementation in the CLVM repo. I
	// don't fully understand the conditions in that while loop. Go might not be
	// able to create a loop like that.
	//
	// https://github.com/Chia-Network/clvm/blob/main/clvm/casts.py
	// make sure the string returned is minimal
	// ie. no leading 00 or ff bytes that are unnecessary
	// while len(r) > 1 and r[0] == (0xFF if r[1] & 0x80 else 0):
	for len(r) > 1 && check(r) {
		r = r[1:]
	}

	return r
}

// @TODO Evaluate if we really need this. I only created it because the original
// code confused me and I thought it would be easier to figure out with a its
// own function. It isn't hurting anything, so I'll leave it for now.
func check(r []byte) bool {
	if r[0] == 0xff {
		return true
	}

	if r[0] == 0x0 {
		return true
	}

	return false
}

// ConvertToMojo converts XCH to Mojo
func ConvertToMojo(amount int) float64 {
	return float64(amount) / Teramojo
}

// ConvertToXCH converts Mojo to XCH
func ConvertToXCH(amount float64) float64 {
	return amount * float64(Teramojo)
}
