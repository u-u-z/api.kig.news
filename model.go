package main

import "gorm.io/gorm"

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
	ID             uint64                 `json:"id" gorm:"primaryKey"`
	Price          float64                `json:"price" gorm:"type:decimal(10,2)"`
	Details        map[string]interface{} `json:"price_details" gorm:"type:jsonb;default:null"`
	UploadUserName string                 `json:"user_id" gorm:"default:null"`
}
