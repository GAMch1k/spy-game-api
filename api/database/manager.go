package database 

import (
	"database/sql"
	"errors"
	"os"
	"log"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var Path string = "./api/database/databases/database.db"
var DBTableName = "games"


type DB_field struct {
	Name string
	Type string
}


func OpenDatabase() *sql.DB {

	log.Println("Opening database file", Path)
	sqliteDatabase, err := sql.Open("sqlite3", Path)

	if err != nil {
		log.Fatal(err.Error())
	}

	return sqliteDatabase
}


func InitDatabase() {
	if _, err := os.Stat(Path); errors.Is(err, os.ErrNotExist) {
		log.Println("Creating database file in", Path)
		
		file, err := os.Create(Path)
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("Database created in", Path)
		} else {
			log.Printf("Database in %s already exists", Path)
	}

	sqliteDatabase := OpenDatabase()
	defer CloseDatabase(sqliteDatabase)

	if !CheckIfTableExists(sqliteDatabase, DBTableName) {
		fields := []DB_field {
			{
				Name: "game_id",
				Type: "INTEGER PRIMARY KEY AUTOINCREMENT",
			},
			{
				Name: "player_id",
				Type: "STRING DEFAULT '[]'",
			},
			{
				Name: "spy_id",
				Type: "STRING DEFAULT ''",
			},
			{
				Name: "topic",
				Type: "STRING DEFAULT ''",
			},
			{
				Name: "Started",
				Type: "INTEGER DEFAULT 0",
			},
		}
		CreateTable(sqliteDatabase, DBTableName, fields)
	}
	
	log.Println("Database fully initialized!")
}


func CloseDatabase(db *sql.DB) {
	db.Close()
	log.Println("Database closed")
}


func CheckIfTableExists(db *sql.DB, name string) bool {
	log.Printf("Checking if table %s exist", name)
	
	query, err := db.Prepare("SELECT name FROM sqlite_master WHERE type='table' AND name=?")
	
	if err != nil {
		log.Fatal(err.Error())
	}
	
	defer query.Close()
	
	var output string
	err = query.QueryRow(name).Scan(&output)
	
	if err == sql.ErrNoRows {
		return false
	}
	
	if err != nil {
		log.Fatal(err.Error())
	}
	
	return true
	
}


func RemoveElementFromFields(slice []DB_field, s int) []DB_field {
	return append(slice[:s], slice[s+1:]...)
}


func CreateTable(db *sql.DB, name string, fields []DB_field) {
	if len(name) <= 0 {
		log.Fatalf("Name of the table %s is empty", name)
	}
	
	if len(fields) <= 0 {
		log.Fatalf("Fields of the table %s is empty", name)
	}
	

	log.Println("Creating new table", name)
	
	createTable := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS "%s" (%s %s)`, name, fields[0].Name, fields[0].Type)
	statement, err := db.Prepare(createTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()

	fields = RemoveElementFromFields(fields, 0)

	log.Println("Created new table:", name)
	log.Println("Creating fields for table", name)
	

	for _, el := range fields {
		createField := fmt.Sprintf("ALTER TABLE %s ADD %s %s", name, el.Name, el.Type)

		statement, err := db.Prepare(createField)
		if err != nil {
			log.Fatal(err.Error())
		}
		statement.Exec()
	}
}


