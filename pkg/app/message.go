package app

type ResponseMessage struct {
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"` // omitempty: if empty, omit this field todo
}
