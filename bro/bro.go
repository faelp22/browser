package bro

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"
)

type browser struct {
	base_url    string
	ssl_verify  bool
	Header      http.Header
	http_client *http.Client
	http_req    *http.Request
}

func NewBrowser(base_url string, ssl_verify bool, header http.Header, timeout int64) *browser {

	new_timeout := time.Duration(10) * time.Second // Default 10 segundos

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Erro ao criar cookiejar")
		fmt.Println(err.Error())
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: ssl_verify},
	}

	if timeout > 0 {
		new_timeout = time.Duration(timeout) * time.Second
	}

	http_client := &http.Client{
		Jar:       jar,
		Transport: tr,
		Timeout:   new_timeout,
	}

	browser := browser{
		base_url:    base_url,
		ssl_verify:  ssl_verify,
		Header:      header,
		http_client: http_client,
	}

	return &browser
}

func (b *browser) do(http_req *http.Request) (*http.Response, error) {
	b.http_req = http_req
	b.updateHeaders()
	resp, err := b.http_client.Do(http_req)
	if err != nil {
		fmt.Println(err.Error())
	}

	return resp, err
}

func (b *browser) Get(url string) (*http.Response, error) {
	if b.base_url != "" {
		url = fmt.Sprintf("%v%v", b.base_url, url)
	}
	http_req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Erro ao criar HTTP GET Request")
		fmt.Println(err.Error())
	}
	return b.do(http_req)
}

func (b *browser) Post(url string, payload io.Reader) (*http.Response, error) {
	if b.base_url != "" {
		url = fmt.Sprintf("%v%v", b.base_url, url)
	}
	http_req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println("Erro ao criar HTTP POST Request")
		fmt.Println(err.Error())
	}
	return b.do(http_req)
}

func (b *browser) Put(url string, payload io.Reader) (*http.Response, error) {
	http_req, err := http.NewRequest("PUT", url, payload)
	if err != nil {
		fmt.Println("Erro ao criar HTTP PUT Request")
		fmt.Println(err.Error())
	}
	return b.do(http_req)
}

func (b *browser) Delete(url string) (*http.Response, error) {
	http_req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println("Erro ao criar HTTP DELETE Request")
		fmt.Println(err.Error())
	}
	return b.do(http_req)
}

func (b *browser) Patch(url string, payload io.Reader) (*http.Response, error) {
	http_req, err := http.NewRequest("PATCH", url, payload)
	if err != nil {
		fmt.Println("Erro ao criar HTTP PATCH Request")
		fmt.Println(err.Error())
	}
	return b.do(http_req)
}

func (b *browser) AddHeader(key, value string) {
	b.Header[key] = append(b.Header[key], value)
}

func (b *browser) CreateJSON(payload []byte) (data_json io.Reader) {
	b.http_req.Header.Add("Content-Type", "	application/json; charset=utf-8")
	data_json = bytes.NewBuffer(payload)
	return
}

func (b *browser) CreateMultipartFormData(fileFieldName, filePath string, fileName string, extraFormFields map[string]string) (buf bytes.Buffer, w *multipart.Writer, err error) {
	b.http_req.Header.Add("Content-Type", w.FormDataContentType())
	w = multipart.NewWriter(&buf)
	defer w.Close()
	var fw io.Writer
	file, err := os.Open(filePath)

	if fw, err = w.CreateFormFile(fileFieldName, fileName); err != nil {
		return
	}
	if _, err = io.Copy(fw, file); err != nil {
		return
	}

	for k, v := range extraFormFields {
		_ = w.WriteField(k, v)
	}

	return
}

func (b *browser) updateHeaders() {
	b.http_req.Header = b.Header
}
