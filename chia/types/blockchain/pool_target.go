// https://raw.githubusercontent.com/Chia-Network/chia-blockchain/46395374ae9d1394925d79a9ed9c025d20f5d216/chia/types/blockchain_format/pool_target.py

package blockchain

// PoolTarget
//
// 		puzzle_hash: bytes32
// 		max_height: uint32  # A max height of 0 means it is valid forever
//
type PoolTarget struct {
	PuzzleHash []byte `json:"puzzle_hash"`
	MaxHeight  uint32 `json:"max_height"`
}
