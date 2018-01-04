package base

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/**
 * TODO: 基础扩展
 */
type Base struct{}

/**
 * TODO: Struct => Map[属性][value|tags]值
 * @param s struct
 * @return map[string]map[string]interface{}
 */
func StructToMap(s interface{}) map[string]map[string]interface{} {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	data := make(map[string]map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = make(map[string]interface{})
		data[t.Field(i).Name]["value"] = v.Field(i).Interface()
		data[t.Field(i).Name]["tags"] = make(map[string]string)
		tags := regexp.MustCompile(`\S+`).FindAllString(string(t.Field(i).Tag), -1)
		for _, tag := range tags {
			temp := strings.Split(tag, ":")
			data[t.Field(i).Name]["tags"].(map[string]string)[temp[0]] = strings.Replace(temp[1], `"`, "", -1)
		}
	}
	return data
}

/**
 * TODO: Map => Struct
 * @param m map[string]interface{}
 * @param s 需要转换的struct
 * @return wrong
 */
func MapToStruct(m map[string]interface{}, s interface{}) error {
	var typeConversion = func(value string, ntype string) (reflect.Value, error) {
		if ntype == "string" {
			return reflect.ValueOf(value), nil
		} else if ntype == "time.Time" {
			t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
			return reflect.ValueOf(t), err
		} else if ntype == "Time" {
			t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
			return reflect.ValueOf(t), err
		} else if ntype == "int" {
			i, err := strconv.Atoi(value)
			return reflect.ValueOf(i), err
		} else if ntype == "int8" {
			i, err := strconv.ParseInt(value, 10, 64)
			return reflect.ValueOf(int8(i)), err
		} else if ntype == "int32" {
			i, err := strconv.ParseInt(value, 10, 64)
			return reflect.ValueOf(int64(i)), err
		} else if ntype == "int64" {
			i, err := strconv.ParseInt(value, 10, 64)
			return reflect.ValueOf(i), err
		} else if ntype == "float32" {
			i, err := strconv.ParseFloat(value, 64)
			return reflect.ValueOf(float32(i)), err
		} else if ntype == "float64" {
			i, err := strconv.ParseFloat(value, 64)
			return reflect.ValueOf(i), err
		}

		//else if .......增加其他一些类型的转换
		return reflect.ValueOf(value), fmt.Errorf("未知的类型：" + ntype)
	}
	var setField = func(obj interface{}, name string, value interface{}) error {
		structValue := reflect.ValueOf(obj).Elem()        //结构体属性值
		structFieldValue := structValue.FieldByName(name) //结构体单个属性值

		if !structFieldValue.IsValid() {
			return fmt.Errorf("No such field: %s in obj", name)
		}

		if !structFieldValue.CanSet() {
			return fmt.Errorf("Cannot set %s field value", name)
		}

		structFieldType := structFieldValue.Type() //结构体的类型
		val := reflect.ValueOf(value)              //map值的反射值

		var err error
		if structFieldType != val.Type() {
			val, err = typeConversion(fmt.Sprintf("%v", value), structFieldValue.Type().Name()) //类型转换
			if err != nil {
				return err
			}
		}

		structFieldValue.Set(val)
		return nil
	}
	for k, v := range m {
		err := setField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

/**
 * TODO: 判断字符串是否为空
 * @param s 字符串
 * @return bool
 */
func StringIsNull(s string) bool {
	return len(s) == 0 || s == ""
}

/**
 * TODO: 清除空字符 unicode码
 */
func StringClear(s string) string {
	sBytes := []byte(s)
	rBytes := make([]byte, 0)
	for _, b := range sBytes {
		if b != 0 {
			rBytes = append(rBytes, b)
		}
	}
	return string(rBytes)
}
