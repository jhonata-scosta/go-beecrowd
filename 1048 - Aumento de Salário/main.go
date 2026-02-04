// Link => https://resources.beecrowd.com/repository/UOJ_1048.html
package main

import "fmt"

func main() {
	var salarioAtual float64
	fmt.Scanf("%f", &salarioAtual)
	switch {
	case salarioAtual <= 400.00:
		reajuste := salarioAtual * 0.15
		novoSalario := salarioAtual + reajuste
		fmt.Printf("Novo salario: %.2f\n", novoSalario)
		fmt.Printf("Reajuste ganho: %.2f\n", reajuste)
		fmt.Println("Em percentual: 15 %")
	case (salarioAtual > 400.00) && (salarioAtual <= 800.00):
		reajuste := salarioAtual * 0.12
		novoSalario := salarioAtual + reajuste
		fmt.Printf("Novo salario: %.2f\n", novoSalario)
		fmt.Printf("Reajuste ganho: %.2f\n", reajuste)
		fmt.Println("Em percentual: 12 %")
	case (salarioAtual > 800.00) && (salarioAtual <= 1200.00):
		reajuste := salarioAtual * 0.10
		novoSalario := salarioAtual + reajuste
		fmt.Printf("Novo salario: %.2f\n", novoSalario)
		fmt.Printf("Reajuste ganho: %.2f\n", reajuste)
		fmt.Println("Em percentual: 10 %")
	case (salarioAtual > 1200.00) && (salarioAtual <= 2000.00):
		reajuste := salarioAtual * 0.07
		novoSalario := salarioAtual + reajuste
		fmt.Printf("Novo salario: %.2f\n", novoSalario)
		fmt.Printf("Reajuste ganho: %.2f\n", reajuste)
		fmt.Println("Em percentual: 7 %")
	case salarioAtual > 2000.00:
		reajuste := salarioAtual * 0.04
		novoSalario := salarioAtual + reajuste
		fmt.Printf("Novo salario: %.2f\n", novoSalario)
		fmt.Printf("Reajuste ganho: %.2f\n", reajuste)
		fmt.Println("Em percentual: 4 %")
	}

}
