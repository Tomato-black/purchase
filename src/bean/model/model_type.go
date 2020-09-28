package model

import "github.com/jinzhu/gorm"

type Supplier struct {
	gorm.Model
	Name string
	Code string
}


