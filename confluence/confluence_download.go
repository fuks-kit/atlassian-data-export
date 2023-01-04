package confluence

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/fuks-kit/atlassian-data-export/common"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type Downloader struct {
	BaseUrl  string
	Username string
	Password string
}

func (downloader Downloader) request(src *url.URL) (resp *http.Response) {
	req, err := http.NewRequest("GET", src.String(), nil)
	if err != nil {
		log.Panic(err)
	}

	authKey := base64.StdEncoding.EncodeToString([]byte(downloader.Username + ":" + downloader.Password))
	req.Header.Set("Authorization", "Basic "+authKey)
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		log.Panic(err)
	}

	return resp
}

func (downloader Downloader) GetMarshall(src *url.URL, obj any) {
	resp := downloader.request(src)
	if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
		log.Panic(err)
	}
}

func (downloader Downloader) FetchList(apiUrl *url.URL) (pages []ContentResult) {

	query := apiUrl.Query()
	query.Set("limit", "200")
	query.Set("start", "0")

	var results ContentResults

	for {
		startAt := results.Start + results.Limit
		query.Set("start", fmt.Sprint(startAt))
		apiUrl.RawQuery = query.Encode()

		// log.Printf("Scrape: apiUrl=%s", apiUrl.String())

		resp := downloader.request(apiUrl)
		if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
			log.Panic(err)
		}

		pages = append(pages, results.Results...)

		if results.Size < results.Limit {
			break
		}
	}

	return
}

func (downloader Downloader) Content() (pages []ContentResult) {
	apiUrl, _ := url.Parse(downloader.BaseUrl + "/rest/api/content?type=page")
	return downloader.FetchList(apiUrl)
}

func (downloader Downloader) GetPDF(pageId string) (byt []byte) {

	apiUrl, _ := url.Parse(downloader.BaseUrl + "/spaces/flyingpdf/pdfpageexport.action?pageId=" + pageId)
	resp := downloader.request(apiUrl)

	byt, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	return byt
}

func (downloader Downloader) GetChildPages(pageId string) (childPages []ContentResult) {
	apiUrl, _ := url.Parse(downloader.BaseUrl + "/rest/api/content/" + pageId + "/child/page")
	return downloader.FetchList(apiUrl)
}

func (downloader Downloader) GetAttachments(pageId string) (results AttachmentResults) {

	apiUrl, _ := url.Parse(downloader.BaseUrl + "/rest/api/content/" + pageId + "/child/attachment")
	resp := downloader.request(apiUrl)

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		log.Panic(err)
	}

	return
}

func (downloader Downloader) GetAttachment(attachment AttachmentResult) (byt []byte) {

	// http://localhost:8090/download/attachments/3834128/Freigaben%20bei%20fuks.pptx?api=v2
	apiUrl, _ := url.Parse(downloader.BaseUrl + attachment.Links.Download)
	resp := downloader.request(apiUrl)

	if resp.StatusCode != 200 {
		log.Printf("(%d) %s", resp.StatusCode, apiUrl)
	}

	byt, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	return byt
}

func (downloader Downloader) Export(exportDir string) {

	pages := downloader.Content()

	//
	// Build Confluence Wiki Structure
	//

	log.Printf("Building Confluence page structure")

	pathIndex, ok := PagePathIndexFromCache()
	if !ok {
		pathIndex = NewPagePathIndex()
		for _, parent := range pages {
			pathIndex.AddPage(parent)
			children := downloader.GetChildPages(parent.Id)
			for _, child := range children {
				pathIndex.AddParent(parent, child)
			}
		}

		common.WriteJSON(pagePathCacheFile, pathIndex)
	}

	//
	// Download Content
	//

	log.Printf("Export Confluence data")

	for inx, page := range pages {
		space := strings.TrimPrefix(page.Expandable.Space, "/rest/api/space/")

		if space == "ADMIN" {
			// Skip confluence sample data
			continue
		}

		dir := filepath.Join(exportDir, space, pathIndex.Filepath(page.Id))
		if err := os.MkdirAll(dir, 0766); err != nil {
			log.Panic(err)
		}

		title := page.Title
		title = strings.ReplaceAll(title, ".", "_")
		title = strings.ReplaceAll(title, "/", "_")
		title = strings.ReplaceAll(title, "\\", "_")
		filename := filepath.Join(dir, title+".pdf")

		log.Printf("(%d/%d) [page.Id=%s] %s", inx, len(pages), page.Id, filename)

		pdf := downloader.GetPDF(page.Id)
		err := os.WriteFile(filename, pdf, 0766)
		if err != nil {
			log.Panic(err)
		}

		attachments := downloader.GetAttachments(page.Id)
		for _, attachment := range attachments.Results {
			log.Println("\t", attachment.Title)

			attachmentsDir := filepath.Join(dir, title+"_attachments")
			if err = os.MkdirAll(attachmentsDir, 0766); err != nil {
				log.Panic(err)
			}

			attachmentTitle := attachment.Title
			attachmentTitle = strings.ReplaceAll(attachmentTitle, "/", "_")
			attachmentTitle = strings.ReplaceAll(attachmentTitle, "\\", "_")
			attachmentFile := filepath.Join(attachmentsDir, attachmentTitle)

			byt := downloader.GetAttachment(attachment)
			err = os.WriteFile(attachmentFile, byt, 0766)
			if err != nil {
				log.Panic(err)
			}
		}
	}
}
