// autoMigrate.go needs to be executed only when it is required

package main

import (
	"gorm.io/gorm"

	"github.com/buker/go-api-starter/internal/config"
	"github.com/buker/go-api-starter/internal/database"
	"github.com/buker/go-api-starter/internal/database/model"
	log "github.com/sirupsen/logrus"
)

// Load all the models
type auth model.Auth
type user model.User

var db *gorm.DB
var errorState int

func main() {
	configureDB := config.Config()
	driver := configureDB.Database.DbDriver

	/*
	** 0 = default/no error
	** 1 = error
	**/
	errorState = 0

	db = database.InitDB()

	// Auto migration
	/*
		- Automatically migrate schema
		- Only create tables with missing columns and missing indexes
		- Will not change/delete any existing columns and their types
	*/

	// Careful! It will drop all the tables!
	dropAllTables()

	// Automatically migrate all the tables
	migrateTables()

	// Manually set foreign keys for MySQL and PostgreSQL
	if driver != "sqlite3" {
		setPkFk()
	}

	if errorState == 0 {
		log.Info("Auto migration is completed!")
	} else {
		log.Error("Auto migration failed!")
	}
}

func dropAllTables() {
	// Careful! It will drop all the tables!
	if err := db.Migrator().DropTable(&user{}, &auth{}); err != nil {
		errorState = 1
		log.WithError(err).Error("Error while dropping all tables")
	} else {
		log.Info("Old tables are deleted!")
	}
}

func migrateTables() {
	configureDB := config.Config()
	driver := configureDB.Database.DbDriver

	if driver == "mysql" {
		// db.Set() --> add table suffix during auto migration
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&auth{},
			&user{}); err != nil {
			errorState = 1
			log.WithError(err).Error("Error while migrating tables")
		} else {
			log.Info("New tables are  migrated successfully!")
		}
	} else {
		if err := db.AutoMigrate(&auth{},
			&user{}); err != nil {
			errorState = 1
			log.WithError(err).Error("Error while migrating tables")
		} else {
			log.Info("New tables are  migrated successfully!")
		}
	}
}

func setPkFk() {
	// Manually set foreign key for MySQL and PostgreSQL
	if err := db.Migrator().CreateConstraint(&auth{}, "Users"); err != nil {
		errorState = 1
		log.WithError(err).Error("Error while creating foreign key for auth table")

		if err := db.Migrator().CreateConstraint(&user{}, "Posts"); err != nil {
			errorState = 1
			log.WithError(err).Error("Error while creating foreign key for user table")
		}
	}
}
