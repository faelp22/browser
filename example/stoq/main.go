package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/faelp22/browser"
	"github.com/faelp22/browser/prepare"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAdmin() *User {
	return &User{
		Username: "admin",
		Password: "supersenha",
	}
}

type Token struct {
	Token string `json:"token"`
}

const USER_TOKEN = "fake-WzD5fqrlaAXLv26bpI0hxvAhDp7T1Bac"

func main() {

	bro := browser.NewBrowser(browser.BrowserConfig{
		BaseURL:   "http://localhost:8080",
		SSLVerify: false,
		Header: http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		},
		Timeout: 3, // 3 segundos
	})

	token := Login(bro)

	bro.AddHeader("Authorization", "Bearer "+token.Token)
	fmt.Println("GetHeader", bro.GetHeader())

	resp, err := bro.Get("/api/v1/products")
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

func Login(bro browser.BrowserCli) *Token {
	url := "/api/v1/user/login"

	user := NewAdmin()

	dados, _ := json.Marshal(user)

	// fmt.Println(string(dados))

	payload := prepare.PrepareJSON(bro, dados)

	resp, err := bro.Post(url, payload)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler Body da request")
		fmt.Println(err.Error())
	}

	// fmt.Println(string(body))

	token := Token{}

	err = json.Unmarshal(body, &token)
	if err != nil {
		fmt.Println("Erro ao ler Body da request")
		fmt.Println(err.Error())
	}

	fmt.Println(token.Token)

	return &token
}
