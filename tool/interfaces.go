/*
* @Describe:
* @Author:   mian
* @Date:     2020/3/19 15:50
 */

package tool

type User struct {
	AnswerCount int `json:"answer_count"`
	FollowCount int `json:"follower_count"`
}

// 答案接口
type answerInt interface {
	VoteNum() int                    // 总赞数
	VotePageNum() int                // 总赞数分为几页
	GetPage(index int) []interface{} // 某一页赞同者的详情
	GetAll() []interface{}           // 所有赞同者的详情
}

// 反反爬虫接口
type antiAntiInt interface {
	SafeGet(ans answerInt, index int) []interface{} // 不触犯反爬虫的get
}

// 并行获取数据
type quickGetInt interface {
	quickGetAll(anti antiAntiInt, ans answerInt) []interface{}
}

// 判断是否为机器人接口
type robotDetectorInt interface {
	IsRobot(u User) bool
}

//日志
type log interface {
	Info(str string)
	Info2(str string) // 二级详细通知
	Debug(str string)
}
