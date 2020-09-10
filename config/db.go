package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Heroku Whitelists DB Connection
const (
	host     = "ec2-34-232-212-164.compute-1.amazonaws.com"
	port     = 5432
	user     = "jjamizitlgywmv"
	password = "6753a5e5cc75e455ffdab7710add7d50cb2c58c1420ad33e594a3d2b5413da74"
	dbname   = "dt403m0agorkn"
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
