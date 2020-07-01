package models


import (
    "log"
    "fmt"
    "time"
    "unsafe"

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

type DbConfig struct{
    dbType string
    dbName string
    user string
    password string
    host string
    tablePrefix string
}

var DConfig *DbConfig

func init() {
    var (
        err error
    )

    DefaultConfig()

    //fmt.Printf("%v",DConfig)
    db, err = gorm.Open(DConfig.dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", 
        DConfig.user, 
        DConfig.password, 
        DConfig.host, 
        DConfig.dbName))

    if err != nil {
        log.Println(err)
    }

    gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
        return DConfig.tablePrefix + defaultTableName;
    }

    db.SingularTable(true)
    db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(100)
}

func DefaultConfig() {
    DConfig = &DbConfig{
        dbType:"mysql",
        dbName:"test",
        user:"root",
        password:"123456",
        host:"122.51.88.27:3306",
        tablePrefix:"test_",
    }

    if unsafe.Sizeof(config.DatabaseSetting) >8 {
        DConfig.dbName = config.DatabaseSetting.MysqlName
        DConfig.user = config.DatabaseSetting.MysqlUser
        DConfig.password = config.DatabaseSetting.MysqlPassword
        DConfig.host = config.DatabaseSetting.MysqlHost
        DConfig.tablePrefix = config.DatabaseSetting.MysqlPrefix
    }
}

func CloseDB() {
    defer db.Close()
}

func (model *Model) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedOn", time.Now().Unix())

    return nil
}

func (model *Model) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("ModifiedOn", time.Now().Unix())
    return nil
}