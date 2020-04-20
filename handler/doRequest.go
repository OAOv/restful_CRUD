package handler

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func ClientDo(req *http.Request) ([]byte, error) {
	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

func createUser(id string, name string, age string) ([]byte, error) {
	jsonStr := []byte("{\"id\":\"" + id + "\",\"name\":\"" + name + "\",\"age\":\"" + age + "\"}")
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8000/user", bytes.NewBuffer(jsonStr))
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func createRecord(recordID string, userID string, subject string, score string) ([]byte, error) {
	jsonStr := []byte("{\"id\":\"" + recordID + "\",\"user_id\":\"" + userID + "\",\"subject\":\"" + subject + "\",\"score\":\"" + score + "\"}")
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8000/record", bytes.NewBuffer(jsonStr))
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func readAllUser() ([]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8000/users", nil)
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func readAllRecord() ([]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8000/records", nil)
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func readOneUser(id string) ([]byte, error) {
	if id == "" {
		id = "-1"
	}
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8000/user/"+id, nil)
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func readOneRecord(recordID string) ([]byte, error) {
	if recordID == "" {
		recordID = "-1"
	}
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8000/record/"+recordID, nil)
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func readOneUserRecords(userID string) ([]byte, error) {
	if userID == "" {
		userID = "-1"
	}
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8000/records/user/"+userID, nil)
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func updateUser(id string, name string, age string) ([]byte, error) {
	if id == "" {
		id = "-1"
	}
	var jsonStr = []byte("{\"id\":\"" + id + "\",\"name\":\"" + name + "\",\"age\":\"" + age + "\"}")
	req, _ := http.NewRequest("PATCH", "http://localhost:8000/user/"+id, bytes.NewBuffer(jsonStr))
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func updateRecord(recordID string, userID string, subject string, score string) ([]byte, error) {
	var jsonStr = []byte("{\"id\":\"" + recordID + "\",\"user_id\":\"" + userID + "\",\"subject\":\"" + subject + "\",\"score\":\"" + score + "\"}")
	req, _ := http.NewRequest("PATCH", "http://localhost:8000/record/"+recordID, bytes.NewBuffer(jsonStr))
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func deleteUser(id string) ([]byte, error) {
	req, _ := http.NewRequest("DELETE", "http://localhost:8000/user/"+id, nil)
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func deleteRecord(recordID string) ([]byte, error) {
	req, _ := http.NewRequest("DELETE", "http://localhost:8000/record/"+recordID, nil)
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

////////////////////////////
////////////////////////////

func DoCreateRequest(id string, name string, age string, recordID string, userID string, subject string, score string) ([]byte, []byte, error) {
	userBody, err = createUser(id, name, age)
	if err != nil {
		return nil, nil, err
	}

	recordBody, err = createRecord(recordID, userID, subject, score)
	if err != nil {
		return nil, nil, err
	}

	if id == "" && userID != "" {
		userBody, err = readAllUser()
		if err != nil {
			return nil, nil, err
		}
	} else if id != "" && userID == "" {
		recordBody, err = readAllRecord()
		if err != nil {
			return nil, nil, err
		}
	}
	return userBody, recordBody, nil
}

func DoReadAllRequest() ([]byte, []byte, error) {
	userBody, err = readAllUser()
	if err != nil {
		return nil, nil, err
	}

	recordBody, err = readAllRecord()
	if err != nil {
		return nil, nil, err
	}

	return userBody, recordBody, err
}

func DoReadOneRequest(id string, recordID string, userID string) ([]byte, []byte, error) {
	userBody, err = readOneUser(id)
	if err != nil {
		return nil, nil, err
	}

	recordBody, err = readOneRecord(recordID)
	if err != nil {
		return nil, nil, err
	}

	if userID != "" {
		recordBody, err = readOneUserRecords(userID)
		if err != nil {
			return nil, nil, err
		}
	}

	if id == "" && (recordID != "" || userID != "") {
		userBody, err = readAllUser()
		if err != nil {
			return nil, nil, err
		}
	} else if id != "" && (recordID == "" && userID == "") {
		recordBody, err = readAllRecord()
		if err != nil {
			return nil, nil, err
		}
	}

	return userBody, recordBody, err
}

func DoUpdateRequest(id string, name string, age string, recordID string, userID string, subject string, score string) ([]byte, []byte, error) {
	if id == "" {
		userBody, err = readAllUser()
	} else {
		userBody, err = updateUser(id, name, age)
	}
	if recordID == "" {
		recordBody, err = readAllRecord()
	} else {
		recordBody, err = updateRecord(recordID, userID, subject, score)
	}
	if err != nil {
		return nil, nil, err
	}
	return userBody, recordBody, nil
}

func DoDeleteRequset(id string, recordID string) ([]byte, []byte, error) {
	userBody, err = deleteUser(id)
	if err != nil {
		return nil, nil, err
	}
	recordBody, err = deleteRecord(recordID)
	if err != nil {
		return nil, nil, err
	}

	return userBody, recordBody, err
}
