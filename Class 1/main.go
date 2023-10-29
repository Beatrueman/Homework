package main

import (
	"awesomeProject/lv1"
	"awesomeProject/lv2"
	"awesomeProject/lv3"
	"awesomeProject/lvx"
	"fmt"
)

func main() {
	fmt.Println(lv1.Add(1, 2))
	fmt.Println(lv2.Squ(2))
	fmt.Println(lv3.Check(9))
	c := lvx.Find()
	fmt.Println("使用rand生成的随机数是:", lvx.Num)
	fmt.Println("这个数是:", c)
}
