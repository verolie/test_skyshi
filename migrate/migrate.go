package migrate

import (
	"fmt"
	"time"

	"github.com/verolie/test-skyshi/multi"
)

func Init() {
	db := multi.SetDatabase()
	db.AutoMigrate(&Activities{})
	db.AutoMigrate(&Todos{})

	sqlDB, _ := db.DB()
	sqlDB.Close()

	fmt.Println("success migrate")
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
