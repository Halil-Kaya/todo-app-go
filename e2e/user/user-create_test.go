package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"todo/app/model"
)

var testurl = "http://localhost:3011/user/create"

func TestUserCreate(t *testing.T) {
	fmt.Println("test çalıştı.")
	user := model.User{Nickname: "asdasd", Password: "qweqweq", FullName: "fafasfas"}
	reqBody, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", testurl, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	fmt.Println("req -> ", json.Unmarshal(body, &data))
	defer resp.Body.Close()

	fmt.Println("err -> ", data)
}
