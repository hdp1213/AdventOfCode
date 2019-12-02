package main


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)


type secrets struct {
	SessionCookie string `json:"session-cookie"`
	VerificationCode string `json:"verification-code"`
}

const baseUrl = "https://adventofcode.com/2019/day/%d/input"
const inputDir = "input"


func loadSecretFile(filename string) (*secrets, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := new(secrets)
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}


func getInput(day int) (string, error) {
	if _, err := os.Stat(inputDir); os.IsNotExist(err) {
		os.Mkdir(inputDir, os.ModePerm)
	}

	inputFile := path.Join(inputDir, fmt.Sprintf("day%02d", day))

	if _, err := os.Stat(inputFile); err == nil {
		return inputFile, nil
	}

	config, err := loadSecretFile("secrets.json")
	if err != nil {
		return "", err
	}

	client := new(http.Client)

	req, err := http.NewRequest("GET", fmt.Sprintf(baseUrl, day), nil)
	if err != nil {
		return "", err
	}

	cookie := new(http.Cookie)
	cookie.Name, cookie.Value = "session", config.SessionCookie
	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(inputFile, body, 0644)
	if err != nil {
		return "", err
	}

	return inputFile, nil
}
