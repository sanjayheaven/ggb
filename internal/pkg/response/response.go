package response

type Response struct {
	HttpCode int         `json:"-"`
	Code     string      `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}
