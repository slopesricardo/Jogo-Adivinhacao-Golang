package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

func main() {
	limparTela()
	ExibirApresentacao()
	numeroSecreto := gerarNumero()

	var chute int
	var tentativa int
	quantidadeTotalDeTentativa := solicitarNivelDeDificuldade()

	limparTela()

	for tentativa = 1; tentativa <= quantidadeTotalDeTentativa; tentativa++ {
		chute = solicitarChute()
		acertou := testarChute(chute, numeroSecreto)

		if acertou {
			exibirMensagemAcerto()
			break
		} else {
			exibirMensagemErro(tentativa)
			mostrarDica(numeroSecreto, chute)
		}
	}
}

func ExibirApresentacao() {
	fmt.Println("*******************************************")
	fmt.Println("*     Bem Vindo ao Jogo da Advinhação     *")
	fmt.Println("*******************************************")
	fmt.Println()
}

func solicitarNivelDeDificuldade() int {
	var opcao int
	fmt.Println("Qual o nível de dificuldade deseja jogar?")
	fmt.Println("1 - Três tentativas")
	fmt.Println("2 - Cinco tentativas")
	fmt.Println("3 - Sete tentativas")
	fmt.Scan(&opcao)

	var tentativas int
	switch opcao {
	case 1:
		tentativas = 3
	case 2:
		tentativas = 5
	case 3:
		tentativas = 7
	default:
		tentativas = 5
	}
	return tentativas
}

func limparTela() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
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
	fmt.Printf("Parabéns, você acertou.\n")
}

func exibirMensagemErro(tentativa int) {
	fmt.Println("Você errou. Essa é a sua tentativa número", tentativa)
}

func mostrarDica(numeroSecreto int, chute int) {
	if chute > numeroSecreto {
		fmt.Printf("Seu chute foi maior que o número secreto.\n\n")
	} else {
		fmt.Printf("Seu chute foi menor que o número secreto.\n\n")
	}
}
