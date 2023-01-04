package confluence

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

const pagePathCacheFile = "confluence.pagePathIndex.json"

type PagePathIndex struct {
	Ids        map[string]string // ChildId --> ParentId
	FolderName map[string]string // Id --> Name
}

func PagePathIndexFromCache() (index PagePathIndex, ok bool) {
	if _, err := os.Stat(pagePathCacheFile); err != nil {
		return PagePathIndex{}, false
	}

	byt, err := os.ReadFile(pagePathCacheFile)
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(byt, &index)
	if err != nil {
		log.Panic(err)
	}

	return index, true
}

func NewPagePathIndex() PagePathIndex {
	return PagePathIndex{
		Ids:        make(map[string]string),
		FolderName: make(map[string]string),
	}
}

func (index PagePathIndex) AddPage(page ContentResult) {
	index.FolderName[page.Id] = page.Title
}

func (index PagePathIndex) AddParent(parent, child ContentResult) {
	index.Ids[child.Id] = parent.Id
	index.FolderName[parent.Id] = parent.Title
	index.FolderName[child.Id] = child.Title
}

func (index PagePathIndex) Filepath(pageId string) (path string) {

	for {
		parent, ok := index.Ids[pageId]
		path = filepath.Join(index.FolderName[pageId], path)

		if ok {
			pageId = parent
		} else {
			break
		}
	}

	return
}
