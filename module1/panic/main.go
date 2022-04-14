package main

import (
	"fmt"
	"time"
)

// panic会是当前线程直接crash。可在系统出现不可恢复的错误时主动调用panic
// 在panic后面的所有代码都将无法执行，因为这个线程已经没有了。
// 可以通过defer在panic后执行一些逻辑，因为defer在调用panic之前已经压栈。
// defer: 保证执行并把控制权交还给接收到panic的函数调用者
// recover: 函数从panic或错误场景中恢复， 只能在defer修饰的函数中使用

/*如果我希望这个函数从错误场景中恢复， 那你就可以通过一个defer的func，然后在里面进行recover。
如果recover成功，这个线程时不会被panic掉的。有点像其他语言的try catch*/
func main() {
	defer func() {
		fmt.Println("defer func is called")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("recover success")
	}()

	fmt.Println("some logic begin")
	time.Sleep(time.Second)
	fmt.Println("some logic end")

	panic("a panic is triggerd")
}

/* 输出结果
some logic begin
some logic end
defer func is called
a panic is triggerd
recover success
根据结果判断：panic相当于抛出错误，程序并没有马上crash，而是执行defer并往外层抛。
直到抛到最上层才会crash或者遇到了defer中的recover进行恢复操作。
*/
