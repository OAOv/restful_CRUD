package types

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

type UserData struct {
	UserList []User `json:"data"`
	Message  string `json:"message"`
}

type Record struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Subject  string `json:"subject"`
	Score    string `json:"score"`
}

type RecordData struct {
	RecordList []Record `json:"data"`
	Message    string   `json:"message"`
}
