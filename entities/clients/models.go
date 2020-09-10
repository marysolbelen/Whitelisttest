package clients

import (
	"Whiteliststest/config"
	"database/sql"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
)

type Client struct {
	Cid       string    `json:"cid"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Contact   string    `json:"contact"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Date      time.Time `json:"dateregistered"`
}
type ClientList struct {
	Area               string `json:"area"`
	Birthday           string `json:"birthday"`
	CID                string `json:"cid"`
	CenterName         string `json:"centername"`
	Contact            string `json:"contact"`
	Flag               string `json:"flag"`
	LengthOfMembership string `json:"lengthofmembership"`
	MemberName         string `json:"membername"`
	NewBranchCode      string `json:"newbranchcode"`
	NewCID             string `json:"newcid"`
	RecognizedDate     string `json:"recognizeddate"`
	SN                 string `json:"sn"`
	Unit               string `json:"unit"`
}
type clientCreds struct {
	Cid           string    `json:"cid"`
	FirstName     string    `json:"firstname"`
	LastName      string    `json:"lastname"`
	Contact       string    `json:"contact"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Date          time.Time `json:"dateregistered"`
	Authenticated bool
}
type Log struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func InsertClient(r *http.Request) (ClientList, error) {
	cl := ClientList{}
	fmt.Println(cl)
	var err error
	_, err = config.DB.Exec("INSERT INTO tblclient(area, birthday, cid, centername, contact, flag, lengthofmembership, membername, newbranchcode, newcid, recognizeddate, sn, unit) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)", cl.Area, cl.Birthday, cl.CID, cl.CenterName, cl.Contact, cl.Flag, cl.LengthOfMembership, cl.MemberName, cl.NewBranchCode, cl.NewCID, cl.RecognizedDate, cl.SN, cl.Unit)
	if err != nil {
		log.Println(err)
		return cl, errors.New("500. Internal Server Error." + err.Error())
	}
	log.Println("New Client Added")
	return cl, nil
}

var Store *sessions.CookieStore

func init() {
	authKeyOne := []byte("::AUTH-ONE::")
	encryptionKeyOne := []byte("::AK&-L@&@Q-SKS-LSQ-aI::")

	Store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15,
		HttpOnly: true,
		SameSite: 2,
	}

	gob.Register(clientCreds{})
}

var granted bool

func AllClients() ([]ClientList, error) {
	//	rows, err := config.DB.Query("SELECT tblLoan.LoanID, tblLoan.firstname, tblLoan.lastname, tblLoan.Contact, tblLoan.amount, tblLoan.Date, tblLoan.Status, tblLoan.AppID, tblRelease.releaseID, tblRelease.appamount, tblRelease.monthly, tblrelease.amortization, tblRelease.branch, tblrelease.appid from tblLoan INNER JOIN tblRelease on tblLoan.LoanID = tblRelease.loanid;")
	rows, err := config.DB.Query("Select area, birthday, cid, centername, contact, flag, lengthofmembership, membername, newbranchcode, newcid, recognizeddate, sn, unit from tblclient;")
	if err != nil {

		log.Println(err)
		return nil, err

	}
	defer rows.Close()
	cls := make([]ClientList, 0)
	for rows.Next() {
		cl := ClientList{}
		err := rows.Scan(&cl.Area, &cl.Birthday, &cl.CID, &cl.CenterName, &cl.Contact, &cl.Flag, &cl.LengthOfMembership, &cl.MemberName, &cl.NewBranchCode, &cl.NewCID, &cl.RecognizedDate, &cl.SN, &cl.Unit)
		if err != nil {
			return nil, err
		}
		cls = append(cls, cl)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cls, nil
}
func CheckClient(w http.ResponseWriter, r *http.Request) ([]ClientList, error) {
	newcid := r.FormValue("newcid")
	//brcode := r.FormValue("brcode")
	//	rows, err := config.DB.Query("SELECT tblLoan.LoanID, tblLoan.firstname, tblLoan.lastname, tblLoan.Contact, tblLoan.amount, tblLoan.Date, tblLoan.Status, tblLoan.AppID, tblRelease.releaseID, tblRelease.appamount, tblRelease.monthly, tblrelease.amortization, tblRelease.branch, tblrelease.appid from tblLoan INNER JOIN tblRelease on tblLoan.LoanID = tblRelease.loanid;")
	rows, err := config.DB.Query("Select area, birthday, cid, centername, contact, flag, lengthofmembership, membername, newbranchcode, newcid, recognizeddate, sn, unit from tblclient where newcid = '" + newcid + "' ;")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	/*
		if rows == nil {

			return nil, err
		}
		w.Write([]byte("Not found"))*/
	defer rows.Close()
	cls := make([]ClientList, 0)
	for rows.Next() {
		cl := ClientList{}
		err := rows.Scan(&cl.Area, &cl.Birthday, &cl.CID, &cl.CenterName, &cl.Contact, &cl.Flag, &cl.LengthOfMembership, &cl.MemberName, &cl.NewBranchCode, &cl.NewCID, &cl.RecognizedDate, &cl.SN, &cl.Unit)
		if err != nil {
			return nil, err
		}
		cls = append(cls, cl)

	}
	/*	if err == sql.ErrNoRows {
			return nil, err
		}
		w.Write([]byte("Not found"))
	*/
	if err = rows.Err(); err != nil {
		return nil, err
	}
	/*if err == sql.ErrNoRows {
		return nil, err
	}
	//cls = append(cls, )
	w.Write([]byte("Not found"))*/
	fmt.Println(cls)
	return cls, nil
}

func LoginClient(w http.ResponseWriter, r *http.Request) ([]Log, error) {
	//	dot, err := dotsql.LoadFromFile("queries.sql")
	//rows, err := dot.Query(config.DB, "search-loan")
	if r.Method == "POST" {
		session, err := Store.Get(r, "cookie-name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil, err
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		result := config.DB.QueryRow("SELECT * FROM tblclient WHERE USERNAME=$1 and PASSWORD=$2", username, password)
		if err != nil {
			// If there is an issue with the database, return a 500 error
			panic(err)

		}

		storedCreds := &clientCreds{}
		err = result.Scan(
			&storedCreds.Cid,
			&storedCreds.FirstName,
			&storedCreds.LastName,
			&storedCreds.Contact,
			&storedCreds.Username,
			&storedCreds.Password,
			&storedCreds.Date,
		)

		if err != nil {
			// If an entry with the username does not exist, send an "Unauthorized"(401) status
			if err == sql.ErrNoRows {
				w.Write([]byte("notExist"))
				return nil, err
			}
			// If the error is of any other type, send a 500 status
			w.WriteHeader(http.StatusInternalServerError)
			return nil, err
		}

		if storedCreds.Password == password {
			granted = true
		} else {
			granted = false
		}

		switch granted {
		case true:
			creds := &clientCreds{
				FirstName:     storedCreds.FirstName,
				LastName:      storedCreds.LastName,
				Contact:       storedCreds.Contact,
				Username:      storedCreds.Username,
				Date:          storedCreds.Date,
				Cid:           storedCreds.Cid,
				Authenticated: true,
			}

			session.Values["user"] = creds

			err = session.Save(r, w)

			if err != nil {
				panic(err)
			}
			http.Redirect(w, r, "/apply/loan", http.StatusMovedPermanently)
			fmt.Println("correct")

		case false:
			fmt.Println("Incorrect")
		}
		return nil, err
	}
	var err error
	return nil, err
}
func GetUser(s *sessions.Session) clientCreds {
	val := s.Values["user"]
	var user = clientCreds{}
	user, ok := val.(clientCreds)
	if !ok {
		return clientCreds{Authenticated: false}
	}

	return user
}
