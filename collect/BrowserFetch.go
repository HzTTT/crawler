package collect

import (
	"fmt"
	"github.com/HzTTT/crawler/proxy"
	"io"
	"net/http"
	"time"
)

type BrowserFetch struct {
	Timeout time.Duration
	Proxy   proxy.ProxyFunc
}

func (b *BrowserFetch) Get(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: b.Timeout,
	}

	if b.Proxy != nil {
		transport := http.DefaultTransport.(*http.Transport)
		transport.Proxy = b.Proxy
		client.Transport = transport
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("new request error : %s", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request error : %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error : %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)

}
