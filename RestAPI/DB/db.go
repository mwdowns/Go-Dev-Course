package db

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

func Client() (*supabase.Client, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	client, err := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"), nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

//var DB *sql.DB
//
//func InitDB() {
//	var err error
//	DB, err = sql.Open("sqlite3", "api.db")
//	if err != nil {
//		panic("could not connect to database")
//	}
//
//	DB.SetMaxOpenConns(10)
//	DB.SetMaxIdleConns(5)
//
//	createTables()
//}
//
//func createTables() {
//	createEventsTable := `
//	CREATE TABLE IF NOT EXISTS events(
//		id INTEGER PRIMARY KEY AUTOINCREMENT,
//		name TEXT NOT NULL,
//		description TEXT NOT NULL,
//		location TEXT NOT NULL,
//		date_time DATETIME NOT NULL,
//		user_id INTEGER
//    )
//	`
//
//	_, err := DB.Exec(createEventsTable)
//	if err != nil {
//		panic("could not create events table")
//	}
//}
