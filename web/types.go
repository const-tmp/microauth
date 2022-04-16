package web

type APIResponse struct {
	OK     bool        `json:"ok"`
	Result interface{} `json:"result,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}
