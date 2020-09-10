package loans

import (
	"Whiteliststest/config"
	"errors"
	"fmt"
	"log"
	"net/http"
)

/*
type Loan struct {
	Lid         string `json:"lid"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Amount      string `json:"amount"`
	PaymentNo   string `json:"paymentno"`
	Purpose     string `json:"purpose"`
	PN          string `json:"pn"`
	DateApplied string `json:"dateapplied"`
	Encoded     string `json:"encoded"`
	Status      string `json:"status"`
	Category    string `json:"category"`
	Pgroup      string `json:"group"`
	Product     string `json:"product"`
	Frequency   string `json:"frequency"`
	Guarantor   string `json:"guarantor"`
	Cid         string `json:"cid"`
}*/
type LoanList struct {
	Lid         string `json:"lid"`
	ClientName  string `json:"firstname"`
	Cid         string `json:"cid"`
	Amount      string `json:"amount"`
	PaymentNo   string `json:"paymentno"`
	Purpose     string `json:"purpose"`
	PN          string `json:"pn"`
	DateApplied string `json:"dateapplied"`
	Encoded     string `json:"encoded"`
	Status      string `json:"status"`
	Category    string `json:"category"`
	Pgroup      string `json:"group"`
	Product     string `json:"product"`
	Frequency   string `json:"frequency"`
	Guarantor   string `json:"guarantor"`
}
type Loans struct {
	Lid         string `json:"lid"`
	Amount      string `json:"amount"`
	PaymentNo   string `json:"paymentno"`
	Purpose     string `json:"purpose"`
	PN          string `json:"pn"`
	DateApplied string `json:"dateapplied"`
	Encoded     string `json:"encoded"`
	Status      string `json:"status"`
	Category    string `json:"category"`
	Pgroup      string `json:"group"`
	Product     string `json:"product"`
	Frequency   string `json:"frequency"`
	Guarantor   string `json:"guarantor"`
	Cid         string `json:"cid"`
}
type Count struct {
	Counts string `json:"count"`
}

func AllLoans() ([]LoanList, error) {
	rows, err := config.DB.Query("Select tblloan.lid, tblclient.membername, tblloan.amount, tblloan.paymentno, tblloan.purpose, tblloan.pn, tblloan.dateapplied, tblloan.encoded, tblloan.status, tblloan.category, tblloan.pgroup, tblloan.product, tblloan.frequency, tblloan.guarantor, tblloan.cid from tblclient inner join tblloan on tblclient.cid = tblloan.cid ;")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	ln := make([]LoanList, 0)
	for rows.Next() {
		lns := LoanList{}
		err := rows.Scan(&lns.Lid, &lns.ClientName, &lns.Amount, &lns.PaymentNo, &lns.Purpose, &lns.PN, &lns.DateApplied, &lns.Encoded, &lns.Status, &lns.Category, &lns.Pgroup, &lns.Product, &lns.Frequency, &lns.Guarantor, &lns.Cid)
		if err != nil {
			return nil, err
		}
		ln = append(ln, lns)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ln, nil
}

func AllUnassignedLoans() ([]LoanList, error) {
	rows, err := config.DB.Query("Select tblloan.lid, tblclient.membername, tblloan.amount, tblloan.paymentno, tblloan.purpose, tblloan.pn, tblloan.dateapplied, tblloan.encoded, tblloan.status, tblloan.category, tblloan.pgroup, tblloan.product, tblloan.frequency, tblloan.guarantor, tblloan.cid from tblclient inner join tblloan on tblclient.cid = tblloan.cid where tblloan.status='Unassigned';")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()
	ln := make([]LoanList, 0)
	for rows.Next() {
		lns := LoanList{}
		err := rows.Scan(&lns.Lid, &lns.ClientName, &lns.Amount, &lns.PaymentNo, &lns.Purpose, &lns.PN, &lns.DateApplied, &lns.Encoded, &lns.Status, &lns.Category, &lns.Pgroup, &lns.Product, &lns.Frequency, &lns.Guarantor, &lns.Cid)
		if err != nil {
			return nil, err
		}
		ln = append(ln, lns)
	}

	if err = rows.Err(); err != nil {
		return nil, err

	}
	return ln, nil

}
func CountAllUnassignedLoans() ([]Count, error) {
	rows, err := config.DB.Query("SELECT COUNT(*) from tblloan where status='Unassigned';")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	ln := make([]Count, 0)
	for rows.Next() {
		lns := Count{}
		err := rows.Scan(&lns.Counts)
		if err != nil {
			return nil, err
		}
		ln = append(ln, lns)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ln, nil

}
func AllAssignedLoans() ([]LoanList, error) {
	rows, err := config.DB.Query("Select tblloan.lid, tblclient.membername, tblloan.amount, tblloan.paymentno, tblloan.purpose, tblloan.pn, tblloan.dateapplied, tblloan.encoded, tblloan.status, tblloan.category, tblloan.pgroup, tblloan.product, tblloan.frequency, tblloan.guarantor, tblloan.cid from tblclient inner join tblloan on tblclient.cid = tblloan.cid where tblloan.status='Assigned';")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()
	ln := make([]LoanList, 0)
	for rows.Next() {
		lns := LoanList{}
		err := rows.Scan(&lns.Lid, &lns.ClientName, &lns.Amount, &lns.PaymentNo, &lns.Purpose, &lns.PN, &lns.DateApplied, &lns.Encoded, &lns.Status, &lns.Category, &lns.Pgroup, &lns.Product, &lns.Frequency, &lns.Guarantor, &lns.Cid)
		if err != nil {
			return nil, err
		}
		ln = append(ln, lns)
	}

	if err = rows.Err(); err != nil {
		return nil, err

	}
	return ln, nil

}

func OneLoan(r *http.Request) (LoanList, error) {
	lns := LoanList{}
	lid := r.FormValue("lid")
	rows := config.DB.QueryRow("Select tblloan.lid, tblclient.membername, tblloan.amount, tblloan.paymentno, tblloan.purpose, tblloan.pn, tblloan.dateapplied, tblloan.encoded, tblloan.status, tblloan.category, tblloan.pgroup, tblloan.product, tblloan.frequency, tblloan.guarantor, tblloan.cid from tblclient inner join tblloan on tblclient.cid = tblloan.cid where tblloan.lid=$1;", lid)
	err := rows.Scan(&lns.Lid, &lns.ClientName, &lns.Amount, &lns.PaymentNo, &lns.Purpose, &lns.PN, &lns.DateApplied, &lns.Encoded, &lns.Status, &lns.Category, &lns.Pgroup, &lns.Product, &lns.Frequency, &lns.Guarantor, &lns.Cid)
	if err != nil {
		return lns, err
	}
	return lns, nil
}

func ConfirmLoan(r *http.Request) (Loans, error) {
	ln := Loans{}
	ln.Lid = r.FormValue("lid")
	ln.Amount = r.FormValue("amount")
	ln.PaymentNo = r.FormValue("paymentno")
	ln.Purpose = r.FormValue("purpose")
	ln.PN = r.FormValue("pn")
	ln.DateApplied = r.FormValue("dateapplied")
	ln.Encoded = r.FormValue("encoded")
	ln.Status = r.FormValue("status")
	ln.Category = r.FormValue("category")
	ln.Pgroup = r.FormValue("pgroup")
	ln.Product = r.FormValue("product")
	ln.Frequency = r.FormValue("frequency")
	ln.Guarantor = r.FormValue("guarantor")
	ln.Cid = r.FormValue("cid")
	var err error
	_, err = config.DB.Exec("UPDATE tblLoan SET lid = $1, amount = $2, paymentno=$3, purpose=$4, pn=$5, dateapplied=$6, encoded=$7, status=$8, category=$9, pgroup=$10, product=$11, frequency=$12, guarantor=$13, cid=$14 WHERE lid=$1;", ln.Lid, ln.Amount, ln.PaymentNo, ln.Purpose, ln.PN, ln.DateApplied, ln.Encoded, ln.Status, ln.Category, ln.Pgroup, ln.Product, ln.Frequency, ln.Guarantor, ln.Cid)

	if err != nil {
		//return ln, err
		log.Println(err)

		return ln, errors.New("500. Internal Server Error." + err.Error())
	}
	log.Println(" Loan Confirmed")
	return ln, nil
}

func InsertLoan(r *http.Request) (Loans, error) {

	ln := Loans{}
	ln.Amount = r.FormValue("amount")
	ln.PaymentNo = r.FormValue("paymentno")
	ln.Purpose = r.FormValue("purpose")
	ln.PN = r.FormValue("pn")
	ln.DateApplied = r.FormValue("dateapplied")
	ln.Encoded = r.FormValue("encoded")
	ln.Status = r.FormValue("status")
	ln.Category = r.FormValue("category")
	ln.Pgroup = r.FormValue("pgroup")
	ln.Product = r.FormValue("product")
	ln.Frequency = r.FormValue("frequency")
	ln.Guarantor = r.FormValue("guarantor")
	ln.Cid = r.FormValue("cid")

	fmt.Println(ln)
	var err error
	_, err = config.DB.Exec("INSERT INTO tblLoan(amount, paymentno, purpose, pn, dateapplied, encoded, status,category, pgroup, product, frequency, guarantor, cid ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)", ln.Amount, ln.PaymentNo, ln.Purpose, ln.PN, ln.DateApplied, ln.Encoded, ln.Status, ln.Category, ln.Pgroup, ln.Product, ln.Frequency, ln.Guarantor, ln.Cid)
	if err != nil {
		log.Println(err)
		return ln, errors.New("500. Internal Server Error." + err.Error())
	}
	log.Println("New Loan Added")
	return ln, nil
}

func SelectQueue(r *http.Request) (LoanList, error) {
	lns := LoanList{}
	lid := r.FormValue("lid")
	rows := config.DB.QueryRow("Select tblloan.lid, tblclient.membername, tblloan.amount, tblloan.paymentno, tblloan.purpose, tblloan.pn, tblloan.dateapplied, tblloan.encoded, tblloan.status, tblloan.category, tblloan.pgroup, tblloan.product, tblloan.frequency, tblloan.guarantor, tblloan.cid from tblclient inner join tblloan on tblclient.cid = tblloan.cid where tblloan.lid=$1;", lid)
	err := rows.Scan(&lns.Lid, &lns.ClientName, &lns.Amount, &lns.PaymentNo, &lns.Purpose, &lns.PN, &lns.DateApplied, &lns.Encoded, &lns.Status, &lns.Category, &lns.Pgroup, &lns.Product, &lns.Frequency, &lns.Guarantor, &lns.Cid)
	fmt.Println("err")
	if err != nil {
		return lns, err
	}
	return lns, nil
}
func SearchDate(r *http.Request) ([]LoanList, error) {

	//	dot, err := dotsql.LoadFromFile("queries.sql")
	//rows, err := dot.Query(config.DB, "search-loan")
	from := r.FormValue("from")
	to := r.FormValue("to")
	rows, err := config.DB.Query("Select tblloan.lid, tblclient.membername,  tblloan.amount, tblloan.paymentno, tblloan.purpose, tblloan.pn, tblloan.dateapplied, tblloan.encoded, tblloan.status, tblloan.category, tblloan.pgroup, tblloan.product, tblloan.frequency, tblloan.guarantor, tblloan.cid from tblclient inner join tblloan on tblclient.cid = tblloan.cid where tblloan.dateapplied BETWEEN '" + from + "' AND '" + to + "' and tblloan.status='Assigned';")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	loan := make([]LoanList, 0)
	for rows.Next() {
		lns := LoanList{}
		err := rows.Scan(&lns.Lid, &lns.ClientName, &lns.Amount, &lns.PaymentNo, &lns.Purpose, &lns.PN, &lns.DateApplied, &lns.Encoded, &lns.Status, &lns.Category, &lns.Pgroup, &lns.Product, &lns.Frequency, &lns.Guarantor, &lns.Cid)
		if err != nil {
			return nil, err
		}
		loan = append(loan, lns)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println(loan)
	return loan, nil
}
