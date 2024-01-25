package parse

import (
	"github.com/HzTTT/crawler/collect"
	"regexp"
)

const cityListRe = `https://www\.douban\.com/group/topic/\d+/`

func ParseURL(contents []byte) collect.ParseResult {
	re := regexp.MustCompile(cityListRe)

	matches := re.FindAllSubmatch(contents, -1)
	result := collect.ParseResult{}

	for _, m := range matches {
		u := string(m[0])
		result.Requests = append(
			result.Requests, &collect.Request{
				Url: u,
				ParseFunc: func(c []byte) collect.ParseResult {
					return GetContent(c, u)
				},
			})
	}
	return result
}

const ContentRe = `<div class="topic-content">[\s\S]*?阳台[\s\S]*?<div`

func GetContent(contents []byte, url string) collect.ParseResult {
	re := regexp.MustCompile(ContentRe)

	ok := re.Match(contents)
	if !ok {
		return collect.ParseResult{
			Items: []interface{}{},
		}
	}

	result := collect.ParseResult{
		Items: []interface{}{url},
	}

	return result
}
