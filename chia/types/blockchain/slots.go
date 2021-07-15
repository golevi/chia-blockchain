// https://github.com/Chia-Network/chia-blockchain/blob/1.2.1/chia/types/blockchain_format/slots.py

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

	// ChallengeChainSPVDF is only present if it isn't the first signage point.
	ChallengeChainSPVDF VDFInfo `json:"challenge_chain_sp_vdf"`
	// ChallengeChainSPSignature []byte  `json:"challenge_chain_sp_signature"`
	ChallengeChainIPVDF VDFInfo `json:"challenge_chain_ip_vdf"`
}
