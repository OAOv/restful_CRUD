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

func DoCreateRequest(id string, name string, age string) ([]byte, error) {
	var jsonStr = []byte("{\"id\":\"" + id + "\",\"name\":\"" + name + "\",\"age\":\"" + age + "\"}")
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8000/user", bytes.NewBuffer(jsonStr))
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func DoReadAllRequest() ([]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8000/users", nil)
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, err
}

func DoReadOneRequest(id string) ([]byte, error) {
	if id == "" {
		id = "-1"
	}
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8000/user/"+id, nil)
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, err
}

func DoUpdateRequest(id string, name string, age string) ([]byte, error) {
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

func DoDeleteRequset(id string) ([]byte, error) {
	req, _ := http.NewRequest("DELETE", "http://localhost:8000/user/"+id, nil)
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}
