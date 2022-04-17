package db

import (
	"fmt"
	"sample_gin_app/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB  *gorm.DB
	err error
)

const (
	DriverName = "mysql"
	UserName   = "sample_gin_user"
	Password   = "passw0rd"
	Protcol    = "tcp(db:3306)"
	Database   = "sample_gin_app_db"
)

func Init() {
	connectTemplate := "%s:%s@%s/%s?parseTime=true"
	connect := fmt.Sprintf(
		connectTemplate,
		UserName,
		Password,
		Protcol,
		Database,
	)

	DB, err = gorm.Open(DriverName, connect)

	if err != nil {
		panic(err)
	}

	autoMigration()

}

func Close() {
	err := DB.Close()

	if err != nil {
		panic(err)
	}
}

func autoMigration() {
	DB.AutoMigrate(&entity.User{})
}
