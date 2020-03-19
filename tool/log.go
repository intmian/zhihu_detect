/*
* @Describe: 日志
* @Author:   mian
* @Date:     2020/3/19 18:13
 */
package tool

import (
	"fmt"
	"time"
)

//str2info 将字符串转换为标准info log
func str2info(s string) string {
	fs := `[info][%s]%s`
	return fmt.Sprintf(fs, time.Now().Format("15:04:05"), s)
}

//Log2 显示所有log
type Log2 struct {
}

func (l Log2) Info(str string) {
	println(str2info(str))
}

func (l Log2) Info2(str string) {
	println(str2info(str))
}

func (l Log2) Debug(str string) {
	// do nothing
}

//Log1 仅显示一级log
type Log1 struct {
}

func (l Log1) Info(str string) {
	println(str2info(str))
}

func (l Log1) Info2(str string) {
	// do nothing
}

func (l Log1) Debug(str string) {
	// do nothing
}

var LOG1 = Log1{}
var LOG2 = Log2{}
