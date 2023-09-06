package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math"
	"time"
)

/**
输入输出都呈时间戳表示
*/

type Timestamp struct {
	time.Time
}

func (t Timestamp) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *Timestamp) Scan(value interface{}) error {
	v, ok := value.(time.Time)
	if ok {
		*t = Timestamp{Time: v}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", value)
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	//output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return []byte("null"), nil
	}
	output := fmt.Sprintf("%d", t.Time.UnixMilli())
	return []byte(output), nil
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	if len(data) < 1 || string(data) == "null" {
		return nil
	}
	if data[0] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		//字符串类型的时间格式：顺序使用612345、RFC3339格式类型
		v, err := time.Parse("2006-01-02 15:04:05", s)
		if err != nil {
			v, err = time.Parse(time.RFC3339, s)
			if err != nil {
				return err
			}
		}
		*t = Timestamp{v}
	} else {
		//时间戳类型
		length := float64(len(data)) - 1
		var x float64
		for _, value := range data {
			tmp := math.Pow(10, length)
			x = x + (float64(value)-48)*tmp
			length--
		}
		*t = Timestamp{time.UnixMilli(int64(x))}
	}
	return nil
}
