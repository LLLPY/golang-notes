package main

import (
	"fmt"
	"time"
)

func main() {

	//获取当前时间
	now := time.Now()
	fmt.Printf("now: %v\n", now)
	year := now.Year()          //年
	month := now.Month()        //月
	day := now.Day()            //日
	hour := now.Hour()          //时
	minute := now.Minute()      //分
	second := now.Second()      //秒
	nsecond := now.Nanosecond() //纳秒
	fmt.Printf("%v-%v-%v:%v:%v:%v:%v\n", year, month, day, hour, minute, second, nsecond)

	//时间戳
	fmt.Printf("now.Unix(): %v\n", now.Unix())           //秒数
	fmt.Printf("now.UnixMicro(): %v\n", now.UnixMicro()) //毫秒
	fmt.Printf("now.UnixNano(): %v\n", now.UnixNano())   //纳秒

	//时间戳转成时间
	fmt.Printf("time.Unix(now.Unix(), 0): %v\n", time.Unix(now.Unix(), 0))

	//时间的运算
	today := now
	tomorrow := today.Add(time.Hour * 24) //在当前时间上增加24小时
	fmt.Printf("today: %v\n", today)
	fmt.Printf("tomorrow: %v\n", tomorrow)

	dif := today.Sub(tomorrow) //今天与昨天相差的时间
	fmt.Printf("dif: %v\n", dif)

	//比较两个时间是否相同
	fmt.Printf("today.Equal(tomorrow): %v\n", today.Equal(tomorrow))

	//判断当前时间是否在目标时间之前
	fmt.Printf("today.Before(tomorrow): %v\n", today.Before(tomorrow))

	//判断当前时间是否在目标时间之后
	fmt.Printf("today.After(tomorrow): %v\n", today.After(tomorrow))

	//定时器
	c := time.Tick(time.Second) //设置一个间隔一秒的定时器
	for i := range c {
		fmt.Printf("now: %v\n", i)             //每隔一秒打印一下当前时间
		if i.After(now.Add(time.Second * 3)) { //3秒后停止
			break
		}
	}

	//Timer
	t := time.NewTimer(time.Second * 2)
	fmt.Printf("now2: %v\n", time.Now())
	<-t.C
	fmt.Printf("now2 after 2 seconds: %v\n", time.Now()) //两秒后执行到这里

	//时间格式化
	fmt.Printf("now.Format(\"2006-01-02 15:04:05.000 Mon Jan\"): %v\n", now.Format("2006/01/02 15:04:05.000 Mon Jan")) //24小时制
	fmt.Printf("now.Format(\"Mon Jan 2006-01-02 3:4:4 PM\"): %v\n", now.Format("Mon Jan 2006-01-02 3:4:4.000 PM"))     //12小时制

	//解析字符串格式的时间
	loc, _ := time.LoadLocation("Asia/Shanghai")

	//第一个参数指定格式，第二参数为字符串格式的时间，第三个参数指定时区
	t2, _ := time.ParseInLocation("2006/01/02 15:04:05.000 Mon Jan", "2022/11/11 22:44:56.882 Mon Nov", loc)
	fmt.Printf("t2: %T\n", t2)
	fmt.Printf("t2: %v\n", t2)

}
