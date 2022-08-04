package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/kimcodell/SH-MH-diary/server/utils"
)

func ConnectToDB() {
	db := getConnector()
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	var name interface{}
	err = db.QueryRow("SELECT * FROM post WHERE id = 1").Scan(&name)
	utils.CatchError(err)
	fmt.Println(name)
	print("success")
}

func getConnector() *sql.DB {
	config := mysql.Config{
		User:      "root",
		Passwd:    "jordan11",
		Addr:      "127.0.0.1:3306",
		Collation: "utf8mb4_general_ci",
		Loc:       time.Local,
		DBName:    "sh-mh-diary",
	}
	connector, err := mysql.NewConnector(&config)
	utils.CatchError(err)
	db := sql.OpenDB(connector)
	return db
}
