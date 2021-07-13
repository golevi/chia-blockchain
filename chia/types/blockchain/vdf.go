// https://raw.githubusercontent.com/Chia-Network/chia-blockchain/main/chia/types/blockchain_format/vdf.py

package blockchain

// VDFInfo
//
// 		challenge: bytes32  # Used to generate the discriminant (VDF group)
// 		number_of_iterations: uint64
// 		output: ClassgroupElement
//
type VDFInfo struct {
	Challenge          [32]byte `json:"challenge"`
	NumberOfIterations uint64   `json:"number_of_iterations"`
	// Output             ClassgroupElement `json:"output"`
}
