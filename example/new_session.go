package main

import (
	"github.com/kucuny/go-metabase/metabase"
)

func main() {
	client := metabase.NewMetabase("http://localhost:3000", "")
	client.SetAuth("kucuny@gmail.com", "qwer1234")
	client.GetSessionKey()
}
