package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Heroku Whiteliststest DB Connection

const (
	host     = "ec2-54-161-58-21.compute-1.amazonaws.com"
	port     = 5432
	user     = "nndlnpvpkykeqw"
	password = "19db4d26d2d0202bb7f1c0063420951dd9d22d747c5cd89466aed3d1faa8f873"
	dbname   = "db42rt0ea031a1"
)

//Heroku DB Connection
/*const (
	host     = "ec2-34-236-215-156.compute-1.amazonaws.com"
	port     = 5432
	user     = "hsyaqxdwsbfitm"
	password = "c770eaa5b154c5a6882c5df91546e60e80da9e3e326d1ab3cbbfe02f522028e7"
	dbname   = "dbkqgrooodfron"
)
*/
//Local DB Connection
/*const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "EMPC"
)*/

var DB *sql.DB

//sslmode=require kapag sa heroku
//sslmode=disable kapag sa local
func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)
	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
