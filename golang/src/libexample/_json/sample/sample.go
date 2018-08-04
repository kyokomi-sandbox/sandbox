package sample

type Media struct {
	DisplayURL    string `json:"display_url"`
	ExpandedURL   string `json:"expanded_url"`
	ID            int    `json:"id"`
	IDStr         string `json:"id_str"`
	MediaURL      string `json:"media_url"`
	MediaURLHTTPS string `json:"media_url_https"`
	Type          string `json:"type"`
	URL           string `json:"url"`
}
