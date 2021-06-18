package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/faelp22/browser/bro"
)

func main() {
	base_url := "https://api.github.com"

	bro := bro.NewBrowser(base_url, false, http.Header{}, 10)
	bro.AddHeader("Content-Type", "application/json; charset=utf-8")

	resp, err := bro.Get("/users/faelp22")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler Body da request")
		fmt.Println(err.Error())
	}

	if resp.StatusCode > 200 {
		fmt.Println("Erro na requisição")
		fmt.Printf("StatusCode: %v\n", resp.StatusCode)
		fmt.Println(string(body))
		os.Exit(1)
	}

	fmt.Println(string(body))

}
