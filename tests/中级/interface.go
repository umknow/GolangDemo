/*
	Go语言的多态 以 接口 的帮助实现
*/

package main

import . "fmt"


func main()  {
	p1 := FixedBilling{"项目一", 5000}
	p2 := FixedBilling{"项目二", 10000}
	p3 := TimeAnaMaterial{"项目三", 100, 40}
	p4 := TimeAnaMaterial{"项目四", 80, 50}
	p5 := Advertisement{"项目五", 10000, 0.32}
	p6 := Advertisement{"项目六", 808889, 0.2}
	ic := []Income{p1, p2, p3, p4, p5, p6}
	Println("公司总收入 >>", caculateNetIncome(ic))
	}

// 计算净收入
func caculateNetIncome(ic []Income) float64 {
	netincome := 0.0
	for _,income := range ic{
		Printf("收入来源:%s, 收入金额:%.2f\n", income.source(), income.caculate())
		netincome += income.caculate()
	}
	return netincome
}


// 定义接口
type Income interface {  // 收入
	caculate() float64
	source() string
}

// 固定账单项目
type FixedBilling struct {  // 固定收入
	projectName string  // 工程项目
	biddedAmount float64  // 项目招标总额
}

// 定时生产项目（定时和材料项目）
type TimeAnaMaterial struct {
	projectName string
	workHours float64  // 工作时长
	hourlyRate float64  // 每小时工资率
}

// 新增广告点击获取的收入来源
type Advertisement struct {
	adName string
	clickCount int
	incomePerclick float64
}

// 固定收入项目
func (f FixedBilling)caculate() float64 {
	return f.biddedAmount
}

func (f FixedBilling)source() string {
	return f.projectName
}

// 定时收入项目
func (t TimeAnaMaterial)caculate() float64 {
	return t.workHours*t.hourlyRate
}

func (t TimeAnaMaterial)source() string {
return t.projectName
}

// 广告收入项目
func (a Advertisement)caculate() float64 {
	return float64(a.clickCount)*a.incomePerclick
}

func (a Advertisement)source() string {
	return a.adName
}
