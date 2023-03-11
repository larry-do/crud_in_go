package datasource

import (
	"CRUD/configuration"
	"CRUD/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	db  *gorm.DB
	err error
)

func Initialize() {
	var datasource, dialect, autoMigrate = loadDatasource("application.properties")

	db, err = gorm.Open(dialect, fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		datasource.Host, datasource.Port, datasource.Dbname,
		datasource.User, datasource.Password,
	))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Database connection created. GORM Dialect %s\n", dialect)

	if autoMigrate {
		migrateToDatabase()
	}
}

func migrateToDatabase() {
	db.AutoMigrate(model.Album{})
	db.AutoMigrate(model.Artist{})
	log.Println("Migrated database.")
}

func loadDatasource(filename string) (configuration.Datasource, string, bool) {
	log.Printf("Read datasource properties from %s\n", filename)
	var err = godotenv.Load(filename)
	if err != nil {
		log.Fatalf("Error loading file %s %s", filename, err)
	}
	var datasource = configuration.Datasource{}
	datasource.Host = os.Getenv("DB_HOST")
	datasource.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	datasource.Dbname = os.Getenv("DB_NAME")
	datasource.User = os.Getenv("DB_USER")
	datasource.Password = os.Getenv("DB_PASSWORD")
	var autoMigrate, _ = strconv.ParseBool(os.Getenv("gorm.auto_migrate"))
	return datasource, os.Getenv("gorm.dialect"), autoMigrate
}

func Connection() *gorm.DB {
	return db
}
