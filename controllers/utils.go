package controllers

import "gorm.io/gorm"

type ResponseStruct struct {
	Status  string `json:"status"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func Response() ResponseStruct {
	return ResponseStruct{
		Status:  "ERROR",
		Data:    nil,
		Message: "Something Went Wrong",
	}
}

var Db *gorm.DB

func SetDB(db *gorm.DB) {
	Db = db
}
