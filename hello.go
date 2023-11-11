package main

import (
	"fmt"
	"os"
)

func main() {

	exibirIntroducao()
	exibirMenu()

	comando := leComando()
	executarComando(comando)
}

func exibirIntroducao() {
	nome := "Alefe"
	versao := 1.1

	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println()
}

func exibirMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)

	return comando
}

func executarComando(comando int) {
	switch comando {
	case 1:
		fmt.Println("Monitorando site...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Comando não encontrado")
		os.Exit(-1)
	}
}
