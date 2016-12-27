package main

import (
	"fmt"
	"github.com/kucuny/go-metabase/metabase"
)

func main() {
	client := metabase.NewMetabase("http://localhost:3000", "")
	client.SetAuth("kucuny@gmail.com", "qwer1234")
	fmt.Println(client.Session.GetSessionKey())
}
