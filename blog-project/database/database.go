package database

import (
	"blog-project/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type databaseI interface {
	Create(value any) error
	FetchAll(value any) error
	FetchByID(value any) error
	Update(value any) error
	Delete(value any) error
}

type Database struct {
	DB *gorm.DB
}

func OpenDBConn() Database {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(new(model.Blog))
	if err != nil {
		panic("Failed to create a table")
	}
	return Database{DB: db}
}

func (d *Database) Create(value any) error {
	return d.DB.Create(value).Error
}

func (d *Database) FetchAll(value any) error {
	return d.DB.Find(value).Error
}

func (d *Database) FetchByID(value any) error {
	return d.DB.First(value).Error
}

func (d *Database) Delete(value any) error {
	return d.DB.Delete(value).Error
}

func (d *Database) Update(value any) error {
	return d.DB.Updates(value).Error
}
