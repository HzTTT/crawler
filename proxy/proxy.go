package proxy

import (
	"errors"
	"net/http"
	"net/url"
	"sync/atomic"
)

type ProxyFunc func(*http.Request) (*url.URL, error)

type roundRobinSwitcher struct {
	proxyURLs []*url.URL
	index     uint32
}

func (r *roundRobinSwitcher) GetProxy(pr *http.Request) (*url.URL, error) {
	index := atomic.AddUint32(&r.index, 1) - 1
	u := r.proxyURLs[int(index)%len(r.proxyURLs)]
	return u, nil
}

func RoundRobinProxyFunc(proxiesURLs ...string) (ProxyFunc, error) {
	if len(proxiesURLs) <= 0 {
		return nil, errors.New("empty proxy list")
	}

	urls, err := parseUrls(proxiesURLs)
	if err != nil {
		return nil, err
	}

	return (&roundRobinSwitcher{proxyURLs: urls, index: 0}).GetProxy, nil

}

func parseUrls(proxiesURLs []string) ([]*url.URL, error) {
	urls := make([]*url.URL, len(proxiesURLs))
	for i, u := range proxiesURLs {
		parseU, err := url.Parse(u)
		if err != nil {
			return nil, err
		}
		urls[i] = parseU
	}
	return urls, nil
}
