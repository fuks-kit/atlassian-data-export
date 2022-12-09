package confluence

type ContentResults struct {
	Results []ContentResult `json:"results"`
	Start   int             `json:"start"`
	Limit   int             `json:"limit"`
	Size    int             `json:"size"`
	Links   struct {
		Self    string `json:"self"`
		Next    string `json:"next"`
		Base    string `json:"base"`
		Context string `json:"context"`
	} `json:"_links"`
}

type ContentResult struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Status     string `json:"status"`
	Title      string `json:"title"`
	Extensions struct {
		Position interface{} `json:"position"`
	} `json:"extensions"`
	Links struct {
		Webui  string `json:"webui"`
		Edit   string `json:"edit"`
		Tinyui string `json:"tinyui"`
		Self   string `json:"self"`
	} `json:"_links"`
	Expandable struct {
		Container    string `json:"container"`
		Metadata     string `json:"metadata"`
		Operations   string `json:"operations"`
		Children     string `json:"children"`
		Restrictions string `json:"restrictions"`
		History      string `json:"history"`
		Ancestors    string `json:"ancestors"`
		Body         string `json:"body"`
		Version      string `json:"version"`
		Descendants  string `json:"descendants"`
		Space        string `json:"space"`
	} `json:"_expandable"`
}

type AttachmentResults struct {
	Results []AttachmentResult `json:"results"`
	Start   int                `json:"start"`
	Limit   int                `json:"limit"`
	Size    int                `json:"size"`
	Links   struct {
		Self    string `json:"self"`
		Base    string `json:"base"`
		Context string `json:"context"`
	} `json:"_links"`
}

type AttachmentResult struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	Status   string `json:"status"`
	Title    string `json:"title"`
	Metadata struct {
		MediaType string `json:"mediaType"`
		Labels    struct {
			Results []interface{} `json:"results"`
			Start   int           `json:"start"`
			Limit   int           `json:"limit"`
			Size    int           `json:"size"`
			Links   struct {
				Self string `json:"self"`
			} `json:"_links"`
		} `json:"labels"`
		Expandable struct {
			Currentuser string `json:"currentuser"`
			Properties  string `json:"properties"`
			Frontend    string `json:"frontend"`
			EditorHtml  string `json:"editorHtml"`
		} `json:"_expandable"`
	} `json:"metadata"`
	Extensions struct {
		MediaType string `json:"mediaType"`
		FileSize  int    `json:"fileSize"`
		Comment   string `json:"comment"`
	} `json:"extensions"`
	Links struct {
		Webui    string `json:"webui"`
		Download string `json:"download"`
		Self     string `json:"self"`
	} `json:"_links"`
	Expandable struct {
		Container    string `json:"container"`
		Operations   string `json:"operations"`
		Children     string `json:"children"`
		Restrictions string `json:"restrictions"`
		History      string `json:"history"`
		Ancestors    string `json:"ancestors"`
		Body         string `json:"body"`
		Version      string `json:"version"`
		Descendants  string `json:"descendants"`
		Space        string `json:"space"`
	} `json:"_expandable"`
}
