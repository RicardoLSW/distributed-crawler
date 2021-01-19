package parser

import (
	"distributed-crawler/model"
	"regexp"

	"distributed-crawler/engine"
)

var userInfoRe = regexp.MustCompile(`<div class="m-btn purple"[^<]*>([^<]+)</div>`)

var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([0-9]+)`)

func parseProfile(contents []byte, name string, url string) engine.ParseResult {
	match := userInfoRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	profile := model.Profile{}
	profile.Name = name
	for i, user := range match {
		switch i {
		case 0:
			profile.Marriage = string(user[1])
			break
		case 1:
			profile.Age = string(user[1])
			break
		case 2:
			profile.Xinzuo = string(user[1])
			break
		case 3:
			profile.Height = string(user[1])
			break
		case 4:
			break
		case 5:
			profile.Income = string(user[1])
			break
		case 6:
			profile.Education = string(user[1])
			break

		}
	}
	id := idUrlRe.FindSubmatch([]byte(url))
	result.Items = append(result.Items, engine.Item{
		Type:    "zhenai",
		Url:     url,
		Id:      string(id[1]),
		Payload: profile,
	})
	return result
}

type ProfileParser struct {
	userName string
}

func (p ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName)
}

func (p ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}
