package browser

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type BrowserCli interface {
	Get(url string) (*http.Response, error)
	Post(url string, payload io.Reader) (*http.Response, error)
	Put(url string, payload io.Reader) (*http.Response, error)
	Delete(url string) (*http.Response, error)
	List(url string) (*http.Response, error)
	Patch(url string, payload io.Reader) (*http.Response, error)
	CopyConfig() BrowserConfig
	AddHeader(key, value string)
	SetHeader(key, value string)
	GetHeader() http.Header
	SetUserAgentName(name string)
}

type BrowserConfig struct {
	BaseURL         string      `json:"base_url"`
	SSLVerify       bool        `json:"ssl_verify"`
	Header          http.Header `json:"header"`
	Timeout         int64       `json:"timeout"`
	TLSClientConfig *tls.Config `json:"-"`
	ProxyURL        string      `json:"proxy_url"`
}

type browser_cli struct {
	BrowserConfig
	http_client *http.Client
	http_req    *http.Request
}

func NewBrowser(bro_conf BrowserConfig) BrowserCli {

	new_timeout := time.Duration(10) * time.Second // Default 10 segundos

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Erro ao criar cookiejar")
		fmt.Println(err.Error())
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !bro_conf.SSLVerify},
	}

	if bro_conf.TLSClientConfig != nil {
		tr.TLSClientConfig = bro_conf.TLSClientConfig
	}

	if bro_conf.ProxyURL != "" {
		proxyURL, _ := url.Parse(bro_conf.ProxyURL)
		tr.Proxy = http.ProxyURL(proxyURL)
	}

	if bro_conf.Timeout > 0 {
		new_timeout = time.Duration(bro_conf.Timeout) * time.Second
	}

	http_client := &http.Client{
		Jar:       jar,
		Transport: tr,
		Timeout:   new_timeout,
	}

	client := browser_cli{
		http_client: http_client,
	}

	client.BrowserConfig = bro_conf

	if bro_conf.Header == nil {
		client.Header = make(http.Header)
	} else {
		client.Header = bro_conf.Header
	}

	client.Header.Set("User-Agent", "Go Browser "+VERSION)

	return &client
}

func (b *browser_cli) do(method, url string, body io.Reader) (*http.Response, error) {

	if b.BaseURL != "" {
		url = fmt.Sprintf("%v%v", b.BaseURL, url)
	}

	http_req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Printf("Erro ao criar HTTP %v Request\n", method)
		fmt.Println(err.Error())
	}

	b.http_req = http_req
	b.http_req.Header = b.Header

	resp, err := b.http_client.Do(http_req)
	if err != nil {
		fmt.Println(err.Error())
	}

	return resp, err
}

func (b *browser_cli) CopyConfig() BrowserConfig {
	return b.BrowserConfig
}
