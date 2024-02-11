package web

type ApiResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Error  interface{} `json:"error"`
}

type ValidationErrorResponse struct {
	Code   int                    `json:"code"`
	Status string                 `json:"status"`
	Errors map[string]interface{} `json:"errors"`
}
