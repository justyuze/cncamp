package main

import (
	"errors"
	"fmt"
)

/* 定义一个错误，有两种方式
*	1. errors.New("error string")
*	2. fmt.Errorf("error string %s", arg)，底层同样是使用的errors.New()
 */
//错误信息最好不要以大写开头，因为通常错误信息都会有前缀 -> Error:
var errorNotFound error = errors.New("not found error")

var err2 error = fmt.Errorf("param %s error", "name")

func main() {
	fmt.Printf("Error: %v \n", errorNotFound)
	fmt.Printf("Error: %v \n", err2)

}
