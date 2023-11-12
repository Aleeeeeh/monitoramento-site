package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
		imprimeLogs()
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
	sites := leSitesDoArquivo()

	fmt.Println("Monitorando site(s)...")
	for posicao, site := range sites {
		fmt.Println("Testando site", posicao, ":", site)
		testaSite(site)
	}

	fmt.Println("")
	fmt.Println("Programa finalizado")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Não foi possível acessar - Status:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')

		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)
		//EOF -  Esse erro indica que chegou ao final do arquivo
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site +
		" - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}

/*arrayTradicional := [4]string[1,2,3,4]
/Slice - []string{1,2,3,4}
Diferença entre os dois é que o tradicional tem a capacidade definida e o slice é dinâmico. Podendo
adicionar mais elementos utilizando o append
*/
