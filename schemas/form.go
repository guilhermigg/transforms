package schemas

import "gorm.io/gorm"

type Form struct {
	gorm.Model
	Title       string
	Description string
	Banner      string
	Owner       int64
	Public      bool
}
