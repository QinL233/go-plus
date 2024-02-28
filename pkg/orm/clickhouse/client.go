package clickhouse

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func Count(db driver.Conn, sql string, args ...any) int {
	var total uint64
	row, err := db.Query(context.Background(), sql, args...)
	if err != nil {
		panic(fmt.Errorf("clickhouse count err:%v", err))
	}
	if row.Next() {
		if err = row.Scan(&total); err != nil {
			panic(fmt.Errorf("search total err2:%v", err))
		}
	}
	return int(total)
}

func One[T any](db driver.Conn, sql string, args ...any) T {
	var one T
	row, err := db.Query(context.Background(), sql, args...)
	if err != nil {
		panic(fmt.Errorf("clickhouse count err:%v", err))
	}
	if row.Next() {
		if err = row.ScanStruct(&one); err != nil {
			panic(fmt.Errorf("search total err2:%v", err))
		}
	}
	return one
}

func List[T any](db driver.Conn, sql string, args ...any) []T {
	var list []T
	rows, err := db.Query(context.Background(), sql, args...)
	if err != nil {
		panic(fmt.Errorf("clickhouse list err:%v", err))
	}
	for rows.Next() {
		var one T
		if err = rows.ScanStruct(&one); err != nil {
			panic(fmt.Errorf("clickhouse list err2:%v", err))
		}
		list = append(list, one)
	}
	return list
}
