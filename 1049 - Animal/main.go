// Link => https://resources.beecrowd.com/repository/UOJ_1049.html
package main

import "fmt"

func main() {
	var a1, a2, a3 string
	fmt.Scanf("%s\n", &a1)
	fmt.Scanf("%s\n", &a2)
	fmt.Scanf("%s\n", &a3)
	switch {
	case a1 == "vertebrado":
		switch {
		case a2 == "ave":
			switch {
			case a3 == "carnivoro":
				fmt.Println("aguia")
			case a3 == "onivoro":
				fmt.Println("pomba")
			}
		case a2 == "mamifero":
			switch {
			case a3 == "onivoro":
				fmt.Println("homem")
			case a3 == "herbivoro":
				fmt.Println("vaca")
			}
		}
	case a1 == "invertebrado":
		switch {
		case a2 == "inseto":
			switch {
			case a3 == "hematofago":
				fmt.Println("pulga")
			case a3 == "herbivoro":
				fmt.Println("lagarta")
			}
		case a2 == "anelideo":
			switch {
			case a3 == "hematofago":
				fmt.Println("sanguessuga")
			case a3 == "onivoro":
				fmt.Println("minhoca")
			}
		}
	}
}
