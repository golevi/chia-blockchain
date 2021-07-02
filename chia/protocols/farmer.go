// https://github.com/Chia-Network/chia-blockchain/blob/main/chia/protocols/farmer_protocol.py

package protocols

// NewSignagePoint
//
// 		challenge_hash: bytes32
// 		challenge_chain_sp: bytes32
// 		reward_chain_sp: bytes32
// 		difficulty: uint64
// 		sub_slot_iters: uint64
// 		signage_point_index: uint8
//
type NewSignagePoint struct {
	ChallengeHash    [32]byte `json:"challenge_hash"`
	ChallengeChainSP [32]byte `json:"challenge_chain_sp"`
	RewardChainSP    [32]byte `json:"reward_chain_sp"`

	Difficulty        uint64 `json:"difficulty"`
	SubSlotIters      uint64 `json:"sub_slot_iters"`
	SignagePointIndex uint8  `json:"signage_point_index"`
}

// @TODO
// DeclareProofOfSpace
//
// 		challenge_hash: bytes32
// 		challenge_chain_sp: bytes32
// 		signage_point_index: uint8
// 		reward_chain_sp: bytes32
// 		proof_of_space: ProofOfSpace
// 		challenge_chain_sp_signature: G2Element
// 		reward_chain_sp_signature: G2Element
// 		farmer_puzzle_hash: bytes32
// 		pool_target: Optional[PoolTarget]
// 		pool_signature: Optional[G2Element]
//
type DeclareProofOfSpace struct {
}
