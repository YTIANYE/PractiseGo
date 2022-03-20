/**
需求：
计算一个数字的各个位数之和，例如数字123，结果为1+2+3=6
随机生成数字进行计算
 */

package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	job *Job
	sum int
}

func main() {

	//创建协程池
	jobChan := make(chan *Job, 8)
	resultChan := make(chan *Result, 8)
	creatPool(20, jobChan, resultChan)

	// 输出结果
	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}

	}(resultChan)

	// 传入 job
	var id int
	for {
		id++
		r_num := rand.Int()
		job := &Job{
			Id:id,
			RandNum: r_num,
		}
		jobChan <- job
	}
}

// 创建协程池
func creatPool(num int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < num; i++ {
		var compute func(chan *Job, chan *Result)
		compute = func(jobChan chan *Job, resultChan chan *Result) {
			for job := range jobChan {
				r_num := job.RandNum
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				res := &Result{
					job: job,
					sum: sum,
				}
				resultChan <- res
			}
		}
		go compute(jobChan, resultChan)
	}
}
