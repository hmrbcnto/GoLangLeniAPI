package models

type Fact struct {
	ID       string `json:"id"`
	Fact     string `json:"fact"`
	AddedBy  string `json:"added_by"`
	EditedBy string `json:"edited_by"`
	Source   string `json:"source"`
}
