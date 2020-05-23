package model

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Content string `json:"content"`
	ImageID uint   `json:"imageID"`
	ItemID  uint   `json:"itemID"`
	UserID  uint   `json:"userId"`
}

type PostImages struct {
	ID    uint   `gorm:"primary_key"`
	Link1 string `json:"link1"`
	Link2 string `json:"link2"`
	Link3 string `json:"link3"`
	Link4 string `json:"link4"`
	Link5 string `json:"link5"`
	Link6 string `json:"link6"`
	Link7 string `json:"link7"`
	Link8 string `json:"link8"`
	Link9 string `json:"link9"`
}

type PostItem struct {
	ID    uint `gorm:"primary_key"`
	Price uint `json:"price"`
}
