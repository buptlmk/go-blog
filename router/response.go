package router

type Response struct {
	State   int         `json:"state"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
