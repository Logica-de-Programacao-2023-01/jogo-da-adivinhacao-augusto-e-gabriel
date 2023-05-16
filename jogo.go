package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numero int
	var tentativas int
	var jogar_de_novo string

	fmt.Println("Bem-vindo ao jogo da adivinhação!")

	for {
		fmt.Print("Digite um número entre 1 e 100: ")
		fmt.Scan(&numero)

		numero_aleatorio := rand.Intn(2) + 1

		for {
			if numero < numero_aleatorio {
				fmt.Printf("O número é maior que %v\n", numero)
				fmt.Print("Digite um número entre 1 e 100: ")
				fmt.Scan(&numero)
				tentativas++
			} else if numero > numero_aleatorio {
				fmt.Printf("O número é menor que %v\n", numero)
				fmt.Print("Digite um número entre 1 e 100: ")
				fmt.Scan(&numero)
				tentativas++
			} else if numero == numero_aleatorio {
				fmt.Println("Parabéns, você acertou!")
				tentativas++
				fmt.Printf("Você utilizou %v tentativas\n", tentativas)
				break
			}
		}
		fmt.Print("Você quer jogar de novo? (s/n):")
		fmt.Scan(&jogar_de_novo)
		if jogar_de_novo == "n" {
			break
		}
	}
}
