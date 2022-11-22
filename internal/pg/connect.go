package pg

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open("host=localhost user=enzo password=password dbname=zion port=5432 sslmode=disable TimeZone=UTC"), &gorm.Config{})
}
