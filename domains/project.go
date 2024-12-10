package domains

type ProjectInsertReq struct {
	Title          string `json:"title"`
	Client         string `json:"client"`
	DateCreation   string `json:"date_creation"`
	DateCompletion string `json:"date_completion"`
	Description    string `json:"description"`
	Position       string `json:"position"`
}

type ProjectFetchRes struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	Client         string `json:"client"`
	DateCreation   string `json:"date_creation"`
	DateCompletion string `json:"date_completion"`
	Description    string `json:"description"`
	Position       string `json:"position"`
}
