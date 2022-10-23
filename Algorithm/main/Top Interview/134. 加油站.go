package main

// 我实现的官方题解
// 我们首先检查第 00 个加油站，并试图判断能否环绕一周；如果不能，就从第一个无法到达的加油站开始继续检查。
/**
执行用时：64 ms, 在所有 Go 提交中击败了83.93%的用户
内存消耗：8.8 MB, 在所有 Go 提交中击败了62.37%的用户
 */
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	i := 0
	for i<n {
		res := 0
		cnt := 0
		for cnt < n {
			j := (i+cnt) % n
			res += gas[j] - cost[j]
			if res < 0 {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		}
		i += cnt + 1
	}
	return -1
}

// 缩短数组长度：我的题解
/**
执行用时：112 ms, 在所有 Go 提交中击败了5.23%的用户
内存消耗：11.3 MB, 在所有 Go 提交中击败了5.10%的用户
*/
func canCompleteCircuit3(gas []int, cost []int) int {
	n := len(gas)
	//
	nums := [][]int{} // 汽油量、index
	for i:=0;i<n;i++{
		num := gas[i] - cost[i]
		if len(nums) > 0 && nums[len(nums)-1][0] >= 0 && num >=0 {
			nums[len(nums)-1][0] += num
		}else {
			nums = append(nums, []int{num, i})
		}
	}
	//
	var canComplete func(index int) bool
	canComplete = func(index int) bool {
		res := nums[index][0]
		if res < 0 {
			return false
		}
		j:= (index+1) % len(nums)
		res += nums[j][0]
		for j != index {
			if res < 0 {
				return false
			}
			j = (j+1) % len(nums)
			res += nums[j][0]
		}
		return true
	}
	//
	for i := range nums {
		if canComplete(i) {
			return nums[i][1]
		}
	}
	return -1
}


// 优化后的暴力方法——超时
func canCompleteCircuit2(gas []int, cost []int) int {
	n := len(gas)
	//
	nums := make([]int, n)
	for i:= range nums {
		nums[i] = gas[i] - cost[i]
	}
	//
	var canComplete func(index int) bool
	canComplete = func(index int) bool {
		res := nums[index]
		if res < 0 {
			return false
		}
		j:= (index+1) % n
		res += nums[j]
		for j != index {
			if res < 0 {
				return false
			}
			j = (j+1) % n
			res += nums[j]
		}
		return true
	}
	//
	for i := range nums {
		if canComplete(i) {
			return i
		}
	}
	return -1
}



// 模拟超时
func canCompleteCircuit1(gas []int, cost []int) int {
	n := len(gas)
	//
	var canComplete func(index int) bool
	canComplete = func(index int) bool {

		qiyou := gas[index] - cost[index]
		j := (index+1) % n
		for qiyou >=0  {
			if j == index {
				return true
			}
			qiyou += gas[j] - cost[j]
			j = (j + 1) % n
		}
		return false
	}
	//
	for i :=  range gas {
		if canComplete(i) {
			return  i
		}
	}
	return -1
}
