package main

import (
	"fmt"
)

func doNop() (int, error) {
	return 12, nil
}

var a int

func main() {
	fmt.Println("init a =", a)
	if true {
		/*
     * 编译问题：
		 * 情况 1：不存在 fmt.Println("a =", a)
		 *   此处的 a 为局部变量，并没用使用全局变量，
		 *   如果去掉 ':' 则会因为 err 未定义出错。
		 * 情况 2：存在代码 fmt.Println("a =", a)，或者简单的 _ = a
		 *   这种情况下 a 是在 if 块内的局部变量，并未改变全局变量，需要小心
		 */
		a = 11
		a, err := doNop() 
    /*
     * 此处是重定义 redeclare，go 语言的一个设计问题。
     * 此处的 a 按官方说法是'Redeclaration does not introduce a new variable; it just assigns a new value to the original.'
     * 对 a 的赋值没有体现在外层，需要注意。 ！！！
     */
		if err != nil {
			fmt.Println("error not nil")
		}
		fmt.Println("call func a =", a)
	}
	fmt.Println("final a =", a)
}
