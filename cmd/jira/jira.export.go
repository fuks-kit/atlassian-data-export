package main

import (
	"flag"
	"github.com/fuks-kit/atlassian-data-export/common"
	"github.com/fuks-kit/atlassian-data-export/jira"
	"log"
)

var (
	baseUrl   = flag.String("url", "http://localhost:8080", "jira base url")
	username  = flag.String("username", "admin", "jira username")
	password  = flag.String("password", "", "jira password")
	exportDir = flag.String("exportDir", "jira-export", "export directory")
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()

	exporter := jira.Exporter{HttpAuth: common.HttpAuth{
		BaseUrl:  *baseUrl,
		Username: *username,
		Password: *password,
	}}

	exporter.Export(*exportDir)

	//exporter.DownloadIssue("FIN-1090", "jira-export")

	//byt, _ := json.MarshalIndent(issue, "", "  ")
	//log.Printf("%s", byt)
	//exporter.DownloadIssueHtml("FIN-1090", "export")
}
