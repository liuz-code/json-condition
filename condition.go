package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	JSON_RULE_AND  = "and"
	JSON_RULE_OR   = "or"
	JSON_RULE_NOT  = "not"
	JSON_RULE_GT   = "gt"
	JSON_RULE_GTE  = "gte"
	JSON_RULE_LT   = "lt"
	JSON_RULE_LTE  = "lte"
	JSON_RULE_LIKE = "like"
)

type JsonCondition struct {
}

func (f *JsonCondition) JsonCheck(data, rule string) (bool, error) {
	var dataMap map[string]interface{}
	var ruleMap map[string]interface{}
	err := json.Unmarshal([]byte(data), &dataMap)
	if err != nil {
		return false, fmt.Errorf("data json parser error. %s", err.Error())
	}
	err = json.Unmarshal([]byte(rule), &ruleMap)
	if err != nil {
		return false, fmt.Errorf("rule json parser error. %s", err.Error())
	}
	return f.Check(dataMap, ruleMap)
}

func (f *JsonCondition) Check(data, rule map[string]interface{}) (bool, error) {
	if len(rule) == 0 {
		fmt.Println("rule map empty. return true")
		return true, nil
	}
	if len(data) == 0 {
		fmt.Println("data empty. return false")
		return false, nil
	}
	for k, v := range rule {
		rv := v.(map[string]interface{})
		var r bool
		var err error
		switch k {
		case JSON_RULE_AND:
			r, err = f.And(data, rv)
			break
		case JSON_RULE_OR:
			r, err = f.Or(data, rv)
			break
		case JSON_RULE_NOT:
			r, err = f.Not(data, rv)
			break
		case JSON_RULE_GT:
			r, err = f.Gt(data, rv)
			break
		case JSON_RULE_GTE:
			r, err = f.Gte(data, rv)
			break
		case JSON_RULE_LT:
			r, err = f.Lt(data, rv)
			break
		case JSON_RULE_LTE:
			r, err = f.Lte(data, rv)
			break
		case JSON_RULE_LIKE:
			r, err = f.Like(data, rv)
			break
		}
		if err != nil {
			fmt.Println(fmt.Sprintf("err: %+v", err.Error()))
		}
		if !r {
			return false, fmt.Errorf("rule type [%s] check fail.", k)
		}
	}
	return true, nil
}

func (f *JsonCondition) And(data, rule map[string]interface{}) (bool, error) {
	for k, v := range rule {
		b := false
		for dk, dv := range data {
			if k == dk {
				if ValueEqCheck(v, dv) {
					b = true
					break
				}
			}
		}
		if !b {
			return false, fmt.Errorf("rule and %s=%s check error.", k, v)
		}
	}
	return true, nil
}
func (f *JsonCondition) Or(data, rule map[string]interface{}) (bool, error) {
	for k, v := range rule {
		for dk, dv := range data {
			if k == dk {
				if ValueEqCheck(v, dv) {
					return true, nil
				}
			}
		}
	}
	return false, fmt.Errorf("rules or %+v check error.", rule)
}
func (f *JsonCondition) Not(data, rule map[string]interface{}) (bool, error) {
	for k, v := range rule {
		b := false
		for dk, dv := range data {
			if k == dk {
				if !ValueEqCheck(v, dv) {
					b = true
					break
				}
			}
		}
		if !b {
			return false, fmt.Errorf("rule not %s=%s check error.", k, v)
		}
	}
	return true, nil
}
func (f *JsonCondition) Gt(data, rule map[string]interface{}) (bool, error) {
	for k, v := range rule {
		b := false
		for dk, dv := range data {
			if k == dk {
				fv, err := ValueToFloat(v)
				if err != nil {
					return false, fmt.Errorf("rule gt %s>%v value type error.", k, v)
				}
				dfv, dErr := ValueToFloat(dv)
				if dErr != nil {
					continue
				}
				if fv < dfv {
					b = true
					break
				}
			}
		}
		if !b {
			return false, fmt.Errorf("rule gt %s>%s check error.", k, v)
		}
	}
	return true, nil
}
func (f *JsonCondition) Gte(data, rule map[string]interface{}) (bool, error) {
	for k, v := range rule {
		b := false
		for dk, dv := range data {
			if k == dk {
				fv, err := ValueToFloat(v)
				if err != nil {
					return false, fmt.Errorf("rule gt %s>=%v value type error.", k, v)
				}
				dfv, dErr := ValueToFloat(dv)
				if dErr != nil {
					continue
				}
				if fv <= dfv {
					b = true
					break
				}
			}
		}
		if !b {
			return false, fmt.Errorf("rule gt %s>=%s check error.", k, v)
		}
	}
	return true, nil
}
func (f *JsonCondition) Lt(data, rule map[string]interface{}) (bool, error) {
	for k, v := range rule {
		b := false
		for dk, dv := range data {
			if k == dk {
				fv, err := ValueToFloat(v)
				if err != nil {
					return false, fmt.Errorf("rule lt %s<%v value type error.", k, v)
				}
				dfv, dErr := ValueToFloat(dv)
				if dErr != nil {
					continue
				}
				if fv > dfv {
					b = true
					break
				}
			}
		}
		if !b {
			return false, fmt.Errorf("rule lt %s<%s check error.", k, v)
		}
	}
	return true, nil
}
func (f *JsonCondition) Lte(data, rule map[string]interface{}) (bool, error) {
	for k, v := range rule {
		b := false
		for dk, dv := range data {
			if k == dk {
				fv, err := ValueToFloat(v)
				if err != nil {
					return false, fmt.Errorf("rule lte %s<=%v value type error.", k, v)
				}
				dfv, dErr := ValueToFloat(dv)
				if dErr != nil {
					continue
				}
				if fv >= dfv {
					b = true
					break
				}
			}
		}
		if !b {
			return false, fmt.Errorf("rule gt %s<=%s check error.", k, v)
		}
	}
	return true, nil
}
func (f *JsonCondition) Like(data, rule map[string]interface{}) (bool, error) {
	for k, v := range rule {
		b := false
		for dk, dv := range data {
			if k == dk {
				fv, err := ValueToString(v)
				if err != nil {
					return false, fmt.Errorf("rule like key:%s, value:%v type error.", k, v)
				}
				dfv, dErr := ValueToString(dv)
				if dErr != nil {
					continue
				}
				if strings.Index(dfv, fv) > -1 {
					b = true
					break
				}
			}
		}
		if !b {
			return false, fmt.Errorf("rule gt %s<=%s check error.", k, v)
		}
	}
	return true, nil
}

func ValueEqCheck(v1, v2 interface{}) bool {
	if v1 == v2 {
		return true
	} else {
		return false
	}
}
func ValueToFloat(v1 interface{}) (float64, error) {
	switch v1.(type) {
	case int:
		// fmt.Println("int")
		return float64(v1.(int)), nil
	case float64:
		// fmt.Println("float64")
		return v1.(float64), nil
	default:
		fmt.Println("default")
		return 0, fmt.Errorf("value type parser error is not float, value:%v", v1)
	}
}
func ValueToString(v1 interface{}) (string, error) {
	switch v1.(type) {
	case string:
		// fmt.Println("string")
		return v1.(string), nil
	default:
		// fmt.Println("default")
		return "", fmt.Errorf("value type parser error is not string, value:%v", v1)
	}
}
