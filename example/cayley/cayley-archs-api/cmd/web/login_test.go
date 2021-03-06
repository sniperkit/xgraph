package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/oren/doc-api"
)

// POST /adminlogin
// expect 200
func TestLogin(t *testing.T) {
	email := "vera@gmail.com"
	password := "112233"

	_ = login(t, email, password)
}

func login(t *testing.T, email string, password string) string {
	url := "http://localhost:3000/adminlogin"

	adminLogin := &admin.EmailAndPassword{
		Email:    email,
		Password: password,
	}

	js, err := json.Marshal(adminLogin)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(js))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		t.Fatal("POST /adminlogin returned", resp.Status)
		return ""
	}

	user := &admin.User{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&user)
	if err != nil {
		t.Fatal("Error in decoding json", err)
		return ""
	}

	return user.JWT
}
