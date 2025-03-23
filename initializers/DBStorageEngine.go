package initializers

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jbarham/gopgsqldriver"
)

type DBStorage struct {
	Db *sql.DB
}

func NewDBStorage(dbConn *sql.DB) *DBStorage {

	return &DBStorage{
		Db: dbConn,
	}
}

var (
	instance *DBStorage
	once     sync.Once
)

func GetConnection(driverName string, url string) (*DBStorage, error) {

	var err error
	once.Do(func() {
		conn, err := sql.Open(driverName, url)
		if err != nil {
			err = err
			return
		}

		instance = &DBStorage{
			Db: conn,
		}
	})
	return instance, err
}
