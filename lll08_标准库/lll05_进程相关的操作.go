package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var mg sync.WaitGroup
var timer *time.Timer

func demo() {
	for {
		timer = time.NewTimer(time.Second * 2)

		//获取当前进程的id
		cur_pid := os.Getpid()
		fmt.Printf("cur_pid: %v\n", cur_pid)
		//获取父进程的id
		p_pid := os.Getppid()
		fmt.Printf("p_pid: %v\n", p_pid)
		<-timer.C
	}

}

func main() {

	// demo()

	//启动一个进程
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}, //指定文件活动对象，标准的输入，输出和错误输出
		Env:   os.Environ(),                               //环境变量
	}

	p, _ := os.StartProcess("/usr/bin/code", []string{"/usr/bin/code"}, attr) //启动一个进程

	fmt.Printf("p.Pid: %v\n", p.Pid)

	//通过进程id查找进程
	p2, _ := os.FindProcess(p.Pid)
	fmt.Printf("p2: %v\n", p2)

	//向进程发送信号
	time.AfterFunc(time.Second*3, func() {
		// p2.Signal(os.Kill) //发送杀死进程的信号
		p2.Kill()
	})

	//等待进程退出，返回进程状态
	ps, _ := p.Wait()
	fmt.Printf("ps: %v\n", ps.String())
}
