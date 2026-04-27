package response

type Response struct {
	Text   string `json:"text" binding:"required"`
	KeyStr string `json:"keystr,omitempty"`
	KeyInt int    `json:"keyint,omitempty"`
}
