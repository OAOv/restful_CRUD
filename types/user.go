package types

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

type UserData struct {
	Data    []User `json:"data"`
	Message string `json:"message"`
}
