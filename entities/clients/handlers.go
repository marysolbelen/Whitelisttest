package clients

import (
	"Whiteliststest/config"
	"encoding/json"
	"net/http"
)

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

/*
func CheckList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	_, err := APIClient()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}
	//	config.TPL.ExecuteTemplate(w, "listofclients", cl)
}*/
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