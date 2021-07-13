package util

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestStandardHash(t *testing.T) {
	data := []struct {
		Input  string
		Output string
	}{
		{
			"Hello",
			"185f8db32271fe25f561a6fc938b2e264306ec304eda518007d1764826381969",
		},
		{
			"Hi",
			"3639efcd08abb273b1619e82e78c29a7df02c1051b1820e99fc395dcaa3326b8",
		},
	}

	for _, d := range data {
		hash := StandardHash([]byte(d.Input))
		output, err := hex.DecodeString(d.Output)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(hash, output) {
			t.Errorf("Expected %x, got %x", output, hash)
		}
	}
}
