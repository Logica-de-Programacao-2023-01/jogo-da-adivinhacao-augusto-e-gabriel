package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numero int
	var jogar_de_novo string
	tentativas := 1
	lista := []int{}

	fmt.Println("Bem-vindo ao jogo da adivinhação!")

	for {
		fmt.Print("Digite um número entre 1 e 100: ")
		fmt.Scan(&numero)

		numero_aleatorio := rand.Intn(100) + 1

		for {
			if numero < numero_aleatorio {
				fmt.Printf("O número é maior que %v\n", numero)
			} else if numero > numero_aleatorio {
				fmt.Printf("O número é menor que %v\n", numero)
			} else if numero == numero_aleatorio {
				fmt.Println("Parabéns, você acertou!")
				fmt.Printf("Você utilizou %v tentativas\n", tentativas)
				lista = append(lista, tentativas)
				tentativas = 1
				break
			}

			fmt.Print("Digite um número entre 1 e 100: ")
			fmt.Scan(&numero)
			tentativas++

		}
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
			break
		}

	}
}
