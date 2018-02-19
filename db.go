package main

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

var (
	Db *sql.DB
)

func InitDb() error {
	var err error
	Db, err = sql.Open(CfgString("db_type"), CfgString("db_conn"))
	if err != nil {
		return err
	}

	if CfgHasKey("db_max_open_conns") {
		Db.SetMaxOpenConns(CfgInt("db_max_open_conns"))
	}

	if CfgHasKey("db_max_idle_conns") {
		Db.SetMaxIdleConns(CfgInt("db_max_idle_conns"))
	}

	if CfgHasKey("db_conn_max_lifetime") {
		Db.SetConnMaxLifetime(time.Duration(CfgInt("db_conn_max_lifetime")) * time.Second)
	}

	if err := Db.Ping(); err != nil {
		return err
	}

	return nil
}
