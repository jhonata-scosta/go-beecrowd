// Link => https://resources.beecrowd.com/repository/UOJ_1044.html
package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Scanf("%d %d", &n1, &n2)
	if (n1%n2 == 0) || (n2%n1 == 0) {
		fmt.Println("Sao Multiplos")
	} else {
		fmt.Println("Nao sao Multiplos")
	}
}
