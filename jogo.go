package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numero int
	var jogar_de_novo string
	replay := true
	lista := []int{}

	fmt.Println("Bem-vindo ao jogo da adivinhação!")
	numero_aleatorio := rand.Intn(5) + 1
	print(numero_aleatorio)

	for tentativas := 1; replay; tentativas++ {
		fmt.Print("Digite um número entre 1 e 100: ")
		fmt.Scan(&numero)
		if numero < 1 || numero > 100 {
			fmt.Println("Numeros aceitos: 1 a 100, tente novamente!")
			fmt.Print("Digite um número entre 1 e 100: ")
			fmt.Scan(&numero)
		}

		if numero < numero_aleatorio {
			fmt.Printf("O número é maior que %v\n", numero)
		} else if numero > numero_aleatorio {
			fmt.Printf("O número é menor que %v\n", numero)
		} else if numero == numero_aleatorio {
			fmt.Println("Parabéns, você acertou!")
			numero_aleatorio = 0
			numero_aleatorio = rand.Intn(5) + 1
			fmt.Printf("Você utilizou %v tentativas\n", tentativas)
			lista = append(lista, tentativas)
			numero = 0
			tentativas = 0
			fmt.Print("Você quer jogar de novo? (s/n):")
			fmt.Scan(&jogar_de_novo)

			if jogar_de_novo != "s" && jogar_de_novo != "n" {
				fmt.Println("Digite 's' para sim e 'n' ou não!")
				fmt.Print("Você quer jogar de novo? (s/n):")
				fmt.Scan(&jogar_de_novo)

			}

			if jogar_de_novo == "n" {
				for i := 0; i < len(lista); i++ {
					fmt.Printf("Na rodada %v você utilizou %v tentativas\n", i+1, lista[i])
				}
				replay = false
			}
		}
	}
}
