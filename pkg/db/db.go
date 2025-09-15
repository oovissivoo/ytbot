package db

import "gorm.io/gorm"

type Db struct {
	*gorm.DB
}
