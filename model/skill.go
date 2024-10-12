package model

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	Topic string
	What  string
}
