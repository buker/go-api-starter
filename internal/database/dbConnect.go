package database

import (
	"database/sql"

	"github.com/buker/go-api-starter/internal/config"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	// Import MySQL database driver
	// _ "github.com/jinzhu/gorm/dialects/mysql"

	// Import PostgreSQL database driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"

	// Import SQLite3 database driver
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/driver/sqlite"

	log "github.com/sirupsen/logrus"
)

// DB global variable to access gorm
var DB *gorm.DB

var sqlDB *sql.DB
var err error

// InitDB - function to initialize db
func InitDB() *gorm.DB {
	var db = DB

	configureDB := config.Config()

	driver := configureDB.Database.DbDriver
	username := configureDB.Database.DbUser
	password := configureDB.Database.DbPass
	database := configureDB.Database.DbName
	host := configureDB.Database.DbHost
	port := configureDB.Database.DbPort
	sslmode := configureDB.Database.DbSslmode
	timeZone := configureDB.Database.DbTimeZone
	maxIdleConns := configureDB.Database.DbMaxIdleConns
	maxOpenConns := configureDB.Database.DbMaxOpenConns
	connMaxLifetime := configureDB.Database.DbConnMaxLifetime
	logLevel := configureDB.Database.DbLogLevel

	switch driver {
	case "postgres":
		dsn := "host=" + host + " port=" + port + " user=" + username + " dbname=" + database + " password=" + password + " sslmode=" + sslmode + " TimeZone=" + timeZone
		sqlDB, err = sql.Open(driver, dsn)
		if err != nil {
			log.WithError(err).Panic()
		}
		sqlDB.SetMaxIdleConns(maxIdleConns)
		sqlDB.SetMaxOpenConns(maxOpenConns)
		sqlDB.SetConnMaxLifetime(connMaxLifetime)

		db, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.LogLevel(logLevel)),
		})
		if err != nil {
			log.WithError(err).Panic()
		}

		if err == nil {
			log.Info("DB connection successful!")
		}

	case "sqlite3":
		db, err = gorm.Open(sqlite.Open(database), &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			log.WithError(err).Panic()
		}

		if err == nil {
			log.Info("DB connection successful!")
		}

	default:
		log.Fatal("The driver " + driver + " is not implemented yet")
	}

	DB = db

	return DB
}

// GetDB - get a connection
func GetDB() *gorm.DB {
	return DB
}
