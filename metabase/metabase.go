package metabase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	headerMetabaseSession string = "X-Metabase-Session"
	defaultContentType    string = "application/json"
)

type Auth struct {
	baseUrl    *url.URL
	email      string `json:"email"`
	password   string `json:"password"`
	sessionKey string
}

type ApiComponent struct {
	activity       *Activity
	card           *Card
	dashboard      *Dashboard
	database       *Database
	dataset        *Dataset
	email          *Email
	field          *Field
	geojson        *Geojson
	gettingStarted *GettingStarted
	label          *Label
	metric         *Metric
	notify         *Notify
	permission     *Permission
	pulse          *Pulse
	revision       *Revision
	segment        *Segment
	session        *Session
	setting        *Setting
	setup          *Setup
	slack          *Slack
	table          *Table
	tiles          *Tiles
	user           *User
	util           *Util
}

type Metabase struct {
	client  *http.Client
	BaseUrl *url.URL
	Auth
	ApiComponent
}

func NewMetabase(baseUrl, sessionKey string) *Metabase {
	return &Metabase{
		client:  &http.Client{},
		BaseUrl: url.Parse(baseUrl),
		Auth: Auth{
			sessionKey: sessionKey,
		},
	}
}

func (m *Metabase) SetAuth(email, password string) {
	m.Auth.email = email
	m.Auth.password = password
}

func (m *Metabase) makeRequest(method, urlStr string, body interface{}) *http.Request {
	rel, err := url.Parse(urlStr)

	if err != nil {
		panic(err)
	}

	header := DefaultHeader
	header.Add(headerMetabaseSession, m.sessionKey)

	return &http.Request{
		Method: method,
		URL:    rel,
		Header: header,
		Body:   body,
	}
}
