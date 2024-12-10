package domains

type ResponseInsert struct {
	RowAffected int `json:"row_affected"`
}

type ResponseFind struct {
	Data  interface{} `json:"data"`
	Page  int         `json:"page"`
	Total int         `json:"total"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
