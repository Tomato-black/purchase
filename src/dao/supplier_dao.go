package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Supplier struct {
	Name string
	Code string
}

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:Cao*012219@tcp(cdb-hex98lc6.bj.tencentcdb.com:10083)/purchase?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//设置全局表名禁用复数
	db.SingularTable(true)
}

func (supplier *Supplier) Save() {
	db.Create(supplier)
}
