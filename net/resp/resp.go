package resp

type Content struct {
	Code    int    `json:"code"`
	Result  bool   `json:"result"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Total   *int64 `json:"total,omitempty"`
}
