package metabase

type MetabaseSession struct {
	MetabaseAuth
	Id string `json:"id"`
}
