package handler

import "github.com/gorilla/mux"

func Router(router *mux.Router) {
	h := &Handler{}

	router.HandleFunc("/users", h.GetUsers).Methods("GET")
	router.HandleFunc("/users", h.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", h.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", h.UpdateUser).Methods("PATCH")
	router.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")
}

func FRouter(router *mux.Router) {
	fh := &FHandler{}

	router.HandleFunc("/", fh.TmplHandler)
}

//users跟user的差別, 單筆資料使用user, 多筆users
//repo跟db層結合
//handler跟(repo+db)中間加一個service層 => handler / service / (repo+db)
//repo層的json.Unmarshal修改到handler, 順便做判斷
//錯誤處理http status, middlewares
//http framework
