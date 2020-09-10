package upload

type Excel struct {
	SN        string `json:"lid" excel:"SN"`
	Area      string `json:"firstname" excel:"Area "`
	Unit      string `json:"lastname" excel:"Unit"`
	Cid       string `json:"amount" excel:"CID"`
	CName     string `json:"paymentno" excel:"Center Name"`
	MemName   string `json:"purpose" excel:"Member's Name"`
	RegDate   string `json:"pn" excel:"Recognized Date"`
	Birthday  string `json:"dateapplied" excel:"Birthday"`
	LengthMem string `json:"encoded" excel:"Length of Membership"`
	NewCode   string `json:"status" excel:"New Branch Code"`
	NewCid    string `json:"category" excel:"New Customer ID"`
	Contact   string `json:"group" excel:"Contact Number"`
}

/*
f, err := excelize.OpenFile("temp-file/upload.xlsx")
	if err != nil {
		println(err.Error())
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Agusan Del Sur 1", "A1")
	if err != nil {
		println(err.Error())
		return
	}
	println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Agusan Del Sur 1")
	for _, row := range rows {
		for _, colCell := range row {
			print(colCell, "\t")
		}
		println()
	}
func ConfirmLoan(r *http.Request) (Excel, error) {
	f, err := excelize.OpenFile("temp-file/upload.xlsx")
	ex := Excel{}
	ex.SN = f.GetCellValue("A1")

	return ex, nil
}

*/
