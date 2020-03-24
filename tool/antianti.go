/*
* @Describe: 一些反爬虫接口
* @Author:   mian
* @Date:     2020/3/24 14:12
 */
package tool

import (
	"math/rand"
	"time"
)

type antiAntiSpider struct {
}

func (a antiAntiSpider) SafeGet(ans answerInt, index int) []interface{} {
	// 在调用前随机停一会
	rand.Seed(time.Now().UnixNano()) // UnixNano()表示纳秒
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return ans.GetPage(index)
}
