package schema

type Request struct {
	Search string `json:"search,omitempty"`
}

// Response schema
type Response struct {
	Code    int    `json:"code,omitempty"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
