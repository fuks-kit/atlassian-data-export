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

	authHttp := common.HttpAuth{
		BaseUrl:  *baseUrl,
		Username: *username,
		Password: *password,
	}

	exporter := jira.Exporter{HttpAuth: authHttp}

	issues := exporter.GetIssueIds()
	//byt, err := json.MarshalIndent(issues, "", "  ")
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//err = os.WriteFile("jira.issues.json", byt, 0755)
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//byt, err := os.ReadFile("jira.issues.json")
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//var issues []string
	//err = json.Unmarshal(byt, &issues)
	//if err != nil {
	//	log.Panic(err)
	//}

	exporter.ExportAll(issues, *exportDir)
}
