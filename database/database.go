package database

import (
	"context"
	"log"
	"sync"

	"github.com/jackc/pgx/v5"
)

var (
	instance *DatabaseContext
	once     sync.Once
)

type DatabaseContext struct {
	conn *pgx.Conn
}

func GetInstance() *DatabaseContext {
	once.Do(func() {
		instance = &DatabaseContext{}
		instance.init()
	})
	return instance
}

func (db *DatabaseContext) init() {
	ctx := context.Background()
	dsn := "host=localhost user=postgres password=postgres dbname=proyecto sslmode=disable"

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	db.conn = conn
}

func (db *DatabaseContext) Close() {
	db.conn.Close(context.Background())
}


// import (
// 	"log"

// 	"example.com/prueba/sqlc"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type DatabaseContext struct {
// 	db  *gorm.DB
// 	err error
// }

// var InstanceDatabase *DatabaseContext

// func (c *DatabaseContext) GetContext() *gorm.DB {
// 	return c.db
// }

// func (c *DatabaseContext) InitDB() {
// 	dataSourceName := "root:root@tcp(localhost:3306)/proyecto?parseTime=True"
// 	c.db, c.err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
// 	if c.err != nil {
// 		log.Fatal("Error al conectar a la base de datos.")
// 	}
// 	err := c.db.AutoMigrate(&sqlc.Connection{}, &sqlc.Site{}, &sqlc.User{}).Error()
// 	if err != "" {
// 		log.Fatal(err)
// 	}
// }

// func GetInstance() *DatabaseContext {
// 	if InstanceDatabase == nil {
// 		InstanceDatabase = &DatabaseContext{}
// 		InstanceDatabase.InitDB()
// 	}
// 	return InstanceDatabase
// }
