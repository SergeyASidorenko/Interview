package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponseStatus(t *testing.T) {
	type JsonRequest struct {
		Id   string
		Text string
	}

	jsonrequest := JsonRequest{Id: "3", Text: "Hello"}
	jsonbody, _ := json.Marshal(jsonrequest)
	ts := httptest.NewServer(GetServer())
	defer ts.Close()
	req, err := http.NewRequest("POST", "http://localhost:8020/md5", bytes.NewBuffer(jsonbody))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Log("Ошибка при обработке запроса сервером")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Log("Ошибка сервера! Код возврата отличен от 200 (OK)")
		t.Fail()
	}
	body, _ := ioutil.ReadAll(resp.Body)

	if len(string(body)) == 0 || len(string(body)) < 19 {
		t.Log("Ошибка! Ответ пустой или не полон!")
		t.Fail()
	}

}
