// https://raw.githubusercontent.com/Chia-Network/chia-blockchain/main/chia/types/blockchain_format/proof_of_space.py

package blockchain

import "github.com/celo-org/celo-bls-go/bls"

// ProofOfSpace
//
// 		challenge: bytes32
// 		pool_public_key: Optional[G1Element]  # Only one of these two should be present
// 		pool_contract_puzzle_hash: Optional[bytes32]
// 		plot_public_key: G1Element
// 		size: uint8
// 		proof: bytes
//
type ProofOfSpace struct {
	Challenge [32]byte `json:"challenge"`

	PoolPublicKey bls.PublicKey `json:"pool_public_key"`
	PlotPublicKey bls.PublicKey `json:"plot_public_key"`

	PoolContractPuzzleHash [32]byte `json:"pool_contract_puzzle_hash"`

	Size uint8 `json:"size"`

	Proof []byte `json:"proof"`
}
