package models

type TopSecretRequest struct {
	Satellites []SateliteInfo `json:"satellites"`
}
