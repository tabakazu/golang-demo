package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Conn *gorm.DB

func init() {
	conn, err := gorm.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}
	Conn = conn
}

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type User struct {
	Model
}

func main() {
	db, err := gorm.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("データベースへの接続に失敗しました")
	}
	fmt.Println("データベースへの接続に成功しました")
	defer db.Close()

	conn := Conn

	if err := conn.AutoMigrate(&User{}).Error; err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("マイグレーションが成功しました")

	if err := conn.Create(&User{}).Error; err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("データ挿入が成功しました")
}
