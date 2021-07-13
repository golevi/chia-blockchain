// https://raw.githubusercontent.com/Chia-Network/chia-blockchain/main/chia/types/blockchain_format/slots.py

package blockchain

// ChallengeBlockInfo
//
// 		proof_of_space: ProofOfSpace
// 		challenge_chain_sp_vdf: Optional[VDFInfo]  # Only present if not the first sp
// 		challenge_chain_sp_signature: G2Element
// 		challenge_chain_ip_vdf: VDFInfo
//
type ChallengeBlockInfo struct {
	ProofOfSpace *ProofOfSpace `json:"proof_of_space"`
}
