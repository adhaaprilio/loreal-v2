package helper

type Paginate struct {
	Page      int `json:"page"`
	PerPage   int `json:"perpage"`
	Total     int `json:"total"`
	TotalPage int `json:"totalpage"`
}

type ResponseParams struct {
	StatusCode int
	Message    string
	Paginate   *Paginate
	Data       any
}

type ResponseData struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Response(params ResponseParams) any {
	response := &ResponseData{
		Message: params.Message,
		Data:    params.Data,
	}
	return response
}
