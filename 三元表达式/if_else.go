/*
参考链接：https://chai2010.cn/advanced-go-programming-book/appendix/appendix-b-gems.html
*/

package main

func If(condition bool, trueVal, falseVal interface{}) interface{} {
    if condition {
        return trueVal
    }
    return falseVal
}

func main() {
    a, b := 2, 3
    max := If(a > b, a, b).(int)
    println(max)
}

