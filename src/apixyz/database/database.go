package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	dbSKS *dbMysql
)

type dbMysql struct {
	dbMs *sql.DB
}

func InitAllConnection() error {

	err := InitConnectionSKS()
	if err != nil {
		return err
	}
	return nil
}

func InitConnectionSKS() error {
	dbSKS = new(dbMysql)
	server := viper.GetString("database.server")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	schema := viper.GetString("database.schema")

	return initConnection(dbSKS, server, user, password, schema, "")
}

func initConnection(dbCon *dbMysql, server, user, pass, scheme, appIntent string) error {
	var db *sql.DB
	db, err := createConnectionMs(server, user, pass, scheme, appIntent)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(10)
	dbCon.dbMs = db
	return nil
}

func createConnectionMs(server, user, pass, scheme, appIntent string) (*sql.DB, error) {
	confinsInfo := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;%s", server, user, pass, scheme, appIntent)
	conn, errdb := sql.Open("mssql", confinsInfo)
	if errdb != nil {
		return nil, errdb
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	return conn, nil
}

func (dms *dbMysql) GetConnection() (*sql.DB, error) {
	if err := dms.dbMs.Ping(); err != nil {
		return nil, err
	}
	return dms.dbMs, nil
}

func GetConnectionSKS() (*sql.DB, error) {
	return dbSKS.GetConnection()
}

func CloseAllConnection() {
	CloseConnectionSKS()
}

func CloseConnectionSKS() {
	logrus.Println("Closing AMS connection...")
	dbSKS.CloseConnection()
}

func (dms *dbMysql) CloseConnection() {
	dms.dbMs.Close()
}
