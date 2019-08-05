package metabase

import (
	"fmt"
)

type SessionComponent struct {
	c *Client
}

type SessionID struct {
	ID string `json:"id",json:"session_id"`
}

func (com *SessionComponent) GetSessionID() (*SessionID, *Response) {
	req, err := com.c.NewRequest(requestPost, "/api/session", com.c.Auth)

	if err != nil {
		return nil, &Response{Response: nil, Err: err}
	}

	sessionKey := new(SessionID)
	resp := com.c.Do(req, sessionKey)

	if resp.Err != nil {
		return nil, resp
	}

	return sessionKey, resp
}

func (com *SessionComponent) DeleteSessionID() *Response {
	url := fmt.Sprintf("/api/session/?session_id=%s", com.c.Auth.SessionID)
	req, err := com.c.NewRequest(requestDelete, url, nil)

	if err != nil {
		return &Response{Response: nil, Err: err}
	}

	sessionKey := new(SessionID)
	resp := com.c.Do(req, sessionKey)

	if resp.Err != nil {
		return resp
	}

	return resp
}
