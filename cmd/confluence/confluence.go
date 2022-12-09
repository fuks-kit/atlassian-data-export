package main

import (
	"flag"
	"github.com/fuks-kit/atlassian-data-export/confluence"
	"log"
)

var (
	baseUrl  = flag.String("url", "http://localhost:8090", "confluence base url")
	username = flag.String("username", "", "confluence username")
	password = flag.String("password", "", "confluence password")
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()

	downloader := confluence.Downloader{
		BaseUrl:  *baseUrl,
		Username: *username,
		Password: *password,
	}

	downloader.Download("export")
}
