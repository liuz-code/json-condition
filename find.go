package jCondition

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func (f *JsonCondition) JsonFind(jsonStr, find string) (any, error) {
	if !IsJson(jsonStr) {
		return nil, fmt.Errorf("Json格式错误")
	}
	jxList := strings.Split(find, ".")
	jxLen := len(jxList)
	var (
		data  = AnyToMap(jsonStr)
		value any
		err   error
	)
	for i := 0; i < jxLen; i++ {
		l := len(jxList[i])
		if l > 2 && string(jxList[i][0]) == "[" && string(jxList[i][l-1]) == "]" {
			numStr := jxList[i][1 : l-1]
			dataList := AnyToArr(value)
			value = dataList[AnyToInt(numStr)]
			data, err = interfaceToMap(value)
			if err != nil {
				continue
			}
		} else {
			if IsHaveKey(data, jxList[i]) {
				value = data[jxList[i]]
				data, err = interfaceToMap(value)
				if err != nil {
					continue
				}
			} else {
				value = nil
			}
		}
	}
	return value, nil
}
func IsJson(str string) bool {
	var tempMap map[string]any
	err := json.Unmarshal([]byte(str), &tempMap)
	return err == nil
}

// AnyToMap any -> map[string]any
func AnyToMap(data any) map[string]any {
	if v, ok := data.(map[string]any); ok {
		return v
	}
	if reflect.ValueOf(data).Kind() == reflect.String {
		dataMap, err := JsonToMap(data.(string))
		if err == nil {
			return dataMap
		}
	}
	return nil
}

// AnyToString any -> string
func AnyToString(data any) string {
	return StringValue(data)
}

// AnyToArr any -> []any
func AnyToArr(data any) []any {
	if v, ok := data.([]any); ok {
		return v
	}
	return nil
}

// AnyToInt any -> int
func AnyToInt(data any) int {
	var t2 int
	switch reflect.TypeOf(data).Kind() {
	case reflect.Uint:
		t2 = int(data.(uint))
	case reflect.Int8:
		t2 = int(data.(int8))
	case reflect.Uint8:
		t2 = int(data.(uint8))
	case reflect.Int16:
		t2 = int(data.(int16))
	case reflect.Uint16:
		t2 = int(data.(uint16))
	case reflect.Int32:
		t2 = int(data.(int32))
	case reflect.Uint32:
		t2 = int(data.(uint32))
	case reflect.Int64:
		t2 = int(data.(int64))
	case reflect.Uint64:
		t2 = int(data.(uint64))
	case reflect.Float32:
		t2 = int(data.(float32))
	case reflect.Float64:
		t2 = int(data.(float64))
	case reflect.String:
		t2, _ = strconv.Atoi(data.(string))
	default:
		t2 = data.(int)
	}
	return t2
}

// StringValue 任何类型返回值字符串形式
func StringValue(i any) string {
	if i == nil {
		return ""
	}
	if reflect.ValueOf(i).Kind() == reflect.String {
		return i.(string)
	}
	var buf bytes.Buffer
	stringValue(reflect.ValueOf(i), 0, &buf)
	return buf.String()
}

func stringValue(v reflect.Value, indent int, buf *bytes.Buffer) {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Struct:
		buf.WriteString("{\n")
		for i := 0; i < v.Type().NumField(); i++ {
			ft := v.Type().Field(i)
			fv := v.Field(i)
			if ft.Name[0:1] == strings.ToLower(ft.Name[0:1]) {
				continue
			}
			if (fv.Kind() == reflect.Ptr || fv.Kind() == reflect.Slice) && fv.IsNil() {
				continue
			}
			buf.WriteString(strings.Repeat(" ", indent+2))
			buf.WriteString(ft.Name + ": ")
			if tag := ft.Tag.Get("sensitive"); tag == "true" {
				buf.WriteString("<sensitive>")
			} else {
				stringValue(fv, indent+2, buf)
			}
			buf.WriteString(",\n")
		}
		buf.WriteString("\n" + strings.Repeat(" ", indent) + "}")

	case reflect.Slice:
		nl, id, id2 := "", "", ""
		if v.Len() > 3 {
			nl, id, id2 = "\n", strings.Repeat(" ", indent), strings.Repeat(" ", indent+2)
		}
		buf.WriteString("[" + nl)
		for i := 0; i < v.Len(); i++ {
			buf.WriteString(id2)
			stringValue(v.Index(i), indent+2, buf)

			if i < v.Len()-1 {
				buf.WriteString("," + nl)
			}
		}
		buf.WriteString(nl + id + "]")

	case reflect.Map:
		buf.WriteString("{\n")
		for i, k := range v.MapKeys() {
			buf.WriteString(strings.Repeat(" ", indent+2))
			buf.WriteString(k.String() + ": ")
			stringValue(v.MapIndex(k), indent+2, buf)
			if i < v.Len()-1 {
				buf.WriteString(",\n")
			}
		}
		buf.WriteString("\n" + strings.Repeat(" ", indent) + "}")

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		buf.WriteString(strconv.FormatInt(v.Int(), 10))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		buf.WriteString(strconv.FormatUint(v.Uint(), 10))

	case reflect.Float32, reflect.Float64:
		result := fmt.Sprintf("%f", v.Float())
		// 去除result末尾的0
		// for strings.HasSuffix(result, "0") {
		result = strings.TrimSuffix(result, "0")
		// }
		// for strings.HasSuffix(result, ".") {
		result = strings.TrimSuffix(result, ".")
		// }
		buf.WriteString(result)

	default:
		format := "%v"
		switch v.Interface().(type) {
		case string:
			format = "%q"
		}
		_, _ = fmt.Fprintf(buf, format, v.Interface())
	}
}

// IsHaveKey map[string]any 是否存在 输入的key
func IsHaveKey[T comparable](data map[T]any, key T) bool {
	_, ok := data[key]
	return ok
}

// interfaceToMap any -> map[string]any
func interfaceToMap(data any) (map[string]any, error) {
	if v, ok := data.(map[string]any); ok {
		return v, nil
	}
	if reflect.ValueOf(data).Kind() == reflect.String {
		return JsonToMap(data.(string))
	}
	return nil, fmt.Errorf("not map type")
}

// JsonToMap json -> map
func JsonToMap(str string) (map[string]any, error) {
	var tempMap map[string]any
	err := json.Unmarshal([]byte(str), &tempMap)
	return tempMap, err
}
