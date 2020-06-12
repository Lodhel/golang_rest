package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"fmt"
)

var db *gorm.DB

func init() {
	fmt.Println("database init runs")
	// Reading the config file
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// This will be needed for PostgreSQL
	username := viper.Get("postgres")
	password := viper.Get("q319546")
	dbName := viper.Get("chief")
	dbHost := viper.Get("localhost")
	dbPort := viper.Get("5432")

	dbUri := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", 
		dbHost, 
		dbPort, 
		username, 
		dbName, 
		password,
	)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}
	db = conn
}