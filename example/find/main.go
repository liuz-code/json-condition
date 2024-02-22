package main

import (
	"fmt"

	jCondition "github.com/liuz-code/json-condition"
)

func main() {
	jsonText := `{
		"reason":"查询成功!",
		"result":{
			"city":"苏州",
			"realtime":{
				"temperature":"17",
				"humidity":"69",
				"info":"阴",
				"wid":"02",
				"direct":"东风",
				"power":"2级",
				"aqi":"30"
			},
			"future":[
				{
					"date":"2021-10-25",
					"temperature":"12\/21℃",
					"weather":"多云",
					"wid":{
						"day":"01",
						"night":"01"
					},
					"direct":"东风"
				},
				{
					"date":"2021-10-26",
					"temperature":"13\/21℃",
					"weather":"多云",
					"wid":{
						"day":"01",
						"night":"01"
					},
					"direct":"东风转东北风"
				},
				{
					"date":"2021-10-27",
					"temperature":"13\/22℃",
					"weather":"多云",
					"wid":{
						"day":"01",
						"night":"01"
					},
					"direct":"东北风"
				}
			]
		},
		"error_code":0
	}`
	path := "result.future.[0].date"
	// path := "result.realtime.humidity"
	var d jCondition.JsonCondition
	node, err := d.JsonFind(jsonText, path)
	fmt.Println("node:", node, ", error:", err)
}
