/**
0-999共1000个结点均匀分布20个虚拟机（0 50 ... 950）
添加虚拟机按照顺序25， 525， 75， 575这样添加，最多添加20个
token计算方式为token每个字符的ASCII和mod999后得到的数，后面的虚拟机上，虚拟机宕机则继续顺延
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// (1)环装链表 初始化
	n := 40
	m := 3 // 节点标号； name； 是否宕机（"false", "true"）；
	servers := make([][]string, n)
	for i := 0; i < n; i++ {
		server := make([]string, m)
		server[0] = strconv.Itoa(i * 25)
		num := i + 1
		if num < 10 {
			server[1] = "redis_0" + strconv.Itoa(num)
		} else {
			server[1] = "redis_" + strconv.Itoa(num)
		}
		server[2] = "false" // 没有宕机
		servers[i] = server
	}

	// （2）计算token数值
	var compToken func(string) int
	compToken = func(token string) int {
		res := 0
		for i := range token {
			res = (res + int(token[i])) % 999
		}
		return res
	}

	// 命令1：返回虚拟机的位置
	var findRedis func(string) int
	findRedis = func(s string) int {
		res := -1
		for i := 0; i < n; i++ {
			if s == servers[i][1] {
				res, _ = strconv.Atoi(servers[i][0])
				break
			}
		}
		return res
	}

	// 命令2：返回token命中的虚拟机
	var setToken func(string) int
	setToken = func(s string) int {
		num := strconv.Itoa(compToken(s))
		res := 0
		if num > servers[n-1][0] {
			return res
		}
		for i := 0; i < n-1; i++ {
			server := servers[i]
			if len(server) == 0 {
				continue
			}
			if server[0] == num {
				res, _ = strconv.Atoi(server[0])
				return res
			} else if server[0] < num && num > servers[i+1][0] {
				res, _ = strconv.Atoi(servers[i+1][0])
				return res
			}
		}
		return res
	}

	// 命令3：存在宕机的虚拟机，返回token命中的虚拟机
	var setToken3 func([]string, string) int
	setToken3 = func(redises []string, token string) int {
		// 添加宕机
		sort.Strings(redises)
		for _, name := range redises {
			for _, server := range servers {
				if name == server[1] {
					server[2] = "true"
				}
			}
		}
		// 计算token位置
		num := strconv.Itoa(compToken(token))
		res := 0
		if num > servers[n-1][0] {
			for i := 0; i < n; i++ {
				server := servers[i]
				if len(server) != 0 && server[2] == "false" {
					return i
				}
			}
		}
		for i := 0; i < n-1; i++ {
			server := servers[i]
			if len(server) == 0 || server[2] == "true" {
				continue
			}
			if server[0] == num {
				res, _ = strconv.Atoi(server[0])
				return res
			} else if server[0] < num && num > servers[i+1][0] {
				res, _ = strconv.Atoi(servers[i+1][0])
				return res
			}
		}
		return res
	}

	// 命令4：添加虚拟机，返回虚拟机的节点标号
	var addRedis func(string) int
	addRedis = func(redis string) int {
		strs := strings.Split(redis, "_")
		num, _ := strconv.Atoi(strs[3])
		var res int
		if num%2 == 1 {
			res = (num/2)*50 + 25
		} else {
			res = (num/2 - 1)*50 + 525
		}
		return res
	}

	// 命令5：添加几个虚拟机，返回token命中的虚拟机的结点标号
	var addRedis5 func(string, string) int
	addRedis5 = func(redis string, token string) int {
		strs := strings.Split(redis, "_")
		num, _ := strconv.Atoi(strs[3])
		for i := 1; i <= num; i++ {

			var jiedian, index int
			if num%2 == 1 {
				jiedian = (num/2)*50 + 25
				index = (num/2) * 2 + 1
			} else {
				jiedian = (num/2-1)*50 + 525
				index = (num/2-1) * 2 + 21
			}
			var name string
			if i < 10 {
				name =strs[0] + "_" + strs[1] + "_0" + strconv.Itoa(i)
			}else {
				name =strs[0] + "_" + strs[1] + "_" + strconv.Itoa(i)
			}


			server := []string{strconv.Itoa(jiedian),name, "false"}
			servers[index] = server
		}
		//
		return setToken(token)
	}

	//  输入命令
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	strs := strings.Split(input.Text(), ":")
	cmd, _ := strconv.Atoi(strs[0])
	content := strs[1]

	if cmd == 1 {
		redisName := content
		res := findRedis(redisName)
		fmt.Println(res)
	} else if cmd == 2 {
		token := content
		res := setToken(token)
		fmt.Println(res)
	} else if cmd == 3 {
		s := strings.Split(content, ";")
		redises := strings.Split(s[0], ",")
		token := s[1]
		res := setToken3(redises, token)
		fmt.Println(res)
	} else if cmd == 4 {
		redis := content
		res := addRedis(redis)
		fmt.Println(res)
	} else if cmd == 5 {
		s := strings.Split(content, ";")
		redis := s[0]
		token := s[1]
		res := addRedis5(redis, token)
		fmt.Println(res)
	}
}
