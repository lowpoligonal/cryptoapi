package response

type Request struct {
	Text string `json:"text" binding:"required"`
	Key  int    `json:"key"`
}
