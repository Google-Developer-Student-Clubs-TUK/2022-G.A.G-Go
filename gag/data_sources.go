package main

import (
	"fmt"

	"gag.com/eclass"
	"gag.com/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dataSources struct {
	DB     *gorm.DB
	Eclass *eclass.Eclass
}

func initDS() (*dataSources, error) {
	// gorm 설정
	dsn := "root:root@tcp(127.0.0.1:3306)/gag?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Device{})
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.Comment{})

	// eclass 설정
	ecl := &eclass.Eclass{}

	return &dataSources{
		DB:     db,
		Eclass: ecl,
	}, nil
}
