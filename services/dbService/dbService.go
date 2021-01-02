package dbService

import (
	. "backend/types"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"reflect"
)

var db *gorm.DB

func Init() {
	dsn, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		dsn = "postgres://postgres:pass@localhost:5432/postgres"
	}
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to postgres")
	}
	db = conn
	types := []interface{}{&About{}}
	for _, val := range types {
		err = db.AutoMigrate(val)
		if err != nil {
			panic(fmt.Sprintf("failed to automigrate %s", reflect.TypeOf(val)))
		}
	}
	return
}

func GetAbout() (items []AboutDTO) {
	db.Model(&About{}).Find(&items)
	return
}
