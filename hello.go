package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibirIntroducao()
	exibirMenu()

	comando := leComando()
	executarComando(comando)
}

func exibirIntroducao() {
	usuario := "Usuário"
	versao := 1.1

	fmt.Println()
	fmt.Println("Olá sr.(a)", usuario)
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
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		fmt.Println("Programa finalizado")
		os.Exit(0)
	default:
		fmt.Println("Comando não encontrado")
		os.Exit(-1)
	}
}

func iniciarMonitoramento() {
	// arrayTradicional := [4]string[1,2,3,4]
	//Slice - Array abaixo que não tem tamanho definido, e pode adicionar elementos conforme necessário
	sites := []string{"https://httpbin.org/status/200",
		"https://www.alura.com.br", "https://www.caelum.com.br", "https://httpbin.org/status/400"}

	fmt.Println("Monitorando site(s)...")
	for posicao, site := range sites {
		fmt.Println("Testando site", posicao, ":", site)
		testaSite(site)
	}

	fmt.Println("")
	fmt.Println("Programa finalizado")
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "está com problemas. Não foi possível acessar - Status:", resp.StatusCode)
	}
}
