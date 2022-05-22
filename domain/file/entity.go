package file

import "time"

type Result struct {
	IsError bool   `json:"is_error"`
	Data    string `json:"data"`
}

type BoxResponse struct {
	TotalCount int       `json:"total_count"`
	Entries    []Entries `json:"entries"`
}

type Entries struct {
	Type       string     `json:"type"`
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	Size       int        `json:"size"`
	SharedLink SharedLink `json:"shared_link"`
}

type ResearchReport struct {
	FileKey       string    `json:"file_key"`
	Title         string    `json:"title"`
	Size          int       `json:"size"`
	CreatedAt     time.Time `json:"created_at"`
	LastUpdatedAt time.Time `json:"last_updated_at"`
	Link          string    `json:"box_shared_link"`
	FileId        string    `json:"box_file_id"`
	TotalDownload int       `json:"total_download"`
}

type SharedLink struct {
	Url string `json:"url"`
}
