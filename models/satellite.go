package models

type SateliteInfo struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}
