package entities

type Report struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	State   string `json:"state"`
}
