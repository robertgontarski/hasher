package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Database interface {
	Connect() error
	Close() error
	GetDb() *sql.DB
}

type DatabaseConf struct {
	Driver       string
	Dns          string
	MaxLifeTime  time.Duration
	MaxOpenConns int
	MaxIdleConns int
}

type MysqlDatabase struct {
	Db *sql.DB
	DatabaseConf
	Database
}

func NewMysqlDatabase(dbConf DatabaseConf) *MysqlDatabase {
	return &MysqlDatabase{
		DatabaseConf: dbConf,
	}
}

func (mdb *MysqlDatabase) Connect() error {
	db, err := sql.Open(mdb.Driver, mdb.Dns)
	if err != nil {
		return err
	}

	db.SetConnMaxLifetime(mdb.MaxLifeTime)
	db.SetMaxOpenConns(mdb.MaxOpenConns)
	db.SetMaxIdleConns(mdb.MaxIdleConns)

	if err := db.Ping(); err != nil {
		return err
	}

	mdb.Db = db

	return nil
}

func (mdb *MysqlDatabase) Close() error {
	if mdb.Db != nil {
		return mdb.Db.Close()
	}

	return fmt.Errorf("connection was not open")
}

func (mdb *MysqlDatabase) GetDb() *sql.DB {
	return mdb.Db
}
