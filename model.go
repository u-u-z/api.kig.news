package main

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id   uint64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type Auth struct {
	gorm.Model
	Key             string `json:"key"`
	RemainingPoints int    `json:"remaining_points"`
}

type Price struct {
	gorm.Model
	// Id is the primary key & auto-incremented
	ID      uint64         `json:"id" gorm:"primaryKey"`
	Time    int64          `json:"time" `
	Price   float64        `json:"price" gorm:"type:decimal(10,2)" validator:"required"`
	Details datatypes.JSON `json:"details" gorm:"type:jsonb;default:null" validate:"max=1024"`
}
