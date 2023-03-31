package multi

import (
	"fmt"
	"os"
	"time"

	setting "github.com/go-sql-driver/mysql"
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

	loc, _ := time.LoadLocation("Asia/Jakarta") // handle any errors!

	c := setting.Config{
		User:      user,
		Passwd:    password,
		DBName:    dbname,
		Addr:      host + ":" + port,
		Net:       "tcp",
		ParseTime: true,
		Loc:       loc,
	}

	dsn2 := c.FormatDSN()
	fmt.Println(dsn2)

	db, err := gorm.Open(mysql.Open(dsn2), &gorm.Config{})

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
