package stat

import "time"


type GetStatRequest struct {
	From time.Time `json:"from"`
	To time.Time `json:"to"`
	By string `json:"by"`
}

type GetStatResponse struct {
	Period string `json:"period"`
	Sum int `json:"sum"`
}
