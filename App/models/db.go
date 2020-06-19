package models

import (
    "log"
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    
    "gintest/config"
)

var db *gorm.DB

type Model struct {
    ID int `gorm:"primary_key" json:"id"`
    CreatedOn int `json:"created_on"`
    ModifiedOn int `json:"modified_on"`
}

func init() {
    var (
        err error
        dbType, dbName, user, password, host, tablePrefix string
    )


    dbType = "mysql"
    dbName = config.MysqlName
    user = config.MysqlUser
    password = config.MysqlPassword
    host = config.MysqlHost
    tablePrefix = config.MysqlPrefix

    db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", 
        user, 
        password, 
        host, 
        dbName))

    if err != nil {
        log.Println(err)
    }

    gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
        return tablePrefix + defaultTableName;
    }

    db.SingularTable(true)
    db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
    defer db.Close()
}