package cmd

import (
	_ "github.com/mattn/go-sqlite3"
)

type databasefunc interface {
	createdatabase()
}


/*
func createdatabase() {
	database, _ := sql.Open("sqlite3", "data.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY), login TEXT, password TEXT")
	statement.Exec()
}*/
