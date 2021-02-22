package clients

import (
	"Whiteliststest/config"
	"encoding/json"
	"fmt"
	"net/http"
)

func UpdateClientForm(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "updateclient", nil)
	config.TPL.ExecuteTemplate(w, "sidebar", nil)
}
func SamplePage(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "testonly", nil)
}

func InsertClientProcess(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	_, err := InsertClient(r)
	if err != nil {
		panic(err)
		//return
	}
	http.Redirect(w, r, "/client/lists", http.StatusMovedPermanently)
}
func DeleteClientProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	err := DeleteClient(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/client/lists", http.StatusSeeOther)
}

func ClientLists(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	cl, err := AllClients()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(w, "listofclients", cl)
	//config.TPL.ExecuteTemplate(w, "sampletable", cl)
}
func GetAllClients(w http.ResponseWriter, r *http.Request) {
	clt, err := AllClients()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Request-Reference-No", "`1e9ac446-8a62-4ae3-852d-c352ceda99b`")
	json.NewEncoder(w).Encode(clt)
}
func CheckList(w http.ResponseWriter, r *http.Request) {
	cl, err := CheckClient(w, r)
	if err != nil {
		http.Error(w, http.StatusText(200)+err.Error(), http.StatusInternalServerError)
		return
	}
	/*	if cl == nil {
		http.Error(w, http.StatusText(400), http.StatusMethodNotAllowed)
	}*/
	//http.Error(w, http.StatusText(400), http.StatusMethodNotAllowed)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Request-Reference-No", "`1e9ac446-8a62-4ae3-852d-c352ceda99b`")
	json.NewEncoder(w).Encode(cl)
}
func SelectedClient(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	cl, err := OneClient(r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	//config.TPL.ExecuteTemplate(w, "sidebar", cl)
	config.TPL.ExecuteTemplate(w, "updateclient", cl)
}
func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	cl, err := SearchClient(w, r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	//config.TPL.ExecuteTemplate(w, "sidebar", cl)
	config.TPL.ExecuteTemplate(w, "sampletable", cl)
}
func UpdateClientProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	_, err := UpdateClient(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/client/lists", http.StatusMovedPermanently)
	fmt.Println("Client Updated")
}
func Login(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "loginclient", nil)
}
func LoginClientProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	_, err := LoginClient(w, r)
	if err != nil {
		//	panic(err)
		return
	}
	http.Redirect(w, r, "/apply/loan", http.StatusMovedPermanently)
}

/*
func ReturnClient(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["key"]
	for _, clt := range All() {
		if clt.cid == key {
			json.NewEncoder(w).Encode(clt)
		}
	}

}
*/
