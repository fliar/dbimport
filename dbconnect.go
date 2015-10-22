package main

import(
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

type DbConnect struct{
	info string
	db *sql.DB
}

func (w *DbConnect)Connect(info string) {
	w.info = info;
	fmt.Println("connecting to ", info)
	db, err := sql.Open("mysql", info)
	if err!= nil {
		fmt.Println("db open failed, param is ", info)
		panic(err.Error())
		os.Exit(0)
	}
	w.db = db
}

func (w DbConnect)AddFile(tablename, filename, filetype, filecontent string) {
	sql := fmt.Sprint(
		"INSERT INTO `", tablename,
		"`(`location`,`type`,`content`) VALUES('", filename,
		"','", filetype,
		"','", filecontent,
		"');")
	stmt, err := w.db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	}
}
