package main

import "fmt"

func main() {
	//----------------and
	// data := "{\"projectId\":\"10001\",\"code\":\"test-push\",\"productCode\":\"cvforce\"}"
	// rule := "{\"and\":{\"projectId\":\"10001\",\"code\":\"test-push\"}}"
	//----------------or
	// data := "{\"projectId\":\"10001\",\"code\":\"test-push\",\"productCode\":\"cvforce\"}"
	// rule := "{\"or\":{\"projectId\":\"10001\",\"code\":\"test-push1111\"}}"
	//----------------not
	// data := "{\"projectId\":\"10001\",\"code\":\"test-push\",\"productCode\":\"cvforce\"}"
	// rule := "{\"not\":{\"projectId\":\"10004\",\"code\":\"test-push1111\"}}"
	//----------------gt
	// data := "{\"age\":15,\"year\":2024}"
	// rule := "{\"gt\":{\"age\":11,\"year\":2023}}"
	//----------------gte
	// data := "{\"age\":15,\"year\":2024}"
	// rule := "{\"gte\":{\"age\":15,\"year\":2023}}"
	//----------------lt
	// data := "{\"age\":15,\"year\":2024}"
	// rule := "{\"lt\":{\"age\":16,\"year\":2025}}"
	//----------------lte
	// data := "{\"age\":15,\"year\":2024}"
	// rule := "{\"lte\":{\"age\":16,\"year\":2024}}"
	//----------------like
	// data := "{\"projectId\":\"10001\",\"code\":\"test-push\",\"name\":\"xiaoming\"}"
	// rule := "{\"like\":{\"name\":\"ao\"}}"
	//----------------and like
	data := "{\"projectId\":\"10001\",\"code\":\"test-push\",\"name\":\"xiaoming\"}"
	rule := "{\"like\":{\"name\":\"ao\"},\"and\":{\"projectId\":\"10001\"}}"
	var d JsonCondition
	b, err := d.JsonCheck(data, rule)
	fmt.Println("check:", b, ", err:", err)
}
