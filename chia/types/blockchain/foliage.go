// https://raw.githubusercontent.com/Chia-Network/chia-blockchain/main/chia/types/blockchain_format/foliage.py

package blockchain

import "github.com/celo-org/celo-bls-go/bls"

// TransactionsInfo
//
// 		# Information that goes along with each transaction block
// 		generator_root: bytes32  # sha256 of the block generator in this block
// 		generator_refs_root: bytes32  # sha256 of the concatenation of the generator ref list entries
// 		aggregated_signature: G2Element
// 		fees: uint64  # This only includes user fees, not block rewards
// 		cost: uint64  # This is the total cost of this block, including CLVM cost, cost of program size and conditions
// 		reward_claims_incorporated: List[Coin]  # These can be in any order
//
type TransactionsInfo struct {
	GeneratorRoot     [32]byte `json:"generator_root"`
	GeneratorRefsRoot [32]byte `json:"generator_refs_root"`

	// AggregatedSignature is a BLS Signature.
	//
	// In the official chia-blockchain Python code they use BLSPY. BLSPY calls a
	// signature a G2Element. A G2Element is a signature (96 bytes). The Go pkg
	// for Celo BLS says a signature is only 48 bytes, so this may cause some
	// issues and will need to be tested.
	//
	// Links
	// * https://pypi.org/project/blspy/
	// * https://pkg.go.dev/github.com/celo-org/celo-bls-go/bls
	// * https://crypto.stanford.edu/~dabo/pubs/papers/BLSmultisig.html
	AggregatedSignature bls.Signature `json:"aggregated_signature"`

	Fees uint64 `json:"fees"`
	Cost uint64 `json:"cost"`

	RewardClaimsIncorporated []Coin `json:"reward_claims_incorporated"`
}

// FoliageTransactionBlock
//
// 		# Information that goes along with each transaction block that is relevant for light clients
// 		prev_transaction_block_hash: bytes32
// 		timestamp: uint64
// 		filter_hash: bytes32
// 		additions_root: bytes32
// 		removals_root: bytes32
// 		transactions_info_hash: bytes32
//
type FoliageTransactionBlock struct {
	PrevTransactionBlockHash [32]byte `json:"prev_transaction_block_hash"`

	Timestamp uint64 `json:"timestamp"`

	FilterHash           [32]byte `json:"filter_hash"`
	AdditionsRoot        [32]byte `json:"additions_root"`
	RemovalsRoot         [32]byte `json:"removals_root"`
	TransactionsInfoHash [32]byte `json:"transactions_info_hash"`
}

// FoliageBlockData
//
// 		# Part of the block that is signed by the plot key
// 		unfinished_reward_block_hash: bytes32
// 		pool_target: PoolTarget
// 		pool_signature: Optional[G2Element]  # Iff ProofOfSpace has a pool pk
// 		farmer_reward_puzzle_hash: bytes32
// 		extension_data: bytes32  # Used for future updates. Can be any 32 byte value initially
//
type FoliageBlockData struct {
	UnfinishedRewardBlockHash [32]byte `json:"unfinished_reward_block_hash"`

	PoolTarget    PoolTarget    `json:"pool_target"`
	PoolSignature bls.Signature `json:"pool_signature"`

	FarmerRewardPuzzleHash [32]byte `json:"farmer_reward_puzzle_hash"`
	ExtensionData          [32]byte `json:"extension_data"`
}

// Foliage
//
// 		# The entire foliage block, containing signature and the unsigned back pointer
// 		# The hash of this is the "header hash". Note that for unfinished blocks, the prev_block_hash
// 		# Is the prev from the signage point, and can be replaced with a more recent block
// 		prev_block_hash: bytes32
// 		reward_block_hash: bytes32
// 		foliage_block_data: FoliageBlockData
// 		foliage_block_data_signature: G2Element
// 		foliage_transaction_block_hash: Optional[bytes32]
// 		foliage_transaction_block_signature: Optional[G2Element]
//
type Foliage struct {
	PrevBlockHash   [32]byte `json:"prev_block_hash"`
	RewardBlockHash [32]byte `json:"reward_block_hash"`

	FoliageBlockData                 FoliageBlockData `json:"foliage_block_data"`
	FoliageBlockDataSignature        bls.Signature    `json:"foliage_block_data_signature"`
	FoliageTransactionBlockHash      [32]byte         `json:"foliage_transaction_block_hash"`
	FoliageTransactionBlockSignature bls.Signature    `json:"foliage_transaction_block_signature"`
}
