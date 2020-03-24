/*
* @Describe: 答案相关接口
* @Author:   mian
* @Date:     2020/3/19 15:48
 */
package tool

import (
	"github.com/kirinlabs/HttpRequest"
	"strconv"
)

var headers map[string]string = map[string]string{ // 带个请求头
	"authority":                 "www.zhihu.com",
	"method":                    "GET",
	"path":                      "/api/v4/answers/1065342258/voters?include=data%5B%2A%5D",
	"scheme":                    "https",
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"accept-encoding":           "utf-8",
	"accept-language":           "zh-CN,zh;q=0.9,en;q=0.8",
	"cache-control":             "max-age=0",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-site":            "none",
	"sec-fetch-user":            "?1",
	"upgrade-insecure-requests": ":",
	"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"}

type answer struct {
	code string // 问题编码
	req  HttpRequest.Request
}

func (a *answer) VoteNum() int {
	url := "https://www.zhihu.com/api/v4/answers/" + a.code + "/voters?include=data%5B%2A%5D"
	respond, _ := a.req.Get(url)
	j := map[string]interface{}{}
	_ = respond.Json(&j)
	r, _ := j["paging"].(map[string]interface{})["totals"].(float64)
	re := int(r)
	return re
}

func (a *answer) VotePageNum() int {
	return a.VoteNum()/10 + 1
}

func (a *answer) GetPage(index int) []interface{} {
	url := "https://www.zhihu.com/api/v4/answers/" + a.code + "/voters?include=data%5B%2A%5D.answer_count%2Carticles_count%2Cfollower_count%2Cis_followed%2Cis_following&offset=" +
		strconv.Itoa(
			(index+1)*10)
	respond, _ := a.req.Get(url)
	j := map[string]interface{}{}
	_ = respond.Json(&j)
	data, ok := j["data"]
	if ok {
		return data.([]interface{})
	} else {
		return []interface{}{}
	}
}

func (a *answer) GetAll() []interface{} {
	votes := []interface{}{}
	for i := 0; i < a.VotePageNum()+1; i++ {
		votes = append(votes, a.GetPage(i)...)
	}
	return votes
}

//NewAnswer 返回以回答编码相关的回答接口
func NewAnswer(code string) *answer {
	a := answer{
		code: code,
		req:  *HttpRequest.NewRequest(),
	}
	a.req.SetHeaders(headers)
	return &a
}

//getUserData 返回一个用户的被关注数、回答数、文章数
func getUserData(user map[string]interface{}) (int, int, int) {
	return int(user["follower_count"].(float64)), int(user["answer_count"].(float64)), int(user["articles_count"].(float64))
}
