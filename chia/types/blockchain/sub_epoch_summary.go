// https://github.com/Chia-Network/chia-blockchain/blob/1.2.1/chia/types/blockchain_format/sub_epoch_summary.py

package blockchain

// SubEpochSummary
//
// 		prev_subepoch_summary_hash: bytes32
// 		reward_chain_hash: bytes32  # hash of reward chain at end of last segment
// 		num_blocks_overflow: uint8  # How many more blocks than 384*(N-1)
// 		new_difficulty: Optional[uint64]  # Only once per epoch (diff adjustment)
// 		new_sub_slot_iters: Optional[uint64]  # Only once per epoch (diff adjustment)
//
type SubEpochSummary struct {
	PrevSubepochSummaryHash []byte `json:"prev_subepoch_summary_hash"`

	// RewardChainHash of reward chain at end of last segment
	RewardChainHash []byte `json:"reward_chain_hash"`

	// NumBlocksOverflow represents how many more blocks than 384*(N-1)
	NumBlocksOverflow uint8 `json:"num_blocks_overflow"`

	// NewDifficulty only once per epoch (diff adjustment)
	NewDifficulty uint64 `json:"new_difficulty"`

	// NewSubSlotIters only once per epoch (diff adjustment)
	NewSubSlotIters uint64 `json:"new_sub_slot_iters"`
}
