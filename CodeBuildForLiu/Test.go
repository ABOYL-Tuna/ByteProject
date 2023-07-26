package main

import (
	"fmt"
	"math"
)

func main() {
	var a = "initial"
	var b, c int = 1, 2 //可同时多定义
	var d = true        //双斜杠进行注释  无需;号进行隔开 var标志变量
	var e float64       //只定义类型
	f := float32(e)  // 大部分情况下与 var f=float(32)是等效的
	g := a + "foo"   // 字符串可以直接加减
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(g)

	const s string = "constant"  // 常量用const修饰
	const h = 500000
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))  //方法的使用 包.方法
}
