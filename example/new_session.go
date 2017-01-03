package main

import (
	"fmt"
	"github.com/kucuny/go-metabase/metabase"
    "reflect"
)

func main() {
	client, _ := metabase.NewMetabase("http://localhost:3001", "331e95d8-c334-4ffd-ace9-066987c94610")
	client.SetAuth("kucuny@gmail.com", "qwer1234")
	sessionKey, resp := client.Session.GetSessionKey()
    fmt.Println(sessionKey, resp)

    client.SetSessionKey(sessionKey.ID)
    response := client.Session.DeleteSessionKey()
    fmt.Println(response.Payload.(*value.Type()).ID)

    value := reflect.ValueOf(response.Payload)
    fmt.Println(value.Type())
}
