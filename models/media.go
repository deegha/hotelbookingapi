package models

type Media struct {
	ID        int   `json:"id"`
	Type      int    `json:"type"`
	URL       string `json:"url"`
}


func (media *Media) setMedia(newMedia Media) {
	media.URL = newMedia.URL
	media.Type = newMedia.Type
}
