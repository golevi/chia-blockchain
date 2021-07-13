package blockchain

import (
	"testing"

	"github.com/celo-org/celo-bls-go/bls"
)

// https://github.com/Chia-Network/bls-signatures/blob/main/src/test.cpp

func TestBLSSignature(t *testing.T) {
	msg := []byte("Hello")

	seed := make([]byte, 32)
	privateKey, err := bls.DeserializePrivateKey(seed)
	if err != nil {
		t.Error(err)
	}

	sig, err := privateKey.SignMessage(msg, nil, false, false)
	if err != nil {
		t.Error(err)
	}

	publicKey, err := privateKey.ToPublic()
	if err != nil {
		t.Error(err)
	}

	err = publicKey.VerifySignature(msg, nil, sig, false, false)
	if err != nil {
		t.Error(err)
	}
}
