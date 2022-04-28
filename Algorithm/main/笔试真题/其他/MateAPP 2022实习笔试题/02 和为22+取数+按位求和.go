/**
第一步：
一个数组长度小于2，返回该数组
数组存在两个值的和为22，返回这两个值的下标，如果存在多组这样的结果，返回下标最小的一组
如果不存在和为22 的两个值，返回数组的前两个数

第二步:
一个数组长度为0 返回0，长度为1， 返回数组第一个值
数组长度：
>= 32 返回第32个值
>= 32/2 (向下取整) 返回第32/2个值
>= 32/4 (向下取整) 返回第32/4个值
...
以此类推

第三步：
取num的绝对值，num的绝对值小于40，则num= 40
对num中每一位上的数相加，如果和为个数，返回结果，
否则num=和，再次重复上述操作
 */

package main

import "fmt"

func main(){
	nums := []int{883, 214, 22, 597, 56, 531, 112, 693, 350, 316, 944, 179, 41}
	fmt.Println(answer02(nums)) // 7
}

/**
* 请勿修改返回值类型
// 883, 214, 22, 597, 56, 531, 112, 693, 350, 316, 944, 179, 41
// 7
*/
func answer02(x []int) int {
	ans1 := step02_1(x)
	ans2 := step02_2(ans1)
	ans3 := step02_3(ans2)
	return ans3
}

// 第一步
func step02_1(nums []int) []int{
	if len(nums) <2{
		return nums
	}

	for i:=1;i<len(nums);i++{
		for j := 0; j<i ;j++{
			if nums[i] + nums[j] == 22{
				return []int{j, i}
			}
		}
	}
	return nums[:2]
}

// 第二步
func step02_2(nums []int) int{
	n := len(nums)

	if n == 0{
		return 0
	}else if n == 1{
		return nums[0]
	}

	i := 32
	if n >= 32 {
		i = 32 - 1
	}else if n >= 32/2{
		i = 32/2-1
	}else if n >= 32/4{
		i = 32/4-1
	}else if n >= 32/8{
		i = 32/8-1
	}else if n >= 32/16{
		i = 32/16-1
	}else{
		i = 1-1
	}
	return nums[i]
}

// 第三步
func step02_3(n int ) int {
	num  := abs(n)
	if num < 40{
		num = 40
	}

	// 处理新的整数
	return plus(num)
}

func plus(n int ) int {
	res := 0
	for n != 0 {
		num := n % 10
		n /= 10
		res += num
	}
	if res >=0 && res <=9{
		return res
	}
	return plus(res)
}

func abs(a int ) int {
	if a < 0 {
		a = -1 * a
	}
	return a
}