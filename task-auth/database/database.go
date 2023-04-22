package database

import (
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	cfg "github.com/vladqstrn/tasker-auth/task-auth/config"
	"github.com/vladqstrn/tasker-auth/task-auth/models"
)

func InitDB() *pg.DB {
	options := &pg.Options{
		User:     cfg.User,
		Password: cfg.Password,
		Addr:     cfg.Host + ":" + cfg.Port,
		Database: cfg.DbName,
	}

	db := pg.Connect(options)

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to database")
	return db
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{&models.User{}} {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
