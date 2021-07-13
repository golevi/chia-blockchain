// https://raw.githubusercontent.com/Chia-Network/chia-blockchain/main/chia/types/blockchain_format/vdf.py

package blockchain

// VDFInfo
//
// 		challenge: bytes32  # Used to generate the discriminant (VDF group)
// 		number_of_iterations: uint64
// 		output: ClassgroupElement
//
type VDFInfo struct {
	Challenge          []byte            `json:"challenge"`
	NumberOfIterations uint64            `json:"number_of_iterations"`
	Output             ClassgroupElement `json:"output"`
}

// VDFProof
//
// 		witness_type: uint8
// 		witness: bytes
// 		normalized_to_identity: bool
//
type VDFProof struct {
	WitnessType          uint8  `json:"witness_type"`
	Witness              []byte `json:"witness"`
	NormalizedToIdentity bool   `json:"normalized_to_identity"`
}
