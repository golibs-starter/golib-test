package golibtest

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type DatabaseTestSuite struct {
	db *gorm.DB
}

func NewDatabaseTestSuite(db *gorm.DB) *DatabaseTestSuite {
	return &DatabaseTestSuite{db: db}
}

func (d *DatabaseTestSuite) TruncateTable(table string)  {
	if err := d.db.Exec(fmt.Sprintf("TRUNCATE TABLE `%s`", table)).Error; err != nil {
		log.Fatalf("Could not truncate table %s: %v", table, err)
	}
}

func (d *DatabaseTestSuite) Insert(model interface{}) {
	if err := d.db.Create(model).Error; err != nil {
		log.Fatalf("Could not create seed data, model: %v, err: %v", model, err)
	}
}

func (d *DatabaseTestSuite) CountWithoutQuery(table string) int64 {
	var count int64
	d.db.Table(table).Count(&count)
	return count
}

func (d *DatabaseTestSuite) CountWithQuery(table string, conditions map[string]interface{}) int64 {
	var count int64
	d.db.Table(table).Where(conditions).Count(&count)
	return count
}
