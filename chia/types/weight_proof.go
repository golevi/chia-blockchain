// https://github.com/Chia-Network/chia-blockchain/blob/1.2.1/chia/types/weight_proof.py

package types

import "github.com/golevi/chia-blockchain/chia/types/blockchain"

// class SubEpochData(Streamable):
// reward_chain_hash: bytes32
// num_blocks_overflow: uint8
// new_sub_slot_iters: Optional[uint64]
// new_difficulty: Optional[uint64]
type SubEpochData struct {
	RewardChainHash   []byte `json:"reward_chain_hash"`
	NumBlocksOverflow uint8  `json:"num_blocks_overflow"`
	NewSubSlotIters   uint64 `json:"new_sub_slot_iters"`
	NewDifficulty     uint64 `json:"new_difficulty"`
}

// # if infused
// proof_of_space: Optional[ProofOfSpace]
// # VDF to signage point
// cc_signage_point: Optional[VDFProof]
// # VDF from signage to infusion point
// cc_infusion_point: Optional[VDFProof]
// icc_infusion_point: Optional[VDFProof]
// cc_sp_vdf_info: Optional[VDFInfo]
// signage_point_index: Optional[uint8]
// # VDF from beginning to end of slot if not infused
// #  from ip to end if infused
// cc_slot_end: Optional[VDFProof]
// icc_slot_end: Optional[VDFProof]
// # info from finished slots
// cc_slot_end_info: Optional[VDFInfo]
// icc_slot_end_info: Optional[VDFInfo]
// cc_ip_vdf_info: Optional[VDFInfo]
// icc_ip_vdf_info: Optional[VDFInfo]
// total_iters: Optional[uint128]
type SubSlotData struct {
	ProofOfSpace blockchain.ProofOfSpace `json:"proof_of_space"`
}
