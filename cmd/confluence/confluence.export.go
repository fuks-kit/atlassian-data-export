package main

import (
	"flag"
	"github.com/fuks-kit/atlassian-data-export/confluence"
	"log"
)

var (
	baseUrl   = flag.String("url", "http://localhost:8090", "confluence base url")
	username  = flag.String("username", "admin", "confluence username")
	password  = flag.String("password", "", "confluence password")
	exportDir = flag.String("exportDir", "confluence-export", "export directory")
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()

	downloader := confluence.Downloader{
		BaseUrl:  *baseUrl,
		Username: *username,
		Password: *password,
	}

	downloader.Export(*exportDir)
	//content := downloader.Content()
	//
	//byt, err := json.MarshalIndent(content, "", "  ")
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//err = os.WriteFile("sample/confluence.content.json", byt, 0755)
	//if err != nil {
	//	log.Panic(err)
	//}
}
