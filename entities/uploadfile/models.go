package uploadfile

import (
	"Whiteliststest/config"
	"log"
)

type UploadFile struct {
	Id       string `json:"id"`
	UserID   string `json:"userid"`
	Filename string `json:"filename"`
	Data     string `json:"date"`
	Datetime string `json:"datetimeuploaded"`
	FileType string `json:"filetype"`
}

func ListofUploadFile() ([]UploadFile, error) {
	rows, err := config.DB.Query("Select id, userid, filename, data, datetimeuploaded, filetype from tblupload")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()
	up := make([]UploadFile, 0)
	for rows.Next() {
		upf := UploadFile{}
		err := rows.Scan(&upf.Id, &upf.UserID, &upf.Filename, &upf.Data, &upf.Datetime, &upf.FileType)
		if err != nil {
			return nil, err
		}
		up = append(up, upf)
	}

	if err = rows.Err(); err != nil {
		return nil, err

	}
	return up, nil

}
