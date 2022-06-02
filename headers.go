package browser

import "net/http"

func (b *browser_cli) AddHeader(key, value string) {
	b.Header.Add(key, value)
}

func (b *browser_cli) SetHeader(key, value string) {
	b.Header.Set(key, value)
}

func (b *browser_cli) GetHeader() http.Header {
	return b.Header
}

func (b *browser_cli) SetUserAgentName(name string) {
	b.Header.Set("User-Agent", name)
}
