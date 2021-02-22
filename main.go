package main

import (
	"Whiteliststest/entities/clients"
	"Whiteliststest/entities/loans"
	"Whiteliststest/entities/upload"
	"Whiteliststest/entities/uploadfile"
	"Whiteliststest/entities/user"
	"Whiteliststest/uploads"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	mux := http.NewServeMux()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	//log.Fatal(http.ListenAndServe(port, mux))
	//Used port := os.Getenv("PORT")
	//"PORT" variable name na ansa .env file
	//kapag nasa local sya nag run, yung 8080 na port ang kukunin nya. yun yung nasa .env na file
	//then kapag sa heroku or other platform naman, kung anong ibinigay nilang port
	//ay maadopt na ni application, so there's no need to manually update the port kapag inilagay mo sa ibang hosting

	log.Println("Server Started...")
	log.Println("Current Port: " + port)
	//upload excel to json
	//mux.HandleFunc("/", uploadexcel.UploadData)
	mux.HandleFunc("/uploadfile", uploads.UploadFile)
	//mux.PathPrefix("/").Handler(http.FileServer(http.Dir("./templates/")))

	mux.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	mux.Handle("/templates/", http.StripPrefix("/templates", http.FileServer(http.Dir("templates"))))

	//end
	//html
	mux.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))
	mux.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("js"))))
	mux.Handle("/json-browse/", http.StripPrefix("/json-browse", http.FileServer(http.Dir("json-browse"))))

	//browser

	mux.HandleFunc("/client/lists", clients.ClientLists)
	mux.HandleFunc("/search/member", clients.Search)
	//mux.HandleFunc("/login/client", clients.Login)
	//mux.HandleFunc("/login/client/process", clients.LoginClientProcess)
	mux.HandleFunc("/loan/lists", loans.LoanLists)
	mux.HandleFunc("/unassigned/loans", loans.ListofUnassigned)
	mux.HandleFunc("/assigned/loans", loans.ListofAssigned)
	mux.HandleFunc("/confirm/loans", loans.SelectedLoan)
	mux.HandleFunc("/confirmation/process", loans.ConfirmationLoanProcess)
	mux.HandleFunc("/apply/loan", loans.ApplyLoan)
	mux.HandleFunc("/apply/loan/process", loans.ApplyLoanProcess)
	//	mux.HandleFunc("/queue/form", loans.QueueForm)
	mux.HandleFunc("/queue/form", loans.Queue)
	mux.HandleFunc("/search/dates", loans.SearchDateProcess)
	mux.HandleFunc("/confirm/message", loans.ConfirmMessage)

	mux.HandleFunc("/", user.LoginUserForm)
	mux.HandleFunc("/login/user/process", user.LoginUserProcess)
	mux.HandleFunc("/register/user", user.RegisterUserForm)
	mux.HandleFunc("/register/user/process", user.RegisterUserProcess)
	mux.HandleFunc("/registertest", user.RegisterUserProcessTest)
	mux.HandleFunc("/password", user.Password)

	//mux.HandleFunc("/update/client", clients.UpdateClientForm)
	mux.HandleFunc("/update/client", clients.SelectedClient)
	mux.HandleFunc("/update/client/process", clients.UpdateClientProcess)
	mux.HandleFunc("/insert/client/process", clients.InsertClientProcess)
	mux.HandleFunc("/delete/client/process", clients.DeleteClientProcess)
	mux.HandleFunc("/samplelogin", user.SampeLogin)
	mux.HandleFunc("/sampleregister", user.SampleRegister)
	//api-json type
	mux.HandleFunc("/all/loans", loans.GetLoans)
	mux.HandleFunc("/insert", loans.Insert)
	mux.HandleFunc("/allclients", clients.GetAllClients)
	//upload
	mux.HandleFunc("/upload/process", upload.UploadFile)
	mux.HandleFunc("/upload/form", upload.UploadForm)
	mux.HandleFunc("/upload/list", uploadfile.Uploadlists)
	mux.HandleFunc("/client", clients.CheckList)
	mux.HandleFunc("/test", clients.SamplePage)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		panic(err)
	}
}
