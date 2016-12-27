package metabase

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"io"
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
	baseUrl *url.URL

	email    string `json:"email"`
	password string `json:"password"`

	sessionKey string
}

type ApiComponent struct {
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
	Client *Client
	*ApiComponent
}

func newClient(baseUrl, sessionKey string) *Client {
	base, err := url.Parse(baseUrl)
	if err != nil {
		return nil
	}

	return &Client{
		client:  http.DefaultClient,
		BaseUrl: base,
		Auth: &Auth{
			sessionKey: sessionKey,
		},
	}
}

func NewMetabase(baseUrl, sessionKey string) *Metabase {
	client := newClient(baseUrl, sessionKey)
	metabaseClient := &Metabase{
		Client:  client,
		ApiComponent: &ApiComponent{
			Session: &SessionComponent{client: client},
		},
	}

	return metabaseClient
}

func (m *Metabase) SetAuth(email, password string) {
	m.Client.Auth.email = email
	m.Client.Auth.password = password
}

func (m *Metabase) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := m.Client.BaseUrl.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", defaultContentType)
	}

	req.Header.Set("User-Agent", defaultUserAgent)
	req.Header.Set(headerMetabaseSession, m.Client.Auth.sessionKey)

	return req, nil
}
