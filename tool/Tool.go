package tool

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
)

// Struct2Map ，将struct转为map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// MD5 将一段字符串转换成md5编码
func MD5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

// TraceStack 返回调用者的堆栈信息
func TraceStack(msg interface{}, level int) ([]byte, error) {
	w := new(bytes.Buffer)
	var err error

	if msg != nil {
		_, err = fmt.Fprintln(w, msg)
		if err != nil {
			return nil, err
		}
	}

	ws := func(str string) {
		if err == nil {
			_, err = w.WriteString(str)
		}
	}

	for i := level; true; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		ws(file)
		ws(":")
		ws(strconv.Itoa(line))
		ws("\n")
	}

	if err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}
