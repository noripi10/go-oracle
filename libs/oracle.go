package libs

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"xorm.io/xorm"

	"github.com/godror/godror"
	_ "github.com/godror/godror"

	"github.com/noripi10/go-oracle/model"
)

func Query() {

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
	// var connectionParams godror.ConnectionParams
	// connectionParams.Username = os.Getenv("USER_NAME")
	// connectionParams.Password = godror.NewPassword(os.Getenv("USER_PASSWORD"))
	// connectionParams.ConnectString = os.Getenv("CONNECTION_STRING")
	// connectionParams.SessionTimeout = 30 * time.Second
	// // P.SetSessionParamOnInit("NLS_NUMERIC_CHARACTERS", ",.")
	// // P.SetSessionParamOnInit("NLS_LANG", "Japanese_JAPAN.JA16SJIS")
	// fmt.Println(godror.NewConnector(connectionParams))
	// db := sql.OpenDB(godror.NewConnector(connectionParams))

	// defer db.Close()

	// CD := os.Getenv("CD")
	// var shainCode, shainName, queryErr = GetShain(db, CD)
	// if queryErr != nil {
	// 	fmt.Println(queryErr)
	// }

	// fmt.Println(shainCode, shainName)

	// 3. xorm + godror
	tnsAdmin := os.Getenv("TNS_ADMIN")
	os.Setenv("TNS_ADMIN", tnsAdmin)
	dataSourceName := os.Getenv("DB_DATA_SOURCE")
	engine, err := xorm.NewEngine("godror", dataSourceName)

	var connectionParams godror.ConnectionParams
	connectionParams.Username = os.Getenv("USER_NAME")
	connectionParams.Password = godror.NewPassword(os.Getenv("USER_PASSWORD"))
	connectionParams.ConnectString = os.Getenv("CONNECTION_STRING")
	connectionParams.SessionTimeout = 30 * time.Second
	// alter session set nls_date_format='yyyy/mm/dd hh24:mi:ss'
	// ORA-01861対応(timeのデフォルト書式に合わせる)
	connectionParams.SetSessionParamOnInit("NLS_DATE_FORMAT", "yyyy-mm-dd hh24:mi:ss")
	connector := godror.NewConnector(connectionParams)
	engine.DB().DB = sql.OpenDB(connector)

	if err != err {
		log.Fatal(err)
	}
	// When you want to manually close the engine, call engine.Close.
	// Generally, we don’t need to do this, because engine will be closed automatically when application exits.
	defer engine.Close()
	engine.ShowSQL(true)

	// Raw Query
	// shains, queryErr := engine.Query("select * from m_shain where shain_code = '11925'")
	// if queryErr != nil {
	// 	log.Fatal(queryErr)
	// }
	// for _, shain := range shains {
	// 	for key, value := range shain {
	// 		fmt.Printf("Key: %v, Value: %v\n", key, string(value))
	// 	}
	// }

	// Sync2
	// syncErr := engine.Sync2(new(model.WkData))
	// if syncErr != nil {
	// 	log.Fatal(syncErr)
	// }

	// Find
	// var wkData []model.WkData
	// findErr := engine.Limit(3).Find(&wkData)
	// if err != nil {
	// 	log.Fatal(findErr)
	// }
	// fmt.Println(wkData)

	wkData := model.WkData{
		Id:         1001,
		Data1:      "hoge",
		InsertDate: time.Now(),
		UpdateDate: time.Now(),
	}
	fmt.Print(wkData)

	// Indert
	session := engine.NewSession()

	affected, insertErr := engine.Insert(&wkData)
	if insertErr != nil {
		session.Rollback()
		log.Fatal(insertErr)
	}

	session.Commit()

	fmt.Println(affected)
}
