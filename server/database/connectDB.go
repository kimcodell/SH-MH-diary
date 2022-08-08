package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/kimcodell/SH-MH-diary/server/utils"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GetConnectedDB() *sql.DB {
	loadEnvError := godotenv.Load()
	utils.CatchError(utils.ErrorParams{Err: loadEnvError, Message: "Fail to load env file."})

	dbUserName := os.Getenv("DB_USER_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSchemaName := os.Getenv("DB_SCHEMA_NAME")
	config := mysql.Config{
		User:                 dbUserName,
		Passwd:               dbPassword,
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		Collation:            "utf8mb4_general_ci",
		Loc:                  time.Local,
		DBName:               dbSchemaName,
		AllowNativePasswords: true,
	}

	// db, dbConnectError := sql.Open("mysql", config.FormatDSN())
	// utils.CatchError(utils.ErrorParams{Err: dbConnectError, Message: "Fail to Open DB"})
	connector, connectionError := mysql.NewConnector(&config)
	utils.CatchError(utils.ErrorParams{Err: connectionError, Message: "Fail to Open DB"})
	fmt.Println("Success to connect")

	db := sql.OpenDB(connector)
	return db
}
