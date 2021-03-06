package metabase

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	headerMetabaseSession string = "X-Metabase-Session"
	defaultContentType    string = "application/json"
	defaultUserAgent      string = "go-metabase"

	requestGet    string = http.MethodGet
	requestPost   string = http.MethodPost
	requestPut    string = http.MethodPut
	requestPatch  string = http.MethodPatch
	requestDelete string = http.MethodDelete
	requestHead   string = http.MethodHead
)

type Auth struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`

	SessionKey string `json:"id,omitempty"`
}

type ApiComponents struct {
	// Activity       *ActivityComponent
	// Card           *CardComponent
	// Dashboard      *DashboardComponent
	// Database       *DatabaseComponent
	// Dataset        *DatasetComponent
	// Email          *EmailComponent
	// Field          *FieldComponent
	// Geojson        *GeojsonComponent
	// GettingStarted *GettingStartedComponent
	// Label          *LabelComponent
	// Metric         *MetricComponent
	// Notify         *NotifyComponent
	// Permission     *PermissionComponent
	// Pulse          *PulseComponent
	// Revision       *RevisionComponent
	// Segment        *SegmentComponent
	Session *SessionComponent
	// Setting        *SettingComponent
	// Setup          *SetupComponent
	// Slack          *SlackComponent
	// Table          *TableComponent
	// Tiles          *TilesComponent
	// User           *UserComponent
	// Util           *UtilComponent
}

type Client struct {
	client  *http.Client
	BaseUrl *url.URL

	*Auth
}

type Metabase struct {
	client *Client
	*ApiComponents
}

type Response struct {
	Response *http.Response
	Err      error
	Payload  interface{}
}

func newClient(baseURL, sessionKey string) (*Client, error) {
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		client:  http.DefaultClient,
		BaseUrl: parsedBaseURL,
		Auth: &Auth{
			SessionKey: sessionKey,
		},
	}, nil
}

func NewMetabase(baseURL, sessionKey string) (*Metabase, error) {
	client, err := newClient(baseURL, sessionKey)

	if err != nil {
		return nil, err
	}

	metabaseClient := &Metabase{
		client: client,
		ApiComponents: &ApiComponents{
			Session: &SessionComponent{c: client},
		},
	}

	return metabaseClient, nil
}

func (m *Metabase) SetAuth(email, password string) {
	m.client.Auth.Email = email
	m.client.Auth.Password = password
}

func (m *Metabase) SetSessionKey(sessionKey string) {
	m.client.Auth.SessionKey = sessionKey
}

func (c *Client) NewRequest(method, urlStr string, v interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseUrl.ResolveReference(rel)

	var buf io.ReadWriter
	if v != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(v)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", defaultUserAgent)

	if v != nil {
		req.Header.Set("Content-Type", defaultContentType)
	}

	if c.Auth.SessionKey != "" {
		req.Header.Set(headerMetabaseSession, c.Auth.SessionKey)
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, payload interface{}) *Response {
	resp, err := c.client.Do(req)

	resultResponse := &Response{
		Response: resp,
		Err:      err,
		Payload:  payload,
	}

	if resultResponse.Err != nil {
		return resultResponse
	}

	defer resultResponse.Response.Body.Close()

	if resultResponse.Payload != nil {
		err = json.NewDecoder(resultResponse.Response.Body).Decode(resultResponse.Payload)
		if err == io.EOF {
			resultResponse.Err = nil
		} else {
			resultResponse.Err = err
		}

	}

	return resultResponse
}
