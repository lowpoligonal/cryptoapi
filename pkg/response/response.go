package response

type Request struct {
	Text   string `json:"text" binding:"required"`
	KeyStr string `json:"keystr,omitempty"`
	KeyInt int    `json:"keyint,omitempty"`
}
