package jira

import (
	"encoding/json"
	"fmt"
	"github.com/fuks-kit/atlassian-data-export/common"
	"log"
	"os"
	"path/filepath"
)

type Exporter struct {
	common.HttpAuth
}

func (exporter Exporter) GetIssueIds() (issuesIds []string) {
	var results SearchResult

	for {
		startAt := results.StartAt + results.MaxResults
		log.Printf("Scrape: startAt=%d", startAt)

		apiUrl := fmt.Sprintf("/rest/api/2/search?startAt=%d", startAt)
		resp := exporter.GetWithBase(apiUrl)
		if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
			log.Panic(err)
		}

		for _, issue := range results.Issues {
			issuesIds = append(issuesIds, issue.Id)
		}

		if results.Total <= results.StartAt+results.MaxResults {
			break
		}
	}

	return
}

func (exporter Exporter) GetIssue(id string) (issue Issue) {

	//
	// http://localhost:8080/rest/api/2/issue/10318
	//

	log.Printf("GetIssue: %s", id)
	exporter.GetMarshal("http://localhost:8080/rest/api/2/issue/"+id, &issue)

	return
}

func (exporter Exporter) DownloadAttachment(attachment IssueAttachment, dir string) {
	if err := os.MkdirAll(dir, 0766); err != nil {
		log.Panic(err)
	}

	log.Printf("DownloadAttachment: %s", attachment.Filename)

	byt := exporter.GetBytes(attachment.Content)
	filename := filepath.Join(dir, attachment.Filename)
	err := os.WriteFile(filename, byt, 0755)
	if err != nil {
		log.Panic(err)
	}
}

func (exporter Exporter) DownloadIssue(id, export string) {
	issue := exporter.GetIssue(id)

	log.Printf("DownloadDoc: %s %s", issue.Fields.Project.Name, id)

	dir := filepath.Join(export, issue.Fields.Project.Name)
	if err := os.MkdirAll(dir, 0766); err != nil {
		log.Panic(err)
	}

	filename := filepath.Join(dir, id+".doc")

	url := fmt.Sprintf("http://localhost:8080/si/jira.issueviews:issue-word/%s/", id)
	byt := exporter.GetBytes(url)
	err := os.WriteFile(filename, byt, 0755)
	if err != nil {
		log.Panic(err)
	}

	attachmentDir := filepath.Join(dir, id+"_attachments")
	for _, attachment := range issue.Fields.Attachment {
		exporter.DownloadAttachment(attachment, attachmentDir)
	}
}

func (exporter Exporter) Export(exportDir string) {
	issues := exporter.GetIssueIds()
	log.Printf("issues=%d", len(issues))
	common.WriteJSON("issues.all.json", issues)

	for _, id := range issues {
		exporter.DownloadIssue(id, exportDir)
	}
}
