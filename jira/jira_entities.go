package jira

type SearchResult struct {
	StartAt    int `json:"startAt"`
	MaxResults int `json:"maxResults"`
	Total      int `json:"total"`
	Issues     []struct {
		Id   string `json:"id"`
		Self string `json:"self"`
		Key  string `json:"key"`
	} `json:"issues"`
}

type Issue struct {
	Id     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields struct {
		Attachment []IssueAttachment `json:"attachment"`
		Project    struct {
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"project"`
	} `json:"fields"`
}

type IssueAttachment struct {
	Self     string `json:"self"`
	Id       string `json:"id"`
	Filename string `json:"filename"`
	Size     int    `json:"size"`
	MimeType string `json:"mimeType"`
	Content  string `json:"content"`
}
