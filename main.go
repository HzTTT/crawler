package main

import (
	"fmt"
	"github.com/HzTTT/crawler/collect"
	"github.com/HzTTT/crawler/log"
	"github.com/HzTTT/crawler/parse"
	"go.uber.org/zap"

	"time"
)

func main() {

	plugin, closer := log.NewStdFilePlugin("./log.txt", zap.InfoLevel)
	defer closer.Close()
	logger := log.NewLogger(plugin)
	logger.Info("log init end")

	var worklist []*collect.Request
	for i := 0; i <= 100; i += 25 {
		str := fmt.Sprintf("https://www.douban.com/group/gz020/discussion?start=%d&type=new", i)
		worklist = append(worklist, &collect.Request{
			Url:       str,
			ParseFunc: parse.ParseURL,
		})
	}

	var f collect.Fetcher = &collect.BrowserFetch{
		Timeout: 3000 * time.Millisecond,
		Proxy:   nil,
	}

	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			body, err := f.Get(item.Url)
			time.Sleep(1 * time.Second)
			if err != nil {
				logger.Error("read content failed",
					zap.Error(err),
				)
				continue
			}
			res := item.ParseFunc(body)
			for _, item := range res.Items {
				logger.Info("result",
					zap.String("get url:", item.(string)))
			}
			worklist = append(worklist, res.Requests...)
		}
	}

}
