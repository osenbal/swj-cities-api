package utils

// Response formater
type Response struct {
	Code    int         `json:"status_code" xml:"status_code"`
	Message interface{} `json:"message" xml:"message"`
	Data    interface{} `json:"data" xml:"data"`
}
