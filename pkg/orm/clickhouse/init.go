package clickhouse

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/QinL233/go-plus/pkg/yaml"
	"log"
	"net"
	"time"
)

var conn driver.Conn

func Init() {
	config := yaml.Config.Orm.Clickhouse
	if config.Host == "" || config.Port == 0 || config.Database == "" || config.Username == "" || config.Password == "" {
		log.Println("clickhouse config is empty !")
		return
	}
	db, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", config.Host, config.Port)},
		Auth: clickhouse.Auth{
			Database: config.Database,
			Username: config.Username,
			Password: config.Password,
		},
		DialContext: func(ctx context.Context, addr string) (net.Conn, error) {
			var d net.Dialer
			return d.DialContext(ctx, "tcp", addr)
		},
		Debug: true,
		Debugf: func(format string, v ...any) {
			fmt.Printf(format, v)
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 500,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		DialTimeout:          time.Second * 30,
		MaxOpenConns:         50,
		MaxIdleConns:         50,
		ConnMaxLifetime:      time.Duration(10) * time.Minute,
		ConnOpenStrategy:     clickhouse.ConnOpenInOrder,
		BlockBufferSize:      10,
		MaxCompressionBuffer: 10240,
		ClientInfo: clickhouse.ClientInfo{ // optional, please see Client info section in the README.md
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "my-app", Version: "0.1"},
			},
		},
	})
	if err != nil {
		log.Fatalf("ping clickhouse err: %v", err)
	}
	if err = db.Ping(context.Background()); err != nil {
		log.Fatalf("ping clickhouse err: %v", err)
	}
	conn = db
}

func Driver() driver.Conn {
	return conn
}
