// Link => https://resources.beecrowd.com/repository/UOJ_1153.html
package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int64
	fac := new(big.Int)
	fmt.Scanln(&n)
	fac.MulRange(1, n)
	fmt.Println(fac)
}
