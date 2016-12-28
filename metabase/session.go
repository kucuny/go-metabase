package metabase

import (
	"fmt"
)

type SessionComponent struct{
	c *Client
}

type SessionKey struct {
	ID string `json:"id",json:"session_id"`
}

func (com *SessionComponent) GetSessionKey() (*SessionKey, *HttpResponse) {
	req, err := com.c.NewRequest(requestPost, "/api/session", com.c.Auth)

	if err != nil {
		return nil, &HttpResponse{Response: nil, Err: err}
	}

	sessionKey := new(SessionKey)
	resp := com.c.Do(req, sessionKey)

	if resp.Err != nil {
		return nil, resp
	}

	return sessionKey, resp
}

func (com *SessionComponent) DeleteSessionKey() *HttpResponse {
	url := fmt.Sprintf("/api/session/?session_id=%s", com.c.Auth.SessionKey)
	req, err := com.c.NewRequest(requestDelete, url, nil)

	if err != nil {
		return &HttpResponse{Response: nil, Err: err}
	}

	sessionKey := new(SessionKey)
	resp := com.c.Do(req, sessionKey)

	if resp.Err != nil {
		return resp
	}

	return resp
}
