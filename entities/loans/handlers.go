package loans

import (
	"Whiteliststest/config"
	"Whiteliststest/entities/user"
	"encoding/json"
	"fmt"
	"net/http"
)

func LoanLists(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	ln, err := AllLoans()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(w, "listofloans", ln)
}

func CountForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	ln, err := CountAllUnassignedLoans()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(w, "unassignedloans", ln)
}

func ListofUnassigned(w http.ResponseWriter, r *http.Request) {
	session, err := user.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := user.GetUser(session)
	if auth := user.Authenticated; !auth {
		session.AddFlash("You don't have access!")
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login/user", http.StatusFound)

		return
	}

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	ln, err := AllUnassignedLoans()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "unassignedloans", ln)

}

func ListofAssigned(w http.ResponseWriter, r *http.Request) {
	session, err := user.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := user.GetUser(session)
	if auth := user.Authenticated; !auth {
		session.AddFlash("You don't have access!")
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login/user", http.StatusFound)

		return
	}

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	ln, err := AllAssignedLoans()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(w, "assignedloans", ln)
	config.TPL.ExecuteTemplate(w, "sidebar", user)
}
func SelectedLoan(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	ln, err := OneLoan(r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(w, "confirmation", ln)
}
func CountUnassigned(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	ln, err := CountAllUnassignedLoans()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(w, "unassignedloans", ln)
}

func ConfirmationLoanProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	_, err := ConfirmLoan(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/assigned/loans", http.StatusMovedPermanently)
	fmt.Println("loan confimation")
}

func ApplyLoan(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "applyloan", nil)
}

func ApplyLoanProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	_, err := InsertLoan(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		fmt.Println(err)
		return

	}
	http.Redirect(w, r, "/unassigned/loans", http.StatusMovedPermanently)
}

func Queue(w http.ResponseWriter, r *http.Request) {
	session, err := user.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := user.GetUser(session)
	if auth := user.Authenticated; !auth {
		session.AddFlash("You don't have access!")
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login/user", http.StatusFound)

		return
	}
	fmt.Println(user)
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	ln, err := SelectQueue(r)

	if err != nil {

		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return

	}

	config.TPL.ExecuteTemplate(w, "queue", ln)
	config.TPL.ExecuteTemplate(w, "sidebar", user)
}
func SearchDateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	ln, err := SearchDate(r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(w, "assignedloans", ln)
}

func ConfirmMessage(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "confirmmessage", nil)
}
func GetLoans(w http.ResponseWriter, r *http.Request) {
	bks, err := AllLoans()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Request-Reference-No", "`1e9ac446-8a62-4ae3-852d-c352ceda99b`")
	json.NewEncoder(w).Encode(bks)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	bks, err := InsertLoan(r)
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Request-Reference-No", "`1e9ac446-8a62-4ae3-852d-c352ceda99b`")
	json.NewEncoder(w).Encode(bks)
}
