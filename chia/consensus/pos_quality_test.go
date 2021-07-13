package consensus

import "testing"

func TestExpectedPlotSize(t *testing.T) {
	data := []struct {
		K    int
		Size uint64
	}{
		{
			32,
			139586437120,
		},
		{
			34,
			592705486848,
		},
		{
			35,
			1219770712064,
		},
		{
			36,
			2508260900864,
		},
	}

	for _, d := range data {
		if ExpectedPlotSize(d.K) != d.Size {
			t.Errorf("Expected %d, got: %d", d.Size, ExpectedPlotSize(d.K))
		}
	}
}
