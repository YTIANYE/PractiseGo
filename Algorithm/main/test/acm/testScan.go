package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	strs := strings.Split(input.Text(), " ")
	fmt.Println("结果：", strs)
}
