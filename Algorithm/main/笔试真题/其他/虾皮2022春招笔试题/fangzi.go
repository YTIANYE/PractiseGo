/*
n皇后问题
*/
package 虾皮2022春招笔试题

/**
 * Note: 类名、方法名、参数名已经指定，请勿修改
 *
 *
 *
 * @param n int整型
 * @return int整型
 */

// 我的题解：通过全部测试样例
func buildHouses(n int) int {
	// 初始化
	allmap := [][]bool{}
	for i := 0; i < n; i++ {
		temp := []bool{}
		for j := 0; j < n; j++ {
			temp = append(temp, false)
		}
		allmap = append(allmap, temp)
	}

	// 检查（i,j）位置是否可以放置皇后
	check := func(i, j int) bool {
		// 竖着
		for t := 0; t < n; t++ {
			if allmap[t][j] == true && i != t {
				return false
			}
		}

		var p, q int
		p, q = i, j
		for p < n && q < n {
			if allmap[p][q] && p != i && q != j {
				return false
			}
			p++
			q++
		}

		p, q = i, j
		for p < n && q >= 0 {
			if allmap[p][q] && p != i && q != j {
				return false
			}
			p++
			q--
		}

		p, q = i, j
		for p >= 0 && q < n {
			if allmap[p][q] && p != i && q != j {
				return false
			}
			p--
			q++
		}

		p, q = i, j
		for p >= 0 && q >= 0 {
			if allmap[p][q] && p != i && q != j {
				return false
			}
			p--
			q--
		}

		return true
	}

	// 递归
	var res int
	var dfs func(int)
	dfs = func(i int) {
		for j := 0; j < n; j++ {
			if check(i, j) {
				if i == n-1 {
					res++
					continue
				}
				allmap[i][j] = true
				dfs(i + 1)
				allmap[i][j] = false
			}
		}
	}
	dfs(0)
	return res
}
