package main

import (
	. "fmt"
	"go-common/app/common/openplatform/random"
	. "math"
	"math/rand"
	"time"
)

func main()  {
	Println("【判断nan】", IsNaN(12.43))
	Println("【取绝对值】", Abs(-12.43))

	Println("【余数进一】", Ceil(-12.43))
	Println("【余数进一】", Ceil(12.43))
	Println("【余数舍一】", Floor(12.43))
	Println("【余数舍一】", Floor(-12.93))

	Println("【舍去余数】", Trunc(12.43))
	Println("【舍去余数】", Trunc(-12.43))

	Println("【比较两个，取最大】", Max(1,2))
	Println("【比较两个，取最小】", Min(1,2))

	Println("【减后与0对比取大】", Dim(-1, -5))
	Println("【减后与0对比取大】", Dim(-1, 19))

	Println("【取模】", Mod(9, 4))
	Println("【平方根】", Sqrt(9))
	Println("【立方根】", Cbrt(3))

	Println("【勾股定理取斜边长】", Hypot(3, 4))
	Println("【次方】", Pow(2,8))
	Println("【对数】", Log(1)) // 1*10^0 == 1
	Println("【对数】", Log2(16)) // 2的4次方 == 16

	Println("【随机uid】random.Uniqid", random.Uniqid(10)) // 字节长度 只支持10~19位
	Println("【随机数】rand.Float", rand.Float64())

	Println("【随机数】rand.Int", rand.Int())
	Println("【随机数】rand.Intn", rand.Intn(5))

	// 获取随机数简单形式
	rand.Seed(time.Now().UnixNano())  // 触发种子 Seed
	Println("【0-10之间的随机数】", rand.Intn(10))
	Println("float", rand.Float64())
	// 5-11之间的随机数
	num := rand.Intn(7)+5
	Println("【5-11之间的随机数】", num)

}
