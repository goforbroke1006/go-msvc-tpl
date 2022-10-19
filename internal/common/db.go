package common

import (
	"crypto/tls"
	"database/sql"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/spf13/viper"
)

func OpenClickhouseSomeDatabaseConn() (*sql.DB, error) {
	var (
		addresses = viper.GetStringSlice("db.ch.some_database.addresses")
		database  = viper.GetString("db.ch.some_database.database")
		username  = viper.GetString("db.ch.some_database.username")
		password  = viper.GetString("db.ch.some_database.password")
	)

	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: addresses,
		Auth: clickhouse.Auth{
			Database: database,
			Username: username,
			Password: password,
		},
		TLS: &tls.Config{
			InsecureSkipVerify: true,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		Compression: &clickhouse.Compression{},
		Debug:       true,
	})
	conn.SetMaxIdleConns(5)
	conn.SetMaxOpenConns(10)
	conn.SetConnMaxLifetime(time.Hour)

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
