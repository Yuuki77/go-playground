package elementsofprograming

import (
	"fmt"

	"github.com/willf/bitset"
)

func countBits(x int) int {

	numBits := 0

	for index := 0; x != 0; index++ {
		numBits += (x & 1)
		x >>= 1
	}
	return numBits
}

// https://qiita.com/drken/items/7c6ff2aa4d8fce1c9361
// 45 = 0b101101
// 25 = 0b011001
// 45 AND 25 = 0b001001
// 各桁の掛け算ほしいものだけを取れる
func bitAnd(x int, y int) int {

	return x & y
}

// https://qiita.com/drken/items/7c6ff2aa4d8fce1c9361
// 45 = 0b101101
// 25 = 0b011001
// 45 OR 25 = 0b111101

func bitOr(x int, y int) int {

	return x | y
}

func bitOperationsTest() {
	const flag0 uint = (1 << 0) // 0000 0000 0000 0001
	const flag1 uint = (1 << 1) // 0000 0000 0000 0010
	const flag2 uint = (1 << 2) // 0000 0000 0000 0100
	const flag3 uint = (1 << 3) // 0000 0000 0000 1000
	const flag4 uint = (1 << 4) // 0000 0000 0001 0000
	const flag5 uint = (1 << 5) // 0000 0000 0010 0000
	const flag6 uint = (1 << 6) // 0000 0000 0100 0000
	const flag7 uint = (1 << 7) // 0000 0000 1000 0000

	// {1, 3, 5} にフラグの立ったビット
	bit := flag1 | flag3 | flag5
	var b bitset.BitSet
	b.Set(bit)
	fmt.Printf("<<{1, 3, 5} = %#b", b.Bytes())

	// 		b.Set(card1)

	// cout<<"{1, 3, 5} = "<<bitset < 8 > (bit)<<endl
}
