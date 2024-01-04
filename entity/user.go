package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint `json:"id" gorm:"primarykey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}


type Todo struct{
	Id        uint `json:"id" gorm:"primarykey"`
	Judul   string `json:"judul"`
	Hari   string `json:"hari"`
	Waktu   string `json:"waktu"`
	Keterangan   string `json:"keterangan"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}