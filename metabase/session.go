package metabase

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Session struct{}

type SessionInfo struct {
	email      string `json:"email",omitempty`
	password   string `json:"password",omitempty`
	sessionKey string `json:"id",omitempty`
}

func (s *Session) GetSessionKey(email, password string) (string, error) {
	req := makeRequest(http.MethodPost, "session")
	payload := SessionInfo{
		email:    email,
		password: password,
	}
	req.Body, _ = json.Marshal(payload)

	client := &http.Client{}

	response, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	result := new(SessionInfo)
	result, err := json.Unmarshal(data, result)

	return result.sessionKey, err
}

func (s *Session) DeleteSessionKey() error {
	req := makeRequest(http.MethodDelete, "session")
}
