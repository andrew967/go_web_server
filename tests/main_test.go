package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestDBConnection(t *testing.T) {
	Dsn := "host=localhost user=andrew password=TiYx9a395%k^ dbname=web_server_go port=5432"
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	assert.Nil(t, err)

	dbconf, err := db.DB()
	defer dbconf.Close()

	err = db.Raw("SELECT 1").Error
	assert.Nil(t, err)

}
