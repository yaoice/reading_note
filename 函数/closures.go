/*
闭包：官方解释（译文）Go函数可以是一个闭包。闭包是一个函数值，它引用了函数体之外的变量。
这个函数可以对这个引用的变量进行访问和赋值；换句话说这个函数被"绑定"在这个变量上。

作用：是缩小变量作用域，降低对全局变量污染的概率
*/

package main

import "fmt"

func add() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	//add()就是一个闭包，并赋值给pos和neg，pos的闭包函数和neg的闭包函数被绑定在各自的sum变量上；
	//两个闭包函数的sum变量之间没有任何关系
	pos, neg := add(), add()
	for i := 0; i < 5; i++ {
		fmt.Println(
			pos(i),
			neg(-i),
		)
	}
}

/*
0 0
1 -1
3 -3
6 -6
10 -10
*/

/*
//闭包实现的斐波那契数列
package main

//fibonacci函数完成核心算法、核心数据存储, 不负责for循环
func fibonacci() func() int {
	b1 := 1
	b2 := 0
	bc := 0
	return func() int {
		bc = b1 + b2
		b1 = b2
		b2 = bc
		return bc
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
*/
