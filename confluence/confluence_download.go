package confluence

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Downloader struct {
	BaseUrl  string
	Username string
	Password string
}

func (downloader Downloader) request(url string) (resp *http.Response) {
	req, err := http.NewRequest("GET", url, nil)
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

func (downloader Downloader) Content() (results ContentResults) {

	apiUrl := downloader.BaseUrl + "/rest/api/content?type=page&start=0&limit=99999"
	resp := downloader.request(apiUrl)

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		log.Panic(err)
	}

	//byt, err := json.MarshalIndent(results, "", "  ")
	//err = os.WriteFile("out.json", byt, 0755)
	//if err != nil {
	//	log.Panic(err)
	//}

	return
}

func (downloader Downloader) GetPDF(pageId string) (byt []byte) {

	apiUrl := downloader.BaseUrl + "/spaces/flyingpdf/pdfpageexport.action?pageId=" + pageId
	resp := downloader.request(apiUrl)

	byt, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	return byt
}

func (downloader Downloader) GetAttachments(pageId string) (results AttachmentResults) {

	apiUrl := downloader.BaseUrl + "/rest/api/content/" + pageId + "/child/attachment"
	resp := downloader.request(apiUrl)

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		log.Panic(err)
	}

	return
}

func (downloader Downloader) GetAttachment(attachment AttachmentResult) (byt []byte) {

	// http://localhost:8090/download/attachments/3834128/Freigaben%20bei%20fuks.pptx?api=v2
	apiUrl := downloader.BaseUrl + attachment.Links.Download
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

	for _, page := range downloader.Content().Results {
		space := strings.TrimPrefix(page.Expandable.Space, "/rest/api/space/")

		if space == "ADMIN" {
			// Skip confluence sample data
			continue
		}

		dir := filepath.Join(exportDir, space)
		if err := os.MkdirAll(dir, 0766); err != nil {
			log.Panic(err)
		}

		log.Printf("%s %s (%s)", space, page.Title, page.Id)
		pdf := downloader.GetPDF(page.Id)

		title := page.Title
		title = strings.ReplaceAll(title, ".", "_")
		title = strings.ReplaceAll(title, "/", "_")
		title = strings.ReplaceAll(title, "\\", "_")

		err := os.WriteFile(filepath.Join(dir, title+".pdf"), pdf, 0766)
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

			byt := downloader.GetAttachment(attachment)
			err = os.WriteFile(filepath.Join(attachmentsDir, attachmentTitle), byt, 0766)
			if err != nil {
				log.Panic(err)
			}
		}
	}
}
