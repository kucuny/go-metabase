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

func (com *SessionComponent) GetSessionKey() (*SessionKey, *Response) {
	req, err := com.c.NewRequest(requestPost, "/api/session", com.c.Auth)

	if err != nil {
		return nil, &Response{Response: nil, Err: err}
	}

	sessionKey := new(SessionKey)
	resp := com.c.Do(req, sessionKey)

	if resp.Err != nil {
		return nil, resp
	}

	return sessionKey, resp
}

func (com *SessionComponent) DeleteSessionKey() *Response {
	url := fmt.Sprintf("/api/session/?session_id=%s", com.c.Auth.SessionKey)
	req, err := com.c.NewRequest(requestDelete, url, nil)

	if err != nil {
		return &Response{Response: nil, Err: err}
	}

	sessionKey := new(SessionKey)
	resp := com.c.Do(req, sessionKey)

	if resp.Err != nil {
		return resp
	}

	return resp
}
