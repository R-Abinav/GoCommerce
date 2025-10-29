package main

import (
	"log"
	"database/sql"
	"github.com/R-Abinav/GoCommerce/cmd/api"
	"github.com/R-Abinav/GoCommerce/config"
	"github.com/R-Abinav/GoCommerce/db"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()

	db, err := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	});

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db);

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB){
	err := db.Ping() //Connects to the DB
	if err != nil{
		log.Fatal(err);
	}

	log.Println("[DB] Successfully connected to the Database!");
}