package handler

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

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
	req, _ := http.NewRequest("POST", "http://localhost:8000/users", bytes.NewBuffer(jsonStr))
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func DoReadAllRequest() ([]byte, error) {
	req, _ := http.NewRequest("GET", "http://localhost:8000/users", nil)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func DoReadOneRequest(id string) ([]byte, error) {
	req, _ := http.NewRequest("GET", "http://localhost:8000/users/"+id, nil)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
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
	req, _ := http.NewRequest("PATCH", "http://localhost:8000/users/"+id, bytes.NewBuffer(jsonStr))
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func DoDeleteRequset(id string) ([]byte, error) {
	req, _ := http.NewRequest("DELETE", "http://localhost:8000/users/"+id, nil)
	body, err := ClientDo(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}
