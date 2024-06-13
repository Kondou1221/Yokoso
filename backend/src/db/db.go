package db

import (
    "fmt"
    "os"
    "log"

    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func NewDB() *gorm.DB {
    err := godotenv.Load()
    if err != nil {
        log.Fatalln(err)
    }
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("MYSQL_USER"),
        os.Getenv("MYSQL_PASSWORD"),
        os.Getenv("MYSQL_HOST"),
        os.Getenv("MYSQL_PORT"),
        os.Getenv("MYSQL_DATABASE"))

	//(postgres.Open(url)) を使って、GORM (gorm.Open) を介してデータベースへの接続
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(url)
		// log.Fatalln(err)
	}
	fmt.Println("Connceted")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}

// func Connect() *sql.DB {
//     err := godotenv.Load()
//     if err != nil {
//         fmt.Println(err.Error())
//     }

//     user := os.Getenv("DB_USER")
//     password := os.Getenv("DB_PASS")
//     host := os.Getenv("DB_HOST")
//     port := os.Getenv("DB_PORT")
//     database_name := os.Getenv("DB_NAME")

//     dbconf := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4"

//     db, err := sql.Open("mysql", dbconf)
//     if err != nil {
//         fmt.Println(err.Error())
//     }
//     return db
// }