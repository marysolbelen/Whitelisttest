package uploads

import (
	"Whiteliststest/config"
	"Whiteliststest/entities/clients"
	"Whiteliststest/entities/user"
	"Whiteliststest/parser"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	//"text/template"
	"log"
	//"time"
)

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
	/*	Area               string `json:"Area"`
		Birthday           string `json:"Birthday"`
		CID                string `json:"CID"`
		CenterName         string `json:"CenterName"`
		Contact            string `json:"Contact"`
		Flag               string `json:"Flag"`
		LengthOfMembership string `json:"LengthOfMembership"`
		MemberName         string `json:"MemberName"`
		NewBranchCode      string `json:"NewBranchCode"`
		NewCID             string `json:"NewCID"`
		RecognizedDate     string `json:"RecognizedDate"`
		SN                 string `json:"SN"`
		Unit               string `json:"Unit"`*/

}

type UploadData struct {
	NilFile, InvalidFile bool
	Clients              map[string]ClientList
}

// func New()(*Salutation){
//     return &( Salutation{} );
// }

// type errUploading struct {
// 	nilFile bool
// 	invalidFile bool
// }

func Cookies(w http.ResponseWriter, req *http.Request) {

	session, err := user.Store.Get(req, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := user.GetUser(session)
	if auth := user.Authenticated; !auth {
		session.AddFlash("You don't have access!")
		err = session.Save(req, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, req, "/login/user", http.StatusFound)

		return
	}
	fmt.Println(user)
	config.TPL.ExecuteTemplate(w, "uploadfile.html", user)
	//	t, _ := template.ParseFiles("./templates/uploadfile.html")
	//	t.Execute(w, user)
}

func UploadFile(w http.ResponseWriter, req *http.Request) {

	// session, err := user.Store.Get(req, "cookie-name")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	//user := user.GetUser(session)
	//uID := user.Uid

	if req.Method == "GET" {
		// ln, err := getAllUploadedFiles()
		// if err != nil {
		// 	http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		config.TPL.ExecuteTemplate(w, "uploadfile", nil)
	} else if req.Method == "POST" {

		// t, _ := template.ParseFiles("./templates/uploadfile.html")
		// t.Execute(w, nil)
		file, handler, err := req.FormFile("uploadfile")
		if err != nil {

			log.Println("Null File: ", err)

			config.TPL.ExecuteTemplate(w, "uploadfile", nil)
			return

		} else {
			fmt.Println("error throws in else statement")
			fmt.Println("handler.Filename", handler.Filename)
			fmt.Printf("Type of handler.Filename:%T\n", handler.Filename)
			fmt.Println("Length:", len(handler.Filename))

			f, err := os.OpenFile("./data/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Println("Open File Error: ", err)
				config.TPL.ExecuteTemplate(w, "uploadfile", nil)
				return
			}
			defer f.Close()
			io.Copy(f, file)
			blobPath := "./data/" + handler.Filename
			var extension = filepath.Ext(blobPath)
			parsedData := ExcelCsvParser(blobPath, extension)
			parsedJson, _ := json.Marshal(parsedData)
			//fmt.Println(string(parsedJson))

			err = os.Remove(blobPath)
			//t, _ := template.ParseFiles("./templates/uploadfile.html")
			//t.Execute(w, string(parsedJson))
			//config.TPL.ExecuteTemplate(w, "uploadfile", string(parsedJson))
			//config.TPL.ExecuteTemplate(w, "sidebar", user)
			//currentTime := time.Now().Local()
			//	currentTime.Format("2006/01/02 15:04")
			//	fmt.Println("YYYY.MM.DD : ", currentTime.Format("2006.01.02 15:04:05"))

			//userid := uID

			/*	_, err = config.DB.Exec("WITH contacts_json (doc) AS (VALUES('" + parsedJson + "'::json))INSERT INTO tblclient (flag, sn, area, unit, cid, centername, membername, rdate, bday, mlength, bcode, newcid, contact) SELECT p.* FROM contacts_json l CROSS JOIN lateral json_populate_recordset(NULL::tblclient, doc) AS p ON conflict (cid) do nothing;")
				fmt.Println("test")
				if err != nil {
					log.Println(err)
				}
			*/
			//	fmt.Println(client)
			//var ClientList error
			//json.Unmarshal([]byte(s), data)
			//fmt.Println(data.Area)
			/*	s := string(parsedJson)
				data := ClientList{}
				err = json.Unmarshal([]byte(s), data)
				fmt.Println(data)
				if err != nil {
					fmt.Println(err.Error())
				}*/

			// _, err = config.DB.Exec("INSERT INTO tblupload(userid, filename, data, datetimeuploaded, filetype) VALUES ($1, $2, $3, $4, $5)", userid, handler.Filename, parsedJson, currentTime.Format("2006/01/02 15:04"), extension)

			// if err != nil {
			// 	log.Println("Insert Data Error: ",err)

			// }

			keysBody := []byte(parsedJson)
			keys := make([]ClientList, 0)
			json.Unmarshal(keysBody, &keys)

			//cl := ClientList{}
			_, err = config.DB.Exec("DELETE FROM tblclient;INSERT INTO tblclient SELECT area, birthday, cid, centername, contact, flag, lengthofmembership, membername, newbranchcode, newcid, recognizeddate, sn, unit FROM json_populate_recordset(NULL::tblclient,  '" + string(keysBody) + "')")
			//	_, err = config.DB.Exec("INSERT INTO tblclient(area, bday, cid, centername, contact, flag, mlength, membername, bcode, newcid, rdate, sn, unit) select * from json_populate_recordset(null::tblclient,  '" + string(keysBody) + "')")
			//_, err = config.DB.Exec("INSERT INTO tblclient(Area, Birthday, CID, CenterName, Contact, Flag, LengthOfMembership, MemberName, NewBranchCode, NewCID, RecognizedDate, SN, Unit) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)", cl.Area, cl.Birthday, cl.CID, cl.CenterName, cl.Contact, cl.Flag, cl.LengthOfMembership, cl.MemberName, cl.NewBranchCode, cl.NewCID, cl.RecognizedDate, cl.SN, cl.Unit)
			//_, err = config.DB.Exec("INSERT INTO tblclient (Area, Birthday, CID, CenterName, Contact, Flag, LengthOfMembership, MemberName, NewBranchCode, NewCID, RecognizedDate, SN, Unit) select * from json_populate_recordset(null::tblclient,  '" + string(keysBody) + "')")

			if err != nil {
				//http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				log.Println("Invalid File: ", err)

				config.TPL.ExecuteTemplate(w, "uploadfile", nil)

			} else {

				ac, err := clients.AllClients()
				if err != nil {
					http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
				}

				config.TPL.ExecuteTemplate(w, "uploadfile", ac)

			}

			//fmt.Println(ClientList.Area)
			//	cl := ClientList{}
			//
			//_, err = config.DB.Exec("INSERT INTO tblclient(area, bday, cid, centername, contact, flag, mlength, membername, bcode, newcid, rdate, sn, unit) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)", cl.Area, cl.Birthday, cl.CID, cl.CenterName, cl.Contact, cl.Flag, cl.LengthOfMembership, cl.MemberName, cl.NewBranchCode, cl.NewCID, cl.RecognizedDate, cl.SN, cl.Unit)
			//	query := `INSERT INTO tblclient(area, bday, cid, centername, contact, flag, mlength, membername, bcode, newcid, rdate, sn, unit) VALUES (:Area,:Birthday,:CID,:CenterName, :Contact, :Flag,:LengthofMembership,:MemberName,:NewBranchCode, :NewCID, :RecognizedDate, :SN, :Unit)`
			//_, err = config.DB.Exec("INSERT INTO tblclient(area, bday, cid, centername, contact, flag, mlength, membername, bcode, newcid, rdate, sn, unit) VALUES ('" + cl.Area + "', '" + cl.Birthday + "', '" + cl.CID + "', '" + cl.CenterName + "', '" + cl.Contact + "', '" + cl.Flag + "', '" + cl.LengthOfMembership + "', '" + cl.MemberName + "', '" + cl.NewBranchCode + "', '" + cl.NewCID + "', '" + cl.RecognizedDate + "', '" + cl.SN + "' ,'" + cl.Unit + "')")

			//insert here

			//	fmt.Println("Client Added")
		}

	}
}

func ExcelCsvParser(blobPath string, blobExtension string) (parsedData []map[string]interface{}) {
	fmt.Println("---------------> We are in product.go")
	if blobExtension == ".csv" {
		fmt.Println("-------We are parsing an csv file.-------------")
		parsedData := parser.ReadCsvFile(blobPath)
		fmt.Printf("Type:%T\n", parsedData)
		return parsedData

	} /*else if blobExtension == ".xlsx" {
		fmt.Println("----------------We are parsing an xlsx file.---------------")
		parsedData := parser.ReadXlsxFile(blobPath)
		return parsedData
	} else if blobExtension == ".xls" {
		fmt.Println("----------------We are parsing an xls file.---------------")
		parsedData := parser.ReadXlsFile(blobPath)
		return parsedData
	}*/
	return parsedData
}

func init() {
	path := "./data"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
		fmt.Println("Created data directory")
	} else {
		fmt.Println("Data directory already exists")
	}
}
