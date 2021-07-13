// https://raw.githubusercontent.com/Chia-Network/chia-blockchain/main/chia/types/blockchain_format/reward_chain_block.py

package blockchain

import "math/big"

// RewardChainBlockUnfinished
//
// 		total_iters: uint128
// 		signage_point_index: uint8
// 		pos_ss_cc_challenge_hash: bytes32
// 		proof_of_space: ProofOfSpace
// 		challenge_chain_sp_vdf: Optional[VDFInfo]  # Not present for first sp in slot
// 		challenge_chain_sp_signature: G2Element
// 		reward_chain_sp_vdf: Optional[VDFInfo]  # Not present for first sp in slot
// 		reward_chain_sp_signature: G2Element
//
type RewardChainBlockUnfinished struct {
	TotalIters           *big.Int `json:"total_iters"`
	SignagePointIndex    uint8    `json:"signage_point_index"`
	POSSSCCChallengeHash []byte   `json:"pos_ss_cc_challenge_hash"`

	ProofOfSpace *ProofOfSpace `json:"proof_of_space"`
}

// RewardChainBlock
//
// 		weight: uint128
// 		height: uint32
// 		total_iters: uint128
// 		signage_point_index: uint8
// 		pos_ss_cc_challenge_hash: bytes32
// 		proof_of_space: ProofOfSpace
// 		challenge_chain_sp_vdf: Optional[VDFInfo]  # Not present for first sp in slot
// 		challenge_chain_sp_signature: G2Element
// 		challenge_chain_ip_vdf: VDFInfo
// 		reward_chain_sp_vdf: Optional[VDFInfo]  # Not present for first sp in slot
// 		reward_chain_sp_signature: G2Element
// 		reward_chain_ip_vdf: VDFInfo
// 		infused_challenge_chain_ip_vdf: Optional[VDFInfo]  # Iff deficit < 16
// 		is_transaction_block: bool
//
type RewardChainBlock struct {
	Weight               *big.Int      `json:"weight"`
	Height               uint32        `json:"height"`
	TotalIters           *big.Int      `json:"total_iters"`
	SignagePointIndex    uint8         `json:"signage_point_index"`
	POSSSCCChallengeHash []byte        `json:"pos_ss_cc_challenge_hash"`
	ProofOfSpace         *ProofOfSpace `json:"proof_of_space"`
	// ChallengeChainSP     *ChallengeChainSP   `json:"challenge_chain_sp"`
	// ChallengeChainIP     *ChallengeChainIP   `json:"challenge_chain_ip"`
	// RewardChainSP        *RewardChainSP      `json:"reward_chain_sp"`
	// RewardChainIP        *RewardChainIP      `json:"reward_chain_ip"`
	// InfusedChallengeIP   *InfusedChallengeIP `json:"infused_challenge_chain_ip"`
	IsTransactionBlock bool `json:"is_transaction_block"`
}
