package main

import (
	"fmt"
	. "time"
)

func main()  {
	now := Now()
	duration := Second // 一秒
	fmt.Println(duration)
	// UTC 为国际时间， local为当地时间
	date := Date(2018, 10,29,23,30,10,0, Local)
	fmt.Println(now.Local())
	fmt.Println("本地时间：", date, "国际统一时间:", date.UTC())
	// 注意的是： 模板的的输入格式必须为一致，见下layout
	t, _ := Parse("2006-01-02 15:04:05", "2019-11-12 23:23:23")
	fmt.Println(t.Local())
	str_now := now.Format("2006-01-02 15:04:05")
	fmt.Printf("format >>%v, T>%T\n", str_now, str_now)

	str_now_s := now.String()
	str_now_u := now.Unix()  // 时间戳s
	str_now_uN := now.UnixNano()  // 精确到纳秒级别的 时间戳
	fmt.Println(str_now_s, "oo", str_now_u, str_now_uN)

	// 解析一个时间段字符串
	d, _ := ParseDuration("1h30m")

	// 推演时间
	fmt.Println("推演时间, *后：", now.Add(d))
	fmt.Println("一年一个月一天之后>>", now.AddDate(1,1, 1))

	Sleep(2*duration)  // 延时

	fmt.Println("耗时：", -now.Sub(Now()))

	// ...
}
