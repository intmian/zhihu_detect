/*
* @Describe:
* @Author:   mian
* @Date:     2020/3/24 13:43
 */

package main

import "zhihu_detect/tool"
func main() {
	a := tool.NewAnswer("1099687485")
	println(a.GetPage(0))
}
