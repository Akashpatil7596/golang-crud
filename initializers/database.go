package initializers

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID    uint `gorm:"primaryKey;default:auto_random()"`
	Name  string
	Age int
	Phone int
}

func ConnectToDB(){
	var err error
	dsn := "root:root@123@tcp(127.0.0.1:3306)/my_sql?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// db,err :=sql.Open("postgres","host=arjuna.db.elephantsql.com user=fygkcrba password=IzoeHiyowdnQGWvlIDUkrrOryQOaqusv dbname=fygkcrba sslmode=disable")

	if err != nil{
		fmt.Println("error",err)
	}

	DB.AutoMigrate(&User{})

	if DB != nil{
		fmt.Println("connected",DB)
	}
}
