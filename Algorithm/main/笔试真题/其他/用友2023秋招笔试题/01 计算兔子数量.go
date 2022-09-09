package main
//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 计算第n个月的兔子总共多少对。
 * @param count int整型 第多少个月后兔子总对数
 * @return long长整型
 */

// 100%
func calculateTotal( count int ) int64 {
	// write code here

	var  tu1, tu2, tu3 int64
	 tu1, tu2, tu3 = 1, 0, 0
	for i:=1;i< count;i++ {
		tu3 += tu2
		tu2 =  tu1
		tu1 = tu3
	}
	return  tu1 + tu2 + tu3
}

/**
1 1
2 1
3 2
4 3
5 5

 */