package main

import (
	"fmt"
	"github.com/HzTTT/crawler/collect"
	"github.com/HzTTT/crawler/log"
	"github.com/HzTTT/crawler/proxy"
	"go.uber.org/zap"

	"time"
)

func main() {
	plugin, closer := log.NewStdFilePlugin("./log.txt", zap.InfoLevel)
	defer closer.Close()
	logger := log.NewLogger(plugin)
	logger.Info("log init end")
	proxyURLs := []string{"http://127.0.0.1:7890"}
	proxyFunc, err := proxy.RoundRobinProxyFunc(proxyURLs...)
	if err != nil {
		return
	}
	url := "https://www.google.com/"
	fetcher := collect.BrowserFetch{
		Timeout: 3000 * time.Millisecond,
		Proxy:   proxyFunc,
	}
	body, err := fetcher.Get(url)
	if err != nil {
		fmt.Printf("fetch url error : %s", err)
		return
	}
	fmt.Printf("body : %s", body)
}
