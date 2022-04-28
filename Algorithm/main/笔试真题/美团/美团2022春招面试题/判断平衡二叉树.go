package main

import "fmt"

import . "awesomeProject1/Algorithm/main/structure/TreeNode"

func main() {
	// a := 0
	// fmt.Scan(&a)
	// fmt.Printf("%d\n", a)
	// fmt.Printf("Hello World!\n")

	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	nums := []int{1, 2, -1, 3, -1, -1, -1, 4, -1}
	root := TreeCreate(nums, 0)
	TreePrint(root)
	fmt.Println()

	res, _ := isBalance(root)
	fmt.Println(res)

}

func isBalance(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	balancel, highl := isBalance(root.Left)
	balancer, highr := isBalance(root.Right)

	high := max(highl, highr) + 1
	if balancel && balancer && abs(highl, highr) <=1 {
		return true, high
	}
	return false, high
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// 原回答
/**
func isBalance(root *TreeNode, high int ) bool {
    if root == nil{
        return true
    }

    if !isBalance(root.Left) ||  !isBalance(root.Right){
        return false
    }

    l := hight(root.Left, high+1)
    r := hight(root.Right, high+1)
    if abs(l, r) > 1{
        return false
    }
    return true
}

func hight(root *TreeNode, high int ) int{
    if root == nil{
        return high
    }

    return max(hight(root.Left, high+1), hight(root.Right, high+1))

}
*/
