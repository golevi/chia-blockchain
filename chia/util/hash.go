// https://github.com/Chia-Network/chia-blockchain/blob/1.2.1/chia/util/hash.py

package util

import "crypto/sha256"

// StandardHash creates a hash.
//
// https://github.com/Chia-Network/chia-blockchain/blob/1.2.1/chia/util/hash.py#L6
// https://github.com/Chia-Network/bls-signatures/blob/main/python-bindings/pythonbindings.cpp#L95
func StandardHash(b []byte) []byte {
	// I'm not sure if the blspy hash256 does anything special. I'm going to
	// have it create a regular SHA 256 hash.
	//
	// return bytes32(blspy.Util.hash256(bytes(b)))

	hasher := sha256.New()
	hasher.Write(b)

	return hasher.Sum(nil)
}
