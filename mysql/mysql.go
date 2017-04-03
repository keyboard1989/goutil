package mysql

import (
	_ "github.com/go-sql-driver/mysql" // register mysql driver in database/sql
	"github.com/jmoiron/sqlx"
)

type Mysql struct {
	mysql *sqlx.DB
}

func NewMysql(dataSource string) Mysql {
	mysqlDB := sqlx.MustConnect("mysql", dataSource).Unsafe()
	mysqlDB.SetMaxIdleConns(1)
	mysqlDB.SetMaxOpenConns(1)

	return Mysql{mysql: mysqlDB}
}

func (m *Mysql) Query(model interface{}, queryStr string) error {
	err := m.mysql.Select(model, queryStr)
	return err
}

func (m *Mysql) Delete() {

}

func (m *Mysql) Add() {

}

func (m *Mysql) Close() {
}
