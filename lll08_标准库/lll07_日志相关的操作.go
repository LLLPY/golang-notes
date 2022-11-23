package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//配置日志的输出前缀
	log.SetPrefix("Log:")

	//配置日志
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	//设置日志输出到文件
	f, _ := os.OpenFile("a.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	defer f.Close()
	log.SetOutput(f)

	//简单打印日志
	log.Print("hello golang...")
	defer fmt.Print("程序结束...")
	//panic日志
	// log.Panic("bye...") //defer会被执行

	//fatal日志
	// log.Fatal("致命错误...") //defer不会被执行

	//自定义logger
	my_logger := log.New(os.Stdout, "My Log:", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
	my_logger.Print("这个是自己定义的logger，一次性配置所有，会方便许多!")

}
