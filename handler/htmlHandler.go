package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/OAOv/restful_CRUD/types"
)

type FHandler struct {
}

var userBody []byte
var recordBody []byte
var err error

var isOne = false
var isGet = true
var searchID string
var searchRecordID string
var searchUserID string

func (fh *FHandler) TmplHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./view/layout.html"))

	var users types.UserData
	var records types.RecordData

	if r.Method != http.MethodPost {
		if !isOne && isGet {
			userBody, recordBody, err = DoReadAllRequest()
		} else if isGet {
			isOne = false
			userBody, recordBody, err = DoReadOneRequest(searchID, searchRecordID, searchUserID)
		}

		log.Println("UserBody Response: " + string(userBody))
		log.Println("recordBody Response: " + string(recordBody))
		json.Unmarshal(userBody, &users)
		json.Unmarshal(recordBody, &records)

		tmpl.Execute(w, struct {
			Title         string
			UserInput     []string
			Operate       []string
			UserData      []types.User
			UserMessage   string
			RecordInput   []string
			RecordTable   []string
			RecordData    []types.Record
			RecordMessage string
		}{
			Title:         "users",
			UserInput:     []string{"ID", "Name", "Age"},
			Operate:       []string{"create", "readAll", "readOne", "update", "delete"},
			UserData:      users.UserList,
			UserMessage:   users.Message,
			RecordInput:   []string{"RecordID", "UserID", "Subject", "Score"},
			RecordTable:   []string{"ID", "UserName", "Subject", "Score"},
			RecordData:    records.RecordList,
			RecordMessage: records.Message,
		})
	} else {
		log.Println("button: " + r.FormValue("btn"))

		switch r.FormValue("btn") {
		case "create":
			isGet = false
			userBody, recordBody, err = DoCreateRequest(r.FormValue("ID"), r.FormValue("Name"), r.FormValue("Age"), r.FormValue("RecordID"), r.FormValue("UserID"), r.FormValue("Subject"), r.FormValue("Score"))

		case "readAll":
			isGet = true

		case "readOne":
			isOne = true
			isGet = true
			searchID = r.FormValue("ID")
			searchRecordID = r.FormValue("RecordID")
			searchUserID = r.FormValue("UserID")

		case "update":
			isGet = false
			userBody, recordBody, err = DoUpdateRequest(r.FormValue("ID"), r.FormValue("Name"), r.FormValue("Age"), r.FormValue("RecordID"), r.FormValue("UserID"), r.FormValue("Subject"), r.FormValue("Score"))

		case "delete":
			isGet = false
			userBody, recordBody, err = DoDeleteRequset(r.FormValue("ID"), r.FormValue("RecordID"))
		}
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
