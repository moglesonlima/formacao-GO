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

const monitoramentos = 2
const delay = 5

func main() {
	exibeIntroducao()
	verificaArquivoSites()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 3:
			listagemSitesMonitorados()
		case 4:
			adicionaSite()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	var nome string

	fmt.Println("Olá, seja bem vindo(a) ao monitor de sites!\nPoderia me informar seu nome?")
	fmt.Scanf("%s", &nome)
	versao := "2.0"
	fmt.Println("Seja bem vindo", nome, "\nEste programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("(1)- Iniciar Monitoramento")
	fmt.Println("(2)- Exibir Logs")
	fmt.Println("(3)- Listar Sites Monitorados")
	fmt.Println("(4)- Adicionar Site")
	fmt.Println("(0)- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
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

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()
	return sites
}

func criaArquivoSites() {

	arquivo, err := os.Create("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString("https://random-status-code.herokuapp.com/\n")
	arquivo.WriteString("https://www.alura.com.br\n")
	arquivo.WriteString("https://www.caelum.com.br")

	arquivo.Close()
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}

func verificaArquivoSites() {
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Arquivo não encontrado, criando arquivo sites.txt...")
		criaArquivoSites()
	}

	arquivo.Close()
}

func listagemSitesMonitorados() {
	sites := leSitesDoArquivo()
	fmt.Println("Sites monitorados: ")
	for i, site := range sites {
		fmt.Println(i, " - ", site)
	}
	fmt.Println("")
}

func adicionaSite() {
	fmt.Println("Digite o site que deseja adicionar: ")
	var site string
	fmt.Scan(&site)

	arquivo, _ := os.OpenFile("sites.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	arquivo.WriteString("\n" + site)
	arquivo.Close()

	fmt.Println("Site - ", site, "- Adicionado com sucesso!")
}
