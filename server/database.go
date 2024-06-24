package server

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"

	_ "github.com/mattn/go-sqlite3"
)

func Init_db() {
	// Open the database:
	_, err := os.Open("./data.db")
	if err != nil {
		fmt.Println(err)
		exec.Command("sqlite3", "data.db")
		fmt.Println("false")
		exec.Command(".tables")
		exec.Command(".exit")
	}
	db := Open_DB()
	defer db.Close()
	// Check si les tables de bases sont créée (WIP)

	// TABLE USERS
	var sts []string = []string{
		`
	CREATE TABLE IF NOT EXISTS users (
		user_id TEXT PRIMARY KEY, 
		username TEXT, 
		email TEXT,
		password TEXT,
		date TEXT,
		session_id TEXT,
		session_expiration TEXT
	);
	`, `
	CREATE TABLE IF NOT EXISTS rss_subscriptions (
		id INT PRIMARY KEY,
		sub_id,
		url TEXT,
		name TEXT,
		FOREIGN KEY(sub_id) REFERENCES users(user_id)
	);
	`,`
	CREATE TABLE IF NOT EXISTS tags (
		sub_id,
		subscription_id,
		title TEXT,
		FOREIGN KEY(sub_id) REFERENCES users(user_id)
		FOREIGN KEY(subscription_id) REFERENCES rss_subscriptions(id)
	);
	`,
	}
	fmt.Println(len(sts))

	for i, s := range sts {
		fmt.Println("debug st : ", i)
		statement, _ := db.Prepare(s)
		statement.Exec()
	}
}

func Open_DB() (DB *sql.DB) {
	DB, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("Can't open the database")
	}
	return DB
}
