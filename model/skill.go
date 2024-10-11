package model

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	What string
}
