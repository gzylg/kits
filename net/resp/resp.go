package resp

type Content struct {
	Code    int         `json:"code"`
	Result  bool        `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Total   *int64      `json:"total,omitempty"`
}
