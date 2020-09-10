package upload

import (
	"Whiteliststest/config"
	"net/http"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "upload", nil)
}
