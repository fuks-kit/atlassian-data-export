package main

import (
	"encoding/json"
	"flag"
	"github.com/fuks-kit/atlassian-data-export/common"
	"github.com/fuks-kit/atlassian-data-export/jira"
	"log"
	"os"
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

	authHttp := common.HttpAuth{
		BaseUrl:  *baseUrl,
		Username: *username,
		Password: *password,
	}

	exporter := jira.Exporter{HttpAuth: authHttp}

	//byt := authHttp.GetBytes(*baseUrl + "/rest/api/2/search?startAt=400")
	//err := os.WriteFile("sample/jira.search.json", byt, 0755)
	//if err != nil {
	//	log.Panic(err)
	//}

	//issues := exporter.GetIssueIds()
	//byt, err := json.MarshalIndent(issues, "", "  ")
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//err = os.WriteFile("jira.issues.json", byt, 0755)
	//if err != nil {
	//	log.Panic(err)
	//}

	byt, err := os.ReadFile("jira.issues.json")
	if err != nil {
		log.Panic(err)
	}

	var issues []string
	err = json.Unmarshal(byt, &issues)
	if err != nil {
		log.Panic(err)
	}

	exporter.ExportAll(issues, *exportDir)

	//exporter.DownloadIssue("FIN-1090", "jira-export")

	//issue := exporter.GetIssue("10131")
	//byt, _ := json.MarshalIndent(issue, "", "  ")
	//log.Printf("%s", byt)
	//exporter.DownloadIssueHtml("FIN-1090", "export")
}
