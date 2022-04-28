/**
主持人排序问题
- 问题背景：我们知道通常一个公司项目是由一个团队来完成的，团队成员之间的协作会遵循一定的开发模式，最具代表性的就是敏捷开发，
敏捷开发会规定一系列的团队固定动作，其中就包括【晨会】这一环节，
晨会中会设置一个很重要的角色来维护晨会的秩序我们称之为【主持人】，主持人是每日轮值，我们的问题就基于此。
- 问题描述：如何选择合适的数据结构来存储主持人的轮值顺序并且满足以下条件：
  - 主持人有【请假】(只能请假一天)权限，请假成功后，排在后面的成员成为主持人；
  - 请假者有【撤销请假】权限，撤销成功后恢复默认主持循序；
  - 本次轮值顺序不会影响到下一轮；
- 参考用例
  - ABCDABCD -> A请假、A撤销请假 -> ABCDABCD
  - ABCDABCD -> A请假、B请假 -> CABDABCD
  - ABCDABCD -> A请假、B请假、C请假、B取消请假 -> BACDABCD
 */

/**
DCBA
DB		CA

 */

package main

import "fmt"

func main(){

	s := "ABCDABCD"
	// order := "aA"  // a 请假 A撤销请假
	// order := "ab"  // a 请假 A撤销请假
	order := "abcB"  // a 请假 A撤销请假
	fmt.Println(qingjia(s, order))


}

func qingjia(s, order string ) string {
	left := []uint8{}
	right := []uint8{}

	n := len(s)

	for i:=n-1;i>=0;i--{
		left = append(left, s[i])
	}

	for i := range order{
		if order[i] >= 'a' && order[i] <= 'd'{
			temp := left[len(left)-1]
			left = left[:len(left)-1]
			right = append(right, temp)
		} else if order[i] >= 'A' && order[i] <= 'D' {
			for j := range right{
				if order[i] == right[j]{
					if j == len(right)-1{
						right = right[:j]
					}else{
						right = append(right[:j], right[j+1:]...)
					}
					break
				}
			}
			left = append(left, order[i])
		}
	}

	// fmt.Println(left)
	// fmt.Println(right)

	///
	res := ""
	f := left[len(left)-1]
	left = left[:len(left)-1]
	res += string(f)

	for i:=0;i<len(right);i++{
		res += string(right[i])
	}
	for i := len(left)-1;i>=0;i--{
		res += string(left[i])
	}

	return res
}


