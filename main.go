package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/noripi10/go-oracle/libs"

	"github.com/joho/godotenv"

	"github.com/godror/godror"
	// _ "github.com/godror/godror"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error read .env")
	}

	// 1.TNSを使った接続方法
	// tnsAdmin := os.Getenv("TNS_ADMIN")
	// os.Setenv("TNS_ADMIN", tnsAdmin)
	// dataSourceName := os.Getenv("DB_DATA_SOURCE")
	// db, err := sql.Open("godror", dataSourceName)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// 2.ConnectionParamsを使った接続方法
	var connectionParams godror.ConnectionParams
	connectionParams.Username = os.Getenv("USER_NAME")
	connectionParams.Password = godror.NewPassword(os.Getenv("USER_PASSWORD"))
	connectionParams.ConnectString = os.Getenv("CONNECTION_STRING")
	connectionParams.SessionTimeout = 30 * time.Second
	// P.SetSessionParamOnInit("NLS_NUMERIC_CHARACTERS", ",.")
	// P.SetSessionParamOnInit("NLS_LANG", "Japanese_JAPAN.JA16SJIS")
	fmt.Println(godror.NewConnector(connectionParams))
	db := sql.OpenDB(godror.NewConnector(connectionParams))

	defer db.Close()

	CD := os.Getenv("CD")
	var shainCode, shainName, queryErr = libs.GetShain(db, CD)
	if queryErr != nil {
		fmt.Println(err)
	}

	fmt.Println(shainCode, shainName)
}
