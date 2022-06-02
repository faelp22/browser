package browser

import (
	"io"
	"net/http"
)

func (b *browser_cli) Get(url string) (*http.Response, error) {
	return b.do("GET", url, nil)
}

func (b *browser_cli) Post(url string, payload io.Reader) (*http.Response, error) {
	return b.do("POST", url, payload)
}

func (b *browser_cli) Put(url string, payload io.Reader) (*http.Response, error) {
	return b.do("PUT", url, payload)
}

func (b *browser_cli) Delete(url string) (*http.Response, error) {
	return b.do("DELETE", url, nil)
}

func (b *browser_cli) List(url string) (*http.Response, error) {
	return b.do("LIST", url, nil)
}

func (b *browser_cli) Patch(url string, payload io.Reader) (*http.Response, error) {
	return b.do("PATCH", url, payload)
}
