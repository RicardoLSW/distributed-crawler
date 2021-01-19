package parser

import (
	"distributed-crawler/config"
	"regexp"

	"distributed-crawler/engine"
)

var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), Parser: NewProfileParser(string(m[2]))})
	}
	matches2 := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches2 {
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), Parser: engine.NewFuncParser(ParseCity, config.ParseCity)})
	}
	return result
}
