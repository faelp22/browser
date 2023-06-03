package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/faelp22/browser"
)

func main() {

	bro := browser.NewBrowser(browser.BrowserConfig{
		BaseURL:   "https://api.github.com",
		SSLVerify: false,
		Header: http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		},
		Timeout: 3, // 3 segundos
		Mode:    browser.DEVELOPER,
	})

	// bro.SetUserAgentName("Meu Cli v-0.1.0")

	// fmt.Println(bro.GetHeader())

	// bro.AddHeader("Content-Type", "application/json; charset=utf-8") // Add append in content
	// bro.SetHeader("Content-Type", "application/json; charset=utf-8") // Set replace the content

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
