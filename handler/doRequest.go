package handler

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

var url = "http://localhost:8000/users"

func ClientDo(req *http.Request) ([]byte, error) {
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}

func DoCreateRequest(name string, age string) ([]byte, error) {
	var jsonStr = []byte("{\"name\":\"" + name + "\",\"age\":\"" + age + "\"}")
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func DoReadAllRequest() ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, err
}

func DoReadOneRequest(id string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url+"/"+id, nil)
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, err
}

func DoUpdateRequest(id string, name string, age string) ([]byte, error) {
	var jsonStr []byte
	if name != "" && age != "" {
		jsonStr = []byte("{\"name\":\"" + name + "\",\"age\":\"" + age + "\"}")
	} else if name != "" && age == "" {
		jsonStr = []byte("{\"name\":\"" + name + "\"}")
	} else if name == "" && age != "" {
		jsonStr = []byte("{\"age\":\"" + age + "\"}")
	}
	req, _ := http.NewRequest("PATCH", url+"/"+id, bytes.NewBuffer(jsonStr))
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func DoDeleteRequset(id string) ([]byte, error) {
	req, _ := http.NewRequest("DELETE", url+"/"+id, nil)
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}
