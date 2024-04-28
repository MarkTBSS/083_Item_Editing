package databases

import (
	"fmt"
	"log"
	"sync"

	"github.com/MarkTBSS/083_Item_Editing/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	*gorm.DB
}

var postgresDatabaseInstance *postgresDatabase
var once sync.Once

func (db *postgresDatabase) Connect() *gorm.DB {
	return postgresDatabaseInstance.DB
}

// Use migrations
func NewPostgresDatabase(conf *config.Database) Database {
	//func NewPostgresDatabase(conf *config.Database) *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s search_path=%s",
			conf.Host,
			conf.User,
			conf.Password,
			conf.DBName,
			conf.Port,
			conf.SSLMode,
			conf.Schema,
		)

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		log.Printf("Connected to database %s", conf.DBName)
		postgresDatabaseInstance = &postgresDatabase{conn}
	})
	// Use migrations
	return postgresDatabaseInstance
	//return postgresDatabaseInstance.DB
}
