package model

import "time"

type CreatedMiddleNames struct {
	ID         uint `gorm:"primary_key"`
	Mr         string
	LName      string
	SurName    string
	CommonName string
	FName      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}
