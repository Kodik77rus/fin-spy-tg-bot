package models

type Market struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	YoohooCode string `json:"yoohooCode"`
	Location   string `json:"location"`
	Country    string `json:"country"`
	Delay      uint   `json:"delay"`
	Hour       string `json:"hour"`
}
