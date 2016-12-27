package metabase

type SessionComponent struct{
	client *Client
}

type SessionKey struct {
	ID string `json:"id"`
}

func (s *SessionComponent) GetSessionKey(email, password string) (string, error) {
	req, err := s.client.NewRequest(requestPost, "/api/session", client.Auth)
	if err != nil {
		return "", err
	}

	var key *SessionKey
	response, err := client.Do(req, key)

	if err != nil {
		return "", err
	}

	return key.ID, nil
}
