package registry

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

type connect func() (Store, error)

var DBConfig struct {
	Host            string
	Password        string
	User            string
	DefaultDatabase string
}

type Store struct {
	DB *gorm.DB
}

type IStores interface {
	Default() Store
}

func (s Store) Default() Store {
	return s
}

func init() {
	DBConfig.Host = os.Getenv("DB_HOST")
	if len(DBConfig.Host) == 0 {
		DBConfig.Host = "127.0.0.1:3306"
	}
	DBConfig.DefaultDatabase = os.Getenv("DB_NAME")
	if len(DBConfig.DefaultDatabase) == 0 {
		DBConfig.DefaultDatabase = "mydb"
	}
	DBConfig.Password = os.Getenv("DB_PASSWORD")
	if len(DBConfig.Password) == 0 {
		DBConfig.Password = ""
	}
	DBConfig.User = os.Getenv("DB_USER")
	if len(DBConfig.User) == 0 {
		DBConfig.User = "root"
	}
}

func NewMySQLConnection() (Store, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.DefaultDatabase))
	if err != nil {
		return Store{}, err
	}
	return Store{DB: db}, nil
}

func RetryConnect(fnc connect, retry int) (Store, error) {
	for i := 0; i < retry; i++ {
		store, err := fnc()
		if err != nil {
			log.Print(err)
			time.Sleep(10 * time.Second)
			continue
		}
		return store, nil
	}
	return Store{}, fmt.Errorf("Failed connect to db")
}
