package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func testeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/teste-requisição" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Opa, seu servidor web está funcionando!")
	fmt.Println("Requisição no servidor web porta 8080 no endpoint teste")
}
func calculadoraHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var data []byte
	type calculo struct {
		Val1     float64 `json:"Val1"`
		Val2     float64 `json:"Val2"`
		Operacao string  `json:"operacao"`
	}
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var operacao = calculo{}
	data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Erro em ioutil.ReadAll ", err)
	}
	err = json.Unmarshal(data, &operacao)
	if err != nil {
		fmt.Println("Erro em json.Unmarshal ", err)
	}
	if operacao.Operacao == "-" {
		total := operacao.Val1 - operacao.Val2
		fmt.Fprint(w, "resultado é ", total)
		fmt.Println("Requisição no servidor web porta 8080 no endpoint calculadora")
		fmt.Println(operacao.Val1, operacao.Operacao, operacao.Val2)
		fmt.Println("resultado é ", total)

		return

	} else if operacao.Operacao == "+" {
		total := operacao.Val1 + operacao.Val2
		fmt.Fprint(w, "resultado é ", total)
		fmt.Println("Requisição no servidor web porta 8080 no endpoint calculadora")
		fmt.Println(operacao.Val1, operacao.Operacao, operacao.Val2)
		fmt.Println("resultado é ", total)

		return

	} else if operacao.Operacao == "/" {
		total := operacao.Val1 / operacao.Val2
		fmt.Fprint(w, "resultado é ", total)
		fmt.Println("Requisição no servidor web porta 8080 no endpoint calculadora")
		fmt.Println(operacao.Val1, operacao.Operacao, operacao.Val2)
		fmt.Println("resultado é ", total)

		return

	} else if operacao.Operacao == "*" {
		total := operacao.Val1 * operacao.Val2

		fmt.Fprint(w, "resultado é ", total)
		fmt.Println("Requisição no servidor web porta 8080 no endpoint calculadora")
		fmt.Println(operacao.Val1, operacao.Operacao, operacao.Val2)
		fmt.Println("resultado é ", total)

		return

	} else {
		fmt.Fprintf(w, "operador não válido, operadores + - / * aceitos")
		return

	}
}
func main() {
	http.HandleFunc("/teste-requisição", testeHandler)
	http.HandleFunc("/calculadora", calculadoraHandler)
	fmt.Printf("Servidor rodando na porta 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
