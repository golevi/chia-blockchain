// https://raw.githubusercontent.com/Chia-Network/chia-blockchain/main/chia/types/blockchain_format/sub_epoch_summary.py

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
}
