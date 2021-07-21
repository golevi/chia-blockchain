package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/golevi/chia-blockchain/chia/consensus"
)

func main() {
	target := uint64(2031616)

	size := 32
	difficulty := 2400

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	start := time.Now()

	x := uint64(0)

	for {
		x++
		difficultyConstantFactor := big.NewInt(1)
		difficultyConstantFactor.SetString("147573952589676412928", 10)

		qualityString, _ := hex.DecodeString("6bc7e96778d56f99640d7d606682543a1449fe84c53d9b3bd764decacc29a10c")

		CCSPOutputHash, _ := hex.DecodeString("396fef5662016a16c8849b58bcbd6362368792f637b2a7a2abd91db2f35b9a80")

		t := r.Uint64()

		bs := make([]byte, 8)
		binary.LittleEndian.PutUint64(bs, t)
		// fmt.Println(bs)

		copy(CCSPOutputHash[0:8], bs[:])

		// fmt.Println(qualityString)
		// fmt.Println(CCSPOutputHash)

		result := consensus.CalculateIterationsQuality(
			*difficultyConstantFactor,
			qualityString,
			size,
			uint64(difficulty),
			CCSPOutputHash,
		)

		if result < target {
			fmt.Println(result, target)

			end := time.Now()

			diff := end.Sub(start)

			fmt.Printf("Duration: %s\n", diff)
			fmt.Printf("X: %d\n", x)
			// return
		}

		// return
	}
	// 3718598
	// 2585665
	// 1929814
	// 1417597
	// 889875
	// 6380416
	// 6253417
	// 6078782
}
