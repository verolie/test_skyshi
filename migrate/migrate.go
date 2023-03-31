package migrate

import (
	"time"

	"github.com/verolie/test-skyshi/multi"
)

const (
	CONST_MYSQL_HOST     string = "localhost"
	CONST_MYSQL_PORT     string = "4040"
	CONST_MYSQL_USER     string = "root"
	CONST_MYSQL_PASSWORD string = "admin"
	CONST_MYSQL_DBNAME   string = "todo4"
)

func Init() {
	db := multi.SetDatabase()
	db.AutoMigrate(&Activities{})
	db.AutoMigrate(&Todos{})
}

type Activities struct {
	Activity_id int `gorm:"primary_key;auto_increment;not_null"`
	Title       string
	Email       string
	CreateAt    time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
type Todos struct {
	Todo_id         int        `gorm:"primary_key;auto_increment;not_null"`
	ActicityGroupId int        `gorm:"index; column:activity_group_id"`
	Activities      Activities `gorm:"ForeignKey:activity_group_id; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Title           string
	Priority        string
	IsActive        int       `gorm:"column:is_active"`
	CreateAt        time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}
