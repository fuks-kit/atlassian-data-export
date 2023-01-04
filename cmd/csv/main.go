package main

import (
	"encoding/json"
	"fmt"
	"github.com/fuks-kit/atlassian-data-export/confluence"
	"log"
	"os"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	byt, err := os.ReadFile("sample/confluence.content.json")
	if err != nil {
		log.Panicln(err)
	}

	spaceIndex := make(map[string]string)
	var content []confluence.ContentResult
	err = json.Unmarshal(byt, &content)
	if err != nil {
		log.Panicln(err)
	}

	for _, data := range content {
		spaceIndex[data.Id] = data.Expandable.Space[len("/rest/api/space/"):]
	}

	index, _ := confluence.PagePathIndexFromCache()
	csvString := "Space,PageName,PagePath,Url\n"

	for pageId := range index.FolderName {
		path := index.Filepath(pageId)
		csvString += fmt.Sprintf("%s,%s,\"%s\",%s\n",
			spaceIndex[pageId], index.FolderName[pageId], path, "http://185.233.107.158/pages/viewpage.action?pageId="+pageId)
	}

	err = os.WriteFile("confluence.csv", []byte(csvString), 0766)
	if err != nil {
		log.Panicln(err)
	}
}
