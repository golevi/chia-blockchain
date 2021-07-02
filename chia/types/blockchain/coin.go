// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/blockchain_format/coin.py

package blockchain

// @TODO
// Coin
//
// This structure is used in the body for the reward and fees genesis coins.
//
// Fields
//
// 		parent_coin_info: bytes32  # down with this sort of thing.
// 		puzzle_hash: bytes32
// 		amount: uint64
//
type Coin struct {
	ParentCoinInfo [32]byte `json:"parent_coin_info"`
	PuzzleHash     [32]byte `json:"puzzle_hash"`

	Amount uint64 `json:"amount"`
}
