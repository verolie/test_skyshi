package multi

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	CONST_MYSQL_HOST     string = "localhost"
	CONST_MYSQL_PORT     string = "4040"
	CONST_MYSQL_USER     string = "root"
	CONST_MYSQL_PASSWORD string = "admin"
	CONST_MYSQL_DBNAME   string = "todo4"
)

func SetDatabase() *gorm.DB {

	host := checkConditionConst("MYSQL_HOST")
	port := checkConditionConst("MYSQL_PORT")
	user := checkConditionConst("MYSQL_USER")
	password := checkConditionConst("MYSQL_PASSWORD")
	dbname := checkConditionConst("MYSQL_DBNAME")

	// loc, _ := time.LoadLocation("Asia/Jakarta") // handle any errors!

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}

func checkConditionConst(constValue string) string {
	value := os.Getenv(constValue)
	if value != "" {
		return value
	} else {
		if constValue == "MYSQL_HOST" {
			return CONST_MYSQL_HOST
		} else if constValue == "MYSQL_PORT" {
			return CONST_MYSQL_PORT
		} else if constValue == "MYSQL_USER" {
			return CONST_MYSQL_USER
		} else if constValue == "MYSQL_PASSWORD" {
			return CONST_MYSQL_PASSWORD
		} else if constValue == "MYSQL_DBNAME" {
			return CONST_MYSQL_DBNAME
		}
	}

	return value
}
