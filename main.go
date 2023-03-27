package main

import (
	"fmt"
	"math/rand"
)

const chances = 3

func main() {
	numeroSecreto := gerarNumero()

	var chute int
	var tentativa int

	for tentativa = 1; tentativa <= chances; tentativa++ {
		chute = solicitarChute()
		acertou := testarChute(chute, numeroSecreto)

		if acertou {
			exibirMensagemAcerto()
			break
		} else {
			exibirMensagemErro(tentativa)
		}
	}
}

func gerarNumero() int {
	return rand.Intn(100)
}

func solicitarChute() int {
	var chute int
	fmt.Print("Digite o seu chute: ")
	fmt.Scan(&chute)
	return chute
}

func testarChute(chute int, numeroSecreto int) bool {
	if chute == numeroSecreto {
		return true
	} else {
		return false
	}
}

func exibirMensagemAcerto() {
	fmt.Println("Parabéns, você acertou.")
}

func exibirMensagemErro(tentativa int) {
	fmt.Println("Você errou. Essa é a sua tentativa ", tentativa)
}
