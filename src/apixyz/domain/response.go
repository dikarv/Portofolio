package domain

type (
	Response struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	Response2 struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Result  interface{} `json:"result"`
	}

	ErrorData struct {
		Status  int   `json:"status"`
		Message error `json:"message"`
	}
)
