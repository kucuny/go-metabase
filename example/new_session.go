package main

import (
	"fmt"
	"github.com/kucuny/go-metabase/metabase"
)

func main() {
	client, _ := metabase.NewMetabase("http://localhost:3001", "331e95d8-c334-4ffd-ace9-066987c94610")
	client.SetAuth("kucuny@gmail.com", "qwer1234")
	sessionKey, resp := client.Session.GetSessionKey()
    fmt.Println(sessionKey, resp)

    client.SetSessionKey(sessionKey.ID)
    fmt.Println(client.Session.DeleteSessionKey())
}
