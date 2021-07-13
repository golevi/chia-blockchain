// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/types/coin_record.py

package types

import "github.com/golevi/chia-blockchain/chia/types/blockchain"

// CoinRecord are values that correspond to a CoinName that are used
// in keeping track of the unspent database.
//
// 		coin: Coin
// 		confirmed_block_index: uint32
// 		spent_block_index: uint32
// 		spent: bool
// 		coinbase: bool
// 		timestamp: uint64  # Timestamp of the block at height confirmed_block_index
//
type CoinRecord struct {
	Coin                blockchain.Coin `json:"coin"`
	ConfirmedBlockIndex uint32          `json:"confirmed_block_index"`
	SpentBlockIndex     uint32          `json:"spent_block_index"`
	Spent               bool            `json:"spent"`
	Coinbase            bool            `json:"coinbase"`
	Timestamp           uint64          `json:"timestamp"`
}

func (cr CoinRecord) Name() string {
	return cr.Coin.StringID()
}
