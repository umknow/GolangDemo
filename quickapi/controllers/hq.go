/*
	行情api
*/
package controllers

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动包
	"strings"
	"time"

	//"github.com/golang/protobuf/protoc-gen-go/testdata/imports/fmt"

	. "fmt"
	//"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

// Operations about hq info
type HqController struct {
	beego.Controller
}

// @Title Get
// @Description find object by objectid
// @Param code string
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (h *HqController) HqDayPrice() { //
	/*分时行情*/

	body := map[string]interface{}{"responseCode": 200, "responseMessage": "成功",
									"responCode": "000", "responMessage": "成功"}
	db, _ := sql.Open("mysql",
		"root:K5pMwOrXgPlRc4MZ@tcp(rm-bp10t403rcn80qaxmio.mysql.rds.aliyuncs.com:3306)/hq?charset=utf8") //&local_infile=1
	defer db.Close()

	Println("【分时行情】：进入函数")
	// 股票号码 --> string
	stockCode := h.GetString("code")[:6]
	// 指定日期
	today := time.Now().Format("20060102") //  15:04:05
	//today := "20191110"
	dt := h.GetString("dt", today)
	// 股票类型 --> int
	stkTyp, _ := strconv.Atoi(h.GetString("stkTyp"))

	// 整合数据
	query_sql := Sprintf("SELECT date AS dt,volume,riseAndFallRate AS ratio, close AS price, high, 	"+
		"low,last_close, name,amount  FROM hq.minute_%s WHERE code = %s and category = %d;", dt, stockCode, stkTyp)

	rows, err := db.Query(query_sql)
	Println(rows, err)
	if checkErr(err) {
		body["responseCode"] = 400
		body["responseMessage"] = "失败"
	} else {

		type hqBar struct {
			dt string
			volume int
			ratio float32
			price float64
			high float64
			low float64
			last_close float64
			name string
			amount uint64
		}

		//cloumsList := make([]interface{}, 0)
		cloums, _ := rows.Columns()
		Println("cloums >>", cloums)
		cloumsLists := new([]hqBar)

		for rows.Next() {

			var dt string
			var volume int
			var ratio float32
			var price float64
			var high float64
			var low float64
			var last_close float64
			var name string
			var amount uint64
			//rows.Columns()
			err := rows.Scan(&dt, &volume, &ratio, &price, &high, &low, &last_close, &name, &amount)
			checkErr(err)
			//tmpDic := map[string]interface{}{
			//	"dt":         strings.Split(dt, " ")[1],
			//	"voluse":     volume,
			//	"ratio":      ratio,
			//	"price":      price,
			//	"high":       high,
			//	"low":        low,
			//	"last_close": last_close,
			//	"name":       name,
			//	"amount":     amount,
			//}

			tmpSlince := hqBar{
				strings.Split(dt, " ")[1],
				volume,
				ratio,
				price,
				high,
				low,
				last_close,
				name,
				amount,
			}
			Println("tmpSlince >>", tmpSlince)

			// 相同的数据结构，append不到全局变量

			*cloumsLists = append(*cloumsLists, tmpSlince)  // 这个地方有点问题
		}
		Println("cloumsList >>", cloumsLists)

		//df := dataframe.LoadStructs(*cloumsLists)
		//Println(df)

		dic := make(map[string]interface{})
		dic["name"] = "wang"
		dic["age"] = 18
		dic["sex"] = "男"
		dic["type"] = stkTyp
		dic["stckCode"] = stockCode
		dic["today"] = today
		dic["shares"] = *cloumsLists

		body["message"] = dic
	}

	h.Data["json"] = body
	h.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param code string
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (h *HqController) HqFiveDayPrice() { // DayPrice
	/*五日行情*/
	Println("【分时行情】：进入函数")
	// 股票号码 --> string
	stockCode := h.GetString("code")[:6]
	// 股票类型 --> int
	stkTyp, _ := strconv.Atoi(h.GetString("stkTyp"))

	// 整合数据

	dic := make(map[string]interface{})
	dic["name"] = "wang yongsheng"
	dic["age"] = 18
	dic["sex"] = "男"
	dic["type"] = stkTyp
	dic["stckCode"] = stockCode
	h.Data["json"] = dic

	h.ServeJSON()
}

func checkErr(err interface{}) bool {
	if err != nil {
		return true
	} else {
		return false
	}
}
