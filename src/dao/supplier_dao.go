package dao

import "purchase/src/dao/conn"

type Supplier struct {
	Name string
	Code string
}

func (supplier *Supplier) Save() {
	conn.DB.Create(supplier)
}
