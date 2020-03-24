/*
* @Describe:
* @Author:   mian
* @Date:     2020/3/24 14:11
 */
package tool

import "time"

func singleGet(userC chan<- []interface{}, ans answerInt, anti antiAntiInt, index int) {
	userC <- anti.SafeGet(ans, index)
}

type parrel struct {
}

func (p *parrel) quickGetAll(anti antiAntiInt, ans answerInt) []interface{} {
	num := ans.VotePageNum()
	userC := make(chan []interface{}, 10) // 稍微带一点缓冲，避免拥塞
	for i := 0; i < num; i++ {
		go singleGet(userC, ans, anti, i)
	}
	allUser := []interface{}{}
	receivedNum := 0
	for receivedNum < num {
		timeout := time.After(30 * time.Second) // 三十秒都没有接收到就判断接不到了
		select {
		case <-timeout:
			break // WARNING：并不是很确定可否正常退出
		case r := <-userC:
			receivedNum++
			allUser = append(allUser, r...)
		}
	}
	return allUser
}
