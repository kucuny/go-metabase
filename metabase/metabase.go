package metabase

type MetabaseAuth struct {
	email      string
	password   string
	sessionKey string
}

type Metabase struct {
	BaseUrl string
	MetabaseAuth
}

func NewMetabase(baseUrl, sessionKey string) {
	return &Metabase{
		BaseUrl: baseUrl,
		&MetabaseAuth{
			sessionKey: sessionKey,
		},
	}
}
