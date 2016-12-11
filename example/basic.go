package main

import (
	"github.com/kucuny/go-metabase"
)

func main() {
	client := metabase.NewMetabase("http://localhost:3000", "5d4d2f49-2b69-4bac-ab16-2e537dc15f6b")
}
