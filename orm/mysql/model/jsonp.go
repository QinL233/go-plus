package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

/**
使用泛型定义通用的json数据结构
*/

type JSONP[T any] struct {
	Data T
}

func (j JSONP[T]) Value() (driver.Value, error) {
	if objectIsBlank(j) {
		return "", nil
	}
	if data, err := json.Marshal(j.Data); err != nil {
		return nil, err
	} else {
		return string(data), nil
	}

	return nil, nil
}

func (j *JSONP[T]) Scan(value interface{}) error {
	if value == nil {
		*j = JSONP[T]{}
		return nil
	}
	u, ok := value.([]uint8)
	if !ok {
		return errors.New("读取db数据结构失败")
	}
	var result T
	err := json.Unmarshal(u, &result)
	if err != nil {
		return err
	}
	*j = JSONP[T]{result}
	return nil
}

func (j JSONP[T]) MarshalJSON() ([]byte, error) {
	if objectIsBlank(j) {
		return []byte("null"), nil
	}
	bytes, err := json.Marshal(j.Data)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (j *JSONP[T]) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("null point exception")
	}
	if bytesIsBlank(data) {
		return nil
	}
	var result T
	err := json.Unmarshal(data, &result)
	if err != nil {
		return err
	}
	*j = JSONP[T]{result}
	return nil
}

//反射判断对象是否为空
func objectIsBlank(data any) bool {
	if data == nil {
		return true
	}
	rs := fmt.Sprintf("%v", data)
	//清空无用得0和false
	if len(rs) > 4 {
		rs = strings.ReplaceAll(rs, "false", "")
		rs = strings.ReplaceAll(rs, "0", "")
		rs = strings.ReplaceAll(rs, " ", "")
	}
	switch rs {
	case "{{}}":
		return true
	case "{[]}":
		return true
	case "{}":
		return true
	case "[]":
		return true
	case "null":
		return true
	case "":
		return true
	default:
		break
	}
	return false
}

//实现判断byte是否为空
func bytesIsBlank(data []byte) bool {
	if data == nil {
		return true
	}
	rs := string(data)
	//清空无用得0和false
	if len(rs) > 4 {
		rs = strings.ReplaceAll(rs, "false", "")
		rs = strings.ReplaceAll(rs, "0", "")
		rs = strings.ReplaceAll(rs, " ", "")
	}
	switch rs {
	case "{{}}":
		return true
	case "{[]}":
		return true
	case "{}":
		return true
	case "[]":
		return true
	case "null":
		return true
	case "":
		return true
	case "\"\"":
		return true
	default:
		break
	}
	return false
}
